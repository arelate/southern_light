package protondb_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"strings"
)

func SummaryUrl(appId string) *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   protonDBHost,
		Path:   strings.Replace(summaryPathTemplate, "{0}", appId, -1),
	}
}
