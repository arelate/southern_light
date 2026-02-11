package steam_grid

type Asset int

const (
	UnknownAsset Asset = iota

	ClientIcon   // 256x256, ar: 1
	ClientTga    // 16x16, ar: 1
	Icon         // 32x32, ar: 1
	Logo         // 184x69, ar: 2.66
	LogoSmall    // 120x45, ar: 2.66
	SmallCapsule // 231x87, ar: 2.66
	Header       // 460x215
	Header2x     // 920x430

	// library_assets_full

	LibraryCapsule   // 300x450
	LibraryCapsule2x // 600x900
	LibraryHero      // 1920x620
	LibraryHero2x    // 3840x1240
	LibraryLogo      // 640x360
	LibraryLogo2x    // 1280x720

	// assets

	MainCapsule       // 616x353
	PageBackground    // 1438x810
	HeroCapsule       // 374x448
	CommunityIcon     // 32x32
	RawPageBackground // 1438x810
)
