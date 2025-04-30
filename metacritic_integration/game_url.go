package metacritic_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"path"
	"strings"
)

func GameUrl(id string) *url.URL {
	gamePath := strings.Replace(gamePathTemplate, "{id}", id, 1)

	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   metacriticHost,
		Path:   gamePath,
	}
}

func ParseId(metacriticUrl string) (string, error) {
	gameUrl, err := url.Parse(metacriticUrl)
	if err != nil {
		return "", err
	}

	_, id := path.Split(gameUrl.Path)
	return id, nil
}
