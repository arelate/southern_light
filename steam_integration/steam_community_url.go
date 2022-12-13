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
