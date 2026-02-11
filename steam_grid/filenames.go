package steam_grid

import "strconv"

func LogoPositionFilename(shortcutId uint32) string {
	return strconv.FormatInt(int64(shortcutId), 10) + ".json"
}

func ImageFilename(shortcutId uint32, imageType Asset) string {
	filename := strconv.FormatInt(int64(shortcutId), 10)
	ext := ".jpg"

	switch imageType {
	case Header:
	// do nothing
	case LibraryCapsule:
		filename += "p"
	case LibraryHero:
		filename += "_hero"
	case LibraryLogo:
		filename += "_logo"
		ext = ".png"
	case ClientIcon:
		filename += "_icon"
		ext = ".png"
	default:
		panic("unknown image type " + imageType.String())
	}
	return filename + ext
}
