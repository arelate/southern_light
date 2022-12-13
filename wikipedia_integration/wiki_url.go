package wikipedia_integration

import "net/url"

func WikiUrl(id string) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   enWikipediaHost,
		Path:   wikiPath + id,
	}

	return u
}
