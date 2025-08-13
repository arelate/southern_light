package steam_integration

import (
	"net/url"
)

const (
	defaultCountryCode = "US"
	defaultLanguage    = "English"
)

func AppDetailsUrl(appId string) *url.URL {
	return appDetailsParamsUrl(appId, defaultCountryCode, defaultLanguage)
}

func appDetailsParamsUrl(appId string, countryCode string, language string) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   StoreHost,
		Path:   appDetailsPath,
	}

	q := u.Query()
	q.Add("appids", appId)
	q.Add("cc", countryCode)
	q.Add("l", language)
	u.RawQuery = q.Encode()

	return u
}
