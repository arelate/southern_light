package steam_integration

import (
	"net/url"
	"strconv"
	"strings"
)

const (
	protonDB          = "www.protondb.com"
	appIdPathTemplate = "/app/{appId}"
)

func ProtonDBUrl(appId uint32) *url.URL {
	return &url.URL{
		Scheme: httpsScheme,
		Host:   protonDB,
		Path:   strings.Replace(appIdPathTemplate, "{appId}", strconv.Itoa(int(appId)), 1),
	}
}
