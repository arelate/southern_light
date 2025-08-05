package steam_integration

import "net/url"

func AppListUrl() *url.URL {
	// https://partner.steamgames.com/doc/webapi/ISteamApps#GetAppList
	return &url.URL{
		Scheme: httpsScheme,
		Host:   ApiHost,
		Path:   getAppsListV2,
	}
}
