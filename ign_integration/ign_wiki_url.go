package ign_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"path/filepath"
)

func WikiPageUrl(slug, page string) *url.URL {
	path := slug
	if page != "" {
		path = filepath.Join(slug, page)
	}

	return southern_light.SuffixIdUrl(ignHost, wikisPath, path)
}

func WikiUrl(slug string) *url.URL {
	return southern_light.SuffixIdUrl(ignHost, wikisPath, slug)
}
