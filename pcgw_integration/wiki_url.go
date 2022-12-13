package pcgw_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

const (
	pageParam  = "page"
	curIdParam = "curid"
)

func WikiGOGUrl(gogId string) *url.URL {
	u := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   pcgwHost,
		Path:   apiGOGPath,
	}

	q := u.Query()
	q.Add(pageParam, gogId)
	u.RawQuery = q.Encode()

	return u
}

func WikiUrl(pageId string) *url.URL {
	u := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   pcgwHost,
		Path:   wikiPath,
	}

	q := u.Query()
	q.Add(curIdParam, pageId)
	u.RawQuery = q.Encode()

	return u
}
