package epic_games

import (
	"encoding/json/v2"
	"io"
	"net/http"
)

const defaultLabel = "Live"

func debugString(reader io.Reader) (string, error) {
	bts, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(bts), err
}

type GameAsset struct {
	AppName       string `json:"appName"`
	LabelName     string `json:"labelName"`
	BuildVersion  string `json:"buildVersion"`
	CatalogItemId string `json:"catalogItemId"`
	Namespace     string `json:"namespace"`
	Metadata      struct {
		InstallationPoolId string `json:"installationPoolId,omitempty"`
		UpdateType         string `json:"update_type,omitempty"`
	} `json:"metadata,omitempty"`
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
		Manifests []struct {
			Uri         string `json:"uri"`
			QueryParams []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"queryParams"`
		} `json:"manifests"`
		IsPreloaded bool `json:"isPreloaded"`
	} `json:"elements"`
}

func GetGameAssets(platform string, token string, client *http.Client) ([]GameAsset, error) {

	ggaUrl := LauncherGameAssetsUrl(platform, defaultLabel)

	resp, err := getResponse(ggaUrl, token, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var gameAssets []GameAsset
	if err = json.UnmarshalRead(resp.Body, &gameAssets); err != nil {
		return nil, err
	}

	return gameAssets, nil
}

func GetGameManifest(namespace, catalogItemId, appName string, platform string, token string, client *http.Client) (*GameManifest, error) {

	lgmUrl := LauncherGameManifestUrl(namespace, catalogItemId, appName, platform, defaultLabel)

	resp, err := getResponse(lgmUrl, token, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var gameManifest GameManifest
	if err = json.UnmarshalRead(resp.Body, &gameManifest); err != nil {
		return nil, err
	}

	return &gameManifest, nil
}

func GetLauncherManifests(platform string, token string, client *http.Client) (*LauncherManifests, error) {

	lmUrl := LauncherManifestsUrl(platform, defaultLabel)

	resp, err := getResponse(lmUrl, token, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var launcherManifests LauncherManifests
	if err = json.UnmarshalRead(resp.Body, &launcherManifests); err != nil {
		return nil, err
	}

	return &launcherManifests, nil
}
