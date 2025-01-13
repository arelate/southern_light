package vangogh_integration

import (
	"github.com/arelate/southern_light/hltb_integration"
	"github.com/boggydigital/kevlar"
	"net/url"
)

type HLTBUrlProvider struct {
	pt  ProductType
	rdx kevlar.ReadableRedux
}

func NewHLTBUrlProvider(pt ProductType, rdx kevlar.ReadableRedux) (*HLTBUrlProvider, error) {
	if err := rdx.MustHave(HLTBBuildIdProperty, HLTBIdProperty); err != nil {
		return nil, err
	}

	return &HLTBUrlProvider{
		pt:  pt,
		rdx: rdx,
	}, nil
}

func (hup *HLTBUrlProvider) GOGIdToHLTBId(gogId string) string {
	if hltbId, ok := hup.rdx.GetLastVal(HLTBIdProperty, gogId); ok {
		return hltbId
	}
	return ""
}

func (hup *HLTBUrlProvider) Url(gogId string) *url.URL {
	switch hup.pt {
	case HLTBRootPage:
		return hltb_integration.RootUrl()
	case HLTBData:
		if buildId, ok := hup.rdx.GetLastVal(HLTBBuildIdProperty, HLTBRootPage.String()); ok {
			if hltbId := hup.GOGIdToHLTBId(gogId); hltbId != "" {
				return hltb_integration.DataUrl(buildId, hltbId)
			}
		}
	}
	return nil
}
