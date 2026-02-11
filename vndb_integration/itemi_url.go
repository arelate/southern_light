package vndb_integration

import (
	"encoding/json/v2"
	"github.com/arelate/southern_light"
	"net/url"
	"strings"
)

func ItemUrl(id string) *url.URL {
	return southern_light.SuffixIdUrl(rootHost, itemPath, id)
}

func VNUrl() *url.URL {
	u := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   apiHost,
		Path:   vnPath,
	}

	return u
}

func VNData(id string, fields ...string) (string, error) {
	data := KanaRequestData{
		Filters: []string{"id", "=", id},
		Fields:  strings.Join(fields, ","),
	}
	sb := &strings.Builder{}
	if err := json.MarshalWrite(sb, data); err != nil {
		return sb.String(), err
	}
	return sb.String(), nil
}
