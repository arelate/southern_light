package steam_integration

import "strconv"

type AppDetails struct {
	Type                string   `json:"type"`
	Name                string   `json:"name"`
	SteamAppId          uint32   `json:"steam_appid"`
	RequiredAge         any      `json:"required_age"` // could be "number" or number
	IsFree              bool     `json:"is_free"`
	ControllerSupport   string   `json:"controller_support"`
	DetailedDescription string   `json:"detailed_description"`
	AboutTheGame        string   `json:"about_the_game"`
	ShortDescription    string   `json:"short_description"`
	SupportedLanguages  string   `json:"supported_languages"`
	HeaderImage         string   `json:"header_image"`
	CapsuleImage        string   `json:"capsule_image"`
	CapsuleImageV5      string   `json:"capsule_imagev5"`
	Website             string   `json:"website"`
	PcRequirements      any      `json:"pc_requirements"`    // could be {Requirements} or []
	MacRequirements     any      `json:"mac_requirements"`   // could be {Requirements} or []
	LinuxRequirements   any      `json:"linux_requirements"` // could be {Requirements} or []
	LegalNotice         string   `json:"legal_notice"`
	Developers          []string `json:"developers"`
	Publishers          []string `json:"publishers"`
	PriceOverview       struct {
		Currency         string `json:"currency"`
		Initial          int    `json:"initial"`
		Final            int    `json:"final"`
		DiscountPercent  int    `json:"discount_percent"`
		InitialFormatted string `json:"initial_formatted"`
		FinalFormatted   string `json:"final_formatted"`
	} `json:"price_overview"`
	Packages      []int `json:"packages"`
	PackageGroups []struct {
		Name                    string `json:"name"`
		Title                   string `json:"title"`
		Description             string `json:"description"`
		SelectionText           string `json:"selection_text"`
		SaveText                string `json:"save_text"`
		DisplayType             any    `json:"display_type"` // could be "number" or number
		IsRecurringSubscription string `json:"is_recurring_subscription"`
		Subs                    []struct {
			PackageId                int    `json:"packageid"`
			PercentSavingsText       string `json:"percent_savings_text"`
			PercentSavings           int    `json:"percent_savings"`
			OptionText               string `json:"option_text"`
			OptionDescription        string `json:"option_description"`
			CanGetFreeLicense        string `json:"can_get_free_license"`
			IsFreeLicense            bool   `json:"is_free_license"`
			PriceInCentsWithDiscount int    `json:"price_in_cents_with_discount"`
		} `json:"subs"`
	} `json:"package_groups"`
	Platforms struct {
		Windows bool `json:"windows"`
		Mac     bool `json:"mac"`
		Linux   bool `json:"linux"`
	} `json:"platforms"`
	Metacritic struct {
		Score int    `json:"score"`
		Url   string `json:"url"`
	} `json:"metacritic"`
	Categories []struct {
		Id          int    `json:"id"`
		Description string `json:"description"`
	} `json:"categories"`
	Genres []struct {
		Id          string `json:"id"`
		Description string `json:"description"`
	} `json:"genres"`
	Screenshots []struct {
		Id            int    `json:"id"`
		PathThumbnail string `json:"path_thumbnail"`
		PathFull      string `json:"path_full"`
	} `json:"screenshots"`
	Recommendations struct {
		Total int `json:"total"`
	} `json:"recommendations"`
	ReleaseDate struct {
		ComingSoon bool   `json:"coming_soon"`
		Date       string `json:"date"`
	} `json:"release_date"`
	SupportInfo struct {
		Url   string `json:"url"`
		Email string `json:"email"`
	} `json:"support_info"`
	Background         string `json:"background"`
	BackgroundRaw      string `json:"background_raw"`
	ContentDescriptors struct {
		Ids   []int   `json:"ids"`
		Notes *string `json:"notes"`
	} `json:"content_descriptors"`
	Ratings struct {
		Esrb   Rating `json:"esrb"`
		Pegi   Rating `json:"pegi"`
		Usk    Rating `json:"usk"`
		Oflc   Rating `json:"oflc"`
		Dejus  Rating `json:"dejus"`
		Csrr   Rating `json:"csrr"`
		Crl    Rating `json:"crl"`
		Nzoflc Rating `json:"nzoflc"`
	} `json:"ratings"`
}

type Rating struct {
	Rating      string `json:"rating"`
	Descriptors string `json:"descriptors"`
	UseAgeGate  string `json:"use_age_gate"`
	RequiredAge string `json:"required_age"`
}

type Requirements struct {
	Minimum     string `json:"minimum"`
	Recommended string `json:"recommended"`
}

type AppDetailsSuccessData struct {
	Success bool       `json:"success"`
	Data    AppDetails `json:"data"`
}

type AppDetailsResponse map[string]AppDetailsSuccessData

func (adr AppDetailsResponse) GetAppDetails() *AppDetails {
	for _, adsd := range adr {
		return &adsd.Data
	}
	return nil
}

func (ad *AppDetails) GetType() string {
	return ad.Type
}

func (ad *AppDetails) GetName() string {
	return ad.Name
}

func (ad *AppDetails) GetRequiredAge() int {
	ra := ad.RequiredAge

	if ira, ok := ra.(int); ok {
		return ira
	} else if fra, yeah := ra.(float64); yeah {
		return int(fra)
	} else if sra, sure := ra.(string); sure {
		if ira64, err := strconv.ParseInt(sra, 10, 32); err == nil {
			return int(ira64)
		}
	}
	return 0
}

func (ad *AppDetails) GetIsFree() bool {
	return ad.IsFree
}

func (ad *AppDetails) GetControllerSupport() string {
	return ad.ControllerSupport
}

func (ad *AppDetails) GetShortDescription() string {
	return ad.ShortDescription
}

func (ad *AppDetails) GetWebsite() string {
	return ad.Website
}

func (ad *AppDetails) GetMetacriticScore() int {
	return ad.Metacritic.Score
}

func (ad *AppDetails) GetMetacriticUrl() string {
	return ad.Metacritic.Url
}

func (ad *AppDetails) GetCategories() []string {
	categories := make([]string, 0)
	for _, category := range ad.Categories {
		categories = append(categories, category.Description)
	}
	return categories
}

func (ad *AppDetails) GetGenres() []string {
	genres := make([]string, 0)
	for _, genre := range ad.Genres {
		genres = append(genres, genre.Description)
	}
	return genres
}

func (ad *AppDetails) GetSupportUrl() string {
	return ad.SupportInfo.Url
}

func (ad *AppDetails) GetSupportEmail() string {
	return ad.SupportInfo.Email
}
