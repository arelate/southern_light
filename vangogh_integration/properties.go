package vangogh_integration

import (
	"github.com/arelate/southern_light/gog_integration"
	"slices"
	"strconv"
	"time"
)

const (
	IdProperty                = "id"
	TitleProperty             = "title"
	DevelopersProperty        = "developers"
	PublishersProperty        = "publishers"
	ImageProperty             = "image"
	VerticalImageProperty     = "vertical-image"
	ScreenshotsProperty       = "screenshots"
	HeroProperty              = "hero"
	LogoProperty              = "logo"
	IconProperty              = "icon"
	IconSquareProperty        = "icon-square"
	BackgroundProperty        = "background"
	RatingProperty            = "rating"
	AggregatedRatingProperty  = "aggregated-rating"
	IncludesGamesProperty     = "includes-games"
	IsIncludedByGamesProperty = "is-included-by-games"
	RequiresGamesProperty     = "requires-games"
	IsRequiredByGamesProperty = "is-required-by-games"
	GenresProperty            = "genres"
	StoreTagsProperty         = "store-tags"
	FeaturesProperty          = "features"
	SeriesProperty            = "series"
	ThemesProperty            = "themes"
	GameModesProperty         = "game-modes"
	TagIdProperty             = "tag"
	TagNameProperty           = "tag-name"
	VideoIdProperty           = "video-id"
	VideoTitleProperty        = "video-title"
	VideoDurationProperty     = "video-duration"
	VideoErrorProperty        = "video-error"
	OperatingSystemsProperty  = "os"
	LanguageCodeProperty      = "lang-code"
	DownloadTypeProperty      = "download-type"
	NoPatchesProperty         = "no-patches"

	//LanguageNameProperty               = "lang-name"
	//NativeLanguageNameProperty         = "native-lang-name"

	SlugProperty                       = "slug"
	GOGReleaseDateProperty             = "gog-release-date"
	GOGOrderDateProperty               = "gog-order-date"
	GlobalReleaseDateProperty          = "global-release-date"
	TypesProperty                      = "types"
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

	//WishlistedProperty                        = "wishlisted"
	//OwnedProperty                             = "owned"

	ProductTypeProperty                       = "product-type"
	InDevelopmentProperty                     = "in-development"
	PreOrderProperty                          = "pre-order"
	ComingSoonProperty                        = "coming-soon"
	BasePriceProperty                         = "base-price"
	PriceProperty                             = "price"
	IsFreeProperty                            = "is-free"
	IsDiscountedProperty                      = "is-discounted"
	DiscountPercentageProperty                = "discount-percentage"
	SteamAppIdProperty                        = "steam-app-id"
	LocalTagsProperty                         = "local-tags"
	SortProperty                              = "sort"
	DescendingProperty                        = "desc"
	SteamReviewScoreDescProperty              = "steam-review-score-desc"
	SteamDeckAppCompatibilityCategoryProperty = "steam-deck-app-compatibility-category"
	DehydratedImageProperty                   = "dehydrated-image"
	DehydratedImageModifiedProperty           = "dehydrated-image-modified"
	DehydratedVerticalImageProperty           = "dehydrated-vertical-image"
	DehydratedVerticalImageModifiedProperty   = "dehydrated-vertical-image-modified"
	RepImageColorProperty                     = "rep-image-color"
	RepVerticalImageColorProperty             = "rep-vertical-image-color"
	SyncEventsProperty                        = "sync-events"
	LastSyncUpdatesProperty                   = "last-sync-updates"
	PcgwPageIdProperty                        = "pcgw-page-id"

	//HltbBuildIdProperty                       = "hltb-next-build"

	HltbIdProperty                  = "hltb-id"
	HltbHoursToCompleteMainProperty = "hltb-comp-main"
	HltbHoursToCompletePlusProperty = "hltb-comp-plus"
	HltbHoursToComplete100Property  = "hltb-comp-100"
	HltbReviewScoreProperty         = "hltb-review-score"
	HltbGenresProperty              = "hltb-genres"
	HltbPlatformsProperty           = "hltb-platforms"
	IgdbIdProperty                  = "igdb-id"
	StrategyWikiIdProperty          = "strategy-wiki-id"
	MobyGamesIdProperty             = "moby-games-id"
	WikipediaIdProperty             = "wikipedia-id"
	WineHQIdProperty                = "winehq-id"
	VndbIdProperty                  = "vndb-id"
	IGNWikiSlugProperty             = "ign-wiki-slug"
	EnginesProperty                 = "engines"
	EnginesBuildsProperty           = "engines-builds"
	ProtonDBTierProperty            = "protondb-tier"
	ProtonDBConfidenceProperty      = "protondb-confidence"

	// new get-data redux properties
	LicencesProperty     = "licences"
	UserWishlistProperty = "user-wishlist"

	IsDemoProperty = "is-demo"

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

	// per-type properties

	TypeErrorMessageProperty = "type-error-message"
	TypeErrorDateProperty    = "type-error-date"
)

