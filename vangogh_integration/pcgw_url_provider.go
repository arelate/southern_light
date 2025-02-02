package vangogh_integration

import (
	"github.com/arelate/southern_light/pcgw_integration"
	"github.com/boggydigital/redux"
	"net/url"
)

type PCGWUrlProvider struct {
	pt  ProductType
	rdx redux.Readable
}

func NewPCGWUrlProvider(pt ProductType, rdx redux.Readable) (*PCGWUrlProvider, error) {
	if err := rdx.MustHave(PCGWPageIdProperty); err != nil {
		return nil, err
	}

	return &PCGWUrlProvider{
		pt:  pt,
		rdx: rdx,
	}, nil
}

func (pcgwup *PCGWUrlProvider) GOGIdToPCGWPageId(gogId string) string {
	if pageId, ok := pcgwup.rdx.GetLastVal(PCGWPageIdProperty, gogId); ok {
		return pageId
	}
	return ""
}

func (pcgwup *PCGWUrlProvider) Url(gogId string) *url.URL {
	switch pcgwup.pt {
	case PCGWPageId:
		return pcgw_integration.PageIdCargoQueryUrl(gogId)
	case PCGWEngine:
		if pageId := pcgwup.GOGIdToPCGWPageId(gogId); pageId != "" {
			return pcgw_integration.EngineCargoQueryUrl(pageId)
		}
	case PCGWExternalLinks:
		if pageId := pcgwup.GOGIdToPCGWPageId(gogId); pageId != "" {
			return pcgw_integration.ParseExternalLinksUrl(pageId)
		}
	}
	return nil
}
