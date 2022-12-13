package igdb_integration

import (
	"net/url"
)

func GameUrl(igdbId string) *url.URL {

	u := &url.URL{
		Scheme: httpsScheme,
		Host:   igdbHost,
		Path:   gamesPath + igdbId,
	}

	return u
}
