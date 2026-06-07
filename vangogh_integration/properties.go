package vangogh_integration

const (
	IdProperty = "id"

	// GOG Properties

	GogLicencesProperty                   = "gog-licences"
	GogUserWishlistProperty               = "gog-user-wishlist"
	GogAccountPageProductsProperty        = "gog-account-page-products" // TODO: rename to GogAccountProductPageProperty
	GogCatalogPageProductsProperty        = "gog-catalog-page-products" // TODO: rename to GogCatalogProductPageProperty
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
	GogTagIdProperty                      = "gog-tag" // TODO: Rename to gog-tag-id
	GogTagNameProperty                    = "gog-tag-name"
	GogSlugProperty                       = "gog-slug"
	GogReleaseDateProperty                = "gog-release-date"
	GogOrderDateProperty                  = "gog-order-date"
	GogGlobalReleaseDateProperty          = "gog-global-release-date"
	GogLocalManualUrlProperty             = "gog-local-manual-url"
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

	LocalTagsProperty = "local-tags" //TODO: Rename to Vangogh*Property

	// download lifecycle properties

	DownloadStatusErrorProperty = "download-status-error" //TODO: Rename to Vangogh*Property
	DownloadQueuedProperty      = "download-queued"       //TODO: Rename to Vangogh*Property
	DownloadStartedProperty     = "download-started"      //TODO: Rename to Vangogh*Property
	DownloadCompletedProperty   = "download-completed"    //TODO: Rename to Vangogh*Property

	// video properties

	VideoIdProperty       = "video-id"       // TODO: Rename to GogYouTubeVideoIdProperty
	VideoTitleProperty    = "video-title"    // TODO: Rename to YouTubeVideoTitleProperty
	VideoDurationProperty = "video-duration" // TODO: Rename to YouTubeVideoDurationProperty
	VideoErrorProperty    = "video-error"    // TODO: Rename to YouTubeVideoErrorProperty

	// downloads properties

	OperatingSystemsProperty = "os"            // TODO: Replace with GogOperatingSystemsProperty, then add generic property for theo
	LanguageCodeProperty     = "lang-code"     // TODO: Replace with GogLanguageCodeProperty, then add generic property for theo
	DownloadTypeProperty     = "download-type" // TODO: Replace with GogDownloadTypeProperty, then add generic property for theo
	NoPatchesProperty        = "no-patches"    // TODO: Replace with GogNoPatchesProperty, then add generic property for theo

	SteamReviewScoreProperty                  = "steam-review-score"
	SteamReviewScoreDescProperty              = "steam-review-score-desc"
	SteamDeckAppCompatibilityCategoryProperty = "steam-deck-app-compatibility-category"
	SteamOsAppCompatibilityCategoryProperty   = "steamos-app-compatibility-category" // TODO: Rename to steam-steamos-app-compatibility-category, no space in SteamOS
	SteamLastCommunityUpdateProperty          = "steam-last-community-update"

	SummaryRatingProperty  = "summary-rating"  //TODO: Rename to Vangogh*Property
	SummaryReviewsProperty = "summary-reviews" //TODO: Rename to Vangogh*Property

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

	CreditsProperty            = "credits"              // TODO: Rename to wikipedia-credits, use Wikipedia Id as key
	HasMultipleCreditsProperty = "has-multiple-credits" // TODO: Deprecate

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

	TypesProperty      = "types"       // TODO: Deprecate
	TopPercentProperty = "top-percent" // TODO: Reduce in OpenCritic data

	// data scheme version

	DataSchemeVersionProperty = "data-scheme-version" // TODO: Rename to vangogh-*

	// sync properties

	SyncEventsProperty      = "sync-events"       // TODO: Rename to vangogh-*
	LastSyncUpdatesProperty = "last-sync-updates" // TODO: Rename to vangogh-*

	// sort properties

	SortProperty       = "sort"
	DescendingProperty = "desc"
)

const (
	DescriptionOverviewKeyValues = "description-overview"
	DescriptionFeaturesKeyValues = "description-features"
	ChangelogKeyValues           = "changelog"
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
		GogCatalogPageProductsProperty,
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
		GogAccountPageProductsProperty,
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
		VideoIdProperty,
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
		DescriptionOverviewKeyValues,
		DescriptionFeaturesKeyValues,
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
		ChangelogKeyValues,
	}
}

func GogGamesDbProperties() []string {
	return []string{
		GogSteamAppIdProperty,
		VideoIdProperty,
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
		SteamOsAppCompatibilityCategoryProperty,
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
		HasMultipleCreditsProperty,
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
		TypesProperty,
		TopPercentProperty,
		SummaryRatingProperty,
		SummaryReviewsProperty,
	}
}

func VideoProperties() []string {
	return []string{
		VideoIdProperty,
		VideoTitleProperty,
		VideoDurationProperty,
		VideoErrorProperty,
	}
}

func LocalProperties() []string {
	return []string{
		LocalTagsProperty,
	}
}

func SyncProperties() []string {
	return []string{
		LastSyncUpdatesProperty,
		SyncEventsProperty,
	}
}

func DownloadsLifecycleProperties() []string {
	return []string{
		GogLocalManualUrlProperty,
		GogManualUrlFilenameProperty,
		GogManualUrlStatusProperty,
		GogManualUrlValidationResultProperty,
		GogManualUrlGeneratedChecksumProperty,
		GogProductValidationResultProperty,
		GogProductGeneratedChecksumProperty,
		GogProductValidationDateProperty,
		DownloadQueuedProperty,
		DownloadStartedProperty,
		DownloadCompletedProperty,
		DownloadStatusErrorProperty,
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

var imageTypeProperties = map[ImageType]string{
	Image:         GogImageProperty,
	VerticalImage: GogVerticalImageProperty,
	Screenshots:   GogScreenshotsProperty,
	Hero:          GogHeroProperty,
	Logo:          GogLogoProperty,
	Icon:          GogIconProperty,
	IconSquare:    GogIconSquareProperty,
	Background:    GogBackgroundProperty,
}

func PropertyFromImageType(it ImageType) string {
	return imageTypeProperties[it]
}
