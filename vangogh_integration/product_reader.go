package vangogh_integration

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/arelate/southern_light/gog_integration"
	"github.com/arelate/southern_light/hltb_integration"
	"github.com/arelate/southern_light/pcgw_integration"
	"github.com/arelate/southern_light/protondb_integration"
	"github.com/arelate/southern_light/steam_integration"
	"github.com/boggydigital/kevlar"
	"golang.org/x/net/html"
	"io"
	"iter"
)

type ProductReader struct {
	productType ProductType
	keyValues   kevlar.KeyValues
}

func NewProductReader(pt ProductType) (*ProductReader, error) {
	dst, err := AbsProductTypeDir(pt)
	if err != nil {
		return nil, err
	}

	kv, err := kevlar.New(dst, kevlar.JsonExt)
	if err != nil {
		return nil, err
	}

	pr := &ProductReader{
		productType: pt,
		keyValues:   kv,
	}

	return pr, nil
}

func (pr *ProductReader) readValue(id string, val interface{}) error {
	spReadCloser, err := pr.keyValues.Get(id)
	if err != nil {
		return err
	}

	if spReadCloser == nil {
		return nil
	}

	defer spReadCloser.Close()

	if err := json.NewDecoder(spReadCloser).Decode(val); err != nil {
		return err
	}

	return nil
}

func (pr *ProductReader) Keys() iter.Seq[string] {
	return pr.keyValues.Keys()
}

func (pr *ProductReader) Has(id string) bool {
	return pr.keyValues.Has(id)
}

func (pr *ProductReader) Get(id string) (io.ReadCloser, error) {
	return pr.keyValues.Get(id)
}

func (pr *ProductReader) GetFromStorage(id string) (io.ReadCloser, error) {
	return pr.keyValues.Get(id)
}

func (pr *ProductReader) Set(id string, data io.Reader) error {
	return pr.keyValues.Set(id, data)
}

func (pr *ProductReader) Cut(id string) error {
	return pr.keyValues.Cut(id)
}

func (pr *ProductReader) Since(ts int64, mts ...kevlar.MutationType) iter.Seq2[string, kevlar.MutationType] {
	return pr.keyValues.Since(ts, mts...)
}

func (pr *ProductReader) CatalogProduct(id string) (catalogProduct *gog_integration.CatalogProduct, err error) {
	err = pr.readValue(id, &catalogProduct)
	return catalogProduct, err
}

func (pr *ProductReader) AccountProduct(id string) (accountProduct *gog_integration.AccountProduct, err error) {
	err = pr.readValue(id, &accountProduct)
	return accountProduct, err
}

func (pr *ProductReader) Details(id string) (details *gog_integration.Details, err error) {
	err = pr.readValue(id, &details)
	return details, err
}

func (pr *ProductReader) ApiProductV2(id string) (apiProductV2 *gog_integration.ApiProductV2, err error) {
	err = pr.readValue(id, &apiProductV2)
	return apiProductV2, err
}

func (pr *ProductReader) CatalogPage(page string) (catalogPage *gog_integration.CatalogPage, err error) {
	err = pr.readValue(page, &catalogPage)
	return catalogPage, err
}

func (pr *ProductReader) AccountPage(page string) (accountPage *gog_integration.AccountPage, err error) {
	err = pr.readValue(page, &accountPage)
	return accountPage, err
}

func (pr *ProductReader) UserWishlist() (userWishlist *gog_integration.UserWishlist, err error) {
	err = pr.readValue(UserWishlist.String(), &userWishlist)
	return userWishlist, err
}

func (pr *ProductReader) Licences() (licences *gog_integration.Licences, err error) {
	err = pr.readValue(Licences.String(), &licences)
	return licences, err
}

func (pr *ProductReader) OrderPage(page string) (orderPage *gog_integration.OrderPage, err error) {
	err = pr.readValue(page, &orderPage)
	return orderPage, err
}

func (pr *ProductReader) Order(id string) (order *gog_integration.Order, err error) {
	err = pr.readValue(id, &order)
	return order, err
}

func (pr *ProductReader) SteamGetAppNewsResponse(id string) (steamAppNewsResponse *steam_integration.GetNewsForAppResponse, err error) {
	err = pr.readValue(id, &steamAppNewsResponse)
	return steamAppNewsResponse, err
}

