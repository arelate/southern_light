package vangogh_integration

import (
	"iter"
	"maps"
	"slices"
	"strconv"

	"github.com/boggydigital/kevlar"
)

type ProductType int

// The types below are intentionally assigned stable sparse constant values:
// - new values in the same group should be added after the last existing value
// - new groups should be added after the last group
// - there is space to add new distribution services above PCGamingWiki

const (
	// cannot use iota, since product type id representation requires stable values
	UnknownProductType ProductType = 0

	// GOG.com product types
	GogCatalogPage     ProductType = 100
	GogAccountPage     ProductType = 101
	GogUserWishlist    ProductType = 102
	GogDetails         ProductType = 103
	GogApiProducts     ProductType = 104
	GogLicences        ProductType = 105
	GogOrderPage       ProductType = 106
	GogUserAccessToken ProductType = 107

	// GamesDB (GOG Galaxy)
	GamesDbGogProducts ProductType = 200

	// Steam product types
	SteamAppDetails              ProductType = 300
	SteamAppNews                 ProductType = 301
	SteamAppReviews              ProductType = 302
	SteamDeckCompatibilityReport ProductType = 303
	SteamAppInfo                 ProductType = 304

	// Epic Games Store
	EgsCatalogItems  ProductType = 400
	EgsGameManifests ProductType = 401
	EgsManifests     ProductType = 402

	// new distributions services should be added here

	// PCGamingWiki product types
	PcgwGogPageId   ProductType = 1000
	PcgwSteamPageId ProductType = 1001
	PcgwRaw         ProductType = 1002

	// Wikipedia product types
	WikipediaRaw ProductType = 1100

	// HLTB product types
	HltbRootPage ProductType = 1200
	HltbData     ProductType = 1201

	// ProtonDB product types
	ProtonDbSummary ProductType = 1300

	// OpenCritic product types
	OpenCriticApiGame ProductType = 1400

	// vangogh/theo product types
	GogChecksums      ProductType = 9000
	GogFilenames      ProductType = 9001
	GogImages         ProductType = 9002
	AvailableProducts ProductType = 9003

	// new data type groups should be added here
)

var productTypeStrings = map[ProductType]string{
	UnknownProductType: "unknown-product-type",

	// GOG.com product types

	GogUserAccessToken: "gog-user-access-token",
	GogLicences:        "gog-licences",
	GogUserWishlist:    "gog-user-wishlist",
	GogCatalogPage:     "gog-catalog-page",
	GogOrderPage:       "gog-order-page",
	GogAccountPage:     "gog-account-page",
	GogApiProducts:     "gog-api-products",
	GogDetails:         "gog-details",

	// GamesDB (GOG Galaxy)

	GamesDbGogProducts: "gamesdb-gog-products",

	// Steam product types

	SteamAppDetails:              "steam-app-details",
	SteamAppNews:                 "steam-app-news",
	SteamAppReviews:              "steam-app-reviews",
	SteamDeckCompatibilityReport: "steam-deck-compatibility-report",
	SteamAppInfo:                 "steam-app-info",

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

	// vangogh/theo product types

	GogChecksums:      "gog-checksums",
	GogFilenames:      "gog-filenames",
	GogImages:         "gog-images",
	AvailableProducts: "available-products",
}

var gogPurchaseProductTypes = []ProductType{
	GogLicences,
	GogUserWishlist,
	GogAccountPage,
	GogDetails,
}

func GogPurchaseProductTypes() []ProductType {
	return gogPurchaseProductTypes
}

func ExtraProductTypes() iter.Seq[ProductType] {
	return func(yield func(ProductType) bool) {
		for pt := range AllProductTypes() {
			if slices.Contains(gogPurchaseProductTypes, pt) {
				continue
			}
			if !yield(pt) {
				return
			}
		}
	}
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

func ProductTypeId(pt ProductType, id string) string {
	return strconv.FormatInt(int64(pt), 10) + "-" + id
}

func (pt ProductType) Ext() string {
	switch pt {
	case PcgwRaw:
		fallthrough
	case WikipediaRaw:
		return kevlar.TxtExt
	default:
		return kevlar.JsonExt
	}
}

func (pt ProductType) ContentType() string {
	switch pt {
	case PcgwRaw:
		fallthrough
	case WikipediaRaw:
		return "text/plain"
	default:
		return "application/json"
	}
}

func ProductTypesCloValues() []string {
	ptsStr := make([]string, 0)
	for pt := range AllProductTypes() {
		ptsStr = append(ptsStr, pt.String())
	}
	return ptsStr
}
