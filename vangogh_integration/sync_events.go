package vangogh_integration

const (
	SyncUnknownKey               = "sync-unknown"
	SyncInterruptedKey           = "sync-interrupted"
	SyncStartKey                 = "sync-start"
	SyncAccountDataKey           = "sync-account-data"
	SyncAccountImagesKey         = "sync-account-images"
	SyncDownloadsKey             = "sync-downloads"
	SyncGenerateMissingChecksums = "sync-generate-missing-checksums"
	SyncCleanupKey               = "sync-cleanup"
	SyncVideoMetadataKey         = "sync-video-metadata"
	SyncBinaries                 = "sync-binaries"
	SyncAdditionalData           = "sync-additional-data"
	SyncAdditionalImagesKey      = "sync-additional-images"
	SyncDescriptionImagesKey     = "sync-description-images"
	SyncBackup                   = "sync-backup"
	SyncCompleteKey              = "sync-complete"
)

var SyncEventsSequence = []string{
	SyncInterruptedKey,
	SyncStartKey,
	SyncAccountDataKey,
	SyncAccountImagesKey,
	SyncDownloadsKey,
	SyncGenerateMissingChecksums,
	SyncCleanupKey,
	SyncVideoMetadataKey,
	SyncBinaries,
	SyncAdditionalData,
	SyncAdditionalImagesKey,
	SyncDescriptionImagesKey,
	SyncBackup,
	SyncCompleteKey, // this should be the last key
}

func NextSyncEvent(completedEvent string) string {
	if completedEvent == SyncCompleteKey {
		return SyncCompleteKey
	}
	for ii, event := range SyncEventsSequence {
		if event == completedEvent {
			if ii == 0 || ii == len(SyncEventsSequence)-1 {
				return SyncUnknownKey
			}
			return SyncEventsSequence[ii+1]
		}
	}
	return SyncUnknownKey
}

var SyncEventsTitles = map[string]string{
	SyncUnknownKey:               "Unknown",
	SyncInterruptedKey:           "Interrupted",
	SyncStartKey:                 "Started sync",
	SyncAccountDataKey:           "Updating account data",
	SyncAccountImagesKey:         "Updating account images",
	SyncDownloadsKey:             "Downloading files",
	SyncGenerateMissingChecksums: "Generating missing checksums",
	SyncCleanupKey:               "Cleaning up downloads",
	SyncVideoMetadataKey:         "Updating video titles",
	SyncBinaries:                 "Updating binaries",
	SyncAdditionalData:           "Updating additional data",
	SyncAdditionalImagesKey:      "Updating additional images",
	SyncDescriptionImagesKey:     "Updating descriptions",
	SyncBackup:                   "Backing up data",
	SyncCompleteKey:              "Sync complete", // this should be the last key
}
