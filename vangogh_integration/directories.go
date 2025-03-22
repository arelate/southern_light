package vangogh_integration

import (
	"fmt"
	"github.com/arelate/southern_light/github_integration"
	"github.com/boggydigital/busan"
	"github.com/boggydigital/pathways"
	"path"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

const DefaultRootDir = "/var/lib/vangogh"

const (
	Backups           pathways.AbsDir = "backups"
	Metadata          pathways.AbsDir = "metadata"
	Input             pathways.AbsDir = "input"
	Output            pathways.AbsDir = "output"
	Images            pathways.AbsDir = "images"
	DescriptionImages pathways.AbsDir = "description_images"
	Downloads         pathways.AbsDir = "downloads"
	Checksums         pathways.AbsDir = "checksums"
	RecycleBin        pathways.AbsDir = "recycle_bin"
	Logs              pathways.AbsDir = "logs"
)

var AllAbsDirs = []pathways.AbsDir{
	Backups,
	Metadata,
	Input,
	Output,
	Images,
	DescriptionImages,
	Downloads,
	Checksums,
	RecycleBin,
	Logs,
}

const (
	Redux          pathways.RelDir = "_redux"
	GitHubReleases pathways.RelDir = "github-releases"
	GitHubAssets   pathways.RelDir = "_github-assets"
	TypeErrors     pathways.RelDir = "_type_errors"
	DLCs           pathways.RelDir = "dlc"
	Extras         pathways.RelDir = "extras"
)

var RelToAbsDirs = map[pathways.RelDir]pathways.AbsDir{
	Redux:          Metadata,
	GitHubReleases: Metadata,
	GitHubAssets:   Downloads,
	TypeErrors:     Metadata,
	DLCs:           Downloads,
	Extras:         Downloads,
}

func AbsImagesDirByImageId(imageId string) (string, error) {
	if imageId == "" {
		return "", fmt.Errorf("imageId cannot be empty")
	}

	imageId = strings.TrimPrefix(imageId, "/")

	if len(imageId) < 2 {
		return "", fmt.Errorf("imageId is too short")
	}

	idp, err := pathways.GetAbsDir(Images)
	return filepath.Join(idp, imageId[0:2]), err
}

func AbsDescriptionImagePath(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("description image path cannot be empty")
	}

	//GOG.com quirk - some item URLs path has multiple slashes
	//e.g. https://items.gog.com//atom_rpg_trudograd/mp4/TGWMap_Night_%281%29.gif.mp4
	//so we need to keep trimming while there is something to trim
	for strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}
	if len(path) < 1 {
		return "", fmt.Errorf("sanitized description image path cannot be empty")
	}

	idp, err := pathways.GetAbsDir(DescriptionImages)
	if err != nil {
		return "", err
	}

	x, _ := utf8.DecodeRuneInString(path)

	return filepath.Join(idp, string(x), path), nil
}

func AbsProductTypeDir(pt ProductType) (string, error) {
	if pt == UnknownProductType {
		return "", fmt.Errorf("no local destination for product type %s", pt)
	}
	amd, err := pathways.GetAbsDir(Metadata)
	if err != nil {
		return "", err
	}
	return filepath.Join(amd, pt.String()), nil
}

func RelProductDownloadsDir(slug string) (string, error) {
	if slug == "" {
		return "", fmt.Errorf("vangogh_urls: empty slug")
	}
	if len(slug) < 1 {
		return "", fmt.Errorf("vangogh_urls: slug is too short")
	}
	return filepath.Join(strings.ToLower(slug[0:1]), slug), nil
}

func AbsProductDownloadsDir(slug string) (string, error) {
	rDir, err := RelProductDownloadsDir(slug)
	if err != nil {
		return rDir, err
	}
	return AbsDownloadDirFromRel(rDir)
}

func AbsDownloadDirFromRel(p string) (string, error) {
	adp, err := pathways.GetAbsDir(Downloads)
	if err != nil {
		return "", err
	}
	return filepath.Join(adp, p), nil
}

func AbsGitHubReleasesDir(ghs *github_integration.GitHubSource, release *github_integration.GitHubRelease) (string, error) {
	assetsDir, err := pathways.GetAbsRelDir(GitHubAssets)
	if err != nil {
		return "", err
	}

	return filepath.Join(assetsDir, ghs.OwnerRepo, busan.Sanitize(release.TagName)), nil
}

func AbsGitHubReleaseAssetPath(ghs *github_integration.GitHubSource, release *github_integration.GitHubRelease, asset *github_integration.GitHubAsset) (string, error) {
	relDir, err := AbsGitHubReleasesDir(ghs, release)
	if err != nil {
		return "", err
	}

	_, fn := path.Split(asset.BrowserDownloadUrl)

	return filepath.Join(relDir, fn), nil
}
