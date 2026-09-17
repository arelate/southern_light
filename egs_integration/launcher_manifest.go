package egs_integration

import (
	"io"
	"net/http"
)

const defaultLabel = "Live"

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

func RequestGameAssets(platform string, token string, client *http.Client) (io.ReadCloser, error) {

	ggaUrl := LauncherGameAssetsUrl(platform, defaultLabel)

	return getResponse(ggaUrl, token, client)
}

func RequestLauncherManifests(platform string, token string, client *http.Client) (io.ReadCloser, error) {

	lmUrl := LauncherManifestsUrl(platform, defaultLabel)

	return getResponse(lmUrl, token, client)
}
