package vangogh_integration

import (
	"github.com/arelate/southern_light/hltb_integration"
	"github.com/boggydigital/redux"
	"net/url"
)

type HltbUrlProvider struct {
	pt  ProductType
	rdx redux.Readable
}

func NewHltbUrlProvider(pt ProductType, rdx redux.Readable) (*HltbUrlProvider, error) {
	//if err := rdx.MustHave(HltbBuildIdProperty, HltbIdProperty); err != nil {
	//	return nil, err
	//}

	return &HltbUrlProvider{
		pt:  pt,
		rdx: rdx,
	}, nil
}

func (hup *HltbUrlProvider) GogIdToHltbId(gogId string) string {
	if hltbId, ok := hup.rdx.GetLastVal(HltbIdProperty, gogId); ok {
		return hltbId
	}
	return ""
}

func (hup *HltbUrlProvider) Url(gogId string) *url.URL {
	switch hup.pt {
	case HltbRootPage:
		return hltb_integration.RootUrl()
	case HltbData:
		//if buildId, ok := hup.rdx.GetLastVal(HltbBuildIdProperty, HltbRootPage.String()); ok {
		//	if hltbId := hup.GogIdToHltbId(gogId); hltbId != "" {
		//		return hltb_integration.DataUrl(buildId, hltbId)
		//	}
		//}
	}
	return nil
}
