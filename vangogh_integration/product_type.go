package vangogh_integration

import (
	"errors"
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

	SteamAppList
	SteamAppDetails
	SteamAppNews
	SteamAppReviews
	SteamDeckCompatibilityReport

	// PCGamingWiki product types

	PcgwGogPageId
	PcgwSteamPageId
	PcgwRaw

	// Wikipedia product types

	WikipediaRaw

	// HLTB product types

	HltbRootPage
	HltbData

	// ProtonDB product types

	ProtonDbSummary

	// OpenCritic product types

	OpenCriticApiGame
)

var productTypeStrings = map[ProductType]string{
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

	SteamAppList:                 "steam-app-list",
	SteamAppDetails:              "steam-app-details",
	SteamAppNews:                 "steam-app-news",
	SteamAppReviews:              "steam-app-reviews",
	SteamDeckCompatibilityReport: "steam-deck-compatibility-report",

	// PCGamingWiki product types

	PcgwSteamPageId: "pcgw-steam-page-id",
	PcgwGogPageId:   "pcgw-gog-page-id",
	PcgwRaw:         "pcgw-raw",

	// Wikipedia product types

	WikipediaRaw: "wikipedia-raw",

	// HLTB product types

	HltbRootPage: "hltb-root-page",
	HltbData:     "hltb-data",

	// ProtonDB product types

	ProtonDbSummary: "protondb-summary",

	// OpenCritic product types

	OpenCriticApiGame: "opencritic-api-game",
}

var productTypePfx = map[ProductType]string{
	UnknownProductType:           "upt",
	UserAccessToken:              "uat",
	Licences:                     "l",
	UserWishlist:                 "uw",
	CatalogPage:                  "cp",
	OrderPage:                    "op",
	AccountPage:                  "ap",
	ApiProducts:                  "api",
	Details:                      "d",
	GamesDbGogProducts:           "ggp",
	SteamAppList:                 "sl",
	SteamAppDetails:              "sd",
	SteamAppNews:                 "sn",
	SteamAppReviews:              "sr",
	SteamDeckCompatibilityReport: "sd",
	PcgwSteamPageId:              "pcs",
	PcgwGogPageId:                "pcg",
	PcgwRaw:                      "pcr",
	WikipediaRaw:                 "wr",
	HltbRootPage:                 "hr",
	HltbData:                     "hd",
	ProtonDbSummary:              "prs",
	OpenCriticApiGame:            "oa",
}

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

func ProductTypeId(pt ProductType, id string) (string, error) {
	pfx, ok := productTypePfx[pt]
	if ok {
		return pfx + id, nil
	}

	return "", errors.New("no prefix for " + pt.String())
}

func ProductTypesCloValues() []string {
	ptsStr := make([]string, 0)
	for pt := range AllProductTypes() {
		ptsStr = append(ptsStr, pt.String())
	}
	return ptsStr
}
