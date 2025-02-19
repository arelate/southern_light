package gog_integration

import (
	"net/url"
	"path"
	"strings"
	"time"
)

type href struct {
	Href string `json:"href"`
}

type name struct {
	Name string `json:"name"`
}

type slug struct {
	Slug string `json:"slug"`
}

type links struct {
	Self                  href   `json:"self"`
	Store                 href   `json:"store"`
	Support               href   `json:"support"`
	Forum                 href   `json:"forum"`
	Icon                  href   `json:"icon"`
	IconSquare            href   `json:"iconSquare"`
	Logo                  href   `json:"logo"`
	BoxArtImage           href   `json:"boxArtImage"`
	BackgroundImage       href   `json:"backgroundImage"`
	GalaxyBackgroundImage href   `json:"galaxyBackgroundImage"`
	IncludesGames         []href `json:"includesGames"`
	IsIncludedInGames     []href `json:"isIncludedInGames"`
	RequiresGames         []href `json:"requiresGames"`
	IsRequiredByGames     []href `json:"isRequiredByGames"`
}

type supportedOperatingSystem struct {
	OperatingSystem struct {
		Name     string `json:"name"`
		Versions string `json:"versions"`
	} `json:"operatingSystem"`
	SystemRequirements []struct {
		Type         string `json:"type"`
		Description  string `json:"description"`
		Requirements []struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"requirements"`
	} `json:"systemRequirements"`
}

type edition struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Bonuses []struct {
		Id int `json:"id"`
		name
		Type struct {
			name
			slug
		} `json:"type"`
	} `json:"bonuses"`
	HasProductCard bool `json:"hasProductCard"`
	Links          struct {
		Self   href `json:"self"`
		Prices struct {
			href
			Templated bool `json:"templated"`
		} `json:"prices"`
	} `json:"_links"`
}

type product struct {
	IsAvailableForSale bool      `json:"isAvailableForSale"`
	IsVisibleInCatalog bool      `json:"isVisibleInCatalog"`
	IsPreorder         bool      `json:"isPreorder"`
	IsVisibleInAccount bool      `json:"isVisibleInAccount"`
	Id                 int       `json:"id"`
	Title              string    `json:"title"`
	IsInstallable      bool      `json:"isInstallable"`
	GlobalReleaseDate  time.Time `json:"globalReleaseDate"`
	HasProductCard     bool      `json:"hasProductCard"`
	GogReleaseDate     time.Time `json:"gogReleaseDate"`
	IsSecret           bool      `json:"isSecret"`
	Links              struct {
		Image struct {
			href
			Templated  bool     `json:"templated"`
			Formatters []string `json:"formatters"`
		} `json:"image"`
		Checkout href `json:"checkout"`
		Prices   struct {
			href
			Templated bool `json:"templated"`
		} `json:"prices"`
	} `json:"_links"`
}

type localization struct {
	Embedded struct {
		Language struct {
			Code string `json:"code"`
			name
		} `json:"language"`
		LocalizationScope struct {
			Type string `json:"type"`
		} `json:"localizationScope"`
	} `json:"_embedded"`
}

type tag struct {
	Id int `json:"id"`
	name
	Level int `json:"level"`
	slug
}

type property struct {
	name
	slug
}

type screenshot struct {
	Links struct {
		Self struct {
			href
			Templated  bool     `json:"templated"`
			Formatters []string `json:"formatters"`
		} `json:"self"`
	} `json:"_links"`
}

type esrbRating struct {
	Category struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"category"`
	ContentDescriptors []struct {
		Descriptor string `json:"descriptor"`
		Id         int    `json:"id"`
	} `json:"contentDescriptors"`
}

type pegiRating struct {
	AgeRating          int `json:"ageRating"`
	ContentDescriptors []struct {
		Descriptor string `json:"descriptor"`
		Id         int    `json:"id"`
	} `json:"contentDescriptors"`
}

type uskRating struct {
	AgeRating int `json:"ageRating"`
	Category  struct {
		Name string `json:"name"`
	} `json:"category"`
}

type brRating struct {
	AgeRating          int `json:"ageRating"`
	ContentDescriptors []struct {
		Descriptor string `json:"descriptor"`
	} `json:"contentDescriptors"`
}

type bonus struct {
	Id int `json:"id"`
	name
	Type struct {
		name
		slug
	} `json:"type"`
}

type video struct {
	Provider    string `json:"provider"`
	VideoId     string `json:"videoId"`
	ThumbnailId string `json:"thumbnailId"`
	Links       struct {
		Self      href `json:"self"`
		Thumbnail href `json:"thumbnail"`
	} `json:"_links"`
}

type ApiProduct struct {
	InDevelopment struct {
		Active bool `json:"active"`
	} `json:"inDevelopment"`
	Copyrights             string `json:"copyrights"`
	IsUsingDosBox          bool   `json:"isUsingDosBox"`
	Description            string `json:"description"`
	Size                   int    `json:"size"`
	Overview               string `json:"overview"`
	FeaturesDescription    string `json:"featuresDescription"`
	AdditionalRequirements string `json:"additionalRequirements"`
	Links                  links  `json:"_links"`
	Embedded               struct {
		Product                   product                    `json:"product"`
		ProductType               string                     `json:"productType"`
		Localizations             []localization             `json:"localizations"`
		Videos                    []video                    `json:"videos"`
		Bonuses                   []bonus                    `json:"bonuses"`
		Screenshots               []screenshot               `json:"screenshots"`
		Publisher                 name                       `json:"publisher"`
		Developers                []name                     `json:"developers"`
		SupportedOperatingSystems []supportedOperatingSystem `json:"supportedOperatingSystems"`
		Tags                      []tag                      `json:"tags"`
		Properties                []property                 `json:"properties"`
		ESRBRating                esrbRating                 `json:"esrbRating"`
		PEGIRating                pegiRating                 `json:"pegiRating"`
		USKRating                 uskRating                  `json:"uskRating"`
		BRRating                  brRating                   `json:"brRating"`
		Features                  []struct {
			Id string `json:"id"`
			name
		} `json:"features"`
		Editions []edition `json:"editions"`
		Series   struct {
			Id int `json:"id"`
			name
		} `json:"series"`
	} `json:"_embedded"`
}

func (ap *ApiProduct) GetId() int { return ap.Embedded.Product.Id }

func (ap *ApiProduct) GetTitle() string {
	return ap.Embedded.Product.Title
}

func (ap *ApiProduct) GetDevelopers() []string {
	devs := make([]string, 0, len(ap.Embedded.Developers))
	for _, dev := range ap.Embedded.Developers {
		devs = append(devs, dev.Name)
	}
	return devs
}

func (ap *ApiProduct) GetPublishers() []string {
	return []string{ap.Embedded.Publisher.Name}
}

func (ap *ApiProduct) GetImage() string {
	return ap.Embedded.Product.Links.Image.Href
}

func (ap *ApiProduct) GetVerticalImage() string {
	return ap.Links.BoxArtImage.Href
}

func (ap *ApiProduct) GetHero() string {
	return ap.Links.GalaxyBackgroundImage.Href
}

func (ap *ApiProduct) GetLogo() string {
	return ap.Links.Logo.Href
}

func (ap *ApiProduct) GetIcon() string { return ap.Links.Icon.Href }

func (ap *ApiProduct) GetIconSquare() string { return ap.Links.IconSquare.Href }

func (ap *ApiProduct) GetBackground() string { return ap.Links.BackgroundImage.Href }

func (ap *ApiProduct) GetScreenshots() []string {
	screenshots := make([]string, 0)
	for _, screenshot := range ap.Embedded.Screenshots {
		screenshots = append(screenshots, screenshot.Links.Self.Href)
	}
	return screenshots
}

func idFromHref(href string) string {
	_, urlPath := path.Split(href)
	parts := strings.Split(urlPath, "?")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

func (ap *ApiProduct) getHrefs(hrefProperty []href) []string {
	hrefs := make([]string, 0, len(hrefProperty))
	for _, hp := range hrefProperty {
		id := idFromHref(hp.Href)
		if id != "" {
			hrefs = append(hrefs, id)
		}
	}
	return hrefs
}

func (ap *ApiProduct) GetGenres() []string {
	genres := make([]string, 0, len(ap.Embedded.Tags))
	for _, tag := range ap.Embedded.Tags {
		genres = append(genres, tag.Name)
	}
	return genres
}

func (ap *ApiProduct) GetFeatures() []string {
	features := make([]string, 0, len(ap.Embedded.Features))
	for _, feature := range ap.Embedded.Features {
		features = append(features, feature.Name)
	}
	return features
}

func (ap *ApiProduct) GetSeries() string {
	return ap.Embedded.Series.Name
}

func (ap *ApiProduct) GetVideoIds() []string {
	videoIds := make([]string, 0, len(ap.Embedded.Videos))
	for _, vid := range ap.Embedded.Videos {
		if vid.Provider != supportedVideoProvider {
			continue
		}
		//1733288799 had video_id specified as "v=...."
		vId := strings.TrimPrefix(vid.VideoId, "v=")
		videoIds = append(videoIds, vId)
	}
	return videoIds
}

func (ap *ApiProduct) GetOperatingSystems() []string {
	os := make([]string, 0, 3)
	for _, sos := range ap.Embedded.SupportedOperatingSystems {
		switch sos.OperatingSystem.Name {
		case windows:
			os = append(os, windows)
		case osx:
			// GOG.com continues to use "osx" name for macOS
			os = append(os, macos)
		case linux:
			os = append(os, linux)
		}
	}
	return os
}

func (ap *ApiProduct) GetIncludesGames() []string {
	return ap.getHrefs(ap.Links.IncludesGames)
}

func (ap *ApiProduct) GetIsIncludedInGames() []string {
	return ap.getHrefs(ap.Links.IsIncludedInGames)
}

func (ap *ApiProduct) GetIsRequiredByGames() []string {
	return ap.getHrefs(ap.Links.IsRequiredByGames)
}

func (ap *ApiProduct) GetRequiresGames() []string {
	return ap.getHrefs(ap.Links.RequiresGames)
}

func (ap *ApiProduct) GetLanguageCodes() []string {
	lcs := make(map[string]bool, 0)
	for _, loc := range ap.Embedded.Localizations {
		lcs[loc.Embedded.Language.Code] = true
	}

	codes := make([]string, 0, len(lcs))
	for code, _ := range lcs {
		codes = append(codes, code)
	}
	return codes
}

func (ap *ApiProduct) GetGlobalRelease() string {
	grd := ap.Embedded.Product.GlobalReleaseDate
	if grd.IsZero() {
		return ""
	}
	if lw, err := time.LoadLocation("Europe/Nicosia"); err == nil {
		tiw := grd.In(lw)
		return tiw.Format("2006.01.02")
	}
	return grd.Format("2006.01.02")
}

func (ap *ApiProduct) GetGOGRelease() string {
	grd := ap.Embedded.Product.GogReleaseDate
	if grd.IsZero() {
		return ""
	}
	//GOG.com uses "1991-12-25T00:00:00+02:00" as "not released" or "unknown release date".
	//GOG.com uses "2000-12-31T23:59:59+02:00" for games that were removed from the store.
	//Considering GOG was founded in 2008 - neither is a real date for practical purposes.
	if grd.Unix() == 693612000 ||
		grd.Unix() == 978299999 {
		return ""
	}

	if lw, err := time.LoadLocation("Europe/Nicosia"); err == nil {
		tiw := ap.Embedded.Product.GogReleaseDate.In(lw)
		return tiw.Format("2006.01.02")
	}
	return grd.Format("2006.01.02")
}

func urlPathFromLink(link string) string {
	if u, err := url.Parse(link); err != nil {
		return ""
	} else {
		return u.Path
	}
}

func (ap *ApiProduct) GetStoreUrl() string {
	return urlPathFromLink(ap.Links.Store.Href)
}

func (ap *ApiProduct) GetForumUrl() string {
	return urlPathFromLink(ap.Links.Forum.Href)
}

func (ap *ApiProduct) GetSupportUrl() string {
	return urlPathFromLink(ap.Links.Support.Href)
}

func (ap *ApiProduct) GetDescriptionOverview() string {
	return ap.Overview
}

func (ap *ApiProduct) GetDescriptionFeatures() string {
	return ap.FeaturesDescription
}

func (ap *ApiProduct) GetProductType() string {
	return ap.Embedded.ProductType
}

func (ap *ApiProduct) GetCopyrights() string {
	return ap.Copyrights
}

// GetStoreTags returns values displayed in Store as "Tags".
// ApiProduct has store tags in Properties
func (ap *ApiProduct) GetStoreTags() []string {
	properties := make([]string, 0, len(ap.Embedded.Properties))
	for _, p := range ap.Embedded.Properties {
		properties = append(properties, p.Name)
	}
	return properties
}

func (ap *ApiProduct) GetAdditionalRequirements() string {
	return ap.AdditionalRequirements
}

func (ap *ApiProduct) GetInDevelopment() bool {
	return ap.InDevelopment.Active
}

func (ap *ApiProduct) GetPreOrder() bool {
	return ap.Embedded.Product.IsPreorder
}
