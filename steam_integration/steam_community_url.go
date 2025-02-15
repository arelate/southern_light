package steam_integration

import (
	"net/url"
	"path"
)

func SteamCommunityUrl(appId string) *url.URL {
	return &url.URL{
		Scheme: httpsScheme,
		Host:   SteamCommunityHost,
		Path:   path.Join(appPath, appId),
	}
}

func SteamGuidesUrl(appId string) *url.URL {
	scu := SteamCommunityUrl(appId)
	scu.Path += guidesPath
	return scu
}
