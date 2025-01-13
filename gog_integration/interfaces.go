package gog_integration

type TotalPagesGetter interface {
	GetTotalPages() int
}

type ProductsGetter interface {
	GetProducts() []IdGetter
}

type IdGetter interface {
	GetId() int
}

type TitleGetter interface {
	GetTitle() string
}

type DevelopersGetter interface {
	GetDevelopers() []string
}

type PublishersGetter interface {
	GetPublishers() []string
}

type ImageGetter interface {
	GetImage() string
}

type VerticalImageGetter interface {
	GetVerticalImage() string
}

type HeroGetter interface {
	GetHero() string
}

type LogoGetter interface {
	GetLogo() string
}

type IconGetter interface {
	GetIcon() string
}

type IconSquareGetter interface {
	GetIconSquare() string
}

type ScreenshotsGetter interface {
	GetScreenshots() []string
}

type RatingGetter interface {
	GetRating() string
}

type GenresGetter interface {
	GetGenres() []string
}

type FeaturesGetter interface {
	GetFeatures() []string
}

type SeriesGetter interface {
	GetSeries() string
}

type accountTagsGetter interface {
	getAccountTags() []accountTag
}

type TagIdsGetter interface {
	GetTagIds() []string
}

type TagNamesGetter interface {
	GetTagNames([]string) map[string]string
}

type VideoIdsGetter interface {
	GetVideoIds() []string
}

type OperatingSystemsGetter interface {
	GetOperatingSystems() []string
}

type IncludesGamesGetter interface {
	GetIncludesGames() []string
}

type IsIncludedInGamesGetter interface {
	GetIsIncludedInGames() []string
}

type IsRequiredByGamesGetter interface {
	GetIsRequiredByGames() []string
}

type RequiresGamesGetter interface {
	GetRequiresGames() []string
}

type LanguagesGetter interface {
	GetLanguages() map[string]string
}

type NativeLanguagesGetter interface {
	GetNativeLanguages() map[string]string
}

type LanguageCodesGetter interface {
	GetLanguageCodes() []string
}

type SlugGetter interface {
	GetSlug() string
}

type GlobalReleaseGetter interface {
	GetGlobalRelease() string
}

type GOGReleaseGetter interface {
	GetGOGRelease() string
}

type StoreUrlGetter interface {
	GetStoreUrl() string
}

type ForumUrlGetter interface {
	GetForumUrl() string
}

type SupportUrlGetter interface {
	GetSupportUrl() string
}

type ChangelogGetter interface {
	GetChangelog() string
}

type DescriptionOverviewGetter interface {
	GetDescriptionOverview() string
}

type DescriptionFeaturesGetter interface {
	GetDescriptionFeatures() string
}

type ProductTypeGetter interface {
	GetProductType() string
}

type CopyrightsGetter interface {
	GetCopyrights() string
}

type StoreTagsGetter interface {
	GetStoreTags() []string
}

type InDevelopmentGetter interface {
	GetInDevelopment() bool
}

type PreOrderGetter interface {
	GetPreOrder() bool
}

type ComingSoonGetter interface {
	GetComingSoon() bool
}

type AdditionalRequirementsGetter interface {
	GetAdditionalRequirements() string
}

type BasePriceGetter interface {
	GetBasePrice() string
}

type PriceGetter interface {
	GetPrice() string
}

type IsFreeGetter interface {
	IsFree() bool
}

type IsDiscountedGetter interface {
	IsDiscounted() bool
}

type DiscountPercentageGetter interface {
	GetDiscountPercentage() int
}

type AggregatedRatingGetter interface {
	GetAggregatedRating() float64
}

type ThemesGetter interface {
	GetThemes() []string
}

type GameModesGetter interface {
	GetGameModes() []string
}
