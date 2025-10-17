package vangogh_integration

type ManualUrlStatus int

const (
	ManualUrlValidated ManualUrlStatus = iota
	ManualUrlDownloaded
	ManualUrlStatusUnknown
	ManualUrlQueued
	ManualUrlDownloading
	ManualUrlDownloadInterrupted
)

var manualUrlStatusStrings = map[ManualUrlStatus]string{
	ManualUrlValidated:           "validated",
	ManualUrlDownloaded:          "downloaded",
	ManualUrlStatusUnknown:       "unknown",
	ManualUrlDownloading:         "downloading",
	ManualUrlQueued:              "download-queued",
	ManualUrlDownloadInterrupted: "download-interrupted",
}

var manualUrlStatusHumanReadableStrings = map[ManualUrlStatus]string{
	ManualUrlValidated:           "Validated",
	ManualUrlDownloaded:          "Downloaded",
	ManualUrlStatusUnknown:       "Unknown",
	ManualUrlDownloading:         "Downloading",
	ManualUrlQueued:              "Download Queued",
	ManualUrlDownloadInterrupted: "Download Interrupted",
}

func (mus ManualUrlStatus) String() string {
	if muss, ok := manualUrlStatusStrings[mus]; ok {
		return muss
	}
	return manualUrlStatusStrings[ManualUrlStatusUnknown]
}

func (mus ManualUrlStatus) HumanReadableString() string {
	return manualUrlStatusHumanReadableStrings[mus]
}

func ParseManualUrlStatus(muss string) ManualUrlStatus {
	for mus, str := range manualUrlStatusStrings {
		if str == muss {
			return mus
		}
	}
	return ManualUrlStatusUnknown
}
