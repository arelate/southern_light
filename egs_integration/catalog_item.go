package egs_integration

import (
	"bytes"
	"encoding/json/v2"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/arelate/southern_light/steam_grid"
	"github.com/arelate/southern_light/vangogh_integration"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/redux"
)

type TypeValue struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type CatalogItem struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	KeyImages   []struct {
		Type         string    `json:"type"`
		Url          string    `json:"url"`
		Md5          string    `json:"md5"`
		Width        int       `json:"width"`
		Height       int       `json:"height"`
		Size         int       `json:"size"`
		UploadedDate time.Time `json:"uploadedDate"`
	} `json:"keyImages"`
	Categories []struct {
		Path string `json:"path"`
	} `json:"categories"`
	Namespace        string               `json:"namespace"`
	Status           string               `json:"status"`
	CreationDate     time.Time            `json:"creationDate"`
	LastModifiedDate time.Time            `json:"lastModifiedDate"`
	CustomAttributes map[string]TypeValue `json:"customAttributes"`
	EntitlementName  string               `json:"entitlementName"`
	EntitlementType  string               `json:"entitlementType"`
	ItemType         string               `json:"itemType"`
	ReleaseInfo      []struct {
		Id       string   `json:"id"`
		AppId    string   `json:"appId"`
		Platform []string `json:"platform"`
	} `json:"releaseInfo"`
	Developer    string   `json:"developer"`
	DeveloperId  string   `json:"developerId"`
	EulaIds      []string `json:"eulaIds"`
	EndOfSupport bool     `json:"endOfSupport"`
	//MainGameItem        CatalogItem   `json:"mainGameItem"`
	MainGameItemList    []CatalogItem `json:"mainGameItemList"`
	DlcItemList         []CatalogItem `json:"dlcItemList"`
	EsrbGameRatingValue string        `json:"esrbGameRatingValue"`
	AgeGatings          struct {
	} `json:"ageGatings"`
	Unsearchable bool `json:"unsearchable"`
}

func RequestCatalogItem(namespace, itemId string, token string, client *http.Client) (io.ReadCloser, error) {

	ciUrl := CatalogItemUrl(namespace, itemId, true, true, "US", "en")

	return getResponse(ciUrl, token, client)
}

func GetCatalogItem(gameAsset *GameAsset, rdx redux.Writeable, force bool) (*CatalogItem, error) {

	catalogItemsDir := vangogh_integration.AbsProductTypeDir(vangogh_integration.EgsCatalogItems)

	kvCatalogItems, err := kevlar.New(catalogItemsDir, kevlar.JsonExt)
	if err != nil {
		return nil, err
	}

	if !kvCatalogItems.Has(gameAsset.CatalogItemId) || force {

		if err = FetchCatalogItem(gameAsset, kvCatalogItems, rdx); err != nil {
			return nil, err
		}
	}

	return ReadLocalCatalogItem(gameAsset.CatalogItemId, kvCatalogItems)
}

func FetchCatalogItem(gameAsset *GameAsset, kvCatalogItems kevlar.KeyValues, rdx redux.Writeable) error {

	client, err := GetClient()
	if err != nil {
		return err
	}

	ptr, err := VerifyToken(client)
	if err != nil {
		return err
	}

	rcCatalogItem, err := RequestCatalogItem(gameAsset.Namespace, gameAsset.CatalogItemId, ptr.AccessToken, client)
	if err != nil {
		return err
	}

	defer rcCatalogItem.Close()

	if err = kvCatalogItems.Set(gameAsset.CatalogItemId, rcCatalogItem); err != nil {
		return err
	}

	return ReduceCatalogItem(gameAsset.AppName, gameAsset.CatalogItemId, kvCatalogItems, rdx)
}

func ReadLocalCatalogItem(catalogItemId string, kvCatalogItems kevlar.KeyValues) (*CatalogItem, error) {

	rcCatalogItem, err := kvCatalogItems.Get(catalogItemId)
	if err != nil {
		return nil, err
	}
	defer rcCatalogItem.Close()

	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, rcCatalogItem); err != nil {
		return nil, err
	}

	jsonCatalogItem := strings.HasPrefix(buf.String(), jsonCatalogItemPfx)

	switch jsonCatalogItem {
	case true:
		var catalogItem CatalogItem
		if err = json.UnmarshalRead(buf, &catalogItem); err != nil {
			return nil, err
		}
		return &catalogItem, nil
	default:
		var catalogItemMap map[string]CatalogItem
		if err = json.UnmarshalRead(buf, &catalogItemMap); err != nil {
			return nil, err
		}
		return new(catalogItemMap[catalogItemId]), nil
	}
}

func ReduceCatalogItem(appName, catalogItemId string, kvCatalogItems kevlar.KeyValues, rdx redux.Writeable) error {

	if err := rdx.MustHave(vangogh_integration.EgsTitleProperty, vangogh_integration.EgsMainGameProperty); err != nil {
		return err
	}

	rcCatalogItem, err := kvCatalogItems.Get(catalogItemId)
	if err != nil {
		return err
	}

	defer rcCatalogItem.Close()

	var catalogItemMap map[string]CatalogItem
	if err = json.UnmarshalRead(rcCatalogItem, &catalogItemMap); err != nil {
		return err
	}

	catalogItem := catalogItemMap[catalogItemId]

	if err = rdx.ReplaceValues(vangogh_integration.EgsTitleProperty, appName, catalogItem.Title); err != nil {
		return err
	}

	if len(catalogItem.MainGameItemList) > 0 {
		var mainGameItems []string

		for _, mainGameItem := range catalogItem.MainGameItemList {
			for _, releaseInfo := range mainGameItem.ReleaseInfo {
				mainGameItems = append(mainGameItems, releaseInfo.AppId)
			}
		}

		if len(mainGameItems) > 0 {
			if err = rdx.ReplaceValues(vangogh_integration.EgsMainGameProperty, appName, mainGameItems...); err != nil {
				return err
			}
		}
	}

	return nil
}

func CatalogItemAssets(catalogItem *CatalogItem) (map[steam_grid.Asset]*url.URL, error) {

	shortcutAssets := make(map[steam_grid.Asset]*url.URL)

	for _, keyImage := range catalogItem.KeyImages {

		var asset steam_grid.Asset

		switch keyImage.Type {
		case "DieselGameBox":
			asset = steam_grid.Header
		case "DieselGameBoxTall":
			asset = steam_grid.LibraryCapsule
		case "DieselGameBoxLogo":
			asset = steam_grid.LibraryLogo
		default:
			return nil, errors.New("unknown key image type: " + keyImage.Type)
		}

		if u, err := url.Parse(keyImage.Url); err == nil {
			shortcutAssets[asset] = u
		} else {
			return nil, err
		}
	}

	return shortcutAssets, nil
}

func CatalogItemDlcGameAssets(osGameAssets map[vangogh_integration.OperatingSystem][]GameAsset, operatingSystem vangogh_integration.OperatingSystem, catalogItem *CatalogItem, force bool) (map[string]string, error) {

	dlcGameAssets := make(map[string]string)

	if len(catalogItem.DlcItemList) == 0 {
		return dlcGameAssets, nil
	}

	for gaOs, gameAssets := range osGameAssets {
		if gaOs != operatingSystem {
			continue
		}

		for _, dlcItem := range catalogItem.DlcItemList {
			for _, releaseInfo := range dlcItem.ReleaseInfo {
				if ContainsGameAsset(releaseInfo.AppId, gameAssets) {
					dlcGameAssets[releaseInfo.AppId] = dlcItem.Title
				}
			}
		}

	}

	return dlcGameAssets, nil
}
