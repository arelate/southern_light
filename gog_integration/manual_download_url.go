package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func ManualDownloadUrl(manualDownload string) *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   WwwGogHost,
		Path:   manualDownload,
	}
}
