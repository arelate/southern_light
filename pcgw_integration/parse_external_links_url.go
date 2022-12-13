package pcgw_integration

import (
	"net/url"
	"strings"
)

const (
	pageIdParam = "pageid"
	propParam   = "prop"
)

func ParseExternalLinksUrl(pageId string) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   pcgwHost,
		Path:   apiPath,
	}

	q := u.Query()
	q.Set(actionParam, "parse")
	q.Set(propParam, strings.Join([]string{"externallinks", "iwlinks"}, "|"))
	q.Set(pageIdParam, pageId)
	q.Set(formatParam, "json")

	u.RawQuery = q.Encode()

	return u
}
