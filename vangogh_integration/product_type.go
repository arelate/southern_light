package vangogh_integration

import (
	"iter"
	"maps"
)

type ProductType int

const (
	UnknownProductType ProductType = iota
	// GOG.com product types
	CatalogPage
	AccountPage
	UserWishlist
	Details
	ApiProducts
	Licences
	OrderPage
	UserAccessToken
	// GamesDB (GOG Galaxy)
	GamesDbGogProducts
	// Steam product types
	SteamAppDetails
	SteamAppNews
	SteamAppReviews
	SteamDeckCompatibilityReport
	// PCGamingWiki product types
	PcgwGogPageId
	PcgwSteamPageId
	PcgwEngine
	PcgwExternalLinks
	// HLTB product types
	HltbRootPage
	HltbData
	// ProtonDB product types
	ProtonDbSummary
	// OpenCritic product types
	OpenCriticApiGame
	OpenCriticApiArticle
	OpenCriticApiRating
)

var (
	productTypeStrings = map[ProductType]string{
		UnknownProductType: "unknown-product-type",

		// GOG.com product types

		UserAccessToken:    "user-access-token",
		Licences:           "licences",
		UserWishlist:       "user-wishlist",
		CatalogPage:        "catalog-page",
		OrderPage:          "order-page",
		AccountPage:        "account-page",
		ApiProducts:        "api-products",
		Details:            "details",
		GamesDbGogProducts: "gamesdb-gog-products",

		// Steam product types

		SteamAppDetails:              "steam-app-details",
		SteamAppNews:                 "steam-app-news",
		SteamAppReviews:              "steam-app-reviews",
		SteamDeckCompatibilityReport: "steam-deck-compatibility-report",

		// PCGamingWiki product types

		PcgwSteamPageId:   "pcgw-steam-page-id",
		PcgwGogPageId:     "pcgw-gog-page-id",
		PcgwExternalLinks: "pcgw-external-links",
		PcgwEngine:        "pcgw-engine",

		// HLTB product types

		HltbRootPage: "hltb-root-page",
		HltbData:     "hltb-data",

		// ProtonDB product types

		ProtonDbSummary: "protondb-summary",

		// OpenCritic product types

		OpenCriticApiGame:    "opencritic-api-game",
		OpenCriticApiArticle: "opencritic-api-article",
		OpenCriticApiRating:  "opencritic-api-rating",
	}
)

func AllProductTypes() iter.Seq[ProductType] {
	return maps.Keys(productTypeStrings)
}

func (pt ProductType) String() string {
	str, ok := productTypeStrings[pt]
	if ok {
		return str
	}

	return productTypeStrings[UnknownProductType]
}

func ParseProductType(productType string) ProductType {
	for pt, str := range productTypeStrings {
		if str == productType {
			return pt
		}
	}
	return UnknownProductType
}
