package vangogh_integration

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	UrlIdParameter                    = "id"
	UrlTitleParameter                 = "title"
	UrlSlugParameter                  = "slug"
	UrlOperatingSystemParameter       = "os"
	UrlLanguageCodeParameter          = "lang-code"
	UrlDownloadTypeParameter          = "download-type"
	UrlNoPatchesParameter             = "no-patches"
	UrlDownloadsLayoutParameter       = "downloads-layout"
	UrlAllParameter                   = "all"
	UrlTestParameter                  = "test"
	UrlWineParameter                  = "wine"
	UrlSteamCmdParameter              = "steamcmd"
	UrlSinceHoursAgoParameter         = "since-hours-ago"
	UrlPropertyParameter              = "property"
	UrlProductTypeParameter           = "product-type"
	UrlForceParameter                 = "force"
	UrlPurchasesParameter             = "purchases"
	UrlExtraParameter                 = "extra"
	UrlRelatedApiProductsParameter    = "related-api-products"
	UrlManualUrlParameter             = "manual-url"
	UrlManualUrlFilterParameter       = "manual-url-filter"
	UrlUpdateDetailsParameter         = "update-details"
	UrlValidateParameter              = "validate"
	UrlChecksumsOnlyParameter         = "checksums-only"
	UrlQueuedParameter                = "queued"
	UrlMissingParameter               = "missing"
	UrlDebugParameter                 = "debug"
	UrlImageTypeParameter             = "image-type"
	UrlCookiesParameter               = "cookies"
	UrlFromParameter                  = "from"
	UrlToParameter                    = "to"
	UrlPortParameter                  = "port"
	UrlInsecureCookiesParameter       = "insecure-cookies"
	UrlStdErrParameter                = "stderr"
	UrlUpdatesOnlyParameter           = "updates-only"
	UrlUsernameParameter              = "username"
	UrlPasswordParameter              = "password"
	UrlNewPasswordParameter           = "new-password"
	UrlRoleParameter                  = "role"
	UrlCreateParameter                = "create"
	UrlDeleteParameter                = "delete"
	UrlChangePasswordParameter        = "change-password"
	UrlListParameter                  = "list"
	UrlValidationStatusParameter      = "validation-status"
	UrlNotValidParameter              = "not-valid"
	UrlNewValueParameter              = "new-property-value"
	UrlValueParameter                 = "value"
	UrlConditionParameter             = "condition"
	UrlSectionParameter               = "section"
	UrlImportCookiesParameter         = "import-cookies"
	UrlSteamParameter                 = "steam"
	UrlEpicGamesParameter             = "epic-games"
	UrlVerboseParameter               = "verbose"
	UrlPurgeParameter                 = "purge"
	UrlOriginParameter                = "origin"
	UrlForIdParameter                 = "for-id"
	UrlHeaderParameter                = "header"
	UrlCapsuleParameter               = "capsule"
	UrlHeroParameter                  = "hero"
	UrlLogoParameter                  = "logo"
	UrlIconParameter                  = "icon"
	UrlPinnedPositionParameter        = "pinned-position"
	UrlWidthPctParameter              = "width-pct"
	UrlHeightPctParameter             = "height-pct"
	UrlRemoveParameter                = "remove"
	UrlWorkDirParameter               = "work-dir"
	UrlTaskParameter                  = "task"
	UrlDefaultLauncherParameter       = "default-launcher"
	UrlExeParameter                   = "exe"
	UrlArgParameter                   = "arg"
	UrlEnvParameter                   = "env"
	UrlProtonRuntimeParameter         = "proton-runtime"
	UrlSteamProtonRuntimeParameter    = "steam-proton-runtime"
	UrlProtonOptionsParameter         = "proton-options"
	UrlInstalledParameter             = "installed"
	UrlDownloadsParameter             = "downloads"
	UrlBackupsParameter               = "backups"
	UrlModParameter                   = "mod"
	UrlProgramParameter               = "program"
	UrlInstallBinaryParameter         = "install-binary"
	UrlAvailableProductsParameter     = "available-products"
	UrlLaunchOptionsParameter         = "launch-options"
	UrlSteamShortcutsParameter        = "steam-shortcuts"
	UrlTasksParameter                 = "tasks"
	UrlUpdateParameter                = "update"
	UrlResetParameter                 = "reset"
	UrlKeepDownloadsParameter         = "keep-downloads"
	UrlNoSteamShortcutParameter       = "no-steam-shortcut"
	UrlNoPresetLaunchOptionsParameter = "no-preset-launch-options"
	UrlSteamAppIdParameter            = "steam-appid"
	UrlUrlParameter                   = "url"
)

