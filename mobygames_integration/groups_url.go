package mobygames_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"strconv"
)

func GroupsUrl(limit, offset int, apiKey string) *url.URL {

	// normalize parameters
	if limit < 1 || limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	u := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   apiMobyGamesHost,
		Path:   groupsPath,
	}

	q := u.Query()
	q.Set("limit", strconv.FormatInt(int64(limit), 10))
	q.Set("offset", strconv.FormatInt(int64(offset), 10))
	q.Set("api_key", apiKey)
	u.RawQuery = q.Encode()

	return u
}
