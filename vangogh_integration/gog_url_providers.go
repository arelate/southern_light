package vangogh_integration

import (
	"github.com/arelate/southern_light/gog_integration"
	"net/url"
)

var gogProductTypeUrlGetters = map[ProductType]func(string) *url.URL{
	CatalogPage:            gog_integration.CatalogPageUrl,
	AccountPage:            gog_integration.AccountPageUrl,
	UserWishlist:           gog_integration.DefaultUserWishlistUrl,
	Details:                gog_integration.DetailsUrl,
	ApiProducts:            gog_integration.ApiProductUrl,
	Licences:               gog_integration.DefaultLicencesUrl,
	OrderPage:              gog_integration.OrdersPageUrl,
	GamesDbGogProducts:     gog_integration.GamesDbGogExternalReleaseUrl,
	GamesDbGenericProducts: gog_integration.GamesDbGenericExternalReleaseUrl,
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
