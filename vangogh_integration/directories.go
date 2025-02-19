package vangogh_integration

import (
	"fmt"
	"github.com/boggydigital/pathways"
	"path/filepath"
	"strings"
)

const DefaultRootDir = "/var/lib/vangogh"

const (
	Backups    pathways.AbsDir = "backups"
	Metadata   pathways.AbsDir = "metadata"
	Input      pathways.AbsDir = "input"
	Output     pathways.AbsDir = "output"
	Images     pathways.AbsDir = "images"
	Items      pathways.AbsDir = "items"
	Downloads  pathways.AbsDir = "downloads"
	Checksums  pathways.AbsDir = "checksums"
	RecycleBin pathways.AbsDir = "recycle_bin"
	Logs       pathways.AbsDir = "logs"
)

var AllAbsDirs = []pathways.AbsDir{
	Backups,
	Metadata,
	Input,
	Output,
	Images,
	Items,
	Downloads,
	Checksums,
	RecycleBin,
	Logs,
}

const (
	Redux      pathways.RelDir = "_redux"
	TypeErrors pathways.RelDir = "_type_errors"
	DLCs       pathways.RelDir = "dlc"
	Extras     pathways.RelDir = "extras"
)

var RelToAbsDirs = map[pathways.RelDir]pathways.AbsDir{
	Redux:      Metadata,
	TypeErrors: Metadata,
	DLCs:       Downloads,
	Extras:     Downloads,
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

func AbsItemPath(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("item path cannot be empty")
	}

	//GOG.com quirk - some item URLs path has multiple slashes
	//e.g. https://items.gog.com//atom_rpg_trudograd/mp4/TGWMap_Night_%281%29.gif.mp4
	//so we need to keep trimming while there is something to trim
	for strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}
	if len(path) < 1 {
		return "", fmt.Errorf("sanitized item path cannot be empty")
	}

	idp, err := pathways.GetAbsDir(Items)
	if err != nil {
		return "", err
	}

	return filepath.Join(idp, path[0:1], path), nil
}

func AbsProductTypeDir(pt ProductType) (string, error) {
	if !IsValidProductType(pt) {
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
