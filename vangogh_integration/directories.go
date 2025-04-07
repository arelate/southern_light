package vangogh_integration

import (
	"errors"
	"fmt"
	"github.com/arelate/southern_light/github_integration"
	"github.com/boggydigital/busan"
	"github.com/boggydigital/pathways"
	"path/filepath"
	"strings"
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

func RelProductDownloadsDir(slug string, dl DownloadsLayout) (string, error) {
	if slug == "" {
		return "", fmt.Errorf("vangogh_urls: empty slug")
	}
	var relDir string
	switch dl {
	case FlatDownloadsLayout:
		relDir = strings.ToLower(slug)
	case ShardedDownloadsLayout:
		relDir = filepath.Join(strings.ToLower(slug[0:1]), slug)
	default:
		return "", errors.New("unsupported downloads layout: " + dl.String())
	}
	return relDir, nil
}

func AbsProductDownloadsDir(slug string, dl DownloadsLayout) (string, error) {
	rDir, err := RelProductDownloadsDir(slug, dl)
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

func AbsGitHubReleasesDir(repo string, release *github_integration.GitHubRelease) (string, error) {
	assetsDir, err := pathways.GetAbsRelDir(GitHubAssets)
	if err != nil {
		return "", err
	}

	return filepath.Join(assetsDir, repo, busan.Sanitize(release.TagName)), nil
}
