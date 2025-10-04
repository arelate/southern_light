package steam_integration

import (
	"net/url"
	"strings"
)

func AssetUrl(appId, assetFilename string) *url.URL {

	path := strings.Replace(assetPathTemplate, "{appId}", appId, 1)
	path = strings.Replace(path, "{asset}", assetFilename, 1)

	return &url.URL{
		Scheme: httpsScheme,
		Host:   SteamStaticHost,
		Path:   path,
	}
}
