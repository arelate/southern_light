package gog_integration

import (
	"net/url"
)

func OnLoginSuccessUrl() *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   WwwGogHost,
		Path:   onLoginSuccessPath,
	}
}
