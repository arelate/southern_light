package egs_integration

import (
	"io"
	"net/http"
	"net/url"
)

const defaultLabel = "Live"

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

type LauncherManifests struct {
	Elements []struct {
		AppName      string `json:"appName"`
		LabelName    string `json:"labelName"`
		BuildVersion string `json:"buildVersion"`
		Hash         string `json:"hash"`
		UseSignedUrl bool   `json:"useSignedUrl"`
		Manifests    []struct {
			Uri         string `json:"uri"`
			QueryParams []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"queryParams"`
		} `json:"manifests"`
		IsPreloaded bool `json:"isPreloaded"`
	} `json:"elements"`
}

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

func GetGameAssets(platform string, token string, client *http.Client) (io.ReadCloser, error) {

	ggaUrl := LauncherGameAssetsUrl(platform, defaultLabel)

	return getResponse(ggaUrl, token, client)
}

func GetGameManifest(namespace, catalogItemId, appName string, platform string, token string, client *http.Client) (io.ReadCloser, error) {

	lgmUrl := LauncherGameManifestUrl(namespace, catalogItemId, appName, platform, defaultLabel)

	return getResponse(lgmUrl, token, client)
}

func GetLauncherManifests(platform string, token string, client *http.Client) (io.ReadCloser, error) {

	lmUrl := LauncherManifestsUrl(platform, defaultLabel)

	return getResponse(lmUrl, token, client)
}
