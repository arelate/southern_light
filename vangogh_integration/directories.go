package vangogh_integration

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/boggydigital/pathways"
)

const (
	rootPathwayDir     = "/var/lib/vangogh"
	setPathwayFilename = "directories.txt"
)

const (
	Backups           pathways.AbsDir = "backups"
	Metadata          pathways.AbsDir = "metadata"
	Input             pathways.AbsDir = "input"
	Output            pathways.AbsDir = "output"
	Images            pathways.AbsDir = "images"
	DescriptionImages pathways.AbsDir = "description_images"
	Downloads         pathways.AbsDir = "downloads"
	Checksums         pathways.AbsDir = "checksums"
	Logs              pathways.AbsDir = "logs"
)

//var AllAbsDirs = []pathways.AbsDir{
//	Backups,
//	Metadata,
//	Input,
//	Output,
//	Images,
//	DescriptionImages,
//	Downloads,
//	Checksums,
//	Logs,
//}

const (
	Redux          pathways.RelDir = "_redux"          // Metadata
	GitHubReleases pathways.RelDir = "github-releases" // Metadata
	Author         pathways.RelDir = "_author"         // Metadata
	WineBinaries   pathways.RelDir = "_wine-binaries"  // Downloads
	DLCs           pathways.RelDir = "dlc"             // Downloads
	Extras         pathways.RelDir = "extras"          // Downloads
)

//var RelToAbsDirs = map[pathways.RelDir]pathways.AbsDir{
//	Redux:          Metadata,
//	GitHubReleases: Metadata,
//	Author:         Metadata,
//	WineBinaries:   Downloads,
//	DLCs:           Downloads,
//	Extras:         Downloads,
//}

var Pwd pathways.Pathway

func AbsImagesDirByImageId(imageId string) (string, error) {
	if imageId == "" {
		return "", fmt.Errorf("imageId cannot be empty")
	}

	imageId = strings.TrimPrefix(imageId, "/")

	if len(imageId) < 2 {
		return "", fmt.Errorf("imageId is too short")
	}

	idp := Pwd.AbsDirPath(Images)
	return filepath.Join(idp, imageId[0:2]), nil
}

func AbsProductTypeDir(pt ProductType) (string, error) {
	if pt == UnknownProductType {
		return "", fmt.Errorf("no local destination for product type %s", pt)
	}
	amd := Pwd.AbsDirPath(Metadata)
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

	relDownloadTypeDir := ""

	switch dt {
	case DLC:
		relDownloadTypeDir = Pwd.AbsRelDirPath(DLCs, Downloads)
	case Extra:
		relDownloadTypeDir = Pwd.AbsRelDirPath(Extras, Downloads)
	default:
		// do nothing - use base product downloads dir
	}

	return filepath.Join(relSlugDir, relDownloadTypeDir), nil
}

func AbsSlugDownloadDir(slug string, dt DownloadType, layout DownloadsLayout) (string, error) {
	rsdtd, err := relSlugDownloadTypeDir(slug, dt, layout)
	if err != nil {
		return "", err
	}

	downloadsDir := Pwd.AbsDirPath(Downloads)
	return filepath.Join(downloadsDir, rsdtd), nil
}
