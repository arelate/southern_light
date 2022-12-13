package steam_integration

import (
	"net/url"
	"path"
	"strconv"
)

func StorePageUrl(appId uint32) *url.URL {
	return &url.URL{
		Scheme: httpsScheme,
		Host:   StoreHost,
		Path:   path.Join(appPath, strconv.Itoa(int(appId))),
	}
}
