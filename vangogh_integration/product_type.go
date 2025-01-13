package vangogh_integration

import (
	"golang.org/x/exp/maps"
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
	ApiProductsV1
	ApiProductsV2
	Licences
	LicenceProducts
	OrderPage
	Orders
	GamesDbProducts
	// Steam product types
	SteamAppList
	SteamAppNews
	SteamReviews
	SteamStorePage
	SteamDeckCompatibilityReport
	// PCGamingWiki product types
	PCGWPageId
	PCGWEngine
	PCGWExternalLinks
	// HLTB product types
	HLTBRootPage
	HLTBData
	// ProtonDB product types
	ProtonDBSummary
)

var productTypeStrings = map[ProductType]string{
	UnknownProductType: "unknown-product-type",
	// GOG.com product types
	CatalogPage:          "catalog-page",
	CatalogProducts:      "catalog-products",
	AccountPage:          "account-page",
	AccountProducts:      "account-products",
	UserWishlist:         "user-wishlist",
	UserWishlistProducts: "user-wishlist-products",
	Details:              "details",
	ApiProductsV1:        "api-products-v1",
	ApiProductsV2:        "api-products-v2",
	Licences:             "licences",
	LicenceProducts:      "licence-products",
	OrderPage:            "order-page",
	Orders:               "orders",
	GamesDbProducts:      "gamesdb-products",
	// Steam product types
	SteamAppList:   "steam-app-list",
	SteamAppNews:   "steam-app-news",
	SteamReviews:   "steam-reviews",
	SteamStorePage: "steam-store-page",
	//SteamAppDetails:              "steam-app-details",
	SteamDeckCompatibilityReport: "steam-deck-compatibility-report",
	// PCGamingWiki product types
	PCGWPageId:        "pcgw-pageid",
	PCGWEngine:        "pcgw-engine",
	PCGWExternalLinks: "pcgw-external-links",
	// HLTB product types
	HLTBRootPage: "hltb-root-page",
	HLTBData:     "hltb-data",
	// ProtonDB product types
	ProtonDBSummary: "protondb-summary",
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

func HLTBArrayProducts() []ProductType {
	return []ProductType{
		HLTBRootPage,
	}
}

func FastPageFetchProducts() []ProductType {
	return []ProductType{
		OrderPage,
	}
}

var gogDetailMainProductTypes = map[ProductType][]ProductType{
	Details: {LicenceProducts, AccountProducts},
	ApiProductsV1: {
		CatalogProducts,
		AccountProducts,
		ApiProductsV2,
	},
	ApiProductsV2: {
		CatalogProducts,
		AccountProducts,
		ApiProductsV2, // includes-games, is-included-in-games, requires-games, is-required-by-games
	},
	GamesDbProducts: {
		CatalogProducts,
	},
}

var steamDetailMainProductTypes = map[ProductType][]ProductType{
	//Steam product types are updated on GOG.com store or account product changes
	SteamAppNews: {
		CatalogProducts,
		AccountProducts,
	},
	SteamReviews: {
		CatalogProducts,
		AccountProducts,
	},
	SteamStorePage: {
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
	PCGWPageId: {
		CatalogProducts,
		AccountProducts,
	},
	PCGWEngine: {
		CatalogProducts,
		AccountProducts,
	},
	PCGWExternalLinks: {
		CatalogProducts,
		AccountProducts,
	},
}

var hltbDetailMainProductTypes = map[ProductType][]ProductType{
	HLTBData: {
		CatalogProducts,
		AccountProducts,
	},
}

var protonDBDetailMainProductTypes = map[ProductType][]ProductType{
	ProtonDBSummary: {
		CatalogProducts,
		AccountProducts,
	},
}

func GOGDetailProducts() []ProductType {
	return maps.Keys(gogDetailMainProductTypes)
}

func SteamArrayProducts() []ProductType {
	return []ProductType{SteamAppList}
}

func SteamDetailProducts() []ProductType {
	return maps.Keys(steamDetailMainProductTypes)
}

func PCGWDetailProducts() []ProductType {
	return maps.Keys(pcgwDetailMainProductTypes)
}

func HLTBDetailProducts() []ProductType {
	return maps.Keys(hltbDetailMainProductTypes)
}

func ProtonDBDetailProducts() []ProductType { return maps.Keys(protonDBDetailMainProductTypes) }

func MainProductTypes(pt ProductType) []ProductType {
	if IsGOGDetailProduct(pt) {
		return gogMainProductTypes(pt)
	} else if IsSteamDetailProduct(pt) {
		return steamMainProductTypes(pt)
	} else if IsPCGWDetailProduct(pt) {
		return pcgwMainProductTypes(pt)
	} else if IsHLTBDetailProduct(pt) {
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
	remote := SteamArrayProducts()
	return append(remote, SteamDetailProducts()...)
}

func PCGWRemoteProducts() []ProductType {
	return PCGWDetailProducts()
}

func HLTBRemoteProducts() []ProductType {
	remote := HLTBArrayProducts()
	return append(remote, HLTBDetailProducts()...)
}

func ProtonDBRemoteProducts() []ProductType {
	return ProtonDBDetailProducts()
}

func LocalProducts() []ProductType {
	lps := maps.Values(splitProductTypes)
	lps = append(lps, GOGDetailProducts()...)
	lps = append(lps, SteamDetailProducts()...)
	lps = append(lps, PCGWRemoteProducts()...)
	lps = append(lps, HLTBRemoteProducts()...)
	lps = append(lps, ProtonDBRemoteProducts()...)

	return lps
}

func RemoteProducts() []ProductType {
	rps := GOGRemoteProducts()
	rps = append(rps, SteamRemoteProducts()...)
	rps = append(rps, PCGWRemoteProducts()...)
	rps = append(rps, HLTBRemoteProducts()...)
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
	CatalogPage:  CatalogProducts,
	AccountPage:  AccountProducts,
	Licences:     LicenceProducts,
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
	ApiProductsV1,
	ApiProductsV2,
	Licences,
	UserWishlist,
	GamesDbProducts,
	SteamAppList,
	SteamAppNews,
	SteamReviews,
	SteamStorePage,
	//SteamAppDetails,
	SteamDeckCompatibilityReport,
	PCGWPageId,
	PCGWEngine,
	PCGWExternalLinks,
	HLTBRootPage,
	HLTBData,
	ProtonDBSummary,
}

var supportedImageTypes = map[ProductType][]ImageType{
	CatalogProducts: {Image, Screenshots},
	AccountProducts: {Image},
	ApiProductsV1:   {Screenshots},
	ApiProductsV2:   {Image, Screenshots, Hero, Logo},
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
