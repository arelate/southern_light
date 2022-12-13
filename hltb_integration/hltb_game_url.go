package hltb_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func GameUrl(hltbId string) *url.URL {
	return southern_light.SuffixIdUrl(hltbHost, gamePath, hltbId)
}
