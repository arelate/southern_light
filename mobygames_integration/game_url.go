package mobygames_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func GameUrl(id string) *url.URL {
	return southern_light.SuffixIdUrl(mobyGamesHost, gamePath, id)
}
