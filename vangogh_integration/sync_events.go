package vangogh_integration

const (
	SyncInterruptedKey       = "sync-interrupted"
	SyncStartKey             = "sync-start"
	SyncDataKey              = "sync-data"
	SyncDescriptionImagesKey = "sync-description-images"
	SyncImagesKey            = "sync-images"
	SyncDehydrateKey         = "sync-dehydrate"
	SyncVideoMetadataKey     = "sync-video-metadata"
	SyncDownloadsKey         = "sync-downloads"
	SyncCleanupKey           = "sync-cleanup"
	SyncWineBinaries         = "sync-wine-binaries"
	SyncBackup               = "sync-backup"
	SyncCompleteKey          = "sync-complete"
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
	SyncCleanupKey,
	SyncWineBinaries,
	SyncBackup,
	SyncCompleteKey, // this should be the last key
}

var CurrentSyncEventForCompleted = map[string]string{
	SyncStartKey:             SyncDataKey,
	SyncDataKey:              SyncDescriptionImagesKey,
	SyncDescriptionImagesKey: SyncImagesKey,
	SyncImagesKey:            SyncDehydrateKey,
	SyncDehydrateKey:         SyncVideoMetadataKey,
	SyncVideoMetadataKey:     SyncDownloadsKey,
	SyncDownloadsKey:         SyncCleanupKey,
	SyncCleanupKey:           SyncWineBinaries,
	SyncWineBinaries:         SyncBackup,
}

var CurrentSyncEventsTitles = map[string]string{
	SyncInterruptedKey:       "Interrupted",
	SyncStartKey:             "Started sync",
	SyncDataKey:              "Updating data",
	SyncDescriptionImagesKey: "Updating description images",
	SyncImagesKey:            "Updating images",
	SyncDehydrateKey:         "Dehydrating images",
	SyncVideoMetadataKey:     "Updating video metadata",
	SyncDownloadsKey:         "Downloading and validating files",
	SyncCleanupKey:           "Cleaning up downloads",
	SyncWineBinaries:         "Updating WINE binaries",
	SyncBackup:               "Backing up data",
	SyncCompleteKey:          "OK", // this should be the last key
}
