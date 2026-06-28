package vangogh_integration

import "github.com/arelate/southern_light/gog_integration"

const (
	IdProperty = "id"

	// GOG Properties

	GogLicencesProperty                   = "gog-licences"
	GogUserWishlistProperty               = "gog-user-wishlist"
	GogAccountProductPageProperty         = "gog-account-product-page"
	GogCatalogProductPageProperty         = "gog-catalog-product-page"
	GogOrderPageProductsProperty          = "gog-order-page-products"
	GogTitleProperty                      = "gog-title"
	GogDevelopersProperty                 = "gog-developers"
	GogPublishersProperty                 = "gog-publishers"
	GogImageProperty                      = "gog-image"
	GogVerticalImageProperty              = "gog-vertical-image"
	GogScreenshotsProperty                = "gog-screenshots"
	GogHeroProperty                       = "gog-hero"
	GogLogoProperty                       = "gog-logo"
	GogIconProperty                       = "gog-icon"
	GogIconSquareProperty                 = "gog-icon-square"
	GogBackgroundProperty                 = "gog-background"
	GogRatingProperty                     = "gog-rating"
	GogProductTypeProperty                = "gog-product-type"
	GogIncludesGamesProperty              = "gog-includes-games"
	GogIsIncludedByGamesProperty          = "gog-is-included-by-games"
	GogRequiresGamesProperty              = "gog-requires-games"
	GogIsRequiredByGamesProperty          = "gog-is-required-by-games"
	GogModifiesGamesProperty              = "gog-modifies-games"
	GogIsModifiedByGamesProperty          = "gog-is-modified-by-games"
	GogEditionsProperty                   = "gog-editions"
	GogRootEditionsProperty               = "gog-root-editions"
	GogStoreTagsProperty                  = "gog-store-tags"
	GogGenresProperty                     = "gog-genres"
	GogFeaturesProperty                   = "gog-features"
	GogSeriesProperty                     = "gog-series"
	GogThemesProperty                     = "gog-themes"
	GogGameModesProperty                  = "gog-game-modes"
	GogTagIdProperty                      = "gog-tag-id"
	GogTagNameProperty                    = "gog-tag-name"
	GogSlugProperty                       = "gog-slug"
	GogReleaseDateProperty                = "gog-release-date"
	GogOrderDateProperty                  = "gog-order-date"
	GogGlobalReleaseDateProperty          = "gog-global-release-date"
	GogManualUrlFilenameProperty          = "gog-manual-url-filename"
	GogManualUrlStatusProperty            = "gog-manual-url-status"
	GogManualUrlValidationResultProperty  = "gog-manual-url-validation-result"
	GogManualUrlGeneratedChecksumProperty = "gog-manual-url-generated-checksum"
	GogProductValidationResultProperty    = "gog-product-validation-result"
	GogProductGeneratedChecksumProperty   = "gog-product-generated-checksum"
	GogProductValidationDateProperty      = "gog-product-validation-date"
	GogStoreUrlProperty                   = "gog-store-url"
	GogForumUrlProperty                   = "gog-forum-url"
	GogSupportUrlProperty                 = "gog-support-url"
	GogAdditionalRequirementsProperty     = "gog-additional-requirements"
	GogCopyrightsProperty                 = "gog-copyrights"
	GogInDevelopmentProperty              = "gog-in-development"
	GogPreOrderProperty                   = "gog-pre-order"
	GogComingSoonProperty                 = "gog-coming-soon"
	GogBasePriceProperty                  = "gog-base-price"
	GogPriceProperty                      = "gog-price"
	GogIsFreeProperty                     = "gog-is-free"
	GogIsDemoProperty                     = "gog-is-demo"
	GogIsModProperty                      = "gog-is-mod"
	GogIsDiscountedProperty               = "gog-is-discounted"
	GogDiscountPercentageProperty         = "gog-discount-percentage"
	GogSteamAppIdProperty                 = "gog-steam-app-id"
	GogBundleNameProperty                 = "gog-bundle-name"

	VangoghLocalTagsProperty = "vangogh-local-tags"

	// download lifecycle properties

	VangoghDownloadStatusErrorProperty = "vangogh-download-status-error"
	VangoghDownloadQueuedProperty      = "vangogh-download-queued"
	VangoghDownloadStartedProperty     = "vangogh-download-started"
	VangoghDownloadCompletedProperty   = "vangogh-download-completed"

	// video properties

	GogYouTubeVideoIdProperty    = "gog-youtube-video-id"
	YouTubeVideoTitleProperty    = "youtube-video-title"
	YouTubeVideoDurationProperty = "youtube-video-duration"
	YouTubeVideoErrorProperty    = "youtube-video-error"

	// downloads properties

	OperatingSystemsProperty = "os"            // TODO: Replace with GogOperatingSystemsProperty, then add generic property for theo
	LanguageCodeProperty     = "lang-code"     // TODO: Replace with GogLanguageCodeProperty, then add generic property for theo
	DownloadTypeProperty     = "download-type" // TODO: Replace with GogDownloadTypeProperty, then add generic property for theo
	NoPatchesProperty        = "no-patches"    // TODO: Replace with GogNoPatchesProperty, then add generic property for theo

	// Steam properties

	SteamTitleProperty                           = "steam-title"
	SteamReviewScoreProperty                     = "steam-review-score"
	SteamReviewScoreDescProperty                 = "steam-review-score-desc"
	SteamDeckAppCompatibilityCategoryProperty    = "steam-deck-app-compatibility-category"
	SteamSteamOsAppCompatibilityCategoryProperty = "steam-steamos-app-compatibility-category"
	SteamLastCommunityUpdateProperty             = "steam-last-community-update"

	// EGS properties

	EgsTitleProperty    = "egs-title"
	EgsMainGameProperty = "egs-main-game"

	VangoghSummaryRatingProperty  = "vangogh-summary-rating"
	VangoghSummaryReviewsProperty = "vangogh-summary-reviews"

	GogPcgwPageIdProperty = "gog-pcgw-page-id"

	// hltb-data properties

	GogHltbIdProperty               = "gog-hltb-id"
	HltbHoursToCompleteMainProperty = "hltb-comp-main"    // TODO: Use HLTB Id as a key
	HltbHoursToCompletePlusProperty = "hltb-comp-plus"    // TODO: Use HLTB Id as a key
	HltbHoursToComplete100Property  = "hltb-comp-100"     // TODO: Use HLTB Id as a key
	HltbReviewScoreProperty         = "hltb-review-score" // TODO: Use HLTB Id as a key
	HltbGenresProperty              = "hltb-genres"       // TODO: Use HLTB Id as a key
	HltbPlatformsProperty           = "hltb-platforms"    // TODO: Use HLTB Id as a key

	// pcgw-raw properties

	GogIgdbIdProperty         = "gog-igdb-id"
	GogStrategyWikiIdProperty = "gog-strategy-wiki-id"
	GogMobyGamesIdProperty    = "gog-moby-games-id"
	GogWikipediaIdProperty    = "gog-wikipedia-id"
	GogWineHqIdProperty       = "gog-winehq-id"
	GogVndbIdProperty         = "gog-vndb-id"
	GogIgnWikiSlugProperty    = "gog-ign-wiki-slug"
	GogOpenCriticIdProperty   = "gog-opencritic-id"
	GogOpenCriticSlugProperty = "gog-opencritic-slug" // TODO: Use OpenCritic Id and migrate back to OpenCriticSlugProperty

	EnginesProperty       = "engines"        // TODO: Rename to pcgw-engines
	EnginesBuildsProperty = "engines-builds" // TODO: Rename to pcgw-engines-builds

	// wikipedia-raw properties

	CreditsProperty = "credits" // TODO: Rename to wikipedia-credits, use Wikipedia Id as key

	CreatorsProperty    = "creators"    // TODO: Rename to wikipedia-*, use Wikipedia Id as key
	DirectorsProperty   = "directors"   // TODO: Rename to wikipedia-*, use Wikipedia Id as key
	ProducersProperty   = "producers"   // TODO: Rename to wikipedia-*, use Wikipedia Id as key
	DesignersProperty   = "designers"   // TODO: Rename to wikipedia-*, use Wikipedia Id as key
	ProgrammersProperty = "programmers" // TODO: Rename to wikipedia-*, use Wikipedia Id as key
	ArtistsProperty     = "artists"     // TODO: Rename to wikipedia-*, use Wikipedia Id as key
	WritersProperty     = "writers"     // TODO: Rename to wikipedia-*, use Wikipedia Id as key
	ComposersProperty   = "composers"   // TODO: Rename to wikipedia-*, use Wikipedia Id as key

	// proton-summary properties

	ProtonDbTierProperty       = "protondb-tier"
	ProtonDbConfidenceProperty = "protondb-confidence"

	// steam-app-details properties

	RequiredAgeProperty       = "required-age"        // TODO: Rename to Steam*Property, Use Steam AppId as a key
	ControllerSupportProperty = "controller-support"  // TODO: Rename to SteamControllerSupportProperty, Use Steam AppId as key
	ShortDescriptionProperty  = "short-description"   // TODO: Rename to Steam*Property, Use Steam AppId as a key
	WebsiteProperty           = "website"             // TODO: Rename to Steam*Property, Use Steam AppId as a key
	MetacriticScoreProperty   = "metacritic-score"    // TODO: Rename to Steam*Property, Use Steam AppId as a key
	MetacriticIdProperty      = "metacritic-id"       // TODO: Rename to Steam*Property, Use Steam AppId as a key
	SteamCategoriesProperty   = "steam-categories"    // TODO: Rename to Steam*Property, Use Steam AppId as a key
	SteamGenresProperty       = "steam-genres"        // TODO: Rename to Steam*Property, Use Steam AppId as a key
	SteamSupportUrlProperty   = "steam-support-url"   // TODO: Rename to Steam*Property, Use Steam AppId as a key
	SteamSupportEmailProperty = "steam-support-email" // TODO: Rename to Steam*Property, Use Steam AppId as a key

	// opencritic properties

	OpenCriticMedianScoreProperty     = "opencritic-median-score"      // TODO: Use OpenCritic Id as key
	OpenCriticTopCriticsScoreProperty = "opencritic-top-critics-score" // TODO: Use OpenCritic Id as key
	OpenCriticPercentileProperty      = "opencritic-percentile"        // TODO: Use OpenCritic Id as key
	OpenCriticTierProperty            = "opencritic-tier"              // TODO: Use OpenCritic Id as key

	// get-data properties

	GetDataErrorDateProperty    = "get-data-error-date"    // Rename to vangogh-*
	GetDataErrorMessageProperty = "get-data-error-message" // Rename to vangogh-*
	GetDataLastUpdatedProperty  = "get-data-last-updated"  // Rename to vangogh-*

	// reduced properties

	GogOwnedProperty = "gog-owned" // TODO: Reduce is proper data types

	TopPercentProperty = "top-percent" // TODO: Reduce in OpenCritic data

	// data scheme version

	VangoghDataSchemeVersionProperty = "vangogh-data-scheme-version"

	// sync properties

	VangoghSyncEventsProperty      = "vangogh-sync-events"
	VangoghLastSyncUpdatesProperty = "vangogh-last-sync-updates"

	// sort properties

	SortProperty       = "sort"
	DescendingProperty = "desc"
)

