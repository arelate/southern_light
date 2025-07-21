package vangogh_integration

type DownloadType int

const (
	AnyDownloadType DownloadType = iota
	Installer
	Movie
	DLC
	Extra
)

var downloadTypeStrings = map[DownloadType]string{
	AnyDownloadType: "any-download-type",
	Installer:       "installer",
	Movie:           "movie",
	DLC:             "downloadable-content",
	Extra:           "extra",
}

func AllDownloadTypes() []DownloadType {
	all := make([]DownloadType, 1, len(downloadTypeStrings))
	// order is important here given this will be used for clo default parameter
	all[0] = AnyDownloadType
	for dt, _ := range downloadTypeStrings {
		if dt == AnyDownloadType {
			continue
		}
		all = append(all, dt)
	}
	return all
}

func (dt DownloadType) String() string {
	str, ok := downloadTypeStrings[dt]
	if ok {
		return str
	}

	return downloadTypeStrings[AnyDownloadType]
}

func ParseDownloadType(downloadType string) DownloadType {
	for dt, str := range downloadTypeStrings {
		if str == downloadType {
			return dt
		}
	}
	return AnyDownloadType
}

func ParseManyDownloadTypes(dtStrings []string) []DownloadType {
	downloadTypes := make([]DownloadType, 0, len(dtStrings))
	for _, dtStr := range dtStrings {
		dt := ParseDownloadType(dtStr)
		downloadTypes = append(downloadTypes, dt)
	}
	return downloadTypes
}

func DownloadTypesCloValues() []string {
	dts := AllDownloadTypes()
	dtsStr := make([]string, 0, len(dts))
	for _, dt := range dts {
		dtsStr = append(dtsStr, dt.String())
	}
	return dtsStr
}
