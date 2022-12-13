package gog_integration

import "net/url"

func UserWishlistUrl() *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   WwwGogHost,
		Path:   userWishlistPath,
	}
}

func DefaultUserWishlistUrl(_ string) *url.URL {
	return UserWishlistUrl()
}
