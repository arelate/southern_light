package gogdb_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func GOGDBUrl(gogId string) *url.URL {
	return southern_light.SuffixIdUrl(gogDBHost, productPath, gogId)
}
