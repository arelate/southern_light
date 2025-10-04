package steam_integration

import "net/url"

// scheme
const (
	httpsScheme = "https"
)

// hosts
const (
	SteamPoweredHost   = "steampowered.com"
	StoreHost          = "store." + SteamPoweredHost
	ApiHost            = "api." + SteamPoweredHost
	SteamCommunityHost = "steamcommunity.com"
	SteamStaticHost    = "shared.akamai.steamstatic.com"
)

// paths
const (
	// Steam Web API paths
	getAppsListV2             = "/ISteamApps/GetAppList/v2"
	getNewsForAppV2           = "/ISteamNews/GetNewsForApp/v2"
	getAppReviewsPathTemplate = "/appreviews/{appId}"

	// Steam Website paths
	appPath = "/app"

	// Store paths
	deckAppCompatibilityReportPath = "/saleaction/ajaxgetdeckappcompatibilityreport"
	appDetailsPath                 = "/api/appdetails"
	guidesPath                     = "/guides"

	assetPathTemplate = "/store_item_assets/steam/apps/{appId}/{asset}"
)

type SteamUrlFunc func(appId string) *url.URL
