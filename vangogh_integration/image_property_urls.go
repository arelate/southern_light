package vangogh_integration

import (
	"errors"
	"net/url"

	"github.com/arelate/southern_light/gog_integration"
)

func ImagePropertyExt(it gog_integration.ImageType) (string, error) {

	var ext string
	var err error

	switch it {
	case gog_integration.Image:
		fallthrough
	case gog_integration.VerticalImage:
		fallthrough
	case gog_integration.Screenshots:
		fallthrough
	case gog_integration.Hero:
		fallthrough
	case gog_integration.Background:
		ext = gog_integration.JpgExt
	case gog_integration.Icon:
		fallthrough
	case gog_integration.IconSquare:
		fallthrough
	case gog_integration.Logo:
		ext = gog_integration.PngExt
	default:
		err = errors.New("no url for unknown image type")
	}

	return ext, err
}

func ImagePropertyUrls(imageIds []string, it gog_integration.ImageType) ([]*url.URL, error) {
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
