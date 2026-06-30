package vangogh_integration

const (
	SyncUnknownKey               = "sync-unknown"
	SyncInterruptedKey           = "sync-interrupted"
	SyncStartKey                 = "sync-start"
	SyncPurchasesDataKey         = "sync-purchases-data"
	SyncPurchasesImagesKey       = "sync-purchases-images"
	SyncDownloadsKey             = "sync-downloads"
	SyncGenerateMissingChecksums = "sync-generate-missing-checksums"
	SyncCleanupKey               = "sync-cleanup"
	SyncVideoMetadataKey         = "sync-video-metadata"
	SyncBinaries                 = "sync-binaries"
	SyncExtraData                = "sync-extra-data"
	SyncExtraImagesKey           = "sync-extra-images"
	SyncDescriptionImagesKey     = "sync-description-images"
	SyncBackup                   = "sync-backup"
	SyncCompleteKey              = "sync-complete"
)

var SyncEventsSequence = []string{
	SyncInterruptedKey,
	SyncStartKey,
	SyncPurchasesDataKey,
	SyncPurchasesImagesKey,
	SyncDownloadsKey,
	SyncGenerateMissingChecksums,
	SyncCleanupKey,
	SyncVideoMetadataKey,
	SyncBinaries,
	SyncExtraData,
	SyncExtraImagesKey,
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
	SyncPurchasesDataKey:         "Updating account data",
	SyncPurchasesImagesKey:       "Updating account images",
	SyncDownloadsKey:             "Downloading files",
	SyncGenerateMissingChecksums: "Generating missing checksums",
	SyncCleanupKey:               "Cleaning up downloads",
	SyncVideoMetadataKey:         "Updating video titles",
	SyncBinaries:                 "Updating binaries",
	SyncExtraData:                "Updating extra data",
	SyncExtraImagesKey:           "Updating extra images",
	SyncDescriptionImagesKey:     "Updating descriptions",
	SyncBackup:                   "Backing up data",
	SyncCompleteKey:              "Sync complete", // this should be the last key
}
