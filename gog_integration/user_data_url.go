package gog_integration

import (
	"net/url"
)

func UserDataUrl() *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   WwwGogHost,
		Path:   userDataPath,
	}
}
