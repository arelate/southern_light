package gog_integration

import (
	"net/url"
	"strconv"

	"github.com/arelate/southern_light"
)

const CatalogPagesProductsLimit = 100

func CatalogPageUrl(searchAfter string) *url.URL {
	return catalogPageUrl(searchAfter, CatalogSortByExternalProductId, true)
}

func catalogPageUrl(searchAfter string, sortBy CatalogSortOrder, desc bool) *url.URL {

	catalogPage := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   catalogHost,
		Path:   catalogPath,
	}

	q := catalogPage.Query()
	if searchAfter != "" {
		q.Add("searchAfter", searchAfter)
	}
	q.Add("limit", strconv.Itoa(CatalogPagesProductsLimit))
	q.Add("order", CatalogOrder(sortBy, desc))
	catalogPage.RawQuery = q.Encode()

	return catalogPage
}
