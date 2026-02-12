package steam_grid

import (
	"net/url"
	"strings"
)

func AssetUrl(steamAppId, imageId string, asset Asset) *url.URL {

	var pathTemplate string
	if apt, ok := assetPathTemplates[asset]; ok {
		pathTemplate = apt
	} else {
		panic("no path template set for asset " + asset.String())
	}

	path := strings.Replace(pathTemplate, "{steamAppId}", steamAppId, 1)
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