const (
	GogDescriptionOverviewKeyValues = "gog-description-overview"
	GogDescriptionFeaturesKeyValues = "gog-description-features"
	GogChangelogKeyValues           = "gog-changelog"
)

const (
	// property values
	TrueValue  = "true"
	FalseValue = "false"
)

var ProductTypeProperties = map[ProductType][]string{
	GogLicences:                  GogLicencesProperties(),
	GogUserWishlist:              GogUserWishlistProperties(),
	GogCatalogPage:               GogCatalogPageProperties(),
	GogOrderPage:                 GogOrderPageProperties(),
	GogAccountPage:               GogAccountPageProperties(),
	GogApiProducts:               GogApiProductProperties(),
	GogDetails:                   GogDetailsProperties(),
	GamesDbGogProducts:           GogGamesDbProperties(),
	SteamAppDetails:              SteamAppDetailsProperties(),
	SteamAppReviews:              SteamAppReviewsProperties(),
	SteamDeckCompatibilityReport: SteamDeckCompatibilityReportProperties(),
	PcgwGogPageId:                PcgwPageIdProperties(),
	PcgwSteamPageId:              PcgwPageIdProperties(),
	PcgwRaw:                      PcgwRawProperties(),
	WikipediaRaw:                 WikipediaRawProperties(),
	HltbData:                     HltbDataProperties(),
	ProtonDbSummary:              ProtonDbSummaryProperties(),
	OpenCriticApiGame:            OpenCriticApiGameProperties(),
}

