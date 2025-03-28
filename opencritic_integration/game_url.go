package opencritic_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"path"
)

func GameUrl(id, slug string) *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   openCriticHost,
		Path:   gamePath + path.Join(id, slug),
	}
}
