package hltb_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func RootUrl() *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   hltbHost,
	}
}
