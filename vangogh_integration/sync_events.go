package vangogh_integration

const (
	SyncStartKey    = "sync-start"
	SyncCompleteKey = "sync-complete"
)

var SyncEventsTitles = map[string]string{
	SyncStartKey:    "Sync Started",
	SyncCompleteKey: "Sync Completed",
}
