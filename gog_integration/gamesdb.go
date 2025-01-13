package gog_integration

import (
	"strconv"
)

type PlatformId struct {
	Id                   string `json:"id"`
	PlatformId           string `json:"platform_id"`
	ExternalId           string `json:"external_id"`
	ReleasePerPlatformId string `json:"release_per_platform_id"`
}

type DefaultEnglish struct {
	Default string `json:"*"`
	EnUS    string `json:"en-US"`
}

type IdNameStringSlug struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type IdNameLocalizedSlug struct {
	Id   string         `json:"id"`
	Name DefaultEnglish `json:"name"`
	Slug string         `json:"slug"`
}

type UrlFormat struct {
	UrlFormat string `json:"url_format"`
}

type Video struct {
	Provider    string `json:"provider"`
	VideoId     string `json:"video_id"`
	ThumbnailId string `json:"thumbnail_id"`
	Name        string `json:"name"`
}

type GamesDbProduct struct {
	Id                        string       `json:"id"`
	GameId                    string       `json:"game_id"`
	PlatformId                string       `json:"platform_id"`
	ExternalId                string       `json:"external_id"`
	DLCsIds                   []string     `json:"dlcs_ids"`
	DLCs                      []PlatformId `json:"dlcs"`
	ParentId                  interface{}  `json:"parent_id"`
	SupportedOperatingSystems []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"supported_operating_systems"`
	AvailableLanguages []struct {
		Code string `json:"code"`
	} `json:"available_languages"`
	FirstReleaseDate string `json:"first_release_date"`
	Game             struct {
		Id                      string                `json:"id"`
		ParentId                interface{}           `json:"parent_id"`
		DLCsIds                 []string              `json:"dlcs_ids"`
		FirstReleaseDate        string                `json:"first_release_date"`
		Releases                []PlatformId          `json:"releases"`
		Title                   DefaultEnglish        `json:"title"`
		SortingTitle            DefaultEnglish        `json:"sorting_title"`
		Type                    string                `json:"type"`
		DevelopersIds           []string              `json:"developers_ids"`
		Developers              []IdNameStringSlug    `json:"developers"`
		PublishersIds           []string              `json:"publishers_ids"`
		Publishers              []IdNameStringSlug    `json:"publishers"`
		GenresIds               []string              `json:"genres_ids"`
		Genres                  []IdNameLocalizedSlug `json:"genres"`
		ThemesIds               []string              `json:"themes_ids"`
		Themes                  []IdNameLocalizedSlug `json:"themes"`
		Screenshots             []UrlFormat           `json:"screenshots"`
		Videos                  []Video               `json:"videos"`
		Artworks                []UrlFormat           `json:"artworks"`
		Summary                 DefaultEnglish        `json:"summary"`
		VisibleInLibrary        bool                  `json:"visible_in_library"`
		AggregatedRating        float64               `json:"aggregated_rating"`
		GameModes               []IdNameStringSlug    `json:"game_modes"`
		HorizontalArtwork       UrlFormat             `json:"horizontal_artwork"`
		Background              UrlFormat             `json:"background"`
		VerticalCover           UrlFormat             `json:"vertical_cover"`
		Cover                   UrlFormat             `json:"cover"`
		Logo                    UrlFormat             `json:"logo"`
		Icon                    UrlFormat             `json:"icon"`
		SquareIcon              UrlFormat             `json:"square_icon"`
		GlobalPopularityAllTime int                   `json:"global_popularity_all_time"`
		GlobalPopularityCurrent int                   `json:"global_popularity_current"`
		Slug                    string                `json:"slug"`
	} `json:"game"`
	Title        DefaultEnglish     `json:"title"`
	SortingTitle DefaultEnglish     `json:"sorting_title"`
	Type         string             `json:"type"`
	Summary      DefaultEnglish     `json:"summary"`
	Videos       []Video            `json:"videos"`
	GameModes    []IdNameStringSlug `json:"game_modes"`
	Icon         UrlFormat          `json:"icon"`
	Logo         UrlFormat          `json:"logo"`
}

func (gdp *GamesDbProduct) GetSteamAppId() uint32 {
	for _, release := range gdp.Game.Releases {
		if release.PlatformId == "steam" {
			if said, err := strconv.ParseInt(release.ExternalId, 32, 10); err == nil {
				return uint32(said)
			}
		}
	}
	return 0
}

func (gdp *GamesDbProduct) GetVideoIds() []string {
	videoIds := make([]string, 0, len(gdp.Videos))

	for _, vid := range gdp.Videos {
		if vid.Provider == "youtube" {
			videoIds = append(videoIds, vid.VideoId)
		}
	}
	return videoIds
}

func (gdp *GamesDbProduct) GetHero() string {
	return gdp.Game.Background.UrlFormat
}

func (gdp *GamesDbProduct) GetVerticalImage() string {
	return gdp.Game.VerticalCover.UrlFormat
}

func (gdp *GamesDbProduct) GetIcon() string {
	return gdp.Game.Icon.UrlFormat
}

func (gdp *GamesDbProduct) GetIconSquare() string {
	return gdp.Game.SquareIcon.UrlFormat
}

func (gdp *GamesDbProduct) GetAggregatedRating() float64 {
	return gdp.Game.AggregatedRating
}

func (gdp *GamesDbProduct) GetThemes() []string {
	themes := make([]string, 0, len(gdp.Game.Themes))
	for _, theme := range gdp.Game.Themes {
		themes = append(themes, theme.Name.Default)
	}
	return themes
}

func (gdp *GamesDbProduct) GetGameModes() []string {
	gameModes := make([]string, 0, len(gdp.Game.GameModes))
	for _, gm := range gdp.Game.GameModes {
		gameModes = append(gameModes, gm.Name)
	}
	return gameModes
}
