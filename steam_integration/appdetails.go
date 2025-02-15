package steam_integration

type AppDetails struct {
	Type                string       `json:"type"`
	Name                string       `json:"name"`
	SteamAppId          uint32       `json:"steam_appid"`
	RequiredAge         string       `json:"required_age"`
	IsFree              bool         `json:"is_free"`
	ControllerSupport   string       `json:"controller_support"`
	DetailedDescription string       `json:"detailed_description"`
	AboutTheGame        string       `json:"about_the_game"`
	ShortDescription    string       `json:"short_description"`
	SupportedLanguages  string       `json:"supported_languages"`
	HeaderImage         string       `json:"header_image"`
	CapsuleImage        string       `json:"capsule_image"`
	CapsuleImageV5      string       `json:"capsule_imagev5"`
	Website             string       `json:"website"`
	PcRequirements      Requirements `json:"pc_requirements"`
	MacRequirements     Requirements `json:"mac_requirements"`
	LinuxRequirements   Requirements `json:"linux_requirements"`
	LegalNotice         string       `json:"legal_notice"`
	Developers          []string     `json:"developers"`
	Publishers          []string     `json:"publishers"`
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
		DisplayType             int    `json:"display_type"`
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

type AppDetailsResponse struct {
	Success bool       `json:"success"`
	Data    AppDetails `json:"data"`
}
