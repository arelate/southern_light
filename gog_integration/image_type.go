package gog_integration

import (
	"errors"
	"maps"
	"net/url"
	"slices"
)

type ImageType int

const (
	UnknownImageType ImageType = iota
	Image                      // 1600x740
	Screenshots                // ...
	VerticalImage              // 342x482
	Hero                       // 2560x683
	Logo                       // 1600x740
	Icon                       // 128x128
	IconSquare                 // 128x128
	Background                 // 2560x655
)

var imageTypeStrings = map[ImageType]string{
	UnknownImageType: "unknown-image-type",
	Image:            "image",
	Screenshots:      "screenshots",
	VerticalImage:    "vertical-image",
	Hero:             "hero",
	Logo:             "logo",
	Icon:             "icon",
	IconSquare:       "icon-square",
	Background:       "background",
}

func (it ImageType) String() string {
	str, ok := imageTypeStrings[it]
	if ok {
		return str
	}

	return imageTypeStrings[UnknownImageType]
}

func ParseImageType(imageType string) ImageType {
	for it, str := range imageTypeStrings {
		if str == imageType {
			return it
		}
	}
	return UnknownImageType
}

func ParseManyImageTypes(imageTypes ...string) []ImageType {
	its := make(map[ImageType]any)
	for _, it := range imageTypes {
		its[ParseImageType(it)] = nil
	}
	return slices.Collect(maps.Keys(its))
}

func AllImageTypes() []ImageType {
	imageTypes := make([]ImageType, 0, len(imageTypeStrings))
	for it, _ := range imageTypeStrings {
		if it == UnknownImageType {
			continue
		}
		imageTypes = append(imageTypes, it)
	}
	return imageTypes
}

func ImageTypesCloValues() []string {
	its := AllImageTypes()
	itsStr := make([]string, 0, len(its))
	for _, it := range its {
		itsStr = append(itsStr, it.String())
	}
	return itsStr
}

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
		ext = JpgExt
	case Icon:
		fallthrough
	case IconSquare:
		fallthrough
	case Logo:
		ext = PngExt
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
		imageUrl, err = ImageUrl(imageId, ext)
		if err != nil {
			return urls, err
		}
		urls = append(urls, imageUrl)
	}

	return urls, nil
}
