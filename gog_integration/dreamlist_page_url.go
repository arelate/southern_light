package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"strconv"
)

const (
	dreamlistLimit           = 50
	globalWishlistVotesCount = "global_wishlist_votes_count"
)

func DreamlistPageUrl(page string) *url.URL {
	dreamlistPage := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   gamesDbHost,
		Path:   gamesDbWishlistedGamesPath,
	}

	q := dreamlistPage.Query()
	q.Add("page", page)
	q.Add("limit", strconv.Itoa(dreamlistLimit))
	q.Add("sort", globalWishlistVotesCount)
	q.Add("order", "desc")
	dreamlistPage.RawQuery = q.Encode()

	return dreamlistPage
}
