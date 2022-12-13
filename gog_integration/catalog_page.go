package gog_integration

type CatalogPage struct {
	Pages        int              `json:"pages"`
	ProductCount int              `json:"productCount"`
	Products     []CatalogProduct `json:"products"`
	Filters      CatalogFilters   `json:"filters"`
}

type CatalogFilters struct {
	ReleaseDateRange struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"releaseDateRange"`
	PriceRange struct {
		Min           float64 `json:"min"`
		Max           float64 `json:"max"`
		Currency      string  `json:"currency"`
		DecimalPlaces int     `json:"decimalPlaces"`
	} `json:"priceRange"`
	Genres          []NameSlug `json:"genres"`
	Languages       []NameSlug `json:"languages"`
	Systems         []NameSlug `json:"systems"`
	Tags            []NameSlug `json:"tags"`
	Discounted      bool       `json:"discounted"`
	Features        []NameSlug `json:"features"`
	ReleaseStatuses []NameSlug `json:"releaseStatuses"`
	Types           []string   `json:"types"`
	FullGenresList  []struct {
		NameSlug
		Level int `json:"level"`
	} `json:"fullGenresList"`
	FullTagsList []NameSlug `json:"fullTagsList"`
}

type NameSlug struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (cp *CatalogPage) GetTotalPages() int {
	return cp.Pages
}

func (cp *CatalogPage) GetProducts() []IdGetter {
	idGetters := make([]IdGetter, 0)
	for _, p := range cp.Products {
		sp := p
		idGetters = append(idGetters, &sp)
	}
	return idGetters
}
