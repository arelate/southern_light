package steam_appinfo

type AppInfo struct {
	AppId    string
	Common   *AppInfoCommon
	Extended *AppInfoExtended
	Config   *AppInfoConfig
	Depots   *AppInfoDepots
	Ufs      *AppInfoUfs
}

type AppInfoCommon struct {
	Name                           string
	NameLocalized                  map[string]string
	Type                           string
	ReleaseState                   string
	Logo                           string
	LogoSmall                      string
	ClientIcon                     string
	LinuxClientIcon                string
	ClientIcns                     string
	ClientTga                      string
	Icon                           string
	OsList                         string
	OsArch                         string
	OsExtended                     string
	Languages                      []string
	ContentDescriptors             []string
	ContentDescriptorsIncludingDlc []string
	SteamDeckCompatibility         *SteamDeckCompatibility
	SteamDeckBlogUrl               string
	ControllerTagWizard            string
	MetacriticName                 string
	ControllerSupport              string
	SmallCapsule                   map[string]string
	HeaderImage                    map[string]string
	LibraryAssets                  *LibraryAssets
	LibraryAssetsFull              *LibraryAssetsFull
	StoreAssetMtime                int64
	Associations                   map[string]string
	WorkshopVisible                string
	PrimaryGenre                   string
	Genres                         []string
	Category                       []string
	SupportedLanguages             map[string]*LanguageSupport
	SteamReleaseDate               int64
	MasterSubsGrantingApp          string
	MetacriticScore                int
	MetacriticUrl                  string
	MetacriticFullUrl              string
	CommunityVisibleStats          int
	CommunityHubVisible            int
	GameId                         int64
	StoreTags                      []string
	ReviewScore                    int
	ReviewPercentage               int
}

type SteamDeckCompatibility struct {
	Category             int
	SteamOsCompatibility int
	TestTimestamp        int64
	TestBuildId          int64
	Tests                []SteamTestResult
	SteamOsTests         []SteamTestResult
	Configuration        map[string]string
}

type LibraryAssets struct {
	LibraryCapsule  string
	LibraryHero     string
	LibraryHeroBlur string
	LibraryLogo     string
	LibraryHeader   string
	LogoPosition    *LogoPosition
}

type LogoPosition struct {
	PinnedPosition string
	WidthPct       float64
	HeighPct       float64
}

type LibraryAssetsFull struct {
	LibraryCapsule  *Image2xAssets
	LibraryHero     *Image2xAssets
	LibraryHeroBlur *Image2xAssets
	LibraryHeader   *Image2xAssets
	LibraryLogo     *LibraryLogoAssetsFull
}

type Image2xAssets struct {
	Image   map[string]string
	Image2x map[string]string
}

type LibraryLogoAssetsFull struct {
	*LogoPosition
	*Image2xAssets
}

type LanguageSupport struct {
	Supported bool
	FullAudio bool
	Subtitles bool
}

type SteamTestResult struct {
	Display int
	Token   string
}

type AppInfoExtended struct {
	DeckResolutionOverride    string
	Developer                 string
	GameDir                   string
	Icon                      string
	InstallScriptOsx          string
	InstallScriptMacOs        string
	Order                     string
	Languages                 string
	LanguagesMac              string
	LanguagesMacOs            string
	NoServers                 string
	PrimaryCache              string
	SourceGame                string
	State                     string
	ThirdPartyCdKey           string
	ValidOsList               string
	VisibleOnlyWhenSubscribed string
	Publisher                 string
	Homepage                  string
	ListOfDlc                 string
	DlcAvailableOnStore       string
}

type AppInfoConfig struct {
	UseMms                            string
	InstallDir                        string
	Launch                            []LaunchOption
	ContentType                       string
	SteamControllerTouchTemplateIndex int
	SteamControllerTouchConfigDetails map[string]SteamControllerTouchConfigDetails
	SteamControllerTemplateIndex      int
	SteamDeckTouchScreen              string
	SteamConfigurator3rdPartyNative   string
	SteamInputManifestPath            string
	CegPublicKey                      string
	CheckGuids                        string
}

type LaunchOption struct {
	Executable     string
	Arguments      string
	WorkingDir     string
	Type           string
	Config         LaunchOptionConfig
	DescriptionLoc map[string]string
	Description    string
}

type LaunchOptionConfig struct {
	OsList string
	OsArch string
}

type SteamControllerTouchConfigDetails struct {
	ControllerType  string
	EnabledBranches string
	UseActionBlock  bool
}

type AppInfoDepots struct {
	BaseLanguages string
	Depots        map[string]Depot
}

type Depot struct {
	Config        DepotConfig
	Manifests     map[string]DepotManifest
	DepotFromApp  string
	SharedInstall bool
}

type DepotConfig struct {
	OsList   string
	Language string
}

type DepotManifest struct {
	Gid      string
	Size     int64
	Download string
}

type AppInfoUfs struct {
	Quota         int64
	MaxNumFiles   int
	SaveFiles     []SaveFiles
	RootOverrides []RootOverrides
}

type SaveFiles struct {
	Root    string
	Path    string
	Pattern string
}

type RootOverrides struct {
	Root           string
	Os             string
	OsCompare      string
	UseInstead     string
	PathTransforms []PathTransform
}

type PathTransform struct {
	Find    string
	Replace string
}
