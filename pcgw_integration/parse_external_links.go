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

func (pel *ParseExternalLinks) GetHowLongToBeatId() string {
	for _, el := range pel.Parse.ExternalLinks {
		if strings.HasPrefix(el, howLongToBeatPrefix) {
			return strings.TrimPrefix(el, howLongToBeatPrefix)
		}
	}
	return ""
}

func (pel *ParseExternalLinks) GetIGDBId() string {
	for _, el := range pel.Parse.ExternalLinks {
		if strings.HasPrefix(el, igdbPrefix) {
			return strings.TrimPrefix(el, igdbPrefix)
		}
	}
	return ""
}
