package protondb_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"strconv"
)

func ProtonDBUrl(steamAppId uint32) *url.URL {
	return southern_light.SuffixIdUrl(
		protonDBHost,
		appPath,
		strconv.FormatInt(int64(steamAppId), 10))
}
