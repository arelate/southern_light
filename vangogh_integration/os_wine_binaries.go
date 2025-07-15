package vangogh_integration

var OsWineBinaries = map[OperatingSystem][]Binary{
	Linux: {
		{
			GitHubOwnerRepo: "Open-Wine-Components/umu-launcher",
			GitHubAssetGlob: "*-zipapp.tar",
		},
		{
			GitHubOwnerRepo: "GloriousEggroll/proton-ge-custom",
			GitHubAssetGlob: "*.tar.gz",
		},
		{
			GitHubOwnerRepo: "Open-Wine-Components/umu-proton",
			GitHubAssetGlob: "*.tar.gz",
		},
	},
	MacOS: {
		{
			Title:       "CrossOver",
			Version:     "25.0.1",
			DownloadUrl: "https://media.codeweavers.com/pub/crossover/cxmac/demo/crossover-25.0.1.zip",
		},
		{
			GitHubOwnerRepo: "Gcenx/macOS_Wine_builds",
			GitHubAssetGlob: "*.tar.xz",
		},
		{
			GitHubOwnerRepo: "3Shain/dxmt",
			GitHubAssetGlob: "*.tar.gz",
		},
	},
	Windows: {
		{
			Title:       "vc_redist.x64.exe",
			Version:     "v14.44.35211.0",
			DownloadUrl: "https://aka.ms/vs/17/release/vc_redist.x64.exe",
		},
		{
			Title:       "vc_redist.x86.exe",
			Version:     "v14.44.35211.0",
			DownloadUrl: "https://aka.ms/vs/17/release/vc_redist.x86.exe",
		},
	},
}
