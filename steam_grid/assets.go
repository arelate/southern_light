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

var assetStrings = map[Asset]string{
	UnknownAsset:      "unknown",
	ClientIcon:        "clienticon",
	ClientTga:         "clienttga",
	Icon:              "icon",
	Logo:              "logo",
	LogoSmall:         "logosmall",
	SmallCapsule:      "smallcapsule",
	Header:            "header",
	Header2x:          "header2x",
	LibraryCapsule:    "librarycapsule",
	LibraryCapsule2x:  "librarycapsule2x",
	LibraryHero:       "libraryhero",
	LibraryHero2x:     "libraryhero2x",
	LibraryLogo:       "librarylogo",
	LibraryLogo2x:     "librarylogo2x",
	MainCapsule:       "maincapsule",
	PageBackground:    "pagebackground",
	HeroCapsule:       "herocapsule",
	CommunityIcon:     "communityicon",
	RawPageBackground: "rawpagebackground",
}

var assetExt = map[Asset]string{
	ClientIcon:        ".ico",
	ClientTga:         ".tga",
	Icon:              ".jpg",
	Logo:              ".jpg",
	LogoSmall:         ".jpg",
	SmallCapsule:      ".jpg",
	Header:            ".jpg",
	Header2x:          ".jpg",
	LibraryCapsule:    ".jpg",
	LibraryCapsule2x:  ".jpg",
	LibraryHero:       ".jpg",
	LibraryHero2x:     ".jpg",
	LibraryLogo:       ".png",
	LibraryLogo2x:     ".png",
	MainCapsule:       ".jpg",
	PageBackground:    ".jpg",
	HeroCapsule:       ".jpg",
	CommunityIcon:     ".jpg",
	RawPageBackground: ".jpg",
}

func (a Asset) String() string {
	if as, ok := assetStrings[a]; ok {
		return as
	}
	return assetStrings[UnknownAsset]
}

func (a Asset) Ext() string {
	if ae, ok := assetExt[a]; ok {
		return ae
	}
	return ""
}

var ShortcutAssets = []Asset{
	Header,
	LibraryCapsule,
	LibraryHero,
	LibraryLogo,
	Icon,
}
