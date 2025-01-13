package vangogh_integration

import (
	"errors"
	"github.com/arelate/southern_light/gog_integration"
	"net/url"
)

func ImagePropertyUrls(imageIds []string, it ImageType) ([]*url.URL, error) {
	urls := make([]*url.URL, 0, len(imageIds))

	var ext string
	switch it {
	case Image:
		fallthrough
	case VerticalImage:
		fallthrough
	case Screenshots:
		fallthrough
	case Hero:
		ext = gog_integration.JpgExt
	case Icon:
		fallthrough
	case IconSquare:
		fallthrough
	case Logo:
		ext = gog_integration.PngExt
	default:
		return nil, errors.New("unknown image type")
	}

	for _, imageId := range imageIds {
		if imageId == "" {
			continue
		}
		imageUrl, err := gog_integration.ImageUrl(imageId, ext)
		if err != nil {
			return urls, err
		}
		urls = append(urls, imageUrl)
	}

	return urls, nil
}

//type DefaultProductUrl func(key string) *url.URL
//
//var defaultProductUrls = map[ProductType]DefaultProductUrl{
//	CatalogPage:   gog_integration.DefaultCatalogPageUrl,
//	AccountPage:   gog_integration.DefaultAccountPageUrl,
//	UserWishlist:  gog_integration.DefaultUserWishlistUrl,
//	Details:       gog_integration.DetailsUrl,
//	ApiProductsV1: gog_integration.ApiProductV1Url,
//	ApiProductsV2: gog_integration.ApiProductV2Url,
//	Licences:      gog_integration.DefaultLicencesUrl,
//	OrderPage:     gog_integration.DefaultOrdersPageUrl,
//
//	SteamAppList: steam_integration.DefaultSteamAppListUrl,
//	// steam data types typically require app level transformation of GOG.com id -> Steam AppID
//	SteamAppNews:   nil,
//	SteamReviews:   nil,
//	SteamStorePage: nil,
//}

//func RemoteProductsUrl(pt ProductType) (ptUrl DefaultProductUrl, err error) {
//	if !IsValidProductType(pt) {
//		return nil, fmt.Errorf("vangogh_urls: no remote source for %s\n", pt)
//	}
//
//	ptUrl, ok := defaultProductUrls[pt]
//	if !ok {
//		err = fmt.Errorf("vangogh_urls: no remote source for %s\n", pt)
//	}
//
//	return ptUrl, err
//}
