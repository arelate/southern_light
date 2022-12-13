package protondb_integration

import (
	"net/url"
	"strconv"
)

func ProtonDBUrl(steamAppId uint32) *url.URL {
	return &url.URL{
		Scheme: httpsScheme,
		Host:   protonDBHost,
		Path:   appPath + strconv.Itoa(int(steamAppId)),
	}
}
