package protondb_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"strconv"
	"strings"
)

func SummaryUrl(steamAppId uint32) *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   protonDBHost,
		Path: strings.Replace(summaryPathTemplate,
			"{0}",
			strconv.FormatInt(int64(steamAppId), 10), -1),
	}
}
