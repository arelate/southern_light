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
)

// paths
const (
	// Steam Web API paths
	iSteamAppsInterfacePath   = "/ISteamApps"
	getAppsList               = iSteamAppsInterfacePath + "/GetAppList"
	getAppsListV2             = getAppsList + "/v2"
	iSteamNewsInterfacePath   = "/ISteamNews"
	getNewsForApp             = iSteamNewsInterfacePath + "/GetNewsForApp"
	getNewsForAppV2           = getNewsForApp + "/v2"
	getAppReviewsPathTemplate = "/appreviews/{appId}"

	// Steam Website paths
	appPath = "/app"

	// Store paths
	deckAppCompatibilityReportPath = "/saleaction/ajaxgetdeckappcompatibilityreport"
	appDetailsPath                 = "/api/appdetails"
	guidesPath                     = "/guides"
)

type SteamUrlFunc func(appId string) *url.URL
