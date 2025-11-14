package gog_integration

import (
	"net/url"

	"github.com/arelate/southern_light"
)

func ManualDownloadUrl(manualDownload string) *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   WwwGogHost,
		Path:   manualDownload,
	}
}
