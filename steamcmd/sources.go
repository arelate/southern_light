package steamcmd

import "github.com/arelate/southern_light/vangogh_integration"

var Urls = map[vangogh_integration.OperatingSystem]string{
	vangogh_integration.Windows: "https://steamcdn-a.akamaihd.net/client/installer/steamcmd.zip",
	vangogh_integration.MacOS:   "https://steamcdn-a.akamaihd.net/client/installer/steamcmd_osx.tar.gz",
	vangogh_integration.Linux:   "https://steamcdn-a.akamaihd.net/client/installer/steamcmd_linux.tar.gz",
}
