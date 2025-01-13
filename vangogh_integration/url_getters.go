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

func OperatingSystemsFromUrl(u *url.URL) []OperatingSystem {
	osStrings := ValuesFromUrl(u, operatingSystemParam)
	return ParseManyOperatingSystems(osStrings)
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
