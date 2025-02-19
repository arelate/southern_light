package vangogh_integration

import (
	"github.com/arelate/southern_light/pcgw_integration"
	"github.com/boggydigital/redux"
	"net/url"
)

type PcgwUrlProvider struct {
	pt  ProductType
	rdx redux.Readable
}

func NewPcgwUrlProvider(pt ProductType, rdx redux.Readable) (*PcgwUrlProvider, error) {
	if err := rdx.MustHave(PcgwPageIdProperty); err != nil {
		return nil, err
	}

	return &PcgwUrlProvider{
		pt:  pt,
		rdx: rdx,
	}, nil
}

func (pcgwup *PcgwUrlProvider) GOGIdToPcgwPageId(gogId string) string {
	if pageId, ok := pcgwup.rdx.GetLastVal(PcgwPageIdProperty, gogId); ok {
		return pageId
	}
	return ""
}

func (pcgwup *PcgwUrlProvider) Url(gogId string) *url.URL {
	switch pcgwup.pt {
	case PcgwPageId:
		return pcgw_integration.PageIdCargoQueryUrl(gogId)
	case PcgwEngine:
		if pageId := pcgwup.GOGIdToPcgwPageId(gogId); pageId != "" {
			return pcgw_integration.EngineCargoQueryUrl(pageId)
		}
	case PcgwExternalLinks:
		if pageId := pcgwup.GOGIdToPcgwPageId(gogId); pageId != "" {
			return pcgw_integration.ParseExternalLinksUrl(pageId)
		}
	}
	return nil
}
