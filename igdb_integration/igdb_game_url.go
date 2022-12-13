package igdb_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func GameUrl(igdbId string) *url.URL {
	return southern_light.SuffixIdUrl(igdbHost, gamesPath, igdbId)
}
