package wine_integration

import (
	"maps"
	"slices"
	"strings"

	"github.com/arelate/southern_light/vangogh_integration"
)

const (
	UmuLauncher        = "Open-Wine-Components/umu-launcher"
	UmuProton          = "Open-Wine-Components/umu-proton"
	ProtonGe           = "GloriousEggroll/proton-ge-custom"
	ProtonCachyOs      = "CachyOS/proton-cachyos"
	ProtonEm           = "Etaash-mathamsetty/Proton"
	CrossOver          = "CrossOver"
	WineMacOs          = "Gcenx/macOS_Wine_builds"
	GamePortingToolkit = "Gcenx/game-porting-toolkit"
	DxMt               = "3Shain/dxmt"
	VcRedistX64        = "Visual C++ Redistributable X64"
	VcRedistX86        = "Visual C++ Redistributable X86"
	DxEndUserRuntime   = "DirectX End-User Runtimes (June 2010)"
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
	SteamProton900:                  "Proton 9.0 (Beta)",
	SteamProton1000:                 "Proton 10.0",
	SteamProtonNext:                 "Proton Next",
	SteamProtonHotfix:               "Proton Hotfix",
	SteamProtonExperimental:         "Proton - Experimental",
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
		GitHubOwnerRepo: ProtonCachyOs,
		GitHubAssetGlob: "*-x86_64.tar.xz",
		OS:              vangogh_integration.Linux,
	},
	{
		GitHubOwnerRepo: ProtonEm,
		GitHubAssetGlob: tarXz,
		OS:              vangogh_integration.Linux,
	},
	{
		Title:   CrossOver,
		Version: "26.1.0",
		// Digest source: https://media.codeweavers.com/pub/crossover/cxmac/demo/crossover-26.1.0.zip.sha256
		Digest:      "sha256:99812be77501995533a9fd4553efa75a4237eac6a20c5639ff43ad46660d9256",
		DownloadUrl: "https://media.codeweavers.com/pub/crossover/cxmac/demo/crossover-26.1.0.zip",
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
		Version: "11/11/2025",
		// Digest source: https://download.visualstudio.microsoft.com/download/pr/c1bd4f2c-3672-468e-8480-7ed419dbb641/90E48ADE404E4576D023ABFA374F323555F233982A8805EA9AC63DCA9491A16B/VC_redist.x64.exe
		// Note the URL path part before the filename
		Digest:      strings.ToLower("sha256:90E48ADE404E4576D023ABFA374F323555F233982A8805EA9AC63DCA9491A16B"),
		DownloadUrl: "https://aka.ms/vc14/vc_redist.x64.exe",
		OS:          vangogh_integration.Windows,
	},
	{
		Title:   VcRedistX86,
		Code:    VcRedistX86Code,
		Version: "11/11/2025",
		// Digest source: https://download.visualstudio.microsoft.com/download/pr/0dd156af-82aa-4812-b524-49c2f894359a/B6AB675F0A27E6600F9726E75DEA08D99C15F8EA4B842A2A1D988FA9529D39B9/VC_redist.x86.exe
		Digest:      strings.ToLower("sha256:B6AB675F0A27E6600F9726E75DEA08D99C15F8EA4B842A2A1D988FA9529D39B9"),
		DownloadUrl: "https://aka.ms/vc14/vc_redist.x86.exe",
		OS:          vangogh_integration.Windows,
	},
	{
		Title:   DxEndUserRuntime,
		Code:    DxEndUserRuntimeCode,
		Version: "9.29.1974.1",
		// Digest source: manual
		Digest:      "sha256:053f76dcbb28802e23341b6a787e3b0791c0fa5c8d4d011b1044172dbf89c73b",
		DownloadUrl: "https://download.microsoft.com/download/8/4/a/84a35bf1-dafe-4ae8-82af-ad2ae20b6b14/directx_Jun2010_redist.exe",
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
