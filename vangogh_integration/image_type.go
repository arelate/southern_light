package vangogh_integration

import (
	"maps"
	"slices"
)

type ImageType int

const (
	UnknownImageType ImageType = iota
	Image
	Screenshots
	VerticalImage
	Hero
	Logo
	Icon
	IconSquare
	Background
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

func IsValidImageType(it ImageType) bool {
	_, ok := imageTypeStrings[it]
	return ok && it != UnknownImageType
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

// starting with empty collection and no image types require auth at the moment
var imageTypeRequiresAuth []ImageType

func ImageTypeFromProperty(property string) ImageType {
	for it, prop := range imageTypeProperties {
		if prop == property {
			return it
		}
	}
	return UnknownImageType
}

// values are result of running issa dehydration on an image set and finding
// sampling rate that produces most compact representation (min sum of string lengths)
var imageTypeDehydrationSamples = map[ImageType]int{
	VerticalImage: 16,
	Image:         20,
}

func ImageTypeDehydrationSamples(it ImageType) int {
	if s, ok := imageTypeDehydrationSamples[it]; ok {
		return s
	}
	return -1
}

func ImageTypesCloValues() []string {
	its := AllImageTypes()
	itsStr := make([]string, 0, len(its))
	for _, it := range its {
		itsStr = append(itsStr, it.String())
	}
	return itsStr
}
