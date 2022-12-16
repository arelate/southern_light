package ign_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func WikiUrl(slug string) *url.URL {
	return southern_light.SuffixIdUrl(ignHost, wikisPath, slug)
}
