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
		fallthrough
	case Background:
		ext = gog_integration.JpgExt
	case Icon:
		fallthrough
	case IconSquare:
		fallthrough
	case Logo:
		ext = gog_integration.PngExt
	default:
		return nil, errors.New("no url for unknown image type")
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
