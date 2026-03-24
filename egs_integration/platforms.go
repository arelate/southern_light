package egs_integration

import "github.com/arelate/southern_light/vangogh_integration"

// https://eoshelp.epicgames.com/s/article/Does-the-Epic-Games-Store-support-Linux-games?language=en_US
var osPlatfoms = map[vangogh_integration.OperatingSystem]string{
	vangogh_integration.Windows: "Windows",
	vangogh_integration.MacOS:   "Mac",
}

func Platform(operatingSystem vangogh_integration.OperatingSystem) string {
	if plat, ok := osPlatfoms[operatingSystem]; ok {
		return plat
	}
	return ""
}
