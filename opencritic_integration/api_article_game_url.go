package opencritic_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"strings"
)

func ApiArticleGameUrl(id string) *url.URL {
	path := strings.Replace(apiArticleGamePathTemplate, "{id}", id, 1)

	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   openCriticApiHost,
		Path:   path,
	}
}
