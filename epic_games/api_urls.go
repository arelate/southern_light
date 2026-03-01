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

func EntitlementsUrl(accountId string, start, count int) *url.URL {
	u := new(url.URL{
		Scheme: httpsScheme,
		Host:   entitlementHost,
		Path:   strings.Replace(entitlementsPathTemplate, "{accountId}", accountId, 1),
	})

	q := u.Query()

	if start > 0 {
		q.Add("start", strconv.Itoa(start))
	}
	if count > 0 {
		q.Add("count", strconv.Itoa(count))
	}

	u.RawQuery = q.Encode()

	return u
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

func LauncherGameAssetsUrl(platform string, label string) *url.URL {
	u := new(url.URL{
		Scheme: httpsScheme,
		Host:   launcherHost,
		Path:   strings.Replace(launcherApiPublicAssetsPathTemplate, "{platform}", platform, 1),
	})

	q := u.Query()

	q.Add("label", label)

	u.RawQuery = q.Encode()

	return u
}

func LauncherGameManifestUrl(namespace, catalogItemId, appName string, platform, label string) *url.URL {

	path := strings.Replace(launcherApiPublicAssetsNamespaceCatalogItemAppLabelPathTemplate, "{namespace}", namespace, 1)
	path = strings.Replace(path, "{catalogItemId}", catalogItemId, 1)
	path = strings.Replace(path, "{appName}", appName, 1)
	path = strings.Replace(path, "{platform}", platform, 1)
	path = strings.Replace(path, "{label}", label, 1)

	return new(url.URL{
		Scheme: httpsScheme,
		Host:   launcherHost,
		Path:   path,
	})
}

func LauncherManifestsUrl(platform, label string) *url.URL {
	u := new(url.URL{
		Scheme: httpsScheme,
		Host:   launcherHost,
		Path:   strings.Replace(launcherApiPublicAssetsV2PlatformLauncherPathTemplate, "{platform}", platform, 1),
	})

	q := u.Query()

	q.Add("label", label)

	u.RawQuery = q.Encode()

	return u
}

func LibraryItemsUrl(includeMetadata bool, cursor string) *url.URL {
	u := new(url.URL{
		Scheme: httpsScheme,
		Host:   libraryHost,
		Path:   libraryItemsPath,
	})

	q := u.Query()

	q.Add("includeMetadata", strconv.FormatBool(includeMetadata))
	if cursor != "" {
		q.Add("cursor", cursor)
	}

	u.RawQuery = q.Encode()

	return u
}
