package vangogh_integration

import (
	"errors"
	"fmt"
	"github.com/boggydigital/redux"
	"net/url"
)

type UrlProvider interface {
	Url(gogId string) *url.URL
}

func NewUrlProvider(pt ProductType, rdx redux.Readable) (UrlProvider, error) {
	if IsGOGProduct(pt) {
		return NewGOGUrlProvider(pt)
	} else if IsSteamProduct(pt) {
		return NewSteamUrlProvider(pt, rdx)
	} else if IsPCGWProduct(pt) {
		return NewPCGWUrlProvider(pt, rdx)
	} else if IsHLTBProduct(pt) {
		return NewHLTBUrlProvider(pt, rdx)
	} else if IsProtonDBProduct(pt) {
		return NewSteamUrlProvider(pt, rdx)
	} else {
		return nil, errors.New(fmt.Sprintf("product type %s is not a url provider", pt))
	}
}
