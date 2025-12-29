package vangogh_integration

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/arelate/southern_light/gog_integration"
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

	downloadsDir := Pwd.AbsDirPath(Downloads)

	relDownloadPath, err := filepath.Rel(downloadsDir, absDownloadPath)
	if err != nil {
		return "", err
	}

	checksumsDir := Pwd.AbsDirPath(Checksums)

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

func AbsCookiePath() string {
	return filepath.Join(Pwd.AbsDirPath(Input), cookiesFilename)
}

func AbsAtomFeedPath() string {
	return filepath.Join(Pwd.AbsDirPath(Output), atomFeedFilename)
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

	idp := Pwd.AbsDirPath(DescriptionImages)

	x, _ := utf8.DecodeRuneInString(path)

	return filepath.Join(idp, string(x), path), nil
}
