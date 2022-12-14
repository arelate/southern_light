package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"strings"
)

func ApiProductV1Url(id string) *url.URL {
	apv1url := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   apiHost,
		Path:   strings.Replace(apiV1PathTemplate, "{id}", id, 1),
	}

	q := apv1url.Query()
	q.Set("expand", strings.Join(expandValues, ","))
	apv1url.RawQuery = q.Encode()

	return apv1url
}
