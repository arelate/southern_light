package steam_integration

type DeckAppCompatibilityReport struct {
	Success int                         `json:"success"`
	Results DeckAppCompatibilityResults `json:"results"`
}

type ResolvedItem struct {
	DisplayType int    `json:"display_type"`
	LocToken    string `json:"loc_token"`
}

type DeckAppCompatibilityResults struct {
	AppID                   uint32         `json:"appid"`
	ResolvedCategory        int            `json:"resolved_category"`
	ResolvedItems           []ResolvedItem `json:"resolved_items"`
	SteamDeckBlogUrl        string         `json:"steam_deck_blog_url"`
	SearchId                interface{}    `json:"search_id"`
	SteamOsResolvedCategory int            `json:"steamos_resolved_category"`
	SteamOsResolvedItems    []ResolvedItem `json:"steamos_resolved_items"`
}

func (dacr *DeckAppCompatibilityReport) IsSuccess() bool {
	return dacr.Success == 1
}

func (dacr *DeckAppCompatibilityReport) SteamDeckString() string {
	return DecodeCategory(dacr.Results.ResolvedCategory)
}

func (dacr *DeckAppCompatibilityReport) GetSteamDeckDisplayTypes() []string {
	ridt := make([]string, 0, len(dacr.Results.ResolvedItems))
	for _, ri := range dacr.Results.ResolvedItems {
		ridt = append(ridt, DecodeCategory(ri.DisplayType-1))
	}
	return ridt
}

func (dacr *DeckAppCompatibilityReport) GetSteamDeckResults() []string {
	rilt := make([]string, 0, len(dacr.Results.ResolvedItems))
	for _, ri := range dacr.Results.ResolvedItems {
		rilt = append(rilt, SteamDeckTrimLocToken(ri.LocToken))
	}
	return rilt
}

func (dacr *DeckAppCompatibilityReport) SteamOsString() string {
	return DecodeCategory(dacr.Results.SteamOsResolvedCategory)
}

func (dacr *DeckAppCompatibilityReport) GetSteamOsDisplayTypes() []string {
	ridt := make([]string, 0, len(dacr.Results.SteamOsResolvedItems))
	for _, ri := range dacr.Results.SteamOsResolvedItems {
		ridt = append(ridt, DecodeCategory(ri.DisplayType-1))
	}
	return ridt
}

func (dacr *DeckAppCompatibilityReport) GetSteamOsResults() []string {
	rilt := make([]string, 0, len(dacr.Results.SteamOsResolvedItems))
	for _, ri := range dacr.Results.SteamOsResolvedItems {
		rilt = append(rilt, SteamOsTrimLocToken(ri.LocToken))
	}
	return rilt
}

func (dacr *DeckAppCompatibilityReport) GetBlogUrl() string {
	return dacr.Results.SteamDeckBlogUrl
}
