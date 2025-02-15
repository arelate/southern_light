package steam_integration

import (
	"net/url"
)

func AppDetailsUrl(appId string) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   StoreHost,
		Path:   appDetailsPath,
	}

	q := u.Query()
	q.Add("appids", appId)
	u.RawQuery = q.Encode()

	return u
}
