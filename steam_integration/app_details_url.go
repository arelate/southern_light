package steam_integration

import (
	"net/url"
	"strconv"
)

func AppDetailsUrl(appId uint32) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   StoreHost,
		Path:   appDetailsPath,
	}

	q := u.Query()
	q.Add("appids", strconv.FormatInt(int64(appId), 10))
	u.RawQuery = q.Encode()

	return u
}
