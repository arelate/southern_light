package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"strings"
)

func GamesDbGogExternalReleaseUrl(id string) *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   gamesDbHost,
		Path:   strings.Replace(gamesDbGogExternalReleasePathTemplate, "{id}", id, 1),
	}
}

func GamesDbGenericExternalReleaseUrl(id string) *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   gamesDbHost,
		Path:   strings.Replace(gamesDbGenericExternalReleasePathTemplate, "{id}", id, 1),
	}
}
