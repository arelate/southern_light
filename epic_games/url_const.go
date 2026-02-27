package epic_games

import "net/url"

const httpsScheme = "https"

const (
	rootHost        = "epicgames.com"
	accountHost     = "account-public-service-prod.ol." + rootHost
	entitlementHost = "entitlement-public-service-prod.ol." + rootHost
	catalogHost     = "catalog-public-service-prod.ol." + rootHost
)

const (
	apiRedirectPath          = "/id/api/redirect"
	accountApiOauthTokenPath = "/account/api/oauth/token"
	entitlementsPathTemplate = "/entitlement/api/account/{accountId}/entitlements"
	catalogItemPathTemplate  = "/catalog/api/shared/namespace/{namespace}/bulk/items"
)

func HostUrl() *url.URL {
	return new(url.URL{Scheme: httpsScheme, Host: rootHost})
}
