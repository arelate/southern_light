package vangogh_integration

import "slices"

type DownloadsManifest struct {
	Id            string                `json:"id"`
	Slug          string                `json:"slug"`
	Title         string                `json:"title"`
	Images        ManifestImages        `json:"images"`
	DownloadLinks ManifestDownloadLinks `json:"download-links,omitempty"`
}

type ManifestImages struct {
	VerticalImage string `json:"vertical-image,omitempty"`
	Image         string `json:"image,omitempty"`
	Hero          string `json:"hero,omitempty"`
	Logo          string `json:"logo,omitempty"`
	Icon          string `json:"icon,omitempty"`
	IconSquare    string `json:"icon-square,omitempty"`
	Background    string `json:"background"`
}

type ManifestDownloadLink struct {
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

type ManifestDownloadLinks []ManifestDownloadLink

func (mdl ManifestDownloadLinks) FilterOperatingSystems(operatingSystems ...OperatingSystem) ManifestDownloadLinks {
	if len(operatingSystems) == 0 {
		return mdl
	}

	if slices.Contains(operatingSystems, AnyOperatingSystem) {
		return mdl
	}

	filteredLinks := make(ManifestDownloadLinks, 0)
	for _, dl := range mdl {
		os := ParseOperatingSystem(dl.OS)
		if os == AnyOperatingSystem || slices.Contains(operatingSystems, os) {
			filteredLinks = append(filteredLinks, dl)
		}
	}

	return filteredLinks
}

func (mdl ManifestDownloadLinks) FilterLanguageCodes(languageCodes ...string) ManifestDownloadLinks {
	if len(languageCodes) == 0 {
		return mdl
	}

	filteredLinks := make(ManifestDownloadLinks, 0)
	for _, dl := range mdl {
		if dl.LanguageCode == "" || slices.Contains(languageCodes, dl.LanguageCode) {
			filteredLinks = append(filteredLinks, dl)
		}
	}

	return filteredLinks
}

func (mdl ManifestDownloadLinks) FilterDownloadTypes(downloadTypes ...DownloadType) ManifestDownloadLinks {
	if len(downloadTypes) == 0 {
		return mdl
	}

	if slices.Contains(downloadTypes, AnyDownloadType) {
		return mdl
	}

	filteredLinks := make(ManifestDownloadLinks, 0)
	for _, dl := range mdl {
		dt := ParseDownloadType(dl.Type)
		if dt == AnyDownloadType || slices.Contains(downloadTypes, dt) {
			filteredLinks = append(filteredLinks, dl)
		}
	}

	return filteredLinks
}
