package vangogh_integration

import (
	"maps"
	"slices"
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

var productTypeStrings = map[ProductType]string{
	UnknownProductType: "unknown-product-type",
	// GOG.com product types
	CatalogPage:            "catalog-page",
	CatalogProducts:        "catalog-products",
	AccountPage:            "account-page",
	AccountProducts:        "account-products",
	UserWishlist:           "user-wishlist",
	UserWishlistProducts:   "user-wishlist-products",
	Details:                "details",
	ApiProducts:            "api-products",
	Licences:               "licences",
	OrderPage:              "order-page",
	Orders:                 "orders",
	UserAccessToken:        "user-access-token",
	DreamlistPage:          "dreamlist-page",
	GamesDbGogProducts:     "gamesdb-gog-products",
	GamesDbGenericProducts: "gamesdb-generic-products",
	// Steam product types
	SteamAppDetails:              "steam-app-details",
	SteamAppNews:                 "steam-app-news",
	SteamAppReviews:              "steam-app-reviews",
	SteamDeckCompatibilityReport: "steam-deck-compatibility-report",
	// PCGamingWiki product types
	PcgwSteamPageId:   "pcgw-steam-page-id",
	PcgwGogPageId:     "pcgw-gog-page-id",
	PcgwEngine:        "pcgw-engine",
	PcgwExternalLinks: "pcgw-external-links",
	// HLTB product types
	HltbRootPage: "hltb-root-page",
	HltbData:     "hltb-data",
	// ProtonDB product types
	ProtonDbSummary: "protondb-summary",
}

// the list is intentionally scoped to very few types we anticipate
// will be interesting to output in human-readable form
var productTypeHumanReadableStrings = map[ProductType]string{
	CatalogProducts:      "store",
	UserWishlistProducts: "wishlist",
	AccountProducts:      "account",
	Details:              "account",
	SteamAppNews:         "news",
}

func (pt ProductType) String() string {
	str, ok := productTypeStrings[pt]
	if ok {
		return str
	}

	return productTypeStrings[UnknownProductType]
}

func (pt ProductType) HumanReadableString() string {
	if hs, ok := productTypeHumanReadableStrings[pt]; ok {
		return hs
	} else {
		return pt.String()
	}
}

func ParseProductType(productType string) ProductType {
	for pt, str := range productTypeStrings {
		if str == productType {
			return pt
		}
	}
	return UnknownProductType
}

func IsValidProductType(pt ProductType) bool {
	_, ok := productTypeStrings[pt]
	return ok && pt != UnknownProductType
}

func GOGPagedProducts() []ProductType {
	return []ProductType{
		CatalogPage,
		AccountPage,
		OrderPage,
	}
}

func GOGArrayProducts() []ProductType {
	return []ProductType{
		Licences,
		UserWishlist,
	}
}

func HltbArrayProducts() []ProductType {
	return []ProductType{
		HltbRootPage,
	}
}

func FastPageFetchProducts() []ProductType {
	return []ProductType{
		OrderPage,
	}
}

var gogDetailMainProductTypes = map[ProductType][]ProductType{
	Details: { /*LicenceProducts, */ AccountProducts},
	ApiProducts: {
		CatalogProducts,
		AccountProducts,
		ApiProducts, // includes-games, is-included-in-games, requires-games, is-required-by-games
	},
	GamesDbGogProducts: {
		CatalogProducts,
	},
}

var steamDetailMainProductTypes = map[ProductType][]ProductType{
	//Steam product types are updated on GOG.com store or account product changes
	SteamAppNews: {
		CatalogProducts,
		AccountProducts,
	},
	SteamAppReviews: {
		CatalogProducts,
		AccountProducts,
	},
	SteamDeckCompatibilityReport: {
		CatalogProducts,
		AccountProducts,
	},
}

var pcgwDetailMainProductTypes = map[ProductType][]ProductType{
	//PCGamingWiki product types are updated on GOG.com store or account product changes
	PcgwSteamPageId: {
		CatalogProducts,
		AccountProducts,
	},
	PcgwGogPageId: {
		CatalogProducts,
		AccountProducts,
	},
	PcgwEngine: {
		CatalogProducts,
		AccountProducts,
	},
	PcgwExternalLinks: {
		CatalogProducts,
		AccountProducts,
	},
}

var hltbDetailMainProductTypes = map[ProductType][]ProductType{
	HltbData: {
		CatalogProducts,
		AccountProducts,
	},
}

var protonDBDetailMainProductTypes = map[ProductType][]ProductType{
	ProtonDbSummary: {
		CatalogProducts,
		AccountProducts,
	},
}

func GOGDetailProducts() []ProductType {
	return slices.Collect(maps.Keys(gogDetailMainProductTypes))
}

func SteamDetailProducts() []ProductType {
	return slices.Collect(maps.Keys(steamDetailMainProductTypes))
}

func PcgwDetailProducts() []ProductType {
	return slices.Collect(maps.Keys(pcgwDetailMainProductTypes))
}

func HltbDetailProducts() []ProductType {
	return slices.Collect(maps.Keys(hltbDetailMainProductTypes))
}

func ProtonDBDetailProducts() []ProductType {
	return slices.Collect(maps.Keys(protonDBDetailMainProductTypes))
}

func MainProductTypes(pt ProductType) []ProductType {
	if IsGOGDetailProduct(pt) {
		return gogMainProductTypes(pt)
	} else if IsSteamDetailProduct(pt) {
		return steamMainProductTypes(pt)
	} else if IsPcgwDetailProduct(pt) {
		return pcgwMainProductTypes(pt)
	} else if IsHltbDetailProduct(pt) {
		return hltbMainProductTypes(pt)
	} else if IsProtonDBDetailProduct(pt) {
		return protonDBMainProductTypes(pt)
	} else {
		return nil
	}
}

func gogMainProductTypes(pt ProductType) []ProductType {
	return gogDetailMainProductTypes[pt]
}

func steamMainProductTypes(pt ProductType) []ProductType {
	return steamDetailMainProductTypes[pt]
}

func pcgwMainProductTypes(pt ProductType) []ProductType {
	return pcgwDetailMainProductTypes[pt]
}

func hltbMainProductTypes(pt ProductType) []ProductType {
	return hltbDetailMainProductTypes[pt]
}

func protonDBMainProductTypes(pt ProductType) []ProductType {
	return protonDBDetailMainProductTypes[pt]
}

func GOGRemoteProducts() []ProductType {
	remote := GOGPagedProducts()
	remote = append(remote, GOGArrayProducts()...)
	return append(remote, GOGDetailProducts()...)
}

func SteamRemoteProducts() []ProductType {
	return SteamDetailProducts()
}

func PcgwRemoteProducts() []ProductType {
	return PcgwDetailProducts()
}

func HltbRemoteProducts() []ProductType {
	remote := HltbArrayProducts()
	return append(remote, HltbDetailProducts()...)
}

func ProtonDBRemoteProducts() []ProductType {
	return ProtonDBDetailProducts()
}

func LocalProducts() []ProductType {
	lps := slices.Collect(maps.Values(splitProductTypes))
	lps = append(lps, GOGDetailProducts()...)
	lps = append(lps, SteamDetailProducts()...)
	lps = append(lps, PcgwRemoteProducts()...)
	lps = append(lps, HltbRemoteProducts()...)
	lps = append(lps, ProtonDBRemoteProducts()...)

	return lps
}

func RemoteProducts() []ProductType {
	rps := GOGRemoteProducts()
	rps = append(rps, SteamRemoteProducts()...)
	rps = append(rps, PcgwRemoteProducts()...)
	rps = append(rps, HltbRemoteProducts()...)
	rps = append(rps, ProtonDBRemoteProducts()...)

	return rps
}

var requireAuth = []ProductType{
	AccountPage,
	UserWishlist,
	Details,
	Licences,
	OrderPage,
}

var splitProductTypes = map[ProductType]ProductType{
	CatalogPage: CatalogProducts,
	AccountPage: AccountProducts,
	//Licences:     LicenceProducts,
	UserWishlist: UserWishlistProducts,
	OrderPage:    Orders,
}

func SplitProductType(pt ProductType) ProductType {
	splitProductType, ok := splitProductTypes[pt]
	if ok {
		return splitProductType
	}

	return UnknownProductType
}

var supportsGetItems = []ProductType{
	Details,
	ApiProducts,
	Licences,
	UserWishlist,
	GamesDbGogProducts,
	SteamAppNews,
	SteamAppReviews,
	SteamAppDetails,
	SteamDeckCompatibilityReport,
	PcgwSteamPageId,
	PcgwGogPageId,
	PcgwEngine,
	PcgwExternalLinks,
	HltbRootPage,
	HltbData,
	ProtonDbSummary,
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

func SupportedPropertiesOnly(pt ProductType, properties []string) []string {
	supported := make([]string, 0, len(properties))
	for _, prop := range properties {
		if IsSupportedProperty(pt, prop) {
			supported = append(supported, prop)
		}
	}
	return supported
}
