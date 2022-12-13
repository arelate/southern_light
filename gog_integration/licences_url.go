package gog_integration

import (
	"net/url"
)

func LicencesUrl() *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   menuHost,
		Path:   licencesPath,
	}
}

func DefaultLicencesUrl(_ string) *url.URL {
	return LicencesUrl()
}
