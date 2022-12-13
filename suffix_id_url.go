package southern_light

import "net/url"

func SuffixIdUrl(host, path, id string) *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   host,
		Path:   path + id,
	}
}
