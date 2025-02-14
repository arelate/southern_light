package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func LicencesUrl() *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   menuHost,
		Path:   accountLicencesPath,
	}
}

func DefaultLicencesUrl(_ string) *url.URL {
	return LicencesUrl()
}
