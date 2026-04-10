package vangogh_integration

const (
	SyncInterruptedKey           = "sync-interrupted"
	SyncStartKey                 = "sync-start"
	SyncPurchasesDataKey         = "sync-purchases-data"
	SyncDescriptionImagesKey     = "sync-description-images"
	SyncImagesKey                = "sync-images"
	SyncVideoMetadataKey         = "sync-video-metadata"
	SyncDownloadsKey             = "sync-downloads"
	SyncGenerateMissingChecksums = "sync-generate-missing-checksums"
	SyncCleanupKey               = "sync-cleanup"
	SyncBinaries                 = "sync-binaries"
	SyncExtraData                = "sync-extra-data"
	SyncBackup                   = "sync-backup"
	SyncCompleteKey              = "sync-complete"
)

var SyncEventsKeys = []string{
	SyncInterruptedKey,
	SyncStartKey,
	SyncPurchasesDataKey,
	SyncDescriptionImagesKey,
	SyncImagesKey,
	SyncVideoMetadataKey,
	SyncDownloadsKey,
	SyncGenerateMissingChecksums,
	SyncCleanupKey,
	SyncBinaries,
	SyncExtraData,
	SyncBackup,
	SyncCompleteKey, // this should be the last key
}

var CurrentSyncEventForCompleted = map[string]string{
	SyncStartKey:                 SyncPurchasesDataKey,
	SyncPurchasesDataKey:         SyncDescriptionImagesKey,
	SyncDescriptionImagesKey:     SyncImagesKey,
	SyncImagesKey:                SyncVideoMetadataKey,
	SyncVideoMetadataKey:         SyncDownloadsKey,
	SyncDownloadsKey:             SyncGenerateMissingChecksums,
	SyncGenerateMissingChecksums: SyncCleanupKey,
	SyncCleanupKey:               SyncBinaries,
	SyncBinaries:                 SyncExtraData,
	SyncExtraData:                SyncBackup,
}

var CurrentSyncEventsTitles = map[string]string{
	SyncInterruptedKey:           "Interrupted",
	SyncStartKey:                 "Started sync",
	SyncPurchasesDataKey:         "Updating account data",
	SyncDescriptionImagesKey:     "Updating descriptions",
	SyncImagesKey:                "Updating images",
	SyncVideoMetadataKey:         "Updating videos",
	SyncDownloadsKey:             "Downloading files",
	SyncGenerateMissingChecksums: "Generating missing checksums",
	SyncCleanupKey:               "Cleaning up downloads",
	SyncBinaries:                 "Updating binaries",
	SyncExtraData:                "Updating extra data",
	SyncBackup:                   "Backing up data",
	SyncCompleteKey:              "Sync complete", // this should be the last key
}
