package steam_integration

import (
	"net/url"
	"strconv"
)

func DeckAppCompatibilityReportUrl(appId uint32) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   StoreHost,
		Path:   deckAppCompatibilityReportPath,
	}

	q := u.Query()
	q.Add("nAppID", strconv.FormatInt(int64(appId), 10))
	u.RawQuery = q.Encode()

	return u
}
