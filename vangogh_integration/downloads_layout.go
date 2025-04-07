package vangogh_integration

type DownloadsLayout int

const (
	UnknownDownloadsLayout DownloadsLayout = iota
	ShardedDownloadsLayout                 // order is important here given this will be used for clo default parameter

	FlatDownloadsLayout
)

const DefaultDownloadsLayout = ShardedDownloadsLayout

var downloadsLayoutsStrings = map[DownloadsLayout]string{
	UnknownDownloadsLayout: "unknown",
	ShardedDownloadsLayout: "sharded",
	FlatDownloadsLayout:    "flat",
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
	for dl := range downloadsLayoutsStrings {
		if dl == UnknownDownloadsLayout {
			continue
		}
		dls = append(dls, dl)
	}
	return dls
}
