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

var CurrentSyncEventsTitles = map[string]string{
	SyncInterruptedKey:       "Interrupted",
	SyncStartKey:             "Started",
	SyncDataKey:              "Updating data",
	SyncDescriptionImagesKey: "Updating description images",
	SyncImagesKey:            "Updating images",
	SyncDehydrateKey:         "Dehydrating images",
	SyncVideoMetadataKey:     "Updating video metadata",
	SyncDownloadsKey:         "Downloading and validating files",
	SyncCleanupKey:           "Cleaning up downloads",
	SyncWineBinaries:         "Updating WINE binaries",
	SyncBackup:               "Backing up data",
	SyncCompleteKey:          "Completed", // this should be the last key
}
