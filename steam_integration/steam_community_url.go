package steam_integration

import (
	"net/url"
	"path"
	"strconv"
)

func SteamCommunityUrl(appId uint32) *url.URL {
	return &url.URL{
		Scheme: httpsScheme,
		Host:   SteamCommunityHost,
		Path:   path.Join(appPath, strconv.Itoa(int(appId))),
	}
}

func SteamGuidesUrl(appId uint32) *url.URL {
	scu := SteamCommunityUrl(appId)
	scu.Path += guidesPath
	return scu
}