func (pr *ProductReader) SteamAppReviews(id string) (steamAppReviews *steam_integration.AppReviews, err error) {
	err = pr.readValue(id, &steamAppReviews)
	return steamAppReviews, err
}

func (pr *ProductReader) SteamDeckAppCompatibilityReport(id string) (deckAppCompatibilityReport *steam_integration.DeckAppCompatibilityReport, err error) {
	err = pr.readValue(id, &deckAppCompatibilityReport)
	// empty results are passed as an empty array [], not a struct
	var ute *json.UnmarshalTypeError
	if errors.As(err, &ute) {
		if ute.Field == "results" && ute.Value == "array" {
			err = nil
		}
	}
	return deckAppCompatibilityReport, err
}

func (pr *ProductReader) ProtonDBSummary(id string) (summary *protondb_integration.Summary, err error) {
	err = pr.readValue(id, &summary)
	return summary, err
}

func (pr *ProductReader) UserWishlistProduct(id string) (userWishlistProduct string, err error) {
	userWishlistProduct, err = id, nil
	return userWishlistProduct, err
}

func (pr *ProductReader) HltbRootPage() (*hltb_integration.RootPage, error) {
	spReadCloser, err := pr.keyValues.Get(HltbRootPage.String())
	if err != nil {
		return nil, err
	}

	if spReadCloser == nil {
		return nil, nil
	}
	defer spReadCloser.Close()

	doc, err := html.Parse(spReadCloser)
	return &hltb_integration.RootPage{Doc: doc}, err
}

func (pr *ProductReader) HltbData(id string) (data *hltb_integration.Data, err error) {
	err = pr.readValue(id, &data)
	return data, err
}

func (pr *ProductReader) PcgwPageId(id string) (ps *pcgw_integration.PageId, err error) {
	err = pr.readValue(id, &ps)
	return ps, err
}

func (pr *ProductReader) PcgwEngine(id string) (e *pcgw_integration.Engine, err error) {
	err = pr.readValue(id, &e)
	return e, err
}

func (pr *ProductReader) PcgwExternalLinks(id string) (pel *pcgw_integration.ParseExternalLinks, err error) {
	err = pr.readValue(id, &pel)
	return pel, err
}

func (pr *ProductReader) GamesDbProduct(id string) (gdp *gog_integration.GamesDbProduct, err error) {
	err = pr.readValue(id, &gdp)
	return gdp, err
}

func (pr *ProductReader) ReadValue(key string) (interface{}, error) {
	switch pr.productType {
	case CatalogProducts:
		return pr.CatalogProduct(key)
	case AccountProducts:
		return pr.AccountProduct(key)
	case Details:
		return pr.Details(key)
	case ApiProducts:
		return pr.ApiProductV2(key)
	case Orders:
		return pr.Order(key)
	case CatalogPage:
		return pr.CatalogPage(key)
	case AccountPage:
		return pr.AccountPage(key)
	case UserWishlist:
		return pr.UserWishlist()
	case UserWishlistProducts:
		return pr.UserWishlistProduct(key)
	case OrderPage:
		return pr.OrderPage(key)
	case Licences:
		return pr.Licences()
	case SteamAppNews:
		return pr.SteamGetAppNewsResponse(key)
	case SteamAppReviews:
		return pr.SteamAppReviews(key)
	case SteamDeckCompatibilityReport:
		return pr.SteamDeckAppCompatibilityReport(key)
	case PcgwPageId:
		return pr.PcgwPageId(key)
	case PcgwEngine:
		return pr.PcgwEngine(key)
	case PcgwExternalLinks:
		return pr.PcgwExternalLinks(key)
	case HltbRootPage:
		return pr.HltbRootPage()
	case HltbData:
		return pr.HltbData(key)
	case ProtonDbSummary:
		return pr.ProtonDBSummary(key)
	case GamesDbGogProducts:
		return pr.GamesDbProduct(key)
	default:
		return nil, fmt.Errorf("vangogh_values: cannot create %s value", pr.productType)
	}
}

func (pr *ProductReader) ProductType() ProductType {
	return pr.productType
}

func (pr *ProductReader) LogModTime(id string) int64 {
	return pr.keyValues.LogModTime(id)
}

func (pr *ProductReader) FileModTime(id string) (int64, error) {
	return pr.keyValues.FileModTime(id)
}

func (pr *ProductReader) Len() int { return pr.keyValues.Len() }
