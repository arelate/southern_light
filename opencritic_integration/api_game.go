package opencritic_integration

import "time"

type ImageSizes struct {
	Og string `json:"og"`
	Xs string `json:"xs"`
	Sm string `json:"sm"`
	Md string `json:"md"`
	Lg string `json:"lg"`
	Xl string `json:"xl"`
	Id string `json:"_id"`
}

type ApiGame struct {
	Images struct {
		Box         ImageSizes   `json:"box"`
		Square      ImageSizes   `json:"square"`
		Masthead    ImageSizes   `json:"masthead"`
		Banner      ImageSizes   `json:"banner"`
		Screenshots []ImageSizes `json:"screenshots"`
	} `json:"images"`
	IsPre2015           bool    `json:"isPre2015"`
	HasLootBoxes        bool    `json:"hasLootBoxes"`
	IsMajorTitle        bool    `json:"isMajorTitle"`
	PercentRecommended  float64 `json:"percentRecommended"`
	NumReviews          int     `json:"numReviews"`
	NumTopCriticReviews int     `json:"numTopCriticReviews"`
	MedianScore         float64 `json:"medianScore"`
	TopCriticScore      float64 `json:"topCriticScore"`
	Percentile          int     `json:"percentile"`
	Tier                string  `json:"tier"`
	Name                string  `json:"name"`
	Platforms           []struct {
		Id          int       `json:"id"`
		Name        string    `json:"name"`
		ShortName   string    `json:"shortName"`
		ImageSrc    string    `json:"imageSrc"`
		ReleaseDate time.Time `json:"releaseDate"`
	} `json:"Platforms"`
	Companies []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"Companies"`
	Genres []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"Genres"`
	Trailers []struct {
		Title         string    `json:"title"`
		VideoId       string    `json:"videoId"`
		ChannelId     string    `json:"channelId"`
		Description   string    `json:"description"`
		ExternalUrl   string    `json:"externalUrl"`
		ChannelTitle  string    `json:"channelTitle"`
		PublishedDate time.Time `json:"publishedDate"`
		IsOpenCritic  bool      `json:"isOpenCritic"`
		IsSpecial     string    `json:"isSpecial"`
	} `json:"trailers"`
	Screenshots []struct {
		Id        string `json:"_id"`
		FullRes   string `json:"fullRes"`
		Thumbnail string `json:"thumbnail"`
	} `json:"screenshots"`
	Id               int       `json:"id"`
	FirstReleaseDate time.Time `json:"firstReleaseDate"`
	Affiliates       []struct {
		ExternalUrl string `json:"externalUrl"`
		Name        string `json:"name"`
	} `json:"Affiliates"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	MastheadScreenshot struct {
		FullRes string `json:"fullRes"`
	} `json:"mastheadScreenshot"`
	Description      string    `json:"description"`
	FirstReviewDate  time.Time `json:"firstReviewDate"`
	LatestReviewDate time.Time `json:"latestReviewDate"`
	BannerScreenshot struct {
		FullRes string `json:"fullRes"`
	} `json:"bannerScreenshot"`
	MonetizationFeatures struct {
		HasLootBoxes bool `json:"hasLootBoxes"`
	} `json:"monetizationFeatures"`
	VerticalLogoScreenshot struct {
		FullRes   string `json:"fullRes"`
		Thumbnail string `json:"thumbnail"`
	} `json:"verticalLogoScreenshot"`
	SquareScreenshot struct {
		FullRes   string `json:"fullRes"`
		Thumbnail string `json:"thumbnail"`
	} `json:"squareScreenshot"`
	ImageMigrationComplete    bool      `json:"imageMigrationComplete"`
	TenthReviewDate           time.Time `json:"tenthReviewDate"`
	BiggestDiscountDollars    float64   `json:"biggestDiscountDollars"`
	BiggestDiscountPercentage float64   `json:"biggestDiscountPercentage"`
	Tags                      []string  `json:"tags"`
	Url                       string    `json:"url"`
}
