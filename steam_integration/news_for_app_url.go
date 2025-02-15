package steam_integration

import (
	"net/url"
)

func NewsForAppUrl(appId string) *url.URL {
	//https://partner.steamgames.com/doc/webapi/ISteamNews#GetNewsForApp
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   ApiHost,
		Path:   getNewsForAppV2,
	}

	q := u.Query()
	q.Add("appid", appId)
	u.RawQuery = q.Encode()

	return u
}
