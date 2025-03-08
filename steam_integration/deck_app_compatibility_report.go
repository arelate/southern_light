package steam_integration

type DeckAppCompatibilityReport struct {
	Success int                         `json:"success"`
	Results DeckAppCompatibilityResults `json:"results"`
}

type DeckAppCompatibilityResults struct {
	AppID            uint32 `json:"appid"`
	ResolvedCategory int    `json:"resolved_category"`
	ResolvedItems    []struct {
		DisplayType int    `json:"display_type"`
		LocToken    string `json:"loc_token"`
	} `json:"resolved_items"`
	SteamDeckBlogUrl string      `json:"steam_deck_blog_url"`
	SearchId         interface{} `json:"search_id"`
}

func (dacr *DeckAppCompatibilityReport) IsSuccess() bool {
	return dacr.Success == 1
}

func (dacr *DeckAppCompatibilityReport) String() string {
	return DecodeCategory(dacr.Results.ResolvedCategory)
}

func (dacr *DeckAppCompatibilityReport) GetDisplayTypes() []string {
	ridt := make([]string, 0, len(dacr.Results.ResolvedItems))
	for _, ri := range dacr.Results.ResolvedItems {
		ridt = append(ridt, DecodeCategory(ri.DisplayType-1))
	}
	return ridt
}

func (dacr *DeckAppCompatibilityReport) GetResults() []string {
	rilt := make([]string, 0, len(dacr.Results.ResolvedItems))
	for _, ri := range dacr.Results.ResolvedItems {
		rilt = append(rilt, TrimLocToken(ri.LocToken))
	}
	return rilt
}

func (dacr *DeckAppCompatibilityReport) GetBlogUrl() string {
	return dacr.Results.SteamDeckBlogUrl
}
