package steam_grid

import (
	"net/url"
	"strings"
)

func AssetUrl(steamAppId, imageId string, asset Asset) *url.URL {

	path := strings.Replace(assetPathTemplate, "{steamAppId}", steamAppId, 1)
	path = strings.Replace(path, "{imageId}", imageId, 1)

	ext := asset.Ext()
	if strings.HasSuffix(imageId, ext) {
		ext = ""
	}

	path = strings.Replace(path, "{ext}", ext, 1)

	return &url.URL{
		Scheme: httpsScheme,
		Host:   steamStaticHost,
		Path:   path,
	}
}
