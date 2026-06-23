package gog_integration

import (
	"maps"
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
