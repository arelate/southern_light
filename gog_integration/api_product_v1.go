package gog_integration

import (
	"net/url"
	"time"
)

type downloadFiles struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Os           string `json:"os"`
	Language     string `json:"language"`
	LanguageFull string `json:"language_full"`
	Version      string `json:"version"`
	Files        []struct {
		Id       string `json:"id"`
		Size     int    `json:"size"`
		Downlink string `json:"downlink"`
	} `json:"files"`
}

type bonusFiles struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Count     int    `json:"count"`
	TotalSize int    `json:"total_size"`
	Files     []struct {
		Id       int    `json:"id"`
		Size     int    `json:"size"`
		Downlink string `json:"downlink"`
	} `json:"files"`
}

type images struct {
	Background          string `json:"background"`
	Logo                string `json:"logo"`
	Logo2X              string `json:"logo2x"`
	Icon                string `json:"icon"`
	SidebarIcon         string `json:"sidebarIcon"`
	SidebarIcon2X       string `json:"sidebarIcon2x"`
	MenuNotificationAv  string `json:"menuNotificationAv"`
	MenuNotificationAv2 string `json:"menuNotificationAv2"`
}

type dlcProduct struct {
	Id           string `json:"id"`
	Link         string `json:"link"`
	ExpandedLink string `json:"expanded_link"`
}

type ApiProductV1 struct {
	Id                         int    `json:"id"`
	Title                      string `json:"title"`
	PurchaseLink               string `json:"purchase_link"`
	Slug                       string `json:"slug"`
	ContentSystemCompatibility struct {
		Windows bool `json:"windows"`
		Osx     bool `json:"osx"`
		Linux   bool `json:"linux"`
	} `json:"content_system_compatibility"`
	//see comment on top of commented out implementation of GetDLCs()
	//that applies here as well - empty maps are passed as [] by GOG
	Languages interface{} `json:"languages"`
	Links     struct {
		PurchaseLink string `json:"purchase_link"`
		ProductCard  string `json:"product_card"`
		Support      string `json:"support"`
		Forum        string `json:"forum"`
	} `json:"links"`
	InDevelopment struct {
		Active bool        `json:"active"`
		Until  interface{} `json:"until"`
	} `json:"in_development"`
	IsSecret      bool   `json:"is_secret"`
	IsInstallable bool   `json:"is_installable"`
	GameType      string `json:"game_type"`
	IsPreOrder    bool   `json:"is_pre_order"`
	ReleaseDate   string `json:"release_date"`
	Images        images `json:"images"`
	//see comment on top of commented out implementation of GetDLCs()
	//that explains why DLCs is an interface and not a struct
	DLCs      interface{} `json:"dlcs"`
	Downloads struct {
		Installers    []downloadFiles `json:"installers"`
		Patches       []downloadFiles `json:"patches"`
		LanguagePacks []downloadFiles `json:"language_packs"`
		BonusContent  []bonusFiles    `json:"bonus_content"`
	} `json:"downloads"`
	ExpandedDLCs []ApiProductV1 `json:"expanded_dlcs"`
	Description  struct {
		Lead             string `json:"lead"`
		Full             string `json:"full"`
		WhatsCoolAboutIt string `json:"whats_cool_about_it"`
	} `json:"description"`
	Screenshots []struct {
		ImageId              string `json:"image_id"`
		FormatterTemplateUrl string `json:"formatter_template_url"`
		FormattedImages      []struct {
			FormatterName string `json:"formatter_name"`
			ImageUrl      string `json:"image_url"`
		} `json:"formatted_images"`
	} `json:"screenshots"`
	Videos []struct {
		VideoUrl     string `json:"video_url"`
		ThumbnailUrl string `json:"thumbnail_url"`
		Provider     string `json:"provider"`
	} `json:"videos"`
	RelatedProducts []ApiProductV1 `json:"related_products"`
	Changelog       string         `json:"changelog"`
}

func (apv1 *ApiProductV1) GetId() int {
	return apv1.Id
}

func (apv1 *ApiProductV1) GetTitle() string {
	return apv1.Title
}

