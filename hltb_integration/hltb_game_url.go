package hltb_integration

import (
	"net/url"
)

func GameUrl(hltbId string) *url.URL {

	u := &url.URL{
		Scheme: httpsScheme,
		Host:   hltbHost,
		Path:   gamePath + hltbId,
	}

	return u
}
