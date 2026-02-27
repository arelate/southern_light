package epic_games

import (
	"net/url"
	"strconv"
	"strings"
)

func ApiRedirectUrl() *url.URL {
	u := new(url.URL{
		Scheme: httpsScheme,
		Host:   rootHost,
		Path:   apiRedirectPath,
	})

	q := u.Query()

	q.Set("clientId", clientId)
	q.Set("responseType", "code")

	u.RawQuery = q.Encode()

	return u
}

func AccountApiOauthTokenUrl() *url.URL {
	return new(url.URL{
		Scheme: httpsScheme,
		Host:   accountHost,
		Path:   accountApiOauthTokenPath,
	})
}

func AccountApiOauthVerifyUrl() *url.URL {
	return new(url.URL{
		Scheme: httpsScheme,
		Host:   accountHost,
		Path:   accountApiOauthVerifyPath,
	})
}

func AccountApiOauthKillUrl(token string) *url.URL {
	return new(url.URL{
		Scheme: httpsScheme,
		Host:   accountHost,
		Path:   strings.Replace(accountApiOauthSessionsKillPathTemplate, "{token}", token, 1),
	})
}

func EntitlementsUrl(accountId string) *url.URL {
	return new(url.URL{
		Scheme: httpsScheme,
		Host:   entitlementHost,
		Path:   strings.Replace(entitlementsPathTemplate, "{accountId}", accountId, 1),
	})
}

func CatalogItemUrl(namespace, itemId string, includeMainGameDetails, includeDlcDetails bool, country, locale string) *url.URL {

	u := new(url.URL{
		Scheme: httpsScheme,
		Host:   catalogHost,
		Path:   strings.Replace(catalogItemPathTemplate, "{namespace}", namespace, 1),
	})

	q := u.Query()

	q.Add("id", itemId)

	if includeMainGameDetails {
		q.Add("includeMainGameDetails", strconv.FormatBool(true))
	}
	if includeDlcDetails {
		q.Add("includeDLCDetails", strconv.FormatBool(true))
	}
	if country != "" {
		q.Add("country", country)
	}
	if locale != "" {
		q.Add("locale", locale)
	}

	u.RawQuery = q.Encode()

	return u
}
