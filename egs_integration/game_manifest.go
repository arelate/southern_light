package egs_integration

import (
	"encoding/json/v2"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/arelate/southern_light/vangogh_integration"
	"github.com/boggydigital/kevlar"
)

type GameManifest struct {
	Elements []struct {
		AppName      string `json:"appName"`
		LabelName    string `json:"labelName"`
		BuildVersion string `json:"buildVersion"`
		Hash         string `json:"hash"`
		UseSignedUrl bool   `json:"useSignedUrl"`
		Metadata     struct {
			InstallationPoolId string `json:"installationPoolId"`
			UpdateType         string `json:"update_type"`
		} `json:"metadata"`
		Manifests   []ManifestUri `json:"manifests"`
		IsPreloaded bool          `json:"isPreloaded"`
	} `json:"elements"`
}

func (gm *GameManifest) Urls() ([]*url.URL, error) {

	manifestUrls := make([]*url.URL, 0)
	for _, element := range gm.Elements {
		for _, manifest := range element.Manifests {

			manifestUrl, err := manifest.Url()
			if err != nil {
				return nil, err
			}

			manifestUrls = append(manifestUrls, manifestUrl)
		}
	}
	return manifestUrls, nil
}

type ManifestUri struct {
	Uri         string `json:"uri"`
	QueryParams []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"queryParams"`
}

func (mu *ManifestUri) Url() (*url.URL, error) {
	manifestUrl, err := url.Parse(mu.Uri)
	if err != nil {
		return nil, err
	}

	q := manifestUrl.Query()

	for _, kv := range mu.QueryParams {
		q.Add(kv.Name, kv.Value)
	}

	manifestUrl.RawQuery = q.Encode()

	return manifestUrl, nil
}

func RequestGameManifest(namespace, catalogItemId, appName string, platform string, token string, client *http.Client) (io.ReadCloser, error) {

	lgmUrl := LauncherGameManifestUrl(namespace, catalogItemId, appName, platform, defaultLabel)

	return getResponse(lgmUrl, token, client)
}

func GetGameManifest(gameAsset *GameAsset, operatingSystem vangogh_integration.OperatingSystem, force bool) (*GameManifest, error) {

	gameManifestsDir := vangogh_integration.AbsProductTypeDir(vangogh_integration.EgsGameManifests)

	kvGameManifests, err := kevlar.New(gameManifestsDir, kevlar.JsonExt)
	if err != nil {
		return nil, err
	}

	osAppNameKey := fmt.Sprintf("%s-%s", gameAsset.AppName, operatingSystem)

	if !kvGameManifests.Has(osAppNameKey) || force {
		if err = FetchGameManifest(osAppNameKey, gameAsset, operatingSystem, kvGameManifests); err != nil {
			return nil, err
		}
	}

	rcGameManifest, err := kvGameManifests.Get(osAppNameKey)
	if err != nil {
		return nil, err
	}
	defer rcGameManifest.Close()

	var gameManifest GameManifest
	if err = json.UnmarshalRead(rcGameManifest, &gameManifest); err != nil {
		return nil, err
	}

	return &gameManifest, nil
}

func FetchGameManifest(key string, gameAsset *GameAsset, operatingSystem vangogh_integration.OperatingSystem, kvGameManifests kevlar.KeyValues) error {

	client, err := GetClient()
	if err != nil {
		return err
	}

	ptr, err := VerifyToken(client)
	if err != nil {
		return err
	}

	rcGameManifest, err := RequestGameManifest(
		gameAsset.Namespace,
		gameAsset.CatalogItemId,
		gameAsset.AppName,
		Platform(operatingSystem),
		ptr.AccessToken, client)
	if err != nil {
		return err
	}

	defer rcGameManifest.Close()

	return kvGameManifests.Set(key, rcGameManifest)
}
