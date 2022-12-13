package mobygames_integration

import "net/url"

func GameUrl(id string) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   mobyGamesHost,
		Path:   gamePath + id,
	}

	return u
}
