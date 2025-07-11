package vangogh_integration

const (
	IdProperty = "id"

	LicencesProperty     = "licences"
	UserWishlistProperty = "user-wishlist"

	AccountPageProductsProperty = "account-page-products"
	CatalogPageProductsProperty = "catalog-page-products"
	OrderPageProductsProperty   = "order-page-products"

	TitleProperty                     = "title"
	DevelopersProperty                = "developers"
	PublishersProperty                = "publishers"
	ImageProperty                     = "image"
	VerticalImageProperty             = "vertical-image"
	ScreenshotsProperty               = "screenshots"
	HeroProperty                      = "hero"
	LogoProperty                      = "logo"
	IconProperty                      = "icon"
	IconSquareProperty                = "icon-square"
	BackgroundProperty                = "background"
	RatingProperty                    = "rating"
	AggregatedRatingProperty          = "aggregated-rating"
	ProductTypeProperty               = "product-type"
	IncludesGamesProperty             = "includes-games"
	IsIncludedByGamesProperty         = "is-included-by-games"
	RequiresGamesProperty             = "requires-games"
	IsRequiredByGamesProperty         = "is-required-by-games"
	ModifiesGamesProperty             = "modifies-games"
	IsModifiedByGamesProperty         = "is-modified-by-games"
	EditionsProperty                  = "editions"
	GenresProperty                    = "genres"
	StoreTagsProperty                 = "store-tags"
	FeaturesProperty                  = "features"
	SeriesProperty                    = "series"
	ThemesProperty                    = "themes"
	GameModesProperty                 = "game-modes"
	TagIdProperty                     = "tag"
	TagNameProperty                   = "tag-name"
	VideoIdProperty                   = "video-id"
	VideoTitleProperty                = "video-title"
	VideoDurationProperty             = "video-duration"
	VideoErrorProperty                = "video-error"
	OperatingSystemsProperty          = "os"
	LanguageCodeProperty              = "lang-code"
	DownloadTypeProperty              = "download-type"
	NoPatchesProperty                 = "no-patches"
	SlugProperty                      = "slug"
	GOGReleaseDateProperty            = "gog-release-date"
	GOGOrderDateProperty              = "gog-order-date"
	GlobalReleaseDateProperty         = "global-release-date"
	LocalManualUrlProperty            = "local-manual-url"
	ManualUrlFilenameProperty         = "manual-url-filename"
	ManualUrlStatusProperty           = "manual-url-status"
	ManualUrlValidationResultProperty = "manual-url-validation-result"
	ProductValidationResultProperty   = "product-validation-result"
	DownloadStatusErrorProperty       = "download-status-error"
	DownloadQueuedProperty            = "download-queued"
	DownloadStartedProperty           = "download-started"
	DownloadCompletedProperty         = "download-completed"
	StoreUrlProperty                  = "store-url"
	ForumUrlProperty                  = "forum-url"
	SupportUrlProperty                = "support-url"
	ChangelogProperty                 = "changelog"
	DescriptionOverviewProperty       = "description-overview"
	DescriptionFeaturesProperty       = "description-features"
	AdditionalRequirementsProperty    = "additional-requirements"
	CopyrightsProperty                = "copyrights"
	InDevelopmentProperty             = "in-development"
	PreOrderProperty                  = "pre-order"
	ComingSoonProperty                = "coming-soon"
	BasePriceProperty                 = "base-price"
	PriceProperty                     = "price"
	IsFreeProperty                    = "is-free"
	IsDemoProperty                    = "is-demo"
	IsModProperty                     = "is-mod"
	IsDiscountedProperty              = "is-discounted"
	DiscountPercentageProperty        = "discount-percentage"
	SteamAppIdProperty                = "steam-app-id"
	LocalTagsProperty                 = "local-tags"

	SteamLastCommunityUpdateProperty = "steam-last-community-update"

	SteamReviewScoreProperty                  = "steam-review-score"
	SteamReviewScoreDescProperty              = "steam-review-score-desc"
	SteamDeckAppCompatibilityCategoryProperty = "steam-deck-app-compatibility-category"
	SteamOsAppCompatibilityCategoryProperty   = "steamos-app-compatibility-category" // no space in SteamOS

	SummaryRatingProperty  = "summary-rating"
	SummaryReviewsProperty = "summary-reviews"

	PcgwPageIdProperty = "pcgw-page-id"

	// hltb-data properties

	HltbIdProperty                  = "hltb-id"
	HltbHoursToCompleteMainProperty = "hltb-comp-main"
	HltbHoursToCompletePlusProperty = "hltb-comp-plus"
	HltbHoursToComplete100Property  = "hltb-comp-100"
	HltbReviewScoreProperty         = "hltb-review-score"
	HltbGenresProperty              = "hltb-genres"
	HltbPlatformsProperty           = "hltb-platforms"

	// pcgw-raw properties

	IgdbIdProperty         = "igdb-id"
	StrategyWikiIdProperty = "strategy-wiki-id"
	MobyGamesIdProperty    = "moby-games-id"
	WikipediaIdProperty    = "wikipedia-id"
	WineHQIdProperty       = "winehq-id"
	VndbIdProperty         = "vndb-id"
	IGNWikiSlugProperty    = "ign-wiki-slug"
	OpenCriticIdProperty   = "opencritic-id"
	OpenCriticSlugProperty = "opencritic-slug"

	EnginesProperty       = "engines"
	EnginesBuildsProperty = "engines-builds"

	// wikipedia-raw properties

	CreditsProperty            = "credits"
	HasMultipleCreditsProperty = "has-multiple-credits"

	CreatorsProperty    = "creators"
	DirectorsProperty   = "directors"
	ProducersProperty   = "producers"
	DesignersProperty   = "designers"
	ProgrammersProperty = "programmers"
	ArtistsProperty     = "artists"
	WritersProperty     = "writers"
	ComposersProperty   = "composers"

	// proton-summary properties

	ProtonDBTierProperty       = "protondb-tier"
	ProtonDBConfidenceProperty = "protondb-confidence"

	// steam-app-details properties

	RequiredAgeProperty       = "required-age"
	ControllerSupportProperty = "controller-support"
	ShortDescriptionProperty  = "short-description"
	WebsiteProperty           = "website"
	MetacriticScoreProperty   = "metacritic-score"
	MetacriticIdProperty      = "metacritic-id"
	SteamCategoriesProperty   = "steam-categories"
	SteamGenresProperty       = "steam-genres"
	SteamSupportUrlProperty   = "steam-support-url"
	SteamSupportEmailProperty = "steam-support-email"

	// opencritic properties

	OpenCriticMedianScoreProperty     = "opencritic-median-score"
	OpenCriticTopCriticsScoreProperty = "opencritic-top-critics-score"
	OpenCriticPercentileProperty      = "opencritic-percentile"
	OpenCriticTierProperty            = "opencritic-tier"

	// get-data properties

	GetDataErrorDateProperty    = "get-data-error-date"
	GetDataErrorMessageProperty = "get-data-error-message"
	GetDataLastUpdatedProperty  = "get-data-last-updated"

	// reduced properties

	OwnedProperty      = "owned"
	TypesProperty      = "types"
	TopPercentProperty = "top-percent"

	// data scheme version

	DataSchemeVersionProperty = "data-scheme-version"

	// sync properties

	SyncEventsProperty      = "sync-events"
	LastSyncUpdatesProperty = "last-sync-updates"

	// dehydrated images properties

	DehydratedImageProperty = "dehydrated-image"
	RepColorProperty        = "rep-color"

	// GitHub releases properties

	GitHubReleasesUpdatedProperty = "github-releases-updated"

	// sort properties

	SortProperty       = "sort"
	DescendingProperty = "desc"
)

