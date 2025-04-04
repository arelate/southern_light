package pcgw_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

const (
	pageParam  = "page"
	appIdParam = "appid"
	curIdParam = "curid"
)

func WikiSteamUrl(steamAppId string) *url.URL {
	u := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   pcgwHost,
		Path:   apiAppIdPath,
	}

	q := u.Query()
	q.Add(appIdParam, steamAppId)
	u.RawQuery = q.Encode()

	return u
}

func WikiGogUrl(gogId string) *url.URL {
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

func WikiRawUrl(pageId string) *url.URL {
	u := WikiUrl(pageId)

	q := u.Query()
	q.Add("action", "raw")
	u.RawQuery = q.Encode()

	return u
}
