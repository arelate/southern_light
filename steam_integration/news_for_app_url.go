package steam_integration

import (
	"net/url"
	"strconv"
)

func NewsForAppUrl(appId uint32) *url.URL {
	//https://partner.steamgames.com/doc/webapi/ISteamNews#GetNewsForApp
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   ApiHost,
		Path:   getNewsForAppV2,
	}

	q := u.Query()
	q.Add("appid", strconv.Itoa(int(appId)))
	u.RawQuery = q.Encode()

	return u
}
