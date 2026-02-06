package vangogh_integration

const (
	SyncInterruptedKey           = "sync-interrupted"
	SyncStartKey                 = "sync-start"
	SyncDataKey                  = "sync-data"
	SyncDescriptionImagesKey     = "sync-description-images"
	SyncImagesKey                = "sync-images"
	SyncDehydrateKey             = "sync-dehydrate"
	SyncVideoMetadataKey         = "sync-video-metadata"
	SyncDownloadsKey             = "sync-downloads"
	SyncGenerateMissingChecksums = "sync-generate-missing-checksums"
	SyncCleanupKey               = "sync-cleanup"
	SyncBinaries                 = "sync-binaries"
	SyncBackup                   = "sync-backup"
	SyncCompleteKey              = "sync-complete"
)

var SyncEventsKeys = []string{
	SyncInterruptedKey,
	SyncStartKey,
	SyncDataKey,
	SyncDescriptionImagesKey,
	SyncImagesKey,
	SyncDehydrateKey,
	SyncVideoMetadataKey,
	SyncDownloadsKey,
	SyncGenerateMissingChecksums,
	SyncCleanupKey,
	SyncBinaries,
	SyncBackup,
	SyncCompleteKey, // this should be the last key
}

var CurrentSyncEventForCompleted = map[string]string{
	SyncStartKey:                 SyncDataKey,
	SyncDataKey:                  SyncDescriptionImagesKey,
	SyncDescriptionImagesKey:     SyncImagesKey,
	SyncImagesKey:                SyncDehydrateKey,
	SyncDehydrateKey:             SyncVideoMetadataKey,
	SyncVideoMetadataKey:         SyncDownloadsKey,
	SyncDownloadsKey:             SyncGenerateMissingChecksums,
	SyncGenerateMissingChecksums: SyncCleanupKey,
	SyncCleanupKey:               SyncBinaries,
	SyncBinaries:                 SyncBackup,
}

var CurrentSyncEventsTitles = map[string]string{
	SyncInterruptedKey:           "Interrupted",
	SyncStartKey:                 "Started sync",
	SyncDataKey:                  "Updating data",
	SyncDescriptionImagesKey:     "Updating descriptions",
	SyncImagesKey:                "Updating images",
	SyncDehydrateKey:             "Processing images",
	SyncVideoMetadataKey:         "Updating videos",
	SyncDownloadsKey:             "Downloading files",
	SyncGenerateMissingChecksums: "Generating missing checksums",
	SyncCleanupKey:               "Cleaning up downloads",
	SyncBinaries:                 "Updating binaries",
	SyncBackup:                   "Backing up data",
	SyncCompleteKey:              "Sync complete", // this should be the last key
}
