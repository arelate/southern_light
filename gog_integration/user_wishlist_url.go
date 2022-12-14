package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func UserWishlistUrl() *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   WwwGogHost,
		Path:   userWishlistPath,
	}
}

func DefaultUserWishlistUrl(_ string) *url.URL {
	return UserWishlistUrl()
}
