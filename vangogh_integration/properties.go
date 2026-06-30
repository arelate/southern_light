package vangogh_integration

import "github.com/arelate/southern_light/gog_integration"

const (
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
	SteamRequiredAgeProperty                     = "steam-required-age"
	SteamControllerSupportProperty               = "steam-controller-support"
	SteamShortDescriptionProperty                = "steam-short-description"
	SteamWebsiteProperty                         = "steam-website"
	SteamMetacriticIdProperty                    = "steam-metacritic-id"
	SteamCategoriesProperty                      = "steam-categories"
	SteamGenresProperty                          = "steam-genres"
	SteamSupportUrlProperty                      = "steam-support-url"
	SteamSupportEmailProperty                    = "steam-support-email"

	// EGS properties

	EgsTitleProperty    = "egs-title"
	EgsMainGameProperty = "egs-main-game"

	// hltb-data properties

	GogHltbIdProperty               = "gog-hltb-id"
	HltbHoursToCompleteMainProperty = "hltb-comp-main"
	HltbHoursToCompletePlusProperty = "hltb-comp-plus"
	HltbHoursToComplete100Property  = "hltb-comp-100"
	HltbReviewScoreProperty         = "hltb-review-score"
	HltbGenresProperty              = "hltb-genres"
	HltbPlatformsProperty           = "hltb-platforms"

	// pcgw-raw properties

	GogPcgwPageIdProperty     = "gog-pcgw-page-id"
	GogIgdbIdProperty         = "gog-igdb-id"
	GogStrategyWikiIdProperty = "gog-strategy-wiki-id"
	GogMobyGamesIdProperty    = "gog-moby-games-id"
	GogWikipediaIdProperty    = "gog-wikipedia-id"
	GogWineHqIdProperty       = "gog-winehq-id"
	GogVndbIdProperty         = "gog-vndb-id"
	GogOpenCriticIdProperty   = "gog-opencritic-id"
	GogMetacriticIdProperty   = "gog-metacritic-id"

	PcgwEnginesProperty       = "pcgw-engines"
	PcgwEnginesBuildsProperty = "pcgw-engines-builds"
	PcgwWebsiteProperty       = "pcgw-website"

	// wikipedia-raw properties

	WikipediaCreatorsProperty    = "wikipedia-creators"
	WikipediaDirectorsProperty   = "wikipedia-directors"
	WikipediaProducersProperty   = "wikipedia-producers"
	WikipediaDesignersProperty   = "wikipedia-designers"
	WikipediaProgrammersProperty = "wikipedia-programmers"
	WikipediaArtistsProperty     = "wikipedia-artists"
	WikipediaWritersProperty     = "wikipedia-writers"
	WikipediaComposersProperty   = "wikipedia-composers"

	// proton-summary properties

	ProtonDbTierProperty       = "protondb-tier"
	ProtonDbConfidenceProperty = "protondb-confidence"

	// metacritic properties

	MetacriticScoreProperty = "metacritic-score"

	// opencritic properties

	OpenCriticMedianScoreProperty     = "opencritic-median-score"
	OpenCriticTopCriticsScoreProperty = "opencritic-top-critics-score"
	OpenCriticPercentileProperty      = "opencritic-percentile"
	OpenCriticTierProperty            = "opencritic-tier"
	OpenCriticSlugProperty            = "opencritic-slug"

	// get-data properties

	GetDataErrorDateProperty    = "get-data-error-date"    // Rename to vangogh-*
	GetDataErrorMessageProperty = "get-data-error-message" // Rename to vangogh-*
	GetDataLastUpdatedProperty  = "get-data-last-updated"  // Rename to vangogh-*

	// reduced properties

	GogOwnedProperty = "gog-owned" // TODO: Reduce using providing data types

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
		SteamRequiredAgeProperty,
		SteamControllerSupportProperty,
		SteamShortDescriptionProperty,
		SteamWebsiteProperty,
		//MetacriticScoreProperty,
		//MetacriticIdProperty,
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
		SteamWebsiteProperty,
		GogVndbIdProperty,

		GogMetacriticIdProperty,
		MetacriticScoreProperty,
		GogOpenCriticIdProperty,
		OpenCriticSlugProperty,
		OpenCriticMedianScoreProperty,

		PcgwEnginesProperty,
		PcgwEnginesBuildsProperty,
		PcgwWebsiteProperty,
	}
}

func WikipediaCreditsProperties() []string {
	return []string{
		WikipediaCreatorsProperty,
		WikipediaDirectorsProperty,
		WikipediaProducersProperty,
		WikipediaDesignersProperty,
		WikipediaProgrammersProperty,
		WikipediaArtistsProperty,
		WikipediaWritersProperty,
		WikipediaComposersProperty,
	}
}

func WikipediaRawProperties() []string {
	// currently it's just the credits properties
	// however since credit properties are used for reduction,
	// we'll keep them separate and wikipedia-raw
	// might add more properties in the future in
	// addition to credits
	return WikipediaCreditsProperties()
}

func HltbDataProperties() []string {
	return []string{
		HltbHoursToCompleteMainProperty,
		HltbHoursToCompletePlusProperty,
		HltbHoursToComplete100Property,
		HltbReviewScoreProperty,
		HltbGenresProperty,
		HltbPlatformsProperty,
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
