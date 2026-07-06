package vangogh_integration

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/boggydigital/camino"
)

const (
	vangoghRootDir      = "/var/lib/vangogh"
	directoriesFilename = "directories.txt"
)

const (
	Backups  camino.AbsDir = iota // vangogh, theo
	Binaries                      // vangogh, theo
	Metadata                      // vangogh, theo
	Output
	Images
	DescriptionImages
	Downloads
	Checksums
	Logs
)

const (
	Redux camino.RelDir = iota
	Author
	Cookies
	GitHubReleases
	Releases
	Runtimes
)

var absDirNames = map[camino.AbsDir]string{
	Backups:           "backups",
	Binaries:          "binaries",
	Metadata:          "metadata",
	Output:            "output",
	Images:            "images",
	DescriptionImages: "description_images",
	Downloads:         "downloads",
	Checksums:         "checksums",
	Logs:              "logs",
}

var relDirNames = map[camino.RelDir]string{
	Redux:          "_redux",
	Author:         "_author",
	Cookies:        "_cookies",
	GitHubReleases: "github-releases",
	Releases:       "releases",
	Runtimes:       "runtimes",
}

var vangoghAbsDirs = []camino.AbsDir{
	Backups,
	Binaries,
	Metadata,
	Output,
	Images,
	DescriptionImages,
	Downloads,
	Checksums,
	Logs,
}

var vangoghRelAbsParents = map[camino.RelDir][]camino.AbsDir{
	Redux:          {Metadata},
	GitHubReleases: {Metadata},
	Author:         {Metadata},
	Releases:       {Binaries},
	Runtimes:       {Binaries},
}

func AbsImagesDirByImageId(imageId string) (string, error) {
	if imageId == "" {
		return "", fmt.Errorf("imageId cannot be empty")
	}

	imageId = strings.TrimPrefix(imageId, "/")

	if len(imageId) < 2 {
		return "", fmt.Errorf("imageId is too short")
	}

	idp := camino.GetAbs(Images)
	return filepath.Join(idp, imageId[0:2]), nil
}

func AbsProductTypeDir(pt ProductType) (string, error) {
	if pt == UnknownProductType {
		return "", fmt.Errorf("no local destination for product type %s", pt)
	}
	amd := camino.GetAbs(Metadata)
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

	var relDownloadTypeDir string

	switch dt {
	case DLC:
		relDownloadTypeDir = "dlc"
	case Extra:
		relDownloadTypeDir = "extras"
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

	downloadsDir := camino.GetAbs(Downloads)
	return filepath.Join(downloadsDir, rsdtd), nil
}

func AbsReduxDir() string {
	return camino.GetRel(Redux, Metadata)
}

func InitVangoghCamino() error {

	var overrides map[string]string

	if _, err := os.Stat(directoriesFilename); err == nil {
		if overrides, err = camino.ReadOverrides(directoriesFilename); err != nil {
			return err
		}
	}

	vangoghAbsDirNames := make(map[camino.AbsDir]string)

	for _, vad := range vangoghAbsDirs {

		var ok bool
		if vangoghAbsDirNames[vad], ok = absDirNames[vad]; !ok {
			return errors.New("vangogh abs dir name not set")
		}
	}

	resolvedVangoghAbsPaths := camino.ResolveAbsPaths(vangoghRootDir, vangoghAbsDirNames, overrides)

	vrds := make(map[camino.RelDir]string)

	for vrp := range vangoghRelAbsParents {
		var ok bool
		if vrds[vrp], ok = relDirNames[vrp]; !ok {
			return errors.New("vangogh rel dir path not set")
		}
	}

	return camino.Register(resolvedVangoghAbsPaths, vrds, vangoghRelAbsParents)
}
