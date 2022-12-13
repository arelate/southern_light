package winehq_integration

import "net/url"

const (
	sClassParam = "sClass"
	iIdParam    = "iId"
)

func WineHQUrl(id string) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   appdbWineHqHost,
		Path:   objectManagerPath,
	}

	q := u.Query()
	q.Set(sClassParam, "application")
	q.Set(iIdParam, id)

	u.RawQuery = q.Encode()

	return u
}
