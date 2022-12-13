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
	strategyWikiPrefix         = "https://strategywiki.org/wiki/"
	mobyGamesPrefix            = "https://www.mobygames.com/game/"
	wikipediaPrefix            = "https://en.wikipedia.org/wiki/"
	wineHQPrefix               = "https://appdb.winehq.org/objectManager.php?sClass=application&iId="
)

type ParseExternalLinks struct {
	Parse struct {
		Title         string   `json:"title"`
		PageId        uint64   `json:"pageid"`
		ExternalLinks []string `json:"externallinks"`
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

func (pel *ParseExternalLinks) extractSuffixId(pfx string) string {
	for _, link := range pel.Parse.ExternalLinks {
		if strings.HasPrefix(link, pfx) {
			return strings.TrimPrefix(link, pfx)
		}
	}
	return ""
}

func (pel *ParseExternalLinks) GetHowLongToBeatId() string {
	return pel.extractSuffixId(howLongToBeatPrefix)
}

func (pel *ParseExternalLinks) GetIGDBId() string {
	return pel.extractSuffixId(igdbPrefix)
}

func (pel *ParseExternalLinks) GetStrategyWikiId() string {
	return pel.extractSuffixId(strategyWikiPrefix)
}

func (pel *ParseExternalLinks) GetMobyGamesId() string {
	return pel.extractSuffixId(mobyGamesPrefix)
}

func (pel *ParseExternalLinks) GetWikipediaId() string {
	return pel.extractSuffixId(wikipediaPrefix)
}

func (pel *ParseExternalLinks) GetWineHQId() string {
	return pel.extractSuffixId(wineHQPrefix)
}
