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
	SearchId                any            `json:"search_id"`
	MachineResolvedCategory int            `json:"machine_resolved_category"`
	MachineResolvedItems    []ResolvedItem `json:"machine_resolved_items"`
	FrameResolvedCategory   int            `json:"frame_resolved_category"`
	FrameResolvedItems      []ResolvedItem `json:"frame_resolved_items"`
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

func (dacr *DeckAppCompatibilityReport) SteamMachineString() string {
	return DecodeCategory(dacr.Results.MachineResolvedCategory)
}

func (dacr *DeckAppCompatibilityReport) GetSteamMachineDisplayTypes() []string {
	ridt := make([]string, 0, len(dacr.Results.MachineResolvedItems))
	for _, ri := range dacr.Results.MachineResolvedItems {
		ridt = append(ridt, DecodeCategory(ri.DisplayType-1))
	}
	return ridt
}

func (dacr *DeckAppCompatibilityReport) GetSteamMachineResults() []string {
	rilt := make([]string, 0, len(dacr.Results.MachineResolvedItems))
	for _, ri := range dacr.Results.MachineResolvedItems {
		rilt = append(rilt, SteamDeckTrimLocToken(ri.LocToken))
	}
	return rilt
}

func (dacr *DeckAppCompatibilityReport) SteamFrameString() string {
	return DecodeCategory(dacr.Results.FrameResolvedCategory)
}

func (dacr *DeckAppCompatibilityReport) GetSteamFrameDisplayTypes() []string {
	ridt := make([]string, 0, len(dacr.Results.FrameResolvedItems))
	for _, ri := range dacr.Results.FrameResolvedItems {
		ridt = append(ridt, DecodeCategory(ri.DisplayType-1))
	}
	return ridt
}

func (dacr *DeckAppCompatibilityReport) GetSteamFrameResults() []string {
	rilt := make([]string, 0, len(dacr.Results.FrameResolvedItems))
	for _, ri := range dacr.Results.FrameResolvedItems {
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
