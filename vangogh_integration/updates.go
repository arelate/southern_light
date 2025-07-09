package vangogh_integration

const (
	UpdatesInstallers    = "files"
	UpdatesReleasedToday = "today"
	UpdatesNewProducts   = "store"
	UpdatesSteamNews     = "steam"
)

var UpdatesShorterTitles = map[string]string{
	UpdatesInstallers:    "Files",
	UpdatesReleasedToday: "Today",
	UpdatesNewProducts:   "New",
	UpdatesSteamNews:     "Steam",
}

var UpdatesLongerTitles = map[string]string{
	UpdatesInstallers:    "Updated Installers",
	UpdatesReleasedToday: "Released Today",
	UpdatesNewProducts:   "New in Store",
	UpdatesSteamNews:     "Steam News",
}
