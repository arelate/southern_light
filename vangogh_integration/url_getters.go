package vangogh_integration

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	downloadTypeParam  = "download-type"
	propertyParam      = "property"
	sinceHoursAgoParam = "since-hours-ago"
)

const (
	defaultLanguageCode = "en"
)

func ValuesFromUrl(u *url.URL, arg string) []string {
	q := u.Query()
	if q.Has(arg) {
		val := q.Get(arg)
		values := strings.Split(val, ",")
		//account for empty strings
		if len(values) == 1 && values[0] == "" {
			values = []string{}
		}
		return values
	}
	return nil
}

func PropertyFromUrl(u *url.URL) string {
	return u.Query().Get(propertyParam)
}

func PropertiesFromUrl(u *url.URL) []string {
	return ValuesFromUrl(u, propertyParam)
}

func ProductTypeFromUrl(u *url.URL) ProductType {
	return ParseProductType(u.Query().Get(ProductTypeProperty))
}

func ProductTypesFromUrl(u *url.URL) []ProductType {

	q := u.Query()

	productTypes := make([]ProductType, 0)
	if !q.Has(ProductTypeProperty) {
		return productTypes
	}
	ptParam := q.Get(ProductTypeProperty)
	pts := strings.Split(ptParam, ",")
	for _, pt := range pts {
		productTypes = append(productTypes, ParseProductType(pt))
	}
	return productTypes
}

func OperatingSystemsFromUrl(u *url.URL) []OperatingSystem {
	osStrings := ValuesFromUrl(u, OperatingSystemsProperty)
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

	ids := ValuesFromUrl(u, IdProperty)

	if slugs := ValuesFromUrl(u, SlugProperty); len(slugs) > 0 {
		slugIds, err := idsFromSlugs(slugs, nil)
		if err != nil {
			return nil, err
		}
		ids = append(ids, slugIds...)
	}

	return ids, nil
}

func SinceFromUrl(u *url.URL) (int64, error) {
	str := u.Query().Get(sinceHoursAgoParam)
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
