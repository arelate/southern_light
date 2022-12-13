package pcgw_integration

import (
	"strconv"
	"strings"
)

const (
	storeSteamPoweredAppPrefix = "https://store.steampowered.com/app/"
	steamCommunityAppPrefix    = "https://steamcommunity.com/app/"
	howLongToBeatPrefix        = "https://howlongtobeat.com/game?id="
	igdbPrefix                 = "https://www.igdb.com/games/"
	mobyGamesPrefix            = "https://www.mobygames.com/game/"
	wineHQPrefix               = "https://appdb.winehq.org/objectManager.php?sClass=application&iId="
	vndbPrefix                 = "https://vndb.org/"

	strategyWikiPrefix = "strategywiki"
	wikipediaPrefix    = "wikipedia"
)

type IWLinks struct {
}

type ParseExternalLinks struct {
	Parse struct {
		Title         string   `json:"title"`
		PageId        uint64   `json:"pageid"`
		ExternalLinks []string `json:"externallinks"`
		IWLinks       []struct {
			Prefix string `json:"prefix"`
			Url    string `json:"url"`
			Id     string `json:"*"`
		} `json:"iwlinks"`
	} `json:"parse"`
}

type SteamAppIdGetter interface {
	GetSteamAppId() uint32
}

type HowLongToBeatIdGetter interface {
	GetHowLongToBeatId() string
}

type IGDBIdGetter interface {
	GetIGDBId() string
}

type StrategyWikiIdGetter interface {
	GetStrategyWikiId() string
}

type MobyGamesIdGetter interface {
	GetMobyGamesId() string
}

type WikipediaIdGetter interface {
	GetWikipediaId() string
}

type WineHQIdGetter interface {
	GetWineHQId() string
}

type VNDBIdGetter interface {
	GetVNDBId() string
}

func extractSteamAppId(link, pfx string) uint32 {
	if strings.HasPrefix(link, pfx) {
		if parts := strings.Split(strings.TrimPrefix(link, pfx), "/"); len(parts) > 0 {
			if sai, err := strconv.ParseInt(parts[0], 10, 32); err == nil {
				return uint32(sai)
			}
		}
	}
	return 0
}

// GetSteamAppId extracts Steam AppId from PCGamingWiki parse externallinks results
// The same data can be obtained with an original cargo query (see cargoquery_url.go),
// however it seems like cargoquery will return all editions AppIds, while externallinks
// seems more focused on a specific product edition (corresponding to that pageId)
func (pel *ParseExternalLinks) GetSteamAppId() uint32 {
	for _, el := range pel.Parse.ExternalLinks {
		if steamAppId := extractSteamAppId(el, storeSteamPoweredAppPrefix); steamAppId > 0 {
			return steamAppId
		}
		if steamAppId := extractSteamAppId(el, steamCommunityAppPrefix); steamAppId > 0 {
			return steamAppId
		}
	}
	return 0
}

func (pel *ParseExternalLinks) externalLinkSuffixId(pfx string) string {
	for _, link := range pel.Parse.ExternalLinks {
		if strings.HasPrefix(link, pfx) {
			return strings.TrimPrefix(link, pfx)
		}
	}
	return ""
}

func (pel *ParseExternalLinks) iwLinkId(pfx string) string {
	for _, iw := range pel.Parse.IWLinks {
		if iw.Prefix == pfx {
			return strings.TrimPrefix(iw.Id, pfx+":")
		}
	}
	return ""
}

func (pel *ParseExternalLinks) GetHowLongToBeatId() string {
	return pel.externalLinkSuffixId(howLongToBeatPrefix)
}

func (pel *ParseExternalLinks) GetIGDBId() string {
	return pel.externalLinkSuffixId(igdbPrefix)
}

func (pel *ParseExternalLinks) GetMobyGamesId() string {
	return pel.externalLinkSuffixId(mobyGamesPrefix)
}

func (pel *ParseExternalLinks) GetWineHQId() string {
	return pel.externalLinkSuffixId(wineHQPrefix)
}

func (pel *ParseExternalLinks) GetVNDBId() string {
	return pel.externalLinkSuffixId(vndbPrefix)
}

func (pel *ParseExternalLinks) GetWikipediaId() string {
	return pel.iwLinkId(wikipediaPrefix)
}

func (pel *ParseExternalLinks) GetStrategyWikiId() string {
	return pel.iwLinkId(strategyWikiPrefix)
}
