package gog_integration

import (
	"net/url"
)

func ReCaptchaUrl() *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   reCaptchaHost,
		Path:   reCaptchaPath}
}
