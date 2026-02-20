package pcgw_integration

import "strings"

type PageId struct {
	Cargoquery []struct {
		Title struct {
			PageID     string `json:"PageID"`
			SteamAppID string `json:"Steam AppID"`
		} `json:"title"`
	} `json:"cargoquery"`
}

type PageIdGetter interface {
	GetPageId() string
}

func (ps *PageId) GetPageId() string {
	for _, cq := range ps.Cargoquery {
		return cq.Title.PageID
	}
	return ""
}

func (ps *PageId) GetSteamAppId() string {
	for _, cq := range ps.Cargoquery {
		saids := strings.SplitSeq(cq.Title.SteamAppID, ",")
		for said := range saids {
			return strings.TrimSpace(said)
		}
	}
	return ""
}
