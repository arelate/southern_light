package gog_integration

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

const (
	formatterPlaceholder = "_{formatter}"
	PngExt               = ".png"
	JpgExt               = ".jpg"
)

func ImageId(imageUrl string) string {
	if imageUrl == "" {
		return imageUrl
	}

	_, fn := path.Split(imageUrl)
	fnSansExt := strings.TrimSuffix(fn, path.Ext(fn))

	return strings.TrimSuffix(fnSansExt, formatterPlaceholder)
}

func ImageUrl(imageId, ext string) (*url.URL, error) {
	if imageId == "" {
		return nil, fmt.Errorf("gog_urls: empty image-id")
	}
	return &url.URL{
		Scheme: HttpsScheme,
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
