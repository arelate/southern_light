package github_integration

import (
	"encoding/json"
	"github.com/boggydigital/kevlar"
	"path"
	"strings"
)

func GetLatestRelease(repo string, kvGitHubReleases kevlar.KeyValues) (*GitHubRelease, error) {

	rcReleases, err := kvGitHubReleases.Get(repo)
	if err != nil {
		return nil, err
	}

	var releases []GitHubRelease
	if err = json.NewDecoder(rcReleases).Decode(&releases); err != nil {
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

func GetReleaseAsset(release *GitHubRelease, assetGlob string) *GitHubAsset {

	if len(release.Assets) == 1 {
		return &release.Assets[0]
	}

	if prefix, suffix, sure := strings.Cut(assetGlob, "*"); sure {
		for _, asset := range release.Assets {
			_, file := path.Split(asset.Name)
			if strings.HasPrefix(file, prefix) && strings.HasSuffix(file, suffix) {
				return &asset
			}
		}
	}

	return nil
}
