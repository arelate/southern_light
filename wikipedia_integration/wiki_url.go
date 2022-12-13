package wikipedia_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func WikiUrl(id string) *url.URL {
	return southern_light.SuffixIdUrl(enWikipediaHost, wikiPath, id)
}
