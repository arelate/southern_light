package wikipedia_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func WikiUrl(id string) *url.URL {
	return southern_light.SuffixIdUrl(enWikipediaHost, wikiPath, id)
}

func WikiRawUrl(id string) *url.URL {
	u := WikiUrl(id)

	q := u.Query()
	q.Add("action", "raw")
	u.RawQuery = q.Encode()

	return u
}