const (
	// property values
	TrueValue  = "true"
	FalseValue = "false"
)

func AllProperties() []string {
	all := []string{IdProperty}
	return append(all, ReduxProperties()...)
}

func IsValidProperty(property string) bool {
	for _, p := range AllProperties() {
		if p == property {
			return true
		}
	}
	return false
}

func TextProperties() []string {
	return []string{
		TitleProperty,
		DevelopersProperty,
		PublishersProperty,
		DescriptionOverviewProperty,
		DescriptionFeaturesProperty,
	}
}

func UrlProperties() []string {
	return []string{
		StoreUrlProperty,
		ForumUrlProperty,
		SupportUrlProperty,
	}
}

func LongTextProperties() []string {
	return []string{
		DescriptionOverviewProperty,
		DescriptionFeaturesProperty,
		ChangelogProperty,
		CopyrightsProperty,
		AdditionalRequirementsProperty,
	}
}

func AllTextProperties() []string {
	return append(TextProperties(),
		IncludesGamesProperty,
		IsIncludedByGamesProperty,
		RequiresGamesProperty,
		IsRequiredByGamesProperty,
		GenresProperty,
		HltbGenresProperty,
		StoreTagsProperty,
		FeaturesProperty,
		SeriesProperty,
		RatingProperty,
		AggregatedRatingProperty,
		TagIdProperty,
		TagNameProperty,
		LocalTagsProperty,
		OperatingSystemsProperty,
		LanguageCodeProperty,
		SlugProperty,
		GlobalReleaseDateProperty,
		GOGOrderDateProperty,
		GOGReleaseDateProperty,
		CopyrightsProperty,
		ThemesProperty,
		GameModesProperty,
	)
}

func VideoProperties() []string {
	return []string{
		VideoIdProperty,
		VideoTitleProperty,
		VideoDurationProperty,
		VideoErrorProperty,
	}
}

func ComputedProperties() []string {
	return []string{
		TypesProperty,
	}
}

func ImageIdProperties() []string {
	return []string{
		ImageProperty,
		ScreenshotsProperty,
		VerticalImageProperty,
		HeroProperty,
		LogoProperty,
		IconProperty,
		IconSquareProperty,
		BackgroundProperty,
	}
}

func AvailabilityProperties() []string {
	return []string{
		InDevelopmentProperty,
		PreOrderProperty,
		ComingSoonProperty,
	}
}

//func AccountStatusProperties() []string {
//	return []string{
//		WishlistedProperty,
//		OwnedProperty,
//	}
//}

func AdvancedProductProperties() []string {
	return []string{
		ProductTypeProperty,
		HltbHoursToCompleteMainProperty,
		HltbHoursToCompletePlusProperty,
		HltbHoursToComplete100Property,
		HltbPlatformsProperty,
		HltbReviewScoreProperty,
	}
}

func EnginesProperties() []string {
	return []string{
		EnginesProperty,
		EnginesBuildsProperty,
	}
}

func PriceProperties() []string {
	return []string{
		BasePriceProperty,
		PriceProperty,
		IsFreeProperty,
		IsDiscountedProperty,
		DiscountPercentageProperty,
		IsDemoProperty,
	}
}

func GOGReducedDataProperties() []string {
	return []string{
		LicencesProperty,
		UserWishlistProperty,
	}
}

func ExternalDataSourcesProperties() []string {
	return []string{
		SteamAppIdProperty,
		SteamReviewScoreDescProperty,
		SteamDeckAppCompatibilityCategoryProperty,
		PcgwPageIdProperty,
		HltbIdProperty,
		IgdbIdProperty,
		StrategyWikiIdProperty,
		MobyGamesIdProperty,
		WikipediaIdProperty,
		WineHQIdProperty,
		VndbIdProperty,
		IGNWikiSlugProperty,
		ProtonDBTierProperty,
		ProtonDBConfidenceProperty,
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
	}
}

func GOGOrderPageProperties() []string {
	return []string{
		GOGOrderDateProperty,
	}
}

