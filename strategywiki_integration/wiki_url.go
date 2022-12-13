package strategywiki_integration

import "net/url"

func WikiUrl(id string) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   strategyWikiHost,
		Path:   wikiPath + id,
	}

	return u
}
