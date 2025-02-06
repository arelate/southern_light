package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"strconv"
)

const limit = 48

func CatalogPageUrl(page string) *url.URL {
	return catalogPageUrl(page, CatalogSortByReleaseDate, true)
}

func catalogPageUrl(page string, sortBy CatalogSortOrder, desc bool) *url.URL {

	catalogPage := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   CatalogHost,
		Path:   catalogPath,
	}

	q := catalogPage.Query()
	q.Add("page", page)
	q.Add("limit", strconv.Itoa(limit))
	q.Add("order", CatalogOrder(sortBy, desc))
	catalogPage.RawQuery = q.Encode()

	return catalogPage
}
