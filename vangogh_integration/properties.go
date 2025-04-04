package vangogh_integration

const (
	IdProperty = "id"

	LicencesProperty     = "licences"
	UserWishlistProperty = "user-wishlist"

	AccountPageProductsProperty = "account-page-products"
	CatalogPageProductsProperty = "catalog-page-products"
	OrderPageProductsProperty   = "order-page-products"

	TitleProperty                      = "title"
	DevelopersProperty                 = "developers"
	PublishersProperty                 = "publishers"
	ImageProperty                      = "image"
	VerticalImageProperty              = "vertical-image"
	ScreenshotsProperty                = "screenshots"
	HeroProperty                       = "hero"
	LogoProperty                       = "logo"
	IconProperty                       = "icon"
	IconSquareProperty                 = "icon-square"
	BackgroundProperty                 = "background"
	RatingProperty                     = "rating"
	AggregatedRatingProperty           = "aggregated-rating"
	ProductTypeProperty                = "product-type"
	IncludesGamesProperty              = "includes-games"
	IsIncludedByGamesProperty          = "is-included-by-games"
	RequiresGamesProperty              = "requires-games"
	IsRequiredByGamesProperty          = "is-required-by-games"
	EditionsProperty                   = "editions"
	GenresProperty                     = "genres"
	StoreTagsProperty                  = "store-tags"
	FeaturesProperty                   = "features"
	SeriesProperty                     = "series"
	ThemesProperty                     = "themes"
	GameModesProperty                  = "game-modes"
	TagIdProperty                      = "tag"
	TagNameProperty                    = "tag-name"
	VideoIdProperty                    = "video-id"
	VideoTitleProperty                 = "video-title"
	VideoDurationProperty              = "video-duration"
	VideoErrorProperty                 = "video-error"
	OperatingSystemsProperty           = "os"
	LanguageCodeProperty               = "lang-code"
	DownloadTypeProperty               = "download-type"
	NoPatchesProperty                  = "no-patches"
	SlugProperty                       = "slug"
	GOGReleaseDateProperty             = "gog-release-date"
	GOGOrderDateProperty               = "gog-order-date"
	GlobalReleaseDateProperty          = "global-release-date"
	LocalManualUrlProperty             = "local-manual-url"
	ManualUrlStatusProperty            = "manual-url-status"
	ManualUrlValidationResultProperty  = "manual-url-validation-result"
	ProductValidationResultProperty    = "product-validation-result"
	ManualUrlGeneratedChecksumProperty = "manual-url-generated-checksum"
	DownloadStatusErrorProperty        = "download-status-error"
	StoreUrlProperty                   = "store-url"
	ForumUrlProperty                   = "forum-url"
	SupportUrlProperty                 = "support-url"
	ChangelogProperty                  = "changelog"
	DescriptionOverviewProperty        = "description-overview"
	DescriptionFeaturesProperty        = "description-features"
	AdditionalRequirementsProperty     = "additional-requirements"
	CopyrightsProperty                 = "copyrights"
	InDevelopmentProperty              = "in-development"
	PreOrderProperty                   = "pre-order"
	ComingSoonProperty                 = "coming-soon"
	BasePriceProperty                  = "base-price"
	PriceProperty                      = "price"
	IsFreeProperty                     = "is-free"
	IsDemoProperty                     = "is-demo"
	IsDiscountedProperty               = "is-discounted"
	DiscountPercentageProperty         = "discount-percentage"
	SteamAppIdProperty                 = "steam-app-id"
	LocalTagsProperty                  = "local-tags"

	SteamReviewScoreProperty                  = "steam-review-score"
	SteamReviewScoreDescProperty              = "steam-review-score-desc"
	SteamDeckAppCompatibilityCategoryProperty = "steam-deck-app-compatibility-category"

	SummaryRatingProperty  = "summary-rating"
	SummaryReviewsProperty = "summary-reviews"

	PcgwPageIdProperty = "pcgw-page-id"

	HltbIdProperty                  = "hltb-id"
	HltbHoursToCompleteMainProperty = "hltb-comp-main"
	HltbHoursToCompletePlusProperty = "hltb-comp-plus"
	HltbHoursToComplete100Property  = "hltb-comp-100"
	HltbReviewScoreProperty         = "hltb-review-score"
	HltbGenresProperty              = "hltb-genres"
	HltbPlatformsProperty           = "hltb-platforms"

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

	ProtonDBTierProperty       = "protondb-tier"
	ProtonDBConfidenceProperty = "protondb-confidence"

	// steam-app-details properties

	RequiredAgeProperty       = "required-age"
	ControllerSupportProperty = "controller-support"
	ShortDescriptionProperty  = "short-description"
	WebsiteProperty           = "website"
	MetacriticScoreProperty   = "metacritic-score"
	MetacriticUrlProperty     = "metacritic-url"
	SteamCategoriesProperty   = "steam-categories"
	SteamGenresProperty       = "steam-genres"
	SteamSupportUrlProperty   = "steam-support-url"
	SteamSupportEmailProperty = "steam-support-email"

	// opencritic properties

	OpenCriticMedianScoreProperty     = "opencritic-median-score"
	OpenCriticTopCriticsScoreProperty = "opencritic-top-critics-score"
	OpenCriticPercentileProperty      = "opencritic-percentile"
	OpenCriticTierProperty            = "opencritic-tier"

	// per-type properties

	TypeErrorMessageProperty = "type-error-message"
	TypeErrorDateProperty    = "type-error-date"

	// reduced properties

	OwnedProperty      = "owned"
	TypesProperty      = "types"
	TopPercentProperty = "top-percent"

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
	PcgwExternalLinks:            PcgwExternalLinksProperties(),
	PcgwEngine:                   PcgwEngineProperties(),
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
		MetacriticUrlProperty,
		SteamCategoriesProperty,
		SteamGenresProperty,
		SteamSupportUrlProperty,
		SteamSupportEmailProperty,
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

func PcgwExternalLinksProperties() []string {
	return []string{
		HltbIdProperty,
		IgdbIdProperty,
		MobyGamesIdProperty,
		VndbIdProperty,
		WikipediaIdProperty,
		StrategyWikiIdProperty,
		OpenCriticIdProperty,
		OpenCriticSlugProperty,
	}
}

func PcgwEngineProperties() []string {
	return []string{
		EnginesProperty,
		EnginesBuildsProperty,
	}
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

func TypeErrorProperties() []string {
	return []string{
		TypeErrorMessageProperty,
		TypeErrorDateProperty,
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

func ValidationResultsProperties() []string {
	return []string{
		LocalManualUrlProperty,
		ManualUrlStatusProperty,
		ManualUrlValidationResultProperty,
		ManualUrlGeneratedChecksumProperty,
		ProductValidationResultProperty,
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
	all = append(all, SteamAppReviewsProperties()...)
	all = append(all, SteamDeckCompatibilityReportProperties()...)
	all = append(all, PcgwPageIdProperties()...)
	all = append(all, PcgwExternalLinksProperties()...)
	all = append(all, PcgwEngineProperties()...)
	all = append(all, HltbDataProperties()...)
	all = append(all, ProtonDbSummaryProperties()...)
	all = append(all, OpenCriticApiGameProperties()...)

	// other properties

	all = append(all, ReducedProperties()...)
	all = append(all, LocalProperties()...)

	all = append(all, VideoProperties()...)
	all = append(all, DehydratedImagesProperties()...)

	all = append(all, SyncProperties()...)
	all = append(all, ValidationResultsProperties()...)
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
