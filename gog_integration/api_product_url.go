package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

var expandValues = []string{
	"downloads",
	"expanded_dlcs",
	"description",
	"screenshots",
	"videos",
	"related_products",
	"changelog"}

func ApiProductUrl(id string) *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   apiHost,
		Path:   apiV2GamesPath + id,
	}
}