const (
	defaultLanguageCode = "en"
)

func ValuesFromUrl(u *url.URL, arg string) []string {
	q := u.Query()
	if q.Has(arg) {
		val := q.Get(arg)
		values := strings.Split(val, ",")
		//account for empty strings
		if len(values) == 1 && values[0] == "" {
			values = []string{}
		}
		return values
	}
	return nil
}

func PropertyFromUrl(u *url.URL) string {
	return u.Query().Get(UrlPropertyParameter)
}

func PropertiesFromUrl(u *url.URL) []string {
	return ValuesFromUrl(u, UrlPropertyParameter)
}

func ProductTypeFromUrl(u *url.URL) ProductType {
	return ParseProductType(u.Query().Get(UrlProductTypeParameter))
}

func ProductTypesFromUrl(u *url.URL) []ProductType {

	q := u.Query()

	productTypes := make([]ProductType, 0)
	if !q.Has(UrlProductTypeParameter) {
		return productTypes
	}
	ptParam := q.Get(UrlProductTypeParameter)
	pts := strings.SplitSeq(ptParam, ",")
	for pt := range pts {
		productTypes = append(productTypes, ParseProductType(pt))
	}
	return productTypes
}

func OperatingSystemsFromUrl(u *url.URL) []OperatingSystem {
	osStrings := ValuesFromUrl(u, UrlOperatingSystemParameter)
	operatingSystems := ParseManyOperatingSystems(osStrings)
	if len(operatingSystems) == 0 {
		operatingSystems = append(operatingSystems, AnyOperatingSystem)
	}
	return operatingSystems
}

func LanguageCodesFromUrl(u *url.URL) []string {
	if langCodes := ValuesFromUrl(u, UrlLanguageCodeParameter); len(langCodes) > 0 {
		return langCodes
	} else {
		return []string{defaultLanguageCode}
	}
}

func DownloadTypesFromUrl(u *url.URL) []DownloadType {
	dtStrings := ValuesFromUrl(u, UrlDownloadTypeParameter)
	return ParseManyDownloadTypes(dtStrings)
}

func IdsFromUrl(u *url.URL) ([]string, error) {

	ids := ValuesFromUrl(u, UrlIdParameter)

	if slugs := ValuesFromUrl(u, UrlSlugParameter); len(slugs) > 0 {
		slugIds, err := idsFromSlugs(slugs, nil)
		if err != nil {
			return nil, err
		}
		ids = append(ids, slugIds...)
	}

	return ids, nil
}

func SinceFromUrl(u *url.URL) (int64, error) {
	str := u.Query().Get(UrlSinceHoursAgoParameter)
	var sha int
	var err error
	if str != "" {
		sha, err = strconv.Atoi(str)
		if err != nil {
			return 0, err
		}
	}
	return time.Now().UTC().Unix() - int64(sha*60*60), err
}

func DownloadsLayoutFromUrl(u *url.URL) DownloadsLayout {
	downloadsLayout := DefaultDownloadsLayout
	q := u.Query()
	if q.Has(UrlDownloadsLayoutParameter) {
		if dl := ParseDownloadsLayout(q.Get(UrlDownloadsLayoutParameter)); dl != UnknownDownloadsLayout {
			downloadsLayout = dl
		}
	}
	return downloadsLayout
}
