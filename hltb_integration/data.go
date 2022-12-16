package hltb_integration

import (
	"fmt"
	"strings"
)

type Data struct {
	PageProps struct {
		Game struct {
			Count int `json:"count"`
			Data  struct {
				Game []struct {
					CountDiscussion   int    `json:"count_discussion"`
					GameId            int    `json:"game_id"`
					GameName          string `json:"game_name"`
					GameNameDate      int    `json:"game_name_date"`
					CountPlaying      int    `json:"count_playing"`
					CountBacklog      int    `json:"count_backlog"`
					CountReplay       int    `json:"count_replay"`
					CountCustom       int    `json:"count_custom"`
					CountComp         int    `json:"count_comp"`
					CountRetired      int    `json:"count_retired"`
					CountReview       int    `json:"count_review"`
					ReviewScore       int    `json:"review_score"`
					GameAlias         string `json:"game_alias"`
					GameImage         string `json:"game_image"`
					GameType          string `json:"game_type"`
					GameParent        int    `json:"game_parent"`
					ProfileSummary    string `json:"profile_summary"`
					ProfileDev        string `json:"profile_dev"`
					ProfilePub        string `json:"profile_pub"`
					ProfilePlatform   string `json:"profile_platform"`
					ProfileGenre      string `json:"profile_genre"`
					ProfileSteam      uint32 `json:"profile_steam"`
					ProfileSteamAlt   uint32 `json:"profile_steam_alt"`
					ProfileItch       int    `json:"profile_itch"`
					ProfileIgn        string `json:"profile_ign"`
					ReleaseWorld      string `json:"release_world"`
					ReleaseNa         string `json:"release_na"`
					ReleaseEu         string `json:"release_eu"`
					ReleaseJp         string `json:"release_jp"`
					RatingEsrb        string `json:"rating_esrb"`
					RatingPegi        string `json:"rating_pegi"`
					RatingCero        string `json:"rating_cero"`
					CompLvlSp         int    `json:"comp_lvl_sp"`
					CompLvlSpd        int    `json:"comp_lvl_spd"`
					CompLvlCo         int    `json:"comp_lvl_co"`
					CompLvlMp         int    `json:"comp_lvl_mp"`
					CompLvlCombine    int    `json:"comp_lvl_combine"`
					CompLvlPlatform   int    `json:"comp_lvl_platform"`
					CompAllCount      int    `json:"comp_all_count"`
					CompAll           int    `json:"comp_all"`
					CompAllL          int    `json:"comp_all_l"`
					CompAllH          int    `json:"comp_all_h"`
					CompAllAvg        int    `json:"comp_all_avg"`
					CompAllMed        int    `json:"comp_all_med"`
					CompMainCount     int    `json:"comp_main_count"`
					CompMain          int    `json:"comp_main"`
					CompMainL         int    `json:"comp_main_l"`
					CompMainH         int    `json:"comp_main_h"`
					CompMainAvg       int    `json:"comp_main_avg"`
					CompMainMed       int    `json:"comp_main_med"`
					CompPlusCount     int    `json:"comp_plus_count"`
					CompPlus          int    `json:"comp_plus"`
					CompPlusL         int    `json:"comp_plus_l"`
					CompPlusH         int    `json:"comp_plus_h"`
					CompPlusAvg       int    `json:"comp_plus_avg"`
					CompPlusMed       int    `json:"comp_plus_med"`
					Comp100Count      int    `json:"comp_100_count"`
					Comp100           int    `json:"comp_100"`
					Comp100L          int    `json:"comp_100_l"`
					Comp100H          int    `json:"comp_100_h"`
					Comp100Avg        int    `json:"comp_100_avg"`
					Comp100Med        int    `json:"comp_100_med"`
					CompSpeedCount    int    `json:"comp_speed_count"`
					CompSpeed         int    `json:"comp_speed"`
					CompSpeedMin      int    `json:"comp_speed_min"`
					CompSpeedMax      int    `json:"comp_speed_max"`
					CompSpeedAvg      int    `json:"comp_speed_avg"`
					CompSpeedMed      int    `json:"comp_speed_med"`
					CompSpeed100Count int    `json:"comp_speed100_count"`
					CompSpeed100      int    `json:"comp_speed100"`
					CompSpeed100Min   int    `json:"comp_speed100_min"`
					CompSpeed100Max   int    `json:"comp_speed100_max"`
					CompSpeed100Avg   int    `json:"comp_speed100_avg"`
					CompSpeed100Med   int    `json:"comp_speed100_med"`
					CountTotal        int    `json:"count_total"`
					InvestedCoCount   int    `json:"invested_co_count"`
					InvestedCo        int    `json:"invested_co"`
					InvestedCoL       int    `json:"invested_co_l"`
					InvestedCoH       int    `json:"invested_co_h"`
					InvestedCoAvg     int    `json:"invested_co_avg"`
					InvestedCoMed     int    `json:"invested_co_med"`
					InvestedMpCount   int    `json:"invested_mp_count"`
					InvestedMp        int    `json:"invested_mp"`
					InvestedMpL       int    `json:"invested_mp_l"`
					InvestedMpH       int    `json:"invested_mp_h"`
					InvestedMpAvg     int    `json:"invested_mp_avg"`
					InvestedMpMed     int    `json:"invested_mp_med"`
					AddedStats        string `json:"added_stats"`
				} `json:"game"`
				Individuality []struct {
					Platform  string `json:"platform"`
					CountComp string `json:"count_comp"`
					CompMain  string `json:"comp_main"`
					CompPlus  string `json:"comp_plus"`
					Comp100   string `json:"comp_100"`
					CompAll   string `json:"comp_all"`
					Compare   string `json:"compare"`
				} `json:"individuality"`
				Relationships []struct {
					GameId       int    `json:"game_id"`
					GameName     string `json:"game_name"`
					GameType     string `json:"game_type"`
					CompMain     int    `json:"comp_main"`
					CompPlus     int    `json:"comp_plus"`
					Comp100      int    `json:"comp_100"`
					CompAll      int    `json:"comp_all"`
					CompAllCount int    `json:"comp_all_count"`
					CountBacklog int    `json:"count_backlog"`
					ReviewScore  int    `json:"review_score"`
				} `json:"relationships"`
				UserReviews struct {
					ReviewCount int    `json:"review_count"`
					Score10     string `json:"score_10"`
					Score20     string `json:"score_20"`
					Score30     string `json:"score_30"`
					Score40     string `json:"score_40"`
					Score50     string `json:"score_50"`
					Score60     string `json:"score_60"`
					Score70     string `json:"score_70"`
					Score80     string `json:"score_80"`
					Score90     string `json:"score_90"`
					Score100    string `json:"score_100"`
				} `json:"userReviews"`
				PlatformData []struct {
					Platform   string `json:"platform"`
					CountComp  int    `json:"count_comp"`
					CountTotal int    `json:"count_total"`
					CompMain   int    `json:"comp_main"`
					CompPlus   int    `json:"comp_plus"`
					Comp100    int    `json:"comp_100"`
					CompLow    int    `json:"comp_low"`
					CompHigh   int    `json:"comp_high"`
				} `json:"platformData"`
			} `json:"data"`
		} `json:"game"`
		IgnWikiSlug string `json:"ignWikiSlug"`
		//PageMetadata struct {
		//	Title       string `json:"title"`
		//	Image       string `json:"image"`
		//	Description string `json:"description"`
		//	Canonical   string `json:"canonical"`
		//	Template    string `json:"template"`
		//} `json:"pageMetadata"`
		//SentryTraceData string `json:"_sentryTraceData"`
		//SentryBaggage   string `json:"_sentryBaggage"`
	} `json:"pageProps"`
	NSSP bool `json:"__N_SSP"`
}

