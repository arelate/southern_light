package vangogh_integration

import (
	"github.com/arelate/southern_light/gog_integration"
	"net/url"
)

var gogProductTypeUrlGetters = map[ProductType]func(string) *url.URL{
	CatalogPage:     gog_integration.DefaultCatalogPageUrl,
	AccountPage:     gog_integration.DefaultAccountPageUrl,
	UserWishlist:    gog_integration.DefaultUserWishlistUrl,
	Details:         gog_integration.DetailsUrl,
	ApiProductsV1:   gog_integration.ApiProductV1Url,
	ApiProductsV2:   gog_integration.ApiProductV2Url,
	Licences:        gog_integration.DefaultLicencesUrl,
	OrderPage:       gog_integration.DefaultOrdersPageUrl,
	GamesDBProducts: gog_integration.GamesDbUrl,
}

type GOGUrlProvider struct {
	pt ProductType
}

func NewGOGUrlProvider(pt ProductType) (*GOGUrlProvider, error) {
	return &GOGUrlProvider{pt: pt}, nil
}

func (gup *GOGUrlProvider) Url(gogId string) *url.URL {
	if ug, ok := gogProductTypeUrlGetters[gup.pt]; ok {
		return ug(gogId)
	} else {
		return nil
	}
}
