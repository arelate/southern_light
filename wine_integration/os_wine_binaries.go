package wine_integration

import (
	"strings"

	"github.com/arelate/southern_light/vangogh_integration"
)

const (
	// GitHub binaries
	UmuLauncher        = "Open-Wine-Components/umu-launcher"
	UmuProton          = "Open-Wine-Components/umu-proton"
	ProtonGeCustom     = "GloriousEggroll/proton-ge-custom"
	ProtonCachyOs      = "CachyOS/proton-cachyos"
	ProtonEm           = "Etaash-mathamsetty/Proton"
	CrossOver          = "CrossOver"
	WineMacOs          = "Gcenx/macOS_Wine_builds"
	GamePortingToolkit = "Gcenx/game-porting-toolkit"
	DxMt               = "3Shain/dxmt"
	Innoextract        = "dscharrer/innoextract"

	// Direct download binaries
	VcRedistX64      = "vc_redist.x64.exe"
	VcRedistX86      = "vc_redist.x86.exe"
	DxEndUserRuntime = "directx_Jun2010_redist.exe"
)

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
		GitHubOwnerRepo: ProtonGeCustom,
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
		Version: "26.2.0",
		// Digest source: https://media.codeweavers.com/pub/crossover/cxmac/demo/crossover-26.2.0.zip.sha256
		Digest:      "sha256:ab39680927a0d9c313cad27fed33dd15a799be40c93bfc432bef70518af10aa4",
		DownloadUrl: "https://media.codeweavers.com/pub/crossover/cxmac/demo/crossover-26.2.0.zip",
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
		// Digest source: https://download.visualstudio.microsoft.com/download/pr/ebdab8e5-1d7b-4d9f-a11b-cbb1720c3b12/843068991DAAA1F73AD9F6239BCE4D0F6A07A51F18C37EA2A867E9BECA71295C/VC_redist.x64.exe
		// Note the URL path part before the filename
		Digest:      strings.ToLower("sha256:843068991DAAA1F73AD9F6239BCE4D0F6A07A51F18C37EA2A867E9BECA71295C"),
		DownloadUrl: "https://aka.ms/vc14/vc_redist.x64.exe",
		OS:          vangogh_integration.Windows,
	},
	{
		Title:   VcRedistX86,
		Code:    VcRedistX86Code,
		Version: "11/11/2025",
		// Digest source: https://download.visualstudio.microsoft.com/download/pr/57eef8ae-a341-46c3-b0bc-c041027b54cd/F0BAB33A302B3CDB2E11113760D016F54FD3D2632C65BA7834FAC4F0ABD7F1A3/VC_redist.x86.exe
		Digest:      strings.ToLower("sha256:F0BAB33A302B3CDB2E11113760D016F54FD3D2632C65BA7834FAC4F0ABD7F1A3"),
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
	{
		GitHubOwnerRepo: Innoextract,
		GitHubAssetGlob: tarXz,
		OS:              vangogh_integration.Linux,
	},
}
