package egs_integration

import "net/url"

const httpsScheme = "https"

const (
	rootHost        = "epicgames.com"
	accountHost     = "account-public-service-prod.ol." + rootHost
	launcherHost    = "launcher-public-service-prod.ol." + rootHost
	entitlementHost = "entitlement-public-service-prod.ol." + rootHost
	catalogHost     = "catalog-public-service-prod.ol." + rootHost
	libraryHost     = "library-service.live.use1a.on." + rootHost
)

const (
	apiRedirectPath                                                 = "/id/api/redirect"
	accountApiOauthTokenPath                                        = "/account/api/oauth/token"
	accountApiOauthVerifyPath                                       = "/account/api/oauth/verify"
	accountApiOauthSessionsKillPathTemplate                         = "/account/api/oauth/sessions/kill/{token}"
	catalogItemPathTemplate                                         = "/catalog/api/shared/namespace/{namespace}/bulk/items"
	entitlementsPathTemplate                                        = "/entitlement/api/account/{accountId}/entitlements"
	launcherApiPublicAssetsPathTemplate                             = "/launcher/api/public/assets/{platform}"
	launcherApiPublicAssetsNamespaceCatalogItemAppLabelPathTemplate = "/launcher/api/public/assets/v2/platform/{platform}/namespace/{namespace}/catalogItem/{catalogItemId}/app/{appName}/label/{label}"
	launcherApiPublicAssetsV2PlatformLauncherPathTemplate           = "/launcher/api/public/assets/v2/platform/{platform}/launcher"
	libraryItemsPath                                                = "/library/api/public/items"
)

func HostUrl() *url.URL {
	return new(url.URL{Scheme: httpsScheme, Host: rootHost})
}
