package egs_integration

import (
	"encoding/json/v2"
	"errors"
	"fmt"
	"slices"

	"github.com/arelate/southern_light/vangogh_integration"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/redux"
)

type GameAsset struct {
	AppName       string `json:"appName"`
	LabelName     string `json:"labelName"`
	BuildVersion  string `json:"buildVersion"`
	CatalogItemId string `json:"catalogItemId"`
	Namespace     string `json:"namespace"`
	Metadata      struct {
		InstallationPoolId string `json:"installationPoolId,omitempty"`
		UpdateType         string `json:"update_type,omitempty"`
	} `json:"metadata"`
	SidecarRvn int    `json:"sidecarRvn,omitempty"`
	AssetId    string `json:"assetId"`
}

func FetchGameAssets(operatingSystem vangogh_integration.OperatingSystem) error {

	client, err := GetClient()
	if err != nil {
		return err
	}

	ptr, err := VerifyToken(client)
	if err != nil {
		return err
	}

	rcGameAssets, err := RequestGameAssets(Platform(operatingSystem), ptr.AccessToken, client)
	if err != nil {
		return err
	}

	defer rcGameAssets.Close()

	egsOsApKey := fmt.Sprintf("%s-%s", vangogh_integration.StoreEpicGames, operatingSystem)

	availableProductsDir := vangogh_integration.AbsProductTypeDir(vangogh_integration.AvailableProducts)

	kvAvailableProducts, err := kevlar.New(availableProductsDir, kevlar.JsonExt)
	if err != nil {
		return err
	}

	return kvAvailableProducts.Set(egsOsApKey, rcGameAssets)
}

func GameAssetOperatingSystems(appName string, force bool) ([]vangogh_integration.OperatingSystem, error) {

	osGameAssets, err := GetGameAssets(force)
	if err != nil {
		return nil, err
	}

	operatingSystems := make([]vangogh_integration.OperatingSystem, 0)

	for sos, gameAssets := range osGameAssets {
		for _, gameAsset := range gameAssets {
			if gameAsset.AppName == appName {
				operatingSystems = append(operatingSystems, sos)
			}
		}
	}

	return operatingSystems, nil
}

func GetGameAssets(update bool) (map[vangogh_integration.OperatingSystem][]GameAsset, error) {

	osGameAssets := make(map[vangogh_integration.OperatingSystem][]GameAsset)

	for _, sos := range SupportedOperatingSystems {
		gameAssets, err := ReadLocalGameAssets(sos)
		if err != nil {
			return nil, err
		}

		if len(gameAssets) == 0 || update {
			if err = FetchGameAssets(sos); err != nil {
				return nil, err
			}

			gameAssets, err = ReadLocalGameAssets(sos)
			if err != nil {
				return nil, err
			}
		}

		osGameAssets[sos] = gameAssets
	}

	return osGameAssets, nil
}

func ReadLocalGameAssets(operatingSystem vangogh_integration.OperatingSystem) ([]GameAsset, error) {

	if !slices.Contains(SupportedOperatingSystems, operatingSystem) {
		return nil, operatingSystem.ErrUnsupported()
	}

	egsOsApKey := fmt.Sprintf("%s-%s", vangogh_integration.StoreEpicGames, operatingSystem)

	availableProductsDir := vangogh_integration.AbsProductTypeDir(vangogh_integration.AvailableProducts)

	kvAvailableProducts, err := kevlar.New(availableProductsDir, kevlar.JsonExt)
	if err != nil {
		return nil, err
	}

	if !kvAvailableProducts.Has(egsOsApKey) {
		return nil, nil
	}

	rcGameAssets, err := kvAvailableProducts.Get(egsOsApKey)
	if err != nil {
		return nil, err
	}
	defer rcGameAssets.Close()

	var gameAssets []GameAsset
	if err = json.UnmarshalRead(rcGameAssets, &gameAssets); err != nil {
		return nil, err
	}

	return gameAssets, nil
}

func availableProductIndex(appName string, availableProducts []vangogh_integration.AvailableProduct) int {
	for ii, ap := range availableProducts {
		if ap.Id == appName {
			return ii
		}
	}
	return -1
}

func GameAssetsAvailableProducts(
	osGameAssets map[vangogh_integration.OperatingSystem][]GameAsset,
	operatingSystem vangogh_integration.OperatingSystem,
	rdx redux.Writeable,
	force bool) ([]vangogh_integration.AvailableProduct, error) {

	availableProducts := make([]vangogh_integration.AvailableProduct, 0)

	for ops, gameAssets := range osGameAssets {

		for _, gameAsset := range gameAssets {

			catalogItem, err := GetCatalogItem(&gameAsset, rdx, force)
			if err != nil {
				return nil, err
			}

			if len(catalogItem.MainGameItemList) > 0 {
				continue
			}

			if index := availableProductIndex(gameAsset.AppName, availableProducts); index != -1 {
				availableProducts[index].OperatingSystems = append(availableProducts[index].OperatingSystems, ops)
			} else {
				ap := vangogh_integration.AvailableProduct{
					Id:               gameAsset.AppName,
					Title:            catalogItem.Title,
					OperatingSystems: []vangogh_integration.OperatingSystem{ops},
				}

				var dlcGameAssets map[string]string
				dlcGameAssets, err = CatalogItemDlcGameAssets(osGameAssets, ops, catalogItem, force)
				if err != nil {
					return nil, err
				}

				ap.Dlc = dlcGameAssets

				availableProducts = append(availableProducts, ap)
			}
		}
	}

	if operatingSystem != vangogh_integration.AnyOperatingSystem {
		osAvailableProducts := make([]vangogh_integration.AvailableProduct, 0, len(availableProducts))
		for _, ap := range availableProducts {
			if slices.Contains(ap.OperatingSystems, operatingSystem) {
				osAvailableProducts = append(osAvailableProducts, ap)
			}
		}
		availableProducts = osAvailableProducts
	}

	return availableProducts, nil
}

func ContainsGameAsset(appName string, gameAssets []GameAsset) bool {
	for _, ga := range gameAssets {
		if ga.AppName == appName {
			return true
		}
	}
	return false
}

func GetGameAsset(appName string, operatingSystem vangogh_integration.OperatingSystem, force bool) (*GameAsset, error) {

	switch appName {
	case eosOverlayGameAsset.AppName:
		return &eosOverlayGameAsset, nil
	case eosHelperGameAsset.AppName:
		return &eosHelperGameAsset, nil
	default:
		// proceed normally
	}

	osGameAssets, err := GetGameAssets(force)
	if err != nil {
		return nil, err
	}

	for sos, gameAssets := range osGameAssets {
		if operatingSystem != vangogh_integration.AnyOperatingSystem && operatingSystem != sos {
			continue
		}
		for _, gameAsset := range gameAssets {
			if gameAsset.AppName == appName {
				return &gameAsset, nil
			}
		}
	}

	return nil, errors.New("game asset not found for appName " + appName)
}