const (
	// property values
	TrueValue  = "true"
	FalseValue = "false"
)

var ProductTypeProperties = map[ProductType][]string{
	Licences:                     GOGLicencesProperties(),
	UserWishlist:                 GOGUserWishlistProperties(),
	CatalogPage:                  GOGCatalogPageProperties(),
	OrderPage:                    GOGOrderPageProperties(),
	AccountPage:                  GOGAccountPageProperties(),
	ApiProducts:                  GOGApiProductProperties(),
	Details:                      GOGDetailsProperties(),
	GamesDbGogProducts:           GOGGamesDbProperties(),
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

func GOGLicencesProperties() []string {
	return []string{
		LicencesProperty,
	}
}

func GOGUserWishlistProperties() []string {
	return []string{
		UserWishlistProperty,
	}
}

func GOGCatalogPageProperties() []string {
	return []string{
		TitleProperty,
		DevelopersProperty,
		PublishersProperty,
		ImageProperty,
		VerticalImageProperty,
		ScreenshotsProperty,
		GenresProperty,
		FeaturesProperty,
		RatingProperty,
		OperatingSystemsProperty,
		SlugProperty,
		GlobalReleaseDateProperty,
		ProductTypeProperty,
		StoreTagsProperty,
		BasePriceProperty,
		PriceProperty,
		IsFreeProperty,
		IsDiscountedProperty,
		DiscountPercentageProperty,
		ComingSoonProperty,
		PreOrderProperty,
		InDevelopmentProperty,
		IsDemoProperty,
		IsModProperty,
		EditionsProperty,
		CatalogPageProductsProperty,
		UserWishlistProperty,
	}
}

func GOGOrderPageProperties() []string {
	return []string{
		GOGOrderDateProperty,
		OrderPageProductsProperty,
	}
}

func GOGAccountPageProperties() []string {
	return []string{
		TagIdProperty,
		TagNameProperty,
		SlugProperty,
		AccountPageProductsProperty,
	}
}

func GOGApiProductProperties() []string {
	return []string{
		TitleProperty,
		DevelopersProperty,
		PublishersProperty,
		LanguageCodeProperty,
		ImageProperty,
		VerticalImageProperty,
		HeroProperty,
		LogoProperty,
		IconProperty,
		IconSquareProperty,
		BackgroundProperty,
		ScreenshotsProperty,
		GenresProperty,
		FeaturesProperty,
		SeriesProperty,
		VideoIdProperty,
		OperatingSystemsProperty,
		RequiresGamesProperty,
		IsRequiredByGamesProperty,
		IncludesGamesProperty,
		IsIncludedByGamesProperty,
		ModifiesGamesProperty,
		IsModifiedByGamesProperty,
		IsModProperty,
		GlobalReleaseDateProperty,
		GOGReleaseDateProperty,
		StoreUrlProperty,
		ForumUrlProperty,
		SupportUrlProperty,
		DescriptionOverviewProperty,
		DescriptionFeaturesProperty,
		ProductTypeProperty,
		CopyrightsProperty,
		StoreTagsProperty,
		AdditionalRequirementsProperty,
		InDevelopmentProperty,
		PreOrderProperty,
	}
}

func GOGDetailsProperties() []string {
	return []string{
		TitleProperty,
		FeaturesProperty,
		TagIdProperty,
		GOGReleaseDateProperty,
		ForumUrlProperty,
		ChangelogProperty,
	}
}

func GOGGamesDbProperties() []string {
	return []string{
		SteamAppIdProperty,
		VideoIdProperty,
		AggregatedRatingProperty,
		ThemesProperty,
		GameModesProperty,
	}
}

func SteamAppDetailsProperties() []string {
	return []string{
		SteamAppIdProperty,
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
		SteamAppIdProperty,
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
	}
}

func PcgwPageIdProperties() []string {
	return []string{
		PcgwPageIdProperty,
		SteamAppIdProperty,
	}
}

func PcgwRawProperties() []string {
	return []string{
		SteamAppIdProperty,
		HltbIdProperty,
		IgdbIdProperty,
		StrategyWikiIdProperty,
		MobyGamesIdProperty,
		WikipediaIdProperty,
		WineHQIdProperty,
		WebsiteProperty,
		VndbIdProperty,

		MetacriticIdProperty,
		MetacriticScoreProperty,
		OpenCriticIdProperty,
		OpenCriticSlugProperty,
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
		IGNWikiSlugProperty,
	}
}

func ProtonDbSummaryProperties() []string {
	return []string{
		ProtonDBTierProperty,
		ProtonDBConfidenceProperty,
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
		OwnedProperty,
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

func DehydratedImagesProperties() []string {
	return []string{
		DehydratedImageProperty,
		RepColorProperty,
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
		LocalManualUrlProperty,
		ManualUrlFilenameProperty,
		ManualUrlStatusProperty,
		ManualUrlValidationResultProperty,
		ProductValidationResultProperty,
		DownloadQueuedProperty,
		DownloadStartedProperty,
		DownloadCompletedProperty,
		DownloadStatusErrorProperty,
	}
}

func ReduxProperties() []string {
	all := make([]string, 0)

	// product type properties (data for those is coming directly from origin data)

	all = append(all, GOGLicencesProperties()...)
	all = append(all, GOGUserWishlistProperties()...)
	all = append(all, GOGCatalogPageProperties()...)
	all = append(all, GOGOrderPageProperties()...)
	all = append(all, GOGAccountPageProperties()...)
	all = append(all, GOGApiProductProperties()...)
	all = append(all, GOGDetailsProperties()...)
	all = append(all, GOGGamesDbProperties()...)
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
	all = append(all, DehydratedImagesProperties()...)

	all = append(all, SyncProperties()...)
	all = append(all, DownloadsLifecycleProperties()...)
	return all
}

var imageTypeProperties = map[ImageType]string{
	Image:         ImageProperty,
	Screenshots:   ScreenshotsProperty,
	VerticalImage: VerticalImageProperty,
	Hero:          HeroProperty,
	Logo:          LogoProperty,
	Icon:          IconProperty,
	IconSquare:    IconSquareProperty,
	Background:    BackgroundProperty,
}

func PropertyFromImageType(it ImageType) string {
	return imageTypeProperties[it]
}
