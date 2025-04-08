package vangogh_integration

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	imageTypeParam       = "image-type"
	operatingSystemParam = OperatingSystemsProperty
	downloadTypeParam    = "download-type"
	productTypeParam     = ProductTypeProperty
	idParam              = IdProperty
	slugParam            = SlugProperty
	propertyParam        = "property"
	sinceHoursAgoParam   = "since-hours-ago"
)

const (
	defaultLanguageCode = "en"
)

func ValueFromUrl(u *url.URL, arg string) string {
	if u == nil {
		return ""
	}

	q := u.Query()
	return q.Get(arg)
}

func ValuesFromUrl(u *url.URL, arg string) []string {
	if FlagFromUrl(u, arg) {
		val := ValueFromUrl(u, arg)
		values := strings.Split(val, ",")
		//account for empty strings
		if len(values) == 1 && values[0] == "" {
			values = []string{}
		}
		return values
	}
	return nil
}

func FlagFromUrl(u *url.URL, arg string) bool {
	if u == nil {
		return false
	}

	q := u.Query()
	return q.Has(arg)
}

func PropertyFromUrl(u *url.URL) string {
	return ValueFromUrl(u, propertyParam)
}

func PropertiesFromUrl(u *url.URL) []string {
	return ValuesFromUrl(u, propertyParam)
}

func ImageTypesFromUrl(u *url.URL) []ImageType {
	imageTypes := ValuesFromUrl(u, imageTypeParam)
	its := make([]ImageType, 0, len(imageTypes))
	for _, imageType := range imageTypes {
		its = append(its, ParseImageType(imageType))
	}
	return its
}

func ProductTypeFromUrl(u *url.URL) ProductType {
	return ParseProductType(ValueFromUrl(u, productTypeParam))
}

func ProductTypesFromUrl(u *url.URL) []ProductType {
	productTypes := make([]ProductType, 0)
	if !u.Query().Has(productTypeParam) {
		return productTypes
	}
	ptParam := u.Query().Get(productTypeParam)
	pts := strings.Split(ptParam, ",")
	for _, pt := range pts {
		productTypes = append(productTypes, ParseProductType(pt))
	}
	return productTypes
}

func OperatingSystemsFromUrl(u *url.URL) []OperatingSystem {
	osStrings := ValuesFromUrl(u, operatingSystemParam)
	return ParseManyOperatingSystems(osStrings)
}

func LanguageCodesFromUrl(u *url.URL) []string {
	if langCodes := ValuesFromUrl(u, LanguageCodeProperty); len(langCodes) > 0 {
		return langCodes
	} else {
		return []string{defaultLanguageCode}
	}
}

func DownloadTypesFromUrl(u *url.URL) []DownloadType {
	dtStrings := ValuesFromUrl(u, downloadTypeParam)
	return ParseManyDownloadTypes(dtStrings)
}

func IdsFromUrl(u *url.URL) ([]string, error) {

	ids := ValuesFromUrl(u, idParam)

	if slugs := ValuesFromUrl(u, slugParam); len(slugs) > 0 {
		slugIds, err := idsFromSlugs(slugs, nil)
		if err != nil {
			return nil, err
		}
		ids = append(ids, slugIds...)
	}

	return ids, nil
}

func SinceFromUrl(u *url.URL) (int64, error) {
	str := ValueFromUrl(u, sinceHoursAgoParam)
	var sha int
	var err error
	if str != "" {
		sha, err = strconv.Atoi(str)
		if err != nil {
			return 0, err
		}
	}
	return time.Now().Unix() - int64(sha*60*60), err
}

func DownloadsLayoutFromUrl(u *url.URL) DownloadsLayout {
	downloadsLayout := DefaultDownloadsLayout
	q := u.Query()
	if q.Has("downloads-layout") {
		if dl := ParseDownloadsLayout(q.Get("downloads-layout")); dl != UnknownDownloadsLayout {
			downloadsLayout = dl
		}
	}
	return downloadsLayout
}
