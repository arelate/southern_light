package vangogh_integration

const (
	UmuLauncher = "Open-Wine-Components/umu-launcher"
	UmuProton   = "Open-Wine-Components/umu-proton"
	ProtonGe    = "GloriousEggroll/proton-ge-custom"
	CrossOver   = "CrossOver"
	WineMacOs   = "Gcenx/macOS_Wine_builds"
	DxMt        = "3Shain/dxmt"
	VcRedistX64 = "vc_redist.x64.exe"
	VcRedistX86 = "vc_redist.x86.exe"
)

const (
	tarGz = "*.tar.gz"
	tarXz = "*.tar.xz"
)

var OsWineBinaries = map[OperatingSystem][]Binary{
	Linux: {
		{
			GitHubOwnerRepo: UmuLauncher,
			GitHubAssetGlob: "*-zipapp.tar",
		},
		{
			GitHubOwnerRepo: ProtonGe,
			GitHubAssetGlob: tarGz,
		},
		{
			GitHubOwnerRepo: UmuProton,
			GitHubAssetGlob: tarGz,
		},
	},
	MacOS: {
		{
			Title:       CrossOver,
			Version:     "25.0.1",
			DownloadUrl: "https://media.codeweavers.com/pub/crossover/cxmac/demo/crossover-25.0.1.zip",
		},
		{
			GitHubOwnerRepo: WineMacOs,
			GitHubAssetGlob: tarXz,
		},
		{
			GitHubOwnerRepo: DxMt,
			GitHubAssetGlob: tarGz,
		},
	},
	Windows: {
		{
			Title:       VcRedistX64,
			Version:     "v14.44.35211.0",
			DownloadUrl: "https://aka.ms/vs/17/release/vc_redist.x64.exe",
		},
		{
			Title:       VcRedistX86,
			Version:     "v14.44.35211.0",
			DownloadUrl: "https://aka.ms/vs/17/release/vc_redist.x86.exe",
		},
	},
}