func GOGAccountPageProperties() []string {
	return []string{
		TagIdProperty,
		TagNameProperty,
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

func TypeErrorProperties() []string {
	return []string{
		TypeErrorMessageProperty,
		TypeErrorDateProperty,
	}
}

func DehydratedImagesProperties() []string {
	return []string{
		DehydratedImageProperty,
		DehydratedVerticalImageProperty,
		RepImageColorProperty,
		RepVerticalImageColorProperty,
	}
}

func SyncProperties() []string {
	return []string{
		LastSyncUpdatesProperty,
		SyncEventsProperty,
	}
}

func StatusValidationResultsProperties() []string {
	return []string{
		LocalManualUrlProperty,
		ManualUrlStatusProperty,
		ManualUrlValidationResultProperty,
		ManualUrlGeneratedChecksumProperty,
		ProductValidationResultProperty,
	}
}

func ReduxProperties() []string {
	all := TextProperties()
	all = append(all, AllTextProperties()...)
	all = append(all, VideoProperties()...)
	all = append(all, ComputedProperties()...)
	all = append(all, ImageIdProperties()...)
	all = append(all, DehydratedImagesProperties()...)
	all = append(all, UrlProperties()...)
	all = append(all, LongTextProperties()...)
	all = append(all, AvailabilityProperties()...)
	//all = append(all, AccountStatusProperties()...)
	all = append(all, AdvancedProductProperties()...)
	all = append(all, PriceProperties()...)
	all = append(all, GOGReducedDataProperties()...)
	all = append(all, ExternalDataSourcesProperties()...)

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

	all = append(all, SyncProperties()...)
	all = append(all, StatusValidationResultsProperties()...)
	all = append(all, EnginesProperties()...)
	return all
}

func DigestibleProperties() []string {
	return []string{
		DevelopersProperty,
		PublishersProperty,
		GenresProperty,
		StoreTagsProperty,
		FeaturesProperty,
		SeriesProperty,
		TagIdProperty,
		LanguageCodeProperty,
		OperatingSystemsProperty,
		SteamReviewScoreDescProperty,
		SteamDeckAppCompatibilityCategoryProperty,
		HltbPlatformsProperty,
		HltbGenresProperty,
		EnginesProperty,
		ProtonDBTierProperty,
		ProtonDBConfidenceProperty,
	}
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

func boolSlice(confirmer func() bool) []string {
	facts := make([]string, 0)
	if confirmer != nil {
		val := FalseValue
		if confirmer() {
			val = TrueValue
		}
		facts = append(facts, val)
	}
	return facts
}

func dateSlice(timestamper func() int64) []string {
	dates := make([]string, 0)
	if timestamper != nil {
		val := timestamper()
		if val > 0 {
			date := time.Unix(val, 0)
			dates = append(dates, date.Format("2006-01-02"))
		}
	}
	return dates
}

func intSlice(integer func() int) []string {
	values := make([]string, 0)
	if integer != nil {
		values = append(values, strconv.FormatInt(int64(integer()), 10))
	}
	return values
}

func floatSlice(floater func() float64) []string {
	values := make([]string, 0)
	if floater != nil {
		values = append(values, strconv.FormatFloat(floater(), 'f', -1, 64))
	}
	return values
}

func uint32Slice(integer func() []uint32) []string {
	values := make([]string, 0)
	if integer != nil {
		intValues := integer()
		for _, intValue := range intValues {
			if intValue > 0 {
				valStr := strconv.FormatUint(uint64(intValue), 10)
				if !slices.Contains(values, valStr) {
					values = append(values, valStr)
				}
			}
		}
	}
	return values
}

func getSlice(stringer func() string) []string {
	values := make([]string, 0)
	if stringer != nil {
		val := stringer()
		if val != "" {
			values = append(values, val)
		}
	}
	return values
}

func getImageIdSlice(stringer func() string) []string {
	strings := getSlice(stringer)
	imageIds := make([]string, 0, len(strings))
	for _, str := range strings {
		imageIds = append(imageIds, gog_integration.ImageId(str))
	}
	return imageIds
}

func getScreenshots(value interface{}) []string {
	screenshotsGetter := value.(gog_integration.ScreenshotsGetter)
	if screenshotsGetter != nil {
		screenshots := screenshotsGetter.GetScreenshots()
		imageIds := make([]string, 0, len(screenshots))
		for _, scr := range screenshots {
			imageIds = append(imageIds, gog_integration.ImageId(scr))
		}
		return imageIds
	}
	return []string{}
}