func GogLicencesProperties() []string {
	return []string{
		GogLicencesProperty,
	}
}

func GogUserWishlistProperties() []string {
	return []string{
		GogUserWishlistProperty,
	}
}

func GogCatalogPageProperties() []string {
	return []string{
		GogTitleProperty,
		GogDevelopersProperty,
		GogPublishersProperty,
		GogImageProperty,
		GogVerticalImageProperty,
		GogScreenshotsProperty,
		GogGenresProperty,
		GogFeaturesProperty,
		GogRatingProperty,
		OperatingSystemsProperty,
		GogSlugProperty,
		GogGlobalReleaseDateProperty,
		GogProductTypeProperty,
		GogStoreTagsProperty,
		GogBasePriceProperty,
		GogPriceProperty,
		GogIsFreeProperty,
		GogIsDiscountedProperty,
		GogDiscountPercentageProperty,
		GogComingSoonProperty,
		GogPreOrderProperty,
		GogInDevelopmentProperty,
		GogIsDemoProperty,
		GogIsModProperty,
		GogEditionsProperty,
		GogRootEditionsProperty,
		GogCatalogProductPageProperty,
		GogUserWishlistProperty,
	}
}

func GogOrderPageProperties() []string {
	return []string{
		GogOrderDateProperty,
		GogOrderPageProductsProperty,
		GogImageProperty,
	}
}

