package steam_integration

import (
	"net/url"
)

func DeckAppCompatibilityReportUrl(appId string) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   StoreHost,
		Path:   deckAppCompatibilityReportPath,
	}

	q := u.Query()
	q.Add("nAppID", appId)
	q.Add("l", "english")
	u.RawQuery = q.Encode()

	return u
}
