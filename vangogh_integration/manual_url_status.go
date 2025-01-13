package vangogh_integration

type ManualUrlStatus int

const (
	ManualUrlStatusUnknown ManualUrlStatus = iota
	ManualUrlQueued
	ManualUrlDownloaded
	ManualUrlValidated
)

var manualUrlStatusStrings = map[ManualUrlStatus]string{
	ManualUrlStatusUnknown: "unknown",
	ManualUrlQueued:        "queued",
	ManualUrlDownloaded:    "downloaded",
	ManualUrlValidated:     "validated",
}

func (mus ManualUrlStatus) String() string {
	if muss, ok := manualUrlStatusStrings[mus]; ok {
		return muss
	}
	return manualUrlStatusStrings[ManualUrlStatusUnknown]
}

func ParseManualUrlStatus(muss string) ManualUrlStatus {
	for mus, str := range manualUrlStatusStrings {
		if str == muss {
			return mus
		}
	}
	return ManualUrlStatusUnknown
}
