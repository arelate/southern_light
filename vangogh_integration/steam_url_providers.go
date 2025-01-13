package vangogh_integration

import (
	"github.com/arelate/southern_light/protondb_integration"
	"github.com/arelate/southern_light/steam_integration"
	"github.com/boggydigital/kevlar"
	"net/url"
	"strconv"
)

var steamProductTypeUrlGetters = map[ProductType]func(uint32) *url.URL{
	SteamAppNews:   steam_integration.NewsForAppUrl,
	SteamReviews:   steam_integration.AppReviewsUrl,
	SteamStorePage: steam_integration.StorePageUrl,
	//SteamAppDetails:              steam_integration.AppDetailsUrl,
	SteamDeckCompatibilityReport: steam_integration.DeckAppCompatibilityReportUrl,
	// ProtonDB product types are using Steam AppID
	ProtonDBSummary: protondb_integration.SummaryUrl,
}

type SteamUrlProvider struct {
	pt  ProductType
	rdx kevlar.ReadableRedux
}

func NewSteamUrlProvider(pt ProductType, rdx kevlar.ReadableRedux) (*SteamUrlProvider, error) {
	if err := rdx.MustHave(SteamAppIdProperty); err != nil {
		return nil, err
	}

	return &SteamUrlProvider{
		pt:  pt,
		rdx: rdx,
	}, nil
}

func (sup *SteamUrlProvider) GOGIdToSteamAppId(gogId string) uint32 {
	if appIdStr, ok := sup.rdx.GetLastVal(SteamAppIdProperty, gogId); ok {
		if appId, err := strconv.ParseUint(appIdStr, 10, 32); err == nil {
			return uint32(appId)
		}
	}
	return 0
}

func (sup *SteamUrlProvider) Url(gogId string) *url.URL {
	switch sup.pt {
	case SteamAppList:
		return steam_integration.AppListUrl()
	default:
		if appId := sup.GOGIdToSteamAppId(gogId); appId > 0 {
			if sug, ok := steamProductTypeUrlGetters[sup.pt]; ok {
				return sug(appId)
			}
		}
	}
	return nil
}
