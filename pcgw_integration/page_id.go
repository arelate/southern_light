package pcgw_integration

type PageId struct {
	Cargoquery []struct {
		Title struct {
			PageID string `json:"PageID"`
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
