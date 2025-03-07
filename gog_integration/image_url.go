package gog_integration

import (
	"fmt"
	"github.com/arelate/southern_light"
	"net/url"
	"path"
	"strings"
)

const (
	underscoreFormatterPlaceholder = "_{formatter}"
	formatterPlaceholder           = "{formatter}"
	PngExt                         = ".png"
	JpgExt                         = ".jpg"
)

func ImageId(imageUrl string) string {
	if imageUrl == "" {
		return imageUrl
	}

	_, fn := path.Split(imageUrl)
	fnSansExt := strings.TrimSuffix(fn, path.Ext(fn))

	fnSansExt = strings.TrimSuffix(fnSansExt, underscoreFormatterPlaceholder)
	fnSansExt = strings.TrimSuffix(fnSansExt, formatterPlaceholder)

	return fnSansExt
}

func ImageIds(imageUrls ...string) []string {
	imageIds := make([]string, 0, len(imageUrls))
	for _, imageUrl := range imageUrls {
		imageIds = append(imageIds, ImageId(imageUrl))
	}
	return imageIds
}

func ImageUrl(imageId, ext string) (*url.URL, error) {
	if imageId == "" {
		return nil, fmt.Errorf("gog_urls: empty image-id")
	}
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   imagesHost,
		Path:   strings.Replace(imagesPathTemplate, "{image_id}", imageId, 1) + ext,
	}, nil
}

func ScreenshotsJpgUrls(screenshots []string) ([]*url.URL, error) {
	scrUrls := make([]*url.URL, 0, len(screenshots))
	for _, scr := range screenshots {
		scrUrl, err := ImageUrl(scr, JpgExt)
		if err != nil {
			return scrUrls, err
		}
		scrUrls = append(scrUrls, scrUrl)
	}
	return scrUrls, nil
}
