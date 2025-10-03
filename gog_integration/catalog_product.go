package gog_integration

import (
	"strconv"
	"strings"
)

const (
	CatalogProductStateDefault     = "default"
	CatalogProductStateComingSoon  = "coming-soon"
	CatalogProductStateEarlyAccess = "early-access"
	CatalogProductStatePreOrder    = "preorder"
)

type CatalogProduct struct {
	Id                    string                 `json:"id"`
	Title                 string                 `json:"title"`
	Slug                  string                 `json:"slug"`
	Developers            []string               `json:"developers"`
	Publishers            []string               `json:"publishers"`
	Features              []NameSlug             `json:"features"`
	Editions              []Edition              `json:"editions"`
	Genres                []NameSlug             `json:"genres"`
	ProductType           string                 `json:"productType"`
	ReleaseDate           *string                `json:"releaseDate"`
	ReviewsRating         int                    `json:"reviewsRating"`
	OperatingSystems      []string               `json:"operatingSystems"`
	CoverHorizontal       string                 `json:"coverHorizontal"`
	CoverVertical         string                 `json:"coverVertical"`
	Screenshots           []string               `json:"screenshots"`
	ProductState          string                 `json:"productState"`
	Tags                  []NameSlug             `json:"tags"`
	UserPreferredLanguage *UserPreferredLanguage `json:"userPreferredLanguage"`
	Price                 CatalogPrice           `json:"price"`
	StoreLink             string                 `json:"storeLink"`
}

type Edition struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	IsRootEdition bool   `json:"isRootEdition"`
}

type AmountCurrency struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type UserPreferredLanguage struct {
	Code    string `json:"code"`
	InAudio bool   `json:"inAudio"`
	InText  bool   `json:"inText"`
}

type CatalogPrice struct {
	Final      string         `json:"final"`
	Base       string         `json:"base"`
	Discount   *string        `json:"discount"`
	FinalMoney AmountCurrency `json:"finalMoney"`
	BaseMoney  AmountCurrency `json:"baseMoney"`
}

func (cp *CatalogProduct) GetTitle() string {
	return cp.Title
}

func (cp *CatalogProduct) GetDevelopers() []string {
	return cp.Developers
}

func (cp *CatalogProduct) GetPublishers() []string {
	return cp.Publishers
}

func (cp *CatalogProduct) GetImage() string {
	return cp.CoverHorizontal
}

func (cp *CatalogProduct) GetVerticalImage() string {
	return cp.CoverVertical
}

func (cp *CatalogProduct) GetScreenshots() []string {
	return cp.Screenshots
}

func (cp *CatalogProduct) GetFeatures() []string {
	return namesSlugsToNames(cp.Features)
}

func (cp *CatalogProduct) GetRating() string {
	return strconv.Itoa(cp.ReviewsRating)
}

func namesSlugsToNames(nss []NameSlug) []string {
	names := make([]string, 0, len(nss))
	for _, ns := range nss {
		names = append(names, ns.Name)
	}
	return names
}

func (cp *CatalogProduct) GetGenres() []string {
	return namesSlugsToNames(cp.Genres)
}

func (cp *CatalogProduct) GetOperatingSystems() []string {
	os := make([]string, 0, 3)
	for _, cpos := range cp.OperatingSystems {
		switch cpos {
		case windows:
			os = append(os, windows)
		case osx:
			// GOG.com continues to use "osx" name for macOS
			os = append(os, macos)
		case linux:
			os = append(os, linux)
		}
	}
	return os
}

func normalizeSlug(slug string) string {
	// GOG.com have been using "_" words delimeter (as in "first_second") for slugs.
	// With catalog-products, however, they've changed to "-" (as in "first-second").
	// This has created inconsistency in how slugs are presented in data,
	// so this function will normalize slugs to established format
	return strings.Replace(slug, "-", "_", -1)
}

func (cp *CatalogProduct) GetSlug() string {
	return normalizeSlug(cp.Slug)
}

func (cp *CatalogProduct) GetGlobalRelease() string {
	if cp.ReleaseDate != nil {
		return *cp.ReleaseDate
	}
	return ""
}

func (cp *CatalogProduct) GetProductType() string {
	return strings.ToUpper(cp.ProductType)
}

func (cp *CatalogProduct) GetStoreTags() []string {
	return namesSlugsToNames(cp.Tags)
}

func (cp *CatalogProduct) GetBasePrice() string {
	return cp.Price.BaseMoney.Amount
}

func (cp *CatalogProduct) GetPrice() string {
	return cp.Price.FinalMoney.Amount
}

func (cp *CatalogProduct) IsFree() bool {
	return cp.Price.FinalMoney.Amount == "0.00"
}

func (cp *CatalogProduct) IsDiscounted() bool {
	return cp.Price.Discount != nil
}

func (cp *CatalogProduct) GetDiscountPercentage() int {
	if ds := cp.Price.Discount; ds != nil {
		da := strings.TrimPrefix(*ds, "-")
		da = strings.TrimSuffix(da, "%")
		if dp, err := strconv.Atoi(da); err == nil {
			return dp
		}
	}
	return 0
}

func (cp *CatalogProduct) GetComingSoon() bool {
	return cp.ProductState == CatalogProductStateComingSoon
}

func (cp *CatalogProduct) GetPreOrder() bool {
	return cp.ProductState == CatalogProductStatePreOrder
}

func (cp *CatalogProduct) GetInDevelopment() bool {
	return cp.ProductState == CatalogProductStateEarlyAccess
}

func (cp *CatalogProduct) GetStoreLink() string { return cp.StoreLink }

func (cp *CatalogProduct) GetEditions() []string {
	eds := make([]string, 0, len(cp.Editions))
	for _, ed := range cp.Editions {
		eds = append(eds, strconv.Itoa(ed.Id))
	}
	return eds
}
