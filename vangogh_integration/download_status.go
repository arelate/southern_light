package vangogh_integration

type DownloadStatus int

const (
	DownloadStatusValidated DownloadStatus = iota
	DownloadStatusDownloaded
	DownloadStatusUnknown
	DownloadStatusQueued
	DownloadStatusDownloading
	DownloadStatusInterrupted
)

var downloadStatusStrings = map[DownloadStatus]string{
	DownloadStatusValidated:   "validated",
	DownloadStatusDownloaded:  "downloaded",
	DownloadStatusUnknown:     "unknown",
	DownloadStatusDownloading: "downloading",
	DownloadStatusQueued:      "download-queued",
	DownloadStatusInterrupted: "download-interrupted",
}

var downloadStatusHumanReadableStrings = map[DownloadStatus]string{
	DownloadStatusValidated:   "Validated",
	DownloadStatusDownloaded:  "Downloaded",
	DownloadStatusUnknown:     "Unknown",
	DownloadStatusDownloading: "Downloading",
	DownloadStatusQueued:      "Download Queued",
	DownloadStatusInterrupted: "Download Interrupted",
}

func (ds DownloadStatus) String() string {
	if muss, ok := downloadStatusStrings[ds]; ok {
		return muss
	}
	return downloadStatusStrings[DownloadStatusUnknown]
}

func (ds DownloadStatus) HumanReadableString() string {
	return downloadStatusHumanReadableStrings[ds]
}

func ParseDownloadStatus(dss string) DownloadStatus {
	for ds, str := range downloadStatusStrings {
		if str == dss {
			return ds
		}
	}
	return DownloadStatusUnknown
}
