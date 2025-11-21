package vangogh_integration

const (
	UpdatesInstallers    = "files"
	UpdatesNewProducts   = "store"
	UpdatesReleasedToday = "today"
	UpdatesSteamNews     = "steam"
)

var UpdatesShorterTitles = map[string]string{
	UpdatesInstallers:    "Files",
	UpdatesNewProducts:   "Store",
	UpdatesReleasedToday: "That day",
	UpdatesSteamNews:     "Steam",
}

var UpdatesLongerTitles = map[string]string{
	UpdatesInstallers:    "Updated Installers",
	UpdatesNewProducts:   "New in Store",
	UpdatesReleasedToday: "Released on {month} {day}",
	UpdatesSteamNews:     "Steam News",
}

var UpdatesOrder = []string{
	UpdatesInstallers,
	UpdatesNewProducts,
	UpdatesReleasedToday,
	UpdatesSteamNews,
}
