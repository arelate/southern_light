package steam_integration

import (
	"net/url"
	"strings"
)

func AppReviewsUrl(appId string) *url.URL {
	//https://partner.steamgames.com/doc/store/getreviews
	aru := &url.URL{
		Scheme: httpsScheme,
		Host:   StoreHost,
		Path:   strings.Replace(getAppReviewsPathTemplate, "{appId}", appId, -1),
	}

	q := aru.Query()
	q.Add("json", "1")
	aru.RawQuery = q.Encode()

	return aru
}
