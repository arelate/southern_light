package gog_integration

type AccountSortOrder string

const (
	AccountSortByTitle        AccountSortOrder = "title"
	AccountSortByPurchaseDate                  = "date_purchased"
	AccountSortByTags                          = "tags"
	AccountSortByOldestFirst                   = "release_date_asc"
	AccountSortByNewestFirst                   = "release_date_asc"
)