func (apv1 *ApiProductV1) GetScreenshots() []string {
	screenshots := make([]string, 0)
	for _, screenshot := range apv1.Screenshots {
		screenshots = append(screenshots, screenshot.ImageId)
	}
	return screenshots
}

func (apv1 *ApiProductV1) GetVideoIds() []string {
	videoIds := make([]string, 0, len(apv1.Videos))
	for _, vid := range apv1.Videos {
		if vid.Provider != supportedVideoProvider {
			continue
		}
		videoIds = append(videoIds, youtubeVideoId(vid.VideoUrl))
	}
	return videoIds
}

func (apv1 *ApiProductV1) GetOperatingSystems() []string {
	os := make([]string, 0, 3)
	if apv1.ContentSystemCompatibility.Windows {
		os = append(os, windows)
	}
	if apv1.ContentSystemCompatibility.Osx {
		os = append(os, macos)
	}
	if apv1.ContentSystemCompatibility.Linux {
		os = append(os, linux)
	}
	return os
}

//GOG.com provides empty `dlcs` as [], while the `dlcs` that has values is an object {...}
//given there is no great way to support this in Go that I'm aware of - ApiProductsV1 won't
//support IsRequiredByGames property (which is effectively DLCs).
//We'll only be able to get that from ApiProductV2 that doesn't have this problem.
//
//func (apv1 *ApiProductV1) GetIsRequiredByGames() []string {
//	dlcs := make([]string, 0, len(apv1.DLCs.Products))
//	for _, prod := range apv1.DLCs.Products {
//		dlcs = append(dlcs, prod.Id)
//	}
//	return dlcs
//}

func (apv1 *ApiProductV1) GetSlug() string {
	return normalizeSlug(apv1.Slug)
}

func (apv1 *ApiProductV1) GetNativeLanguages() map[string]string {
	nativeLanguages := make(map[string]string, 0)

	if apv1.Languages == nil {
		return nativeLanguages
	}

	langs, ok := apv1.Languages.(map[string]interface{})
	if !ok {
		return nativeLanguages
	}

	for langCode, nativeInterface := range langs {
		nativeName, ok := nativeInterface.(string)
		if !ok {
			continue
		}
		nativeLanguages[langCode] = nativeName
	}

	return nativeLanguages
}

func (apv1 *ApiProductV1) GetGOGRelease() string {
	if apv1.ReleaseDate == "" {
		return ""
	}
	//see comment for ApiProductV2.GetGOGRelease that explains those special dates
	if apv1.ReleaseDate == "1991-12-25T00:00:00+0200" ||
		apv1.ReleaseDate == "2000-12-31T23:59:59+02:00" {
		return ""
	}
	if lw, err := time.LoadLocation("Europe/Nicosia"); err == nil {
		if t, err := time.Parse("2006-01-02T15:04:05-0700", apv1.ReleaseDate); err == nil {
			return t.In(lw).Format("2006.01.02")
		}
	}
	return ""
}

func urlPathFromLink(link string) string {
	if u, err := url.Parse(link); err != nil {
		return ""
	} else {
		return u.Path
	}
}

func (apv1 *ApiProductV1) GetStoreUrl() string {
	return urlPathFromLink(apv1.Links.ProductCard)
}

func (apv1 *ApiProductV1) GetForumUrl() string {
	return urlPathFromLink(apv1.Links.Forum)
}

func (apv1 *ApiProductV1) GetSupportUrl() string {
	return urlPathFromLink(apv1.Links.Support)
}

func (apv1 *ApiProductV1) GetChangelog() string {
	return apv1.Changelog
}

func (apv1 *ApiProductV1) GetDescriptionOverview() string {
	return apv1.Description.Full
}

func (apv1 *ApiProductV1) GetDescriptionFeatures() string {
	return apv1.Description.WhatsCoolAboutIt
}

func (apv1 *ApiProductV1) GetInDevelopment() bool {
	return apv1.InDevelopment.Active
}

func (apv1 *ApiProductV1) GetPreOrder() bool {
	return apv1.IsPreOrder
}
