package vangogh_integration

type ManualUrlStatus int

const (
	ManualUrlStatusUnknown ManualUrlStatus = iota
	ManualUrlQueued
	ManualUrlDownloading
	ManualUrlDownloadInterrupted
	ManualUrlDownloaded
	ManualUrlValidated
)

var manualUrlStatusStrings = map[ManualUrlStatus]string{
	ManualUrlStatusUnknown:       "unknown",
	ManualUrlQueued:              "download-queued",
	ManualUrlDownloading:         "downloading",
	ManualUrlDownloadInterrupted: "download-interrupted",
	ManualUrlDownloaded:          "downloaded",
	ManualUrlValidated:           "validated",
}

var manualUrlStatusHumanReadableStrings = map[ManualUrlStatus]string{
	ManualUrlStatusUnknown:       "Unknown",
	ManualUrlQueued:              "Download Queued",
	ManualUrlDownloading:         "Downloading",
	ManualUrlDownloadInterrupted: "Download Interrupted",
	ManualUrlDownloaded:          "Downloaded",
	ManualUrlValidated:           "Validated",
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
