package vangogh_integration

const (
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
	SyncInterruptedKey       = "sync-interrupted"
)

var SyncEventsKeys = []string{
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
	SyncCompleteKey,
	SyncInterruptedKey,
}

var SyncEventsTitles = map[string]string{
	SyncStartKey:             "Started",
	SyncDataKey:              "Data updated",
	SyncDescriptionImagesKey: "Description images updated",
	SyncImagesKey:            "Images updated",
	SyncDehydrateKey:         "Images dehydrated",
	SyncVideoMetadataKey:     "Video metadata updated",
	SyncDownloadsKey:         "Downloads updated",
	SyncCleanupKey:           "Download cleaned up",
	SyncWineBinaries:         "WINE binaries updated",
	SyncBackup:               "Data backed up",
	SyncCompleteKey:          "Completed",
	SyncInterruptedKey:       "Interrupted",
}
