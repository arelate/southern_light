package vangogh_integration

import (
	"errors"
	"net/url"

	"github.com/arelate/southern_light/gog_integration"
)

func ImagePropertyExt(it ImageType) (string, error) {

	var ext string
	var err error

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
		err = errors.New("no url for unknown image type")
	}

	return ext, err
}

func ImagePropertyUrls(imageIds []string, it ImageType) ([]*url.URL, error) {
	urls := make([]*url.URL, 0, len(imageIds))

	ext, err := ImagePropertyExt(it)
	if err != nil {
		return nil, err
	}

	for _, imageId := range imageIds {
		if imageId == "" {
			continue
		}
		var imageUrl *url.URL
		imageUrl, err = gog_integration.ImageUrl(imageId, ext)
		if err != nil {
			return urls, err
		}
		urls = append(urls, imageUrl)
	}

	return urls, nil
}
