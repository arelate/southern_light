package vndb_integration

import "net/url"

func ItemUrl(id string) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   vndbHost,
		Path:   itemPath + id,
	}

	return u
}
