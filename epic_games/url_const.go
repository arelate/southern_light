package epic_games

import "net/url"

const httpsScheme = "https"

const (
	epicGamesHost = "www.epicgames.com"
)

const (
	apiRedirectPath = "/id/api/redirect"
)

func HostUrl() *url.URL {
	return new(url.URL{Scheme: httpsScheme, Host: epicGamesHost})
}
