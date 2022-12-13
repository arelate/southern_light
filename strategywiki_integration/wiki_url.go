package strategywiki_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func WikiUrl(id string) *url.URL {
	return southern_light.SuffixIdUrl(strategyWikiHost, wikiPath, id)
}
