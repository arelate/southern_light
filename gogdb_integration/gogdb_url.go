package gogdb_integration

import "net/url"

func GOGDBUrl(gogId string) *url.URL {
	return &url.URL{
		Scheme: httpsScheme,
		Host:   gogDBHost,
		Path:   productPath + gogId,
	}
}
