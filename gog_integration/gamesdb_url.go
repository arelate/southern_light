package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"strings"
)

func GamesDbUrl(id string) *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   gamesDbHost,
		Path:   strings.Replace(gamesDbPathTemplate, "{id}", id, 1),
	}
}
