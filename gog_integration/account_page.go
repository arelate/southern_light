// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

type accountTag struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ProductCount string `json:"productCount"`
}

type AccountPage struct {
	TotalPages      int    `json:"totalPages"`
	Page            int    `json:"page"`
	SortBy          string `json:"sortBy"`
	TotalProducts   int    `json:"totalProducts"`
	ProductsPerPage int    `json:"productsPerPage"`
	// TODO: find data examples where ContentSystemCompatibility is not empty and create a type from that
	ContentSystemCompatibility interface{}      `json:"contentSystemCompatibility"`
	MoviesCount                int              `json:"moviesCount"`
	Tags                       []accountTag     `json:"tags"`
	Products                   []AccountProduct `json:"products"`
	UpdatedProductsCount       int              `json:"updatedProductsCount"`
	HiddenUpdatedProductsCount int              `json:"hiddenUpdatedProductsCount"`
	AppliedFilters             struct {
		Tags []struct {
			Id           string `json:"id"`
			Name         string `json:"name"`
			ProductCount string `json:"productCount"`
		} `json:"tags"`
	} `json:"appliedFilters"`
	HasHiddenProducts bool `json:"hasHiddenProducts"`
}

func (app *AccountPage) GetTotalPages() int {
	return app.TotalPages
}

func (app *AccountPage) GetProducts() []IdGetter {
	idGetters := make([]IdGetter, 0)
	for _, ap := range app.Products {
		ap := ap
		idGetters = append(idGetters, &ap)
	}
	return idGetters
}

func (app *AccountPage) getAccountTags() []accountTag {
	return app.Tags
}

func (app *AccountPage) GetTagIds() []string {
	return getTagIds(app)
}

func (app *AccountPage) GetTagNames(tagIds []string) map[string]string {
	return getTagNames(tagIds, app)
}
