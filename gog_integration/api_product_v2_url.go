package gog_integration

import (
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

func ApiProductV2Url(id string) *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   apiHost,
		Path:   apiV2GamesPath + id,
	}
}
