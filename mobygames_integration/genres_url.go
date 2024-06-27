package mobygames_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func GenresUrl(apiKey string) *url.URL {
	u := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   apiMobyGamesHost,
		Path:   genresPath,
	}

	q := u.Query()
	q.Set("api_key", apiKey)
	u.RawQuery = q.Encode()

	return u
}
