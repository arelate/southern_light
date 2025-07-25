package vangogh_integration

import (
	"github.com/arelate/southern_light/github_integration"
	"github.com/boggydigital/kevlar"
)

type Binary struct {
	Title           string
	Version         string
	Digest          string
	DownloadUrl     string
	GitHubOwnerRepo string
	GitHubAssetGlob string
}

func (bin *Binary) String() string {
	switch bin.Title {
	case "":
		return bin.GitHubOwnerRepo
	default:
		return bin.Title
	}
}

func (bin *Binary) GetDownloadUrl(kvGitHubReleases kevlar.KeyValues) (string, error) {
	binaryUrl := bin.DownloadUrl

	if binaryUrl != "" {
		return binaryUrl, nil
	}

	latestRelease, err := github_integration.GetLatestRelease(bin.GitHubOwnerRepo, kvGitHubReleases)
	if err != nil {
		return "", err
	}

	latestAsset := github_integration.GetReleaseAsset(latestRelease, bin.GitHubAssetGlob)

	return latestAsset.BrowserDownloadUrl, nil
}

func (bin *Binary) GetVersion(kvGitHubReleases kevlar.KeyValues) (string, error) {
	binaryVersion := bin.Version
	if binaryVersion != "" {
		return binaryVersion, nil
	}

	latestRelease, err := github_integration.GetLatestRelease(bin.GitHubOwnerRepo, kvGitHubReleases)
	if err != nil {
		return "", err
	}

	return latestRelease.TagName, nil
}

func (bin *Binary) GetDigest(kvGitHubReleases kevlar.KeyValues) (string, error) {
	digest := bin.Digest

	if digest != "" {
		return digest, nil
	}

	latestRelease, err := github_integration.GetLatestRelease(bin.GitHubOwnerRepo, kvGitHubReleases)
	if err != nil {
		return "", err
	}

	latestAsset := github_integration.GetReleaseAsset(latestRelease, bin.GitHubAssetGlob)

	if latestAsset.Digest != nil {
		return *latestAsset.Digest, nil
	}

	return "", nil
}
