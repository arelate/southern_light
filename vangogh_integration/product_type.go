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
	CatalogProducts
	AccountPage
	AccountProducts
	UserWishlist
	UserWishlistProducts
	Details
	ApiProducts
	Licences
	OrderPage
	Orders
	UserAccessToken
	DreamlistPage
	// GamesDB (GOG Galaxy)
	GamesDbGogProducts
	GamesDbGenericProducts
	// Steam product types
	SteamAppDetails
	SteamAppNews
	SteamAppReviews
	SteamDeckCompatibilityReport
	// PCGamingWiki product types
	PcgwSteamPageId
	PcgwGogPageId
	PcgwEngine
	PcgwExternalLinks
	// HLTB product types
	HltbRootPage
	HltbData
	// ProtonDB product types
	ProtonDbSummary
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

var supportedImageTypes = map[ProductType][]ImageType{
	CatalogProducts: {Image, Screenshots},
	AccountProducts: {Image},
	ApiProducts:     {Image, Screenshots, Hero, Logo},
}

func ProductTypesSupportingImageType(imageType ImageType) []ProductType {
	pts := make([]ProductType, 0)
	for pt, its := range supportedImageTypes {
		for _, it := range its {
			if it == imageType {
				pts = append(pts, pt)
				break
			}
		}
	}
	return pts
}
