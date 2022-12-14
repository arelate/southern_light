package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"strconv"
)

const limit = 48

func DefaultCatalogPageUrl(page string) *url.URL {
	return CatalogPageUrl(page, CatalogSortByReleaseDate, true)
}

func CatalogPageUrl(page string, sortBy CatalogSortOrder, desc bool) *url.URL {

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
