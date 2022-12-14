package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func UserDataUrl() *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   WwwGogHost,
		Path:   userDataPath,
	}
}
