package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func UserAccessTokenUrl() *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   apiHost,
		Path:   userAccessTokenPath,
	}
}
