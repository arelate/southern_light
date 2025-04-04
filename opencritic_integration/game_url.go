package opencritic_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"strings"
)

func GameUrl(id, slug string) *url.URL {

	path := strings.Replace(gamePathTemplate, "{id}", id, 1)
	path = strings.Replace(path, "{slug}", slug, 1)

	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   openCriticHost,
		Path:   path,
	}
}
