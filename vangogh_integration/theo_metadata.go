package vangogh_integration

import "golang.org/x/exp/slices"

type TheoMetadata struct {
	Id            string            `json:"id"`
	Slug          string            `json:"slug"`
	Title         string            `json:"title"`
	Images        TheoImages        `json:"images"`
	DownloadLinks TheoDownloadLinks `json:"download-links,omitempty"`
}

type TheoImages struct {
	VerticalImage string `json:"vertical-image,omitempty"`
	Image         string `json:"image,omitempty"`
	Hero          string `json:"hero,omitempty"`
	Logo          string `json:"logo,omitempty"`
	Icon          string `json:"icon,omitempty"`
	IconSquare    string `json:"icon-square,omitempty"`
}

type TheoDownloadLink struct {
	ManualUrl        string `json:"manual-url"`
	Name             string `json:"name"`
	Status           string `json:"status"`
	ValidationResult string `json:"validation-result"`
	LocalFilename    string `json:"local-filename"`
	Md5              string `json:"md5"`
	OS               string `json:"os"`
	Type             string `json:"type"`
	LanguageCode     string `json:"language-code"`
	Version          string `json:"version"`
	EstimatedBytes   int    `json:"estimated-bytes"`
}

type TheoDownloadLinks []TheoDownloadLink

func (tdl TheoDownloadLinks) FilterOperatingSystems(operatingSystems ...OperatingSystem) TheoDownloadLinks {
	if len(operatingSystems) == 0 {
		return tdl
	}

	if slices.Contains(operatingSystems, AnyOperatingSystem) {
		return tdl
	}

	filteredLinks := make(TheoDownloadLinks, 0)
	for _, dl := range tdl {
		os := ParseOperatingSystem(dl.OS)
		if os == AnyOperatingSystem || slices.Contains(operatingSystems, os) {
			filteredLinks = append(filteredLinks, dl)
		}
	}

	return filteredLinks
}

func (tdl TheoDownloadLinks) FilterLanguageCodes(languageCodes ...string) TheoDownloadLinks {
	if len(languageCodes) == 0 {
		return tdl
	}

	filteredLinks := make(TheoDownloadLinks, 0)
	for _, dl := range tdl {
		if dl.LanguageCode == "" || slices.Contains(languageCodes, dl.LanguageCode) {
			filteredLinks = append(filteredLinks, dl)
		}
	}

	return filteredLinks
}

func (tdl TheoDownloadLinks) FilterDownloadTypes(downloadTypes ...DownloadType) TheoDownloadLinks {
	if len(downloadTypes) == 0 {
		return tdl
	}

	if slices.Contains(downloadTypes, AnyDownloadType) {
		return tdl
	}

	filteredLinks := make(TheoDownloadLinks, 0)
	for _, dl := range tdl {
		dt := ParseDownloadType(dl.Type)
		if dt == AnyDownloadType || slices.Contains(downloadTypes, dt) {
			filteredLinks = append(filteredLinks, dl)
		}
	}

	return filteredLinks
}
