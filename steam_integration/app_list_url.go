package steam_integration

import (
	"net/url"
)

func AppListUrl() *url.URL {
	// https://partner.steamgames.com/doc/webapi/ISteamApps#GetAppList
	return &url.URL{
		Scheme: httpsScheme,
		Host:   ApiHost,
		Path:   getAppsListV2,
	}
}

// DefaultSteamAppListUrl is a vangogh_local_data specific wrapper of steam_integration URL func
func DefaultSteamAppListUrl(_ string) *url.URL {
	return AppListUrl()
}
