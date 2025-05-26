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
	DLCs           pathways.RelDir = "dlc"
	Extras         pathways.RelDir = "extras"
)

var RelToAbsDirs = map[pathways.RelDir]pathways.AbsDir{
	Redux:          Metadata,
	GitHubReleases: Metadata,
	GitHubAssets:   Downloads,
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

func relSlugDownloadTypeDir(slug string, dt DownloadType, layout DownloadsLayout) (string, error) {
	if slug == "" {
		return "", fmt.Errorf("vangogh_urls: empty slug")
	}
	// this is required to address https://github.com/arelate/vangogh/issues/96
	// for both sharded and flat downloads layouts:
	// - for sharded layouts slug would need to be different from the shard ("a_" would be sharded into "a")
	// to avoid issues when performing relayout from sharded layout and collision of "a" (shard) -> "a" (slug)
	// - for flat layout this is required to prevent potential collision of "a" (slug) -> "a" (shard)
	if len(slug) == 1 {
		slug = fmt.Sprintf("%s_", slug)
	}
	var relSlugDir string
	switch layout {
	case FlatDownloadsLayout:
		relSlugDir = strings.ToLower(slug)
	case ShardedDownloadsLayout:
		shard := strings.ToLower(slug[0:1])
		relSlugDir = filepath.Join(shard, slug)
	default:
		return "", errors.New("unsupported downloads layout: " + layout.String())
	}

	var err error
	relDownloadTypeDir := ""

	switch dt {
	case DLC:
		relDownloadTypeDir, err = pathways.GetRelDir(DLCs)
	case Extra:
		relDownloadTypeDir, err = pathways.GetRelDir(Extras)
	default:
		// do nothing - use base product downloads dir
	}

	if err != nil {
		return "", err
	}

	return filepath.Join(relSlugDir, relDownloadTypeDir), nil
}

func AbsSlugDownloadDir(slug string, dt DownloadType, layout DownloadsLayout) (string, error) {
	if rsdtd, err := relSlugDownloadTypeDir(slug, dt, layout); err == nil {

		downloadsDir, err := pathways.GetAbsDir(Downloads)
		if err != nil {
			return "", err
		}

		return filepath.Join(downloadsDir, rsdtd), nil

	} else {
		return "", err
	}
}

func AbsGitHubReleasesDir(repo string, release *github_integration.GitHubRelease) (string, error) {
	assetsDir, err := pathways.GetAbsRelDir(GitHubAssets)
	if err != nil {
		return "", err
	}

	return filepath.Join(assetsDir, repo, busan.Sanitize(release.TagName)), nil
}
