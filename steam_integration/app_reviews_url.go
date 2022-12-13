package steam_integration

import (
	"net/url"
	"strconv"
	"strings"
)

func AppReviewsUrl(appId uint32) *url.URL {
	//https://partner.steamgames.com/doc/store/getreviews
	aru := &url.URL{
		Scheme: httpsScheme,
		Host:   StoreHost,
		Path:   strings.Replace(getAppReviewsPathTemplate, "{appId}", strconv.Itoa(int(appId)), -1),
	}

	q := aru.Query()
	q.Add("json", "1")
	aru.RawQuery = q.Encode()

	return aru
}
