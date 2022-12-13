package gog_integration

import (
	"net/http"
	"net/url"
)

func wishlistOperation(
	httpClient *http.Client,
	ids []string,
	wishlistUrl func(string) *url.URL) error {

	for _, id := range ids {
		wUrl := wishlistUrl(id)
		resp, err := httpClient.Get(wUrl.String())
		if err != nil {
			return err
		}

		if err := resp.Body.Close(); err != nil {
			return err
		}
	}

	return nil
}

func AddToWishlist(
	httpClient *http.Client,
	ids ...string) error {
	return wishlistOperation(httpClient, ids, AddToWishlistUrl)
}

func RemoveFromWishlist(
	httpClient *http.Client,
	ids ...string) error {
	return wishlistOperation(httpClient, ids, RemoveFromWishlistUrl)
}