func GogAccountPageProperties() []string {
	return []string{
		GogTagIdProperty,
		GogTagNameProperty,
		GogImageProperty,
		GogSlugProperty,
		GogAccountProductPageProperty,
	}
}

func GogApiProductProperties() []string {
	return []string{
		GogTitleProperty,
		GogDevelopersProperty,
		GogPublishersProperty,
		LanguageCodeProperty,
		GogImageProperty,
		GogVerticalImageProperty,
		GogScreenshotsProperty,
		GogHeroProperty,
		GogLogoProperty,
		GogIconProperty,
		GogIconSquareProperty,
		GogBackgroundProperty,
		GogGenresProperty,
		GogFeaturesProperty,
		GogSeriesProperty,
		GogYouTubeVideoIdProperty,
		OperatingSystemsProperty,
		GogRequiresGamesProperty,
		GogIsRequiredByGamesProperty,
		GogIncludesGamesProperty,
		GogIsIncludedByGamesProperty,
		GogModifiesGamesProperty,
		GogIsModifiedByGamesProperty,
		GogIsModProperty,
		GogGlobalReleaseDateProperty,
		GogReleaseDateProperty,
		GogStoreUrlProperty,
		GogForumUrlProperty,
		GogSupportUrlProperty,
		GogProductTypeProperty,
		GogCopyrightsProperty,
		GogStoreTagsProperty,
		GogAdditionalRequirementsProperty,
		GogInDevelopmentProperty,
		GogPreOrderProperty,
	}
}