type HoursToCompleteMainGetter interface {
	GetHoursToCompleteMain() string
}

type HoursToCompletePlusGetter interface {
	GetHoursToCompletePlus() string
}

type HoursToComplete100Getter interface {
	GetHoursToComplete100() string
}

type ReviewScoreGetter interface {
	GetReviewScore() int
}

type PlatformsGetter interface {
	GetPlatforms() []string
}

type IGNWikiSlugGetter interface {
	GetIGNWikiSlug() string
}

func (d *Data) fmtSecondsToHours(style string) string {
	seconds := 0
	if gd := d.PageProps.Game.Data.Game; len(gd) > 0 {
		switch style {
		case "main":
			seconds = gd[0].CompMain
		case "plus":
			seconds = gd[0].CompPlus
		case "100":
			seconds = gd[0].Comp100
		}
	}

	if seconds > 0 {
		return fmt.Sprintf("%07.1f", float64(seconds)/3600.0)
	}
	return ""
}

func (d *Data) GetHoursToCompleteMain() string {
	return d.fmtSecondsToHours("main")
}

func (d *Data) GetHoursToCompletePlus() string {
	return d.fmtSecondsToHours("plus")
}

func (d *Data) GetHoursToComplete100() string {
	return d.fmtSecondsToHours("100")
}

func (d *Data) GetSteamAppId() uint32 {
	if gd := d.PageProps.Game.Data.Game; len(gd) > 0 {
		return gd[0].ProfileSteam
	}
	return 0
}

func (d *Data) GetReviewScore() int {
	if gd := d.PageProps.Game.Data.Game; len(gd) > 0 {
		return gd[0].ReviewScore
	}
	return 0
}

func (d *Data) GetGenres() []string {
	if gd := d.PageProps.Game.Data.Game; len(gd) > 0 {
		return strings.Split(gd[0].ProfileGenre, ", ")
	}
	return nil
}

func (d *Data) GetPlatforms() []string {
	if gd := d.PageProps.Game.Data.Game; len(gd) > 0 {
		return strings.Split(gd[0].ProfilePlatform, ", ")
	}
	return nil
}

func (d *Data) GetGlobalRelease() string {
	if gd := d.PageProps.Game.Data.Game; len(gd) > 0 {
		return strings.Replace(gd[0].ReleaseWorld, "-", ".", -1)
	}
	return ""
}

func (d *Data) GetIGNWikiSlug() string {
	return d.PageProps.IgnWikiSlug
}
