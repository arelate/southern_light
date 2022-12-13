package pcgw_integration

type Cargo struct {
	Cargoquery []struct {
		Title struct {
			PageID string `json:"PageID"`
		} `json:"title"`
	} `json:"cargoquery"`
}

type PageIdGetter interface {
	GetPageId() string
}

func (c *Cargo) GetPageId() string {
	for _, cq := range c.Cargoquery {
		return cq.Title.PageID
	}
	return ""
}
