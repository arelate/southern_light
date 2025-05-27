package vangogh_integration

import "slices"

type ProductDetails struct {
	Id               string               `json:"id"`
	Slug             string               `json:"slug"`
	SteamAppId       string               `json:"steam-app-id,omitempty"`
	Title            string               `json:"title"`
	ProductType      string               `json:"product-type"`
	OperatingSystems []OperatingSystem    `json:"operating-systems"`
	Developers       []string             `json:"developers"`
	Publishers       []string             `json:"publishers"`
	Images           ProductImages        `json:"images"`
	DownloadLinks    ProductDownloadLinks `json:"download-links,omitempty"`
	IncludesGames    []string             `json:"includes-games,omitempty"`
	RequiresGames    []string             `json:"requires-games,omitempty"`
}

type ProductImages struct {
	VerticalImage string `json:"vertical-image,omitempty"`
	Image         string `json:"image,omitempty"`
	Hero          string `json:"hero,omitempty"`
	Logo          string `json:"logo,omitempty"`
	Icon          string `json:"icon,omitempty"`
	IconSquare    string `json:"icon-square,omitempty"`
	Background    string `json:"background,omitempty"`
}

type ProductDownloadLink struct {
	ManualUrl        string           `json:"manual-url"`
	Name             string           `json:"name"`
	Status           string           `json:"status"`
	ValidationResult ValidationResult `json:"validation-result"`
	LocalFilename    string           `json:"local-filename"`
	Md5              string           `json:"md5"`
	OperatingSystem  OperatingSystem  `json:"os"`
	Type             DownloadType     `json:"type"`
	LanguageCode     string           `json:"language-code"`
	Version          string           `json:"version"`
	EstimatedBytes   int              `json:"estimated-bytes"`
}

type ProductDownloadLinks []ProductDownloadLink

func (mdl ProductDownloadLinks) FilterOperatingSystems(operatingSystems ...OperatingSystem) ProductDownloadLinks {
	if len(operatingSystems) == 0 {
		return mdl
	}

	if slices.Contains(operatingSystems, AnyOperatingSystem) {
		return mdl
	}

	filteredLinks := make(ProductDownloadLinks, 0)
	for _, dl := range mdl {
		if dl.OperatingSystem == AnyOperatingSystem || slices.Contains(operatingSystems, dl.OperatingSystem) {
			filteredLinks = append(filteredLinks, dl)
		}
	}

	return filteredLinks
}

func (mdl ProductDownloadLinks) FilterLanguageCodes(languageCodes ...string) ProductDownloadLinks {
	if len(languageCodes) == 0 {
		return mdl
	}

	filteredLinks := make(ProductDownloadLinks, 0)
	for _, dl := range mdl {
		if dl.LanguageCode == "" || slices.Contains(languageCodes, dl.LanguageCode) {
			filteredLinks = append(filteredLinks, dl)
		}
	}

	return filteredLinks
}

func (mdl ProductDownloadLinks) FilterDownloadTypes(downloadTypes ...DownloadType) ProductDownloadLinks {
	if len(downloadTypes) == 0 {
		return mdl
	}

	if slices.Contains(downloadTypes, AnyDownloadType) {
		return mdl
	}

	filteredLinks := make(ProductDownloadLinks, 0)
	for _, dl := range mdl {
		if dl.Type == AnyDownloadType || slices.Contains(downloadTypes, dl.Type) {
			filteredLinks = append(filteredLinks, dl)
		}
	}

	return filteredLinks
}
