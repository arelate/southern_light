package vangogh_integration

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/arelate/southern_light/gog_integration"
	"github.com/boggydigital/pathways"
)

const (
	xmlExt           = ".xml"
	cookiesFilename  = "cookies_gog_com.json"
	atomFeedFilename = "atom.xml"
)

var validatedExtensions = map[string]bool{
	".exe": true,
	".bin": true,
	".dmg": true,
	".pkg": true,
	".sh":  true,
}

func AbsChecksumPath(absDownloadPath string) (string, error) {

	downloadsDir, err := pathways.GetAbsDir(Downloads)
	if err != nil {
		return "", err
	}

	relDownloadPath, err := filepath.Rel(downloadsDir, absDownloadPath)
	if err != nil {
		return "", err
	}

	checksumsDir, err := pathways.GetAbsDir(Checksums)
	if err != nil {
		return "", err
	}

	return filepath.Join(checksumsDir, relDownloadPath+xmlExt), nil
}

func absLocalVideoPath(videoId string, videoDir string, ext string) string {
	videoPath := filepath.Join(videoDir, videoId+ext)

	if _, err := os.Stat(videoPath); err == nil {
		return videoPath
	}

	return ""

}

func AbsLocalImagePath(imageId string) (string, error) {
	exts := []string{gog_integration.JpgExt, gog_integration.PngExt}
	idp, err := AbsImagesDirByImageId(imageId)
	if err != nil {
		return "", err
	}
	for _, ext := range exts {
		aip := filepath.Join(idp, imageId+ext)
		if _, err := os.Stat(aip); err == nil {
			return aip, nil
		} else {
			continue
		}
	}
	return "", err
}

func AbsCookiePath() (string, error) {
	ifdp, err := pathways.GetAbsDir(Input)
	return filepath.Join(ifdp, cookiesFilename), err
}

func AbsAtomFeedPath() (string, error) {
	ofdp, err := pathways.GetAbsDir(Output)
	return filepath.Join(ofdp, atomFeedFilename), err
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
