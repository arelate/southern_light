package protondb_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func ProtonDBUrl(appId string) *url.URL {
	return southern_light.SuffixIdUrl(protonDBHost, appPath, appId)
}
