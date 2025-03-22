package github_integration

import (
	"encoding/json"
	"github.com/boggydigital/kevlar"
	"path"
	"strings"
)

type GitHubSource struct {
	OwnerRepo string
	AssetGlob string
	Unpack    []string
}

var UmuLauncher = &GitHubSource{
	OwnerRepo: "Open-Wine-Components/umu-launcher",
	AssetGlob: "umu-launcher-*-zipapp.tar",
}

var UmuProton = &GitHubSource{
	OwnerRepo: "Open-Wine-Components/umu-proton",
	AssetGlob: "UMU-Proton-*.tar.gz",
}

func (ghs *GitHubSource) GetLatestRelease(kvGitHubReleases kevlar.KeyValues) (*GitHubRelease, error) {

	rcReleases, err := kvGitHubReleases.Get(ghs.OwnerRepo)
	if err != nil {
		return nil, err
	}

	var releases []GitHubRelease
	if err := json.NewDecoder(rcReleases).Decode(&releases); err != nil {
		rcReleases.Close()
		return nil, err
	}

	if err = rcReleases.Close(); err != nil {
		return nil, err
	}

	var latestRelease *GitHubRelease
	if len(releases) > 0 {
		latestRelease = &releases[0]
	}

	return latestRelease, nil
}

func (ghs *GitHubSource) GetAsset(release *GitHubRelease) *GitHubAsset {

	if len(release.Assets) == 1 {
		return &release.Assets[0]
	}

	if prefix, suffix, ok := strings.Cut(ghs.AssetGlob, "*"); ok {
		for _, asset := range release.Assets {
			_, file := path.Split(asset.Name)
			if strings.HasPrefix(file, prefix) && strings.HasSuffix(file, suffix) {
				return &asset
			}
		}
	}

	return nil
}
