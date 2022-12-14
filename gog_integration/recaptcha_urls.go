package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func ReCaptchaUrl() *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   reCaptchaHost,
		Path:   reCaptchaPath}
}
