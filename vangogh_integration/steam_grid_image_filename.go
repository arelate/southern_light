package vangogh_integration

import (
	"strconv"
)

func SteamGridImageFilename(shortcutId uint32, imageType ImageType) string {
	filename := strconv.FormatInt(int64(shortcutId), 10)
	ext := ".jpg"

	switch imageType {
	case Image:
	// do nothing
	case VerticalImage:
		filename += "p"
	case Hero:
		filename += "_hero"
	case Logo:
		filename += "_logo"
		ext = ".png"
	case Icon:
		fallthrough
	case IconSquare:
		filename += "_icon"
		ext = ".png"
	default:
		panic("unknown image type " + imageType.String())
	}
	return filename + ext
}

func SteamGridLogoPositionFilename(shortcutId uint32) string {
	return strconv.FormatInt(int64(shortcutId), 10) + ".json"
}
