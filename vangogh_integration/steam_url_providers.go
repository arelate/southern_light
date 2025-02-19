package vangogh_integration

import (
	"github.com/arelate/southern_light/protondb_integration"
	"github.com/arelate/southern_light/steam_integration"
	"github.com/boggydigital/redux"
	"net/url"
)

var steamProductTypeUrlGetters = map[ProductType]func(string) *url.URL{
	SteamAppNews:    steam_integration.NewsForAppUrl,
	SteamAppReviews: steam_integration.AppReviewsUrl,
	//SteamAppDetails:              steam_integration.AppDetailsUrl,
	SteamDeckCompatibilityReport: steam_integration.DeckAppCompatibilityReportUrl,
	// ProtonDB product types are using Steam AppID
	ProtonDBSummary: protondb_integration.SummaryUrl,
}

type SteamUrlProvider struct {
	pt  ProductType
	rdx redux.Readable
}

func NewSteamUrlProvider(pt ProductType, rdx redux.Readable) (*SteamUrlProvider, error) {
	if err := rdx.MustHave(SteamAppIdProperty); err != nil {
		return nil, err
	}

	return &SteamUrlProvider{
		pt:  pt,
		rdx: rdx,
	}, nil
}

func (sup *SteamUrlProvider) GOGIdToSteamAppId(gogId string) string {
	if appId, ok := sup.rdx.GetLastVal(SteamAppIdProperty, gogId); ok {
		return appId
	}
	return ""
}

func (sup *SteamUrlProvider) Url(gogId string) *url.URL {
	switch sup.pt {
	default:
		if appId := sup.GOGIdToSteamAppId(gogId); appId != "" {
			if sug, ok := steamProductTypeUrlGetters[sup.pt]; ok {
				return sug(appId)
			}
		}
	}
	return nil
}
