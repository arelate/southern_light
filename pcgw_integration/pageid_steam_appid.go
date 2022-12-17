package pcgw_integration

type PageIdSteamAppId struct {
	Cargoquery []struct {
		Title struct {
			PageID string `json:"PageID"`
		} `json:"title"`
	} `json:"cargoquery"`
}

type PageIdGetter interface {
	GetPageId() string
}

func (ps *PageIdSteamAppId) GetPageId() string {
	for _, cq := range ps.Cargoquery {
		return cq.Title.PageID
	}
	return ""
}
