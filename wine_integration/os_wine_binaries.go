package wine_integration

import (
	"maps"
	"slices"

	"github.com/arelate/southern_light/vangogh_integration"
)

const (
	UmuLauncher        = "Open-Wine-Components/umu-launcher"
	UmuProton          = "Open-Wine-Components/umu-proton"
	ProtonGe           = "GloriousEggroll/proton-ge-custom"
	CrossOver          = "CrossOver"
	WineMacOs          = "Gcenx/macOS_Wine_builds"
	GamePortingToolkit = "Gcenx/game-porting-toolkit"
	DxMt               = "3Shain/dxmt"
	VcRedistX64        = "Visual C++ Redistributable X64"
	VcRedistX86        = "Visual C++ Redistributable X86"
)

const (
	SteamProton316                  = "proton-3.16"
	SteamProton370                  = "proton-3.7"
	SteamProton411                  = "proton-4.11"
	SteamProton420                  = "proton-4.2"
	SteamProton500                  = "proton-5.0"
	SteamProton513                  = "proton-5.13"
	SteamProton630                  = "proton-6.3"
	SteamProton700                  = "proton-7.0"
	SteamProton800                  = "proton-8.0"
	SteamProton900                  = "proton-9.0"
	SteamProton1000                 = "proton-10.0"
	SteamProtonNext                 = "proton-next"
	SteamProtonHotfix               = "proton-hotfix"
	SteamProtonExperimental         = "proton-experimental"
	SteamProtonBattlEyeRuntime      = "proton-battleye-runtime"
	SteamProtonEasyAntiCheatRuntime = "proton-easy-anti-cheat-runtime"
)

var SteamProtonDirectories = map[string]string{
	SteamProton316:                  "Proton 3.16",
	SteamProton370:                  "Proton 3.7",
	SteamProton411:                  "Proton 4.11",
	SteamProton420:                  "Proton 4.2",
	SteamProton500:                  "Proton 5.0",
	SteamProton513:                  "Proton 5.13",
	SteamProton630:                  "Proton 6.3",
	SteamProton700:                  "Proton 7.0",
	SteamProton800:                  "Proton 8.0",
	SteamProton900:                  "Proton 9.0",
	SteamProton1000:                 "Proton 10.0",
	SteamProtonNext:                 "Proton Next",
	SteamProtonHotfix:               "Proton Hotfix",
	SteamProtonExperimental:         "Proton Experimental",
	SteamProtonBattlEyeRuntime:      "Proton BattlEye Runtime",
	SteamProtonEasyAntiCheatRuntime: "Proton Easy Anti-Cheat Runtime",
}

const (
	tarGz = "*.tar.gz"
	tarXz = "*.tar.xz"
)

var OsWineBinaries = []Binary{
	{
		GitHubOwnerRepo: UmuLauncher,
		GitHubAssetGlob: "*-zipapp.tar",
		OS:              vangogh_integration.Linux,
	},
	{
		GitHubOwnerRepo: ProtonGe,
		GitHubAssetGlob: tarGz,
		OS:              vangogh_integration.Linux,
	},
	{
		GitHubOwnerRepo: UmuProton,
		GitHubAssetGlob: tarGz,
		OS:              vangogh_integration.Linux,
	},
	{
		Title:   CrossOver,
		Version: "25.1.0",
		// Digest source: https://media.codeweavers.com/pub/crossover/cxmac/demo/crossover-25.1.0.zip.sha256
		Digest:      "sha256:d3f6425f6d8778a32244125769d38275d19f24681a078e08a1a863d802d8e675",
		DownloadUrl: "https://media.codeweavers.com/pub/crossover/cxmac/demo/crossover-25.1.0.zip",
		OS:          vangogh_integration.MacOS,
	},
	{
		GitHubOwnerRepo: WineMacOs,
		GitHubAssetGlob: tarXz,
		OS:              vangogh_integration.MacOS,
	},
	{
		GitHubOwnerRepo: GamePortingToolkit,
		GitHubAssetGlob: tarXz,
		OS:              vangogh_integration.MacOS,
	},
	{
		GitHubOwnerRepo: DxMt,
		GitHubAssetGlob: tarGz,
		OS:              vangogh_integration.MacOS,
	},
	{
		Title:   VcRedistX64,
		Code:    VcRedistX64Code,
		Version: "v14.44.35211.0",
		// Digest source: https://download.visualstudio.microsoft.com/download/pr/7ebf5fdb-36dc-4145-b0a0-90d3d5990a61/CC0FF0EB1DC3F5188AE6300FAEF32BF5BEEBA4BDD6E8E445A9184072096B713B/VC_redist.x64.exe
		// Note the URL path part before the filename
		Digest:      "sha256:cc0ff0eb1dc3f5188ae6300faef32bf5beeba4bdd6e8e445a9184072096b713b",
		DownloadUrl: "https://aka.ms/vs/17/release/vc_redist.x64.exe",
		OS:          vangogh_integration.Windows,
	},
	{
		Title:   VcRedistX86,
		Code:    VcRedistX86Code,
		Version: "v14.44.35211.0",
		// Digest source: https://download.visualstudio.microsoft.com/download/pr/7ebf5fdb-36dc-4145-b0a0-90d3d5990a61/0C09F2611660441084CE0DF425C51C11E147E6447963C3690F97E0B25C55ED64/VC_redist.x86.exe
		Digest:      "sha256:0c09f2611660441084ce0df425c51c11e147e6447963c3690f97e0b25c55ed64",
		DownloadUrl: "https://aka.ms/vs/17/release/vc_redist.x86.exe",
		OS:          vangogh_integration.Windows,
	},
}

func AllSteamProtonRuntimes() []string {
	return slices.Collect(maps.Keys(SteamProtonDirectories))
}

func IsSteamProtonRuntime(runtime string) bool {
	_, ok := SteamProtonDirectories[runtime]
	return ok
}
