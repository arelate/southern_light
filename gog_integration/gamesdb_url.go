package gog_integration

import (
	"net/url"
	"strings"

	"github.com/arelate/southern_light"
)

func GamesDbGenericUrl(id string) *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   gamesDbHost,
		Path:   strings.Replace(gamesDbGenericExternalReleasesPathTemplate, "{id}", id, 1),
	}
}

func GamesDbGogUrl(id string) *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   gamesDbHost,
		Path:   strings.Replace(gamesDbGogExternalReleasesPathTemplate, "{id}", id, 1),
	}
}

func GamesDbSteamUrl(appId string) *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   gamesDbHost,
		Path:   strings.Replace(gamesDbSteamExternalReleasesPathTemplate, "{appId}", appId, 1),
	}
}

func GamesDbEpicUrl(appName string) *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   gamesDbHost,
		Path:   strings.Replace(gamesDbEpicExternalReleasesPathTemplate, "{appName}", appName, 1),
	}
}