func GogApiProductsKeyValues() []string {
	return []string{
		GogDescriptionOverviewKeyValues,
		GogDescriptionFeaturesKeyValues,
	}
}

func GogDetailsProperties() []string {
	return []string{
		GogTitleProperty,
		GogFeaturesProperty,
		GogTagIdProperty,
		GogReleaseDateProperty,
		GogForumUrlProperty,
		OperatingSystemsProperty,
		GogBackgroundProperty,
	}
}

func GogDetailsKeyValues() []string {
	return []string{
		GogChangelogKeyValues,
	}
}

func GogGamesDbProperties() []string {
	return []string{
		GogSteamAppIdProperty,
		GogYouTubeVideoIdProperty,
		GogThemesProperty,
		GogGameModesProperty,
	}
}

func SteamAppDetailsProperties() []string {
	return []string{
		GogSteamAppIdProperty,
		RequiredAgeProperty,
		ControllerSupportProperty,
		ShortDescriptionProperty,
		WebsiteProperty,
		MetacriticScoreProperty,
		MetacriticIdProperty,
		SteamCategoriesProperty,
		SteamGenresProperty,
		SteamSupportUrlProperty,
		SteamSupportEmailProperty,
	}
}

func SteamAppNewsProperties() []string {
	return []string{
		GogSteamAppIdProperty,
		SteamLastCommunityUpdateProperty,
	}
}

func SteamAppReviewsProperties() []string {
	return []string{
		SteamReviewScoreProperty,
		SteamReviewScoreDescProperty,
	}
}

func SteamDeckCompatibilityReportProperties() []string {
	return []string{
		SteamDeckAppCompatibilityCategoryProperty,
		SteamSteamOsAppCompatibilityCategoryProperty,
	}
}

func PcgwPageIdProperties() []string {
	return []string{
		GogPcgwPageIdProperty,
		GogSteamAppIdProperty,
	}
}

func PcgwRawProperties() []string {
	return []string{
		GogSteamAppIdProperty,
		GogHltbIdProperty,
		GogIgdbIdProperty,
		GogStrategyWikiIdProperty,
		GogMobyGamesIdProperty,
		GogWikipediaIdProperty,
		GogWineHqIdProperty,
		WebsiteProperty,
		GogVndbIdProperty,

		MetacriticIdProperty,
		MetacriticScoreProperty,
		GogOpenCriticIdProperty,
		GogOpenCriticSlugProperty,
		OpenCriticMedianScoreProperty,

		EnginesProperty,
		EnginesBuildsProperty,
	}
}

func CreditsProperties() []string {
	return []string{
		CreditsProperty,
		CreatorsProperty,
		DirectorsProperty,
		ProducersProperty,
		DesignersProperty,
		ProgrammersProperty,
		ArtistsProperty,
		WritersProperty,
		ComposersProperty,
	}
}

func WikipediaRawProperties() []string {
	// currently it's just the credits properties
	// however since credit properties are used for reduction,
	// we'll keep them separate and wikipedia-raw
	// might add more properties in the future in
	// addition to credits
	return CreditsProperties()
}

func HltbDataProperties() []string {
	return []string{
		HltbHoursToCompleteMainProperty,
		HltbHoursToCompletePlusProperty,
		HltbHoursToComplete100Property,
		HltbReviewScoreProperty,
		HltbGenresProperty,
		HltbPlatformsProperty,
		GogIgnWikiSlugProperty,
	}
}

func ProtonDbSummaryProperties() []string {
	return []string{
		ProtonDbTierProperty,
		ProtonDbConfidenceProperty,
	}
}

