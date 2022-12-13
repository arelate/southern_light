package gog_integration

import (
	"net/url"
)

func ManualDownloadUrl(manualDownload string) *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   WwwGogHost,
		Path:   manualDownload,
	}
}
