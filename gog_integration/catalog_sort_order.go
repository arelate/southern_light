package gog_integration

type CatalogSortOrder string

const (
	Ascending  = "asc"
	Descending = "desc"
)

const (
	CatalogSortByTitle       CatalogSortOrder = "title"
	CatalogSortByDiscount    CatalogSortOrder = "discount"
	CatalogSortByReleaseDate CatalogSortOrder = "storeReleaseDate"
	CatalogSortByRating      CatalogSortOrder = "reviewsRating"
	CatalogSortByPrice       CatalogSortOrder = "price"
	CatalogSortByBestselling CatalogSortOrder = "bestselling"
)

func CatalogOrder(sortBy CatalogSortOrder, desc bool) string {
	order := ""
	if desc {
		order = "desc:"
	} else {
		order = "asc:"
	}
	order += string(sortBy)

	return order
}