func OpenCriticApiGameProperties() []string {
	return []string{
		OpenCriticMedianScoreProperty,
		OpenCriticTopCriticsScoreProperty,
		OpenCriticPercentileProperty,
		OpenCriticTierProperty,
	}
}

func GetDataProperties() []string {
	return []string{
		GetDataErrorDateProperty,
		GetDataErrorMessageProperty,
		GetDataLastUpdatedProperty,
	}
}

func ReducedProperties() []string {
	return []string{
		GogOwnedProperty,
		TopPercentProperty,
		VangoghSummaryRatingProperty,
		VangoghSummaryReviewsProperty,
	}
}

func VideoProperties() []string {
	return []string{
		GogYouTubeVideoIdProperty,
		YouTubeVideoTitleProperty,
		YouTubeVideoDurationProperty,
		YouTubeVideoErrorProperty,
	}
}

func LocalProperties() []string {
	return []string{
		VangoghLocalTagsProperty,
	}
}

func SyncProperties() []string {
	return []string{
		VangoghLastSyncUpdatesProperty,
		VangoghSyncEventsProperty,
	}
}

func DownloadsLifecycleProperties() []string {
	return []string{
		GogManualUrlFilenameProperty,
		GogManualUrlStatusProperty,
		GogManualUrlValidationResultProperty,
		GogManualUrlGeneratedChecksumProperty,
		GogProductValidationResultProperty,
		GogProductGeneratedChecksumProperty,
		GogProductValidationDateProperty,
		VangoghDownloadQueuedProperty,
		VangoghDownloadStartedProperty,
		VangoghDownloadCompletedProperty,
		VangoghDownloadStatusErrorProperty,
	}
}

func ReduxProperties() []string {
	all := make([]string, 0)

	// product type properties (data for those is coming directly from origin data)

	all = append(all, GogLicencesProperties()...)
	all = append(all, GogUserWishlistProperties()...)
	all = append(all, GogCatalogPageProperties()...)
	all = append(all, GogOrderPageProperties()...)
	all = append(all, GogAccountPageProperties()...)
	all = append(all, GogApiProductProperties()...)
	all = append(all, GogDetailsProperties()...)
	all = append(all, GogGamesDbProperties()...)
	all = append(all, SteamAppDetailsProperties()...)
	all = append(all, SteamAppNewsProperties()...)
	all = append(all, SteamAppReviewsProperties()...)
	all = append(all, SteamDeckCompatibilityReportProperties()...)
	all = append(all, PcgwPageIdProperties()...)
	all = append(all, PcgwRawProperties()...)
	all = append(all, WikipediaRawProperties()...)
	all = append(all, HltbDataProperties()...)
	all = append(all, ProtonDbSummaryProperties()...)
	all = append(all, OpenCriticApiGameProperties()...)

	// other properties

	all = append(all, ReducedProperties()...)
	all = append(all, LocalProperties()...)

	all = append(all, VideoProperties()...)

	all = append(all, SyncProperties()...)
	all = append(all, DownloadsLifecycleProperties()...)
	return all
}

func DataKeyValues() []string {
	all := make([]string, 0)

	all = append(all, GogDetailsKeyValues()...)
	all = append(all, GogApiProductsKeyValues()...)

	return all
}

var imageTypeProperties = map[gog_integration.ImageType]string{
	gog_integration.Image:         GogImageProperty,
	gog_integration.VerticalImage: GogVerticalImageProperty,
	gog_integration.Screenshots:   GogScreenshotsProperty,
	gog_integration.Hero:          GogHeroProperty,
	gog_integration.Logo:          GogLogoProperty,
	gog_integration.Icon:          GogIconProperty,
	gog_integration.IconSquare:    GogIconSquareProperty,
	gog_integration.Background:    GogBackgroundProperty,
}

func PropertyFromImageType(it gog_integration.ImageType) string {
	return imageTypeProperties[it]
}
