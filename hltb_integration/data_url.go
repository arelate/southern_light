package hltb_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"strings"
)

const (
	gameIdParam = "gameId"
)

func DataUrl(buildId, id string) *url.URL {
	u := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   hltbHost,
		Path: strings.Replace(
			strings.Replace(
				dataPathTemplate, "{build}", buildId, -1),
			"{id}", id, -1),
	}

	q := u.Query()
	q.Set(gameIdParam, id)
	u.RawQuery = q.Encode()

	return u
}
