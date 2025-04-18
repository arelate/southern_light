package vangogh_integration

type DownloadsLayout int

const (
	UnknownDownloadsLayout DownloadsLayout = iota
	ShardedDownloadsLayout                 // order is important here given this will be used for clo default parameter
	FlatDownloadsLayout
)

const DefaultDownloadsLayout = FlatDownloadsLayout

var downloadsLayoutsStrings = map[DownloadsLayout]string{
	UnknownDownloadsLayout: "unknown",
	FlatDownloadsLayout:    "flat",
	ShardedDownloadsLayout: "sharded",
}

func (dl DownloadsLayout) String() string {
	if dls, ok := downloadsLayoutsStrings[dl]; ok {
		return dls
	}
	return downloadsLayoutsStrings[UnknownDownloadsLayout]
}

func ParseDownloadsLayout(downloadsLayout string) DownloadsLayout {
	for dl, dls := range downloadsLayoutsStrings {
		if dls == downloadsLayout {
			return dl
		}
	}
	return UnknownDownloadsLayout
}

func AllDownloadsLayouts() []DownloadsLayout {
	dls := make([]DownloadsLayout, 0, len(downloadsLayoutsStrings)-1)
	dls = append(dls, DefaultDownloadsLayout)
	for dl := range downloadsLayoutsStrings {
		if dl == UnknownDownloadsLayout ||
			dl == DefaultDownloadsLayout {
			continue
		}
		dls = append(dls, dl)
	}
	return dls
}
