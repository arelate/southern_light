package vangogh_integration

type Binary struct {
	Title             string
	DownloadUrl       string
	GitHubOwnerRepo   string
	GitHubReleaseGlob string
}

var OsLaunchers = map[OperatingSystem][]Binary{
	Linux: {
		{
			GitHubOwnerRepo:   "Open-Wine-Components/umu-launcher",
			GitHubReleaseGlob: "*-zipapp.tar",
		},
	},
}

var OsRuntimes = map[OperatingSystem][]Binary{
	Linux: {
		{
			GitHubOwnerRepo:   "GloriousEggroll/proton-ge-custom",
			GitHubReleaseGlob: "*.tar.gz",
		},
		{
			GitHubOwnerRepo:   "Open-Wine-Components/umu-proton",
			GitHubReleaseGlob: "*.tar.gz",
		},
	},
	MacOS: {
		{
			Title:       "CrossOver",
			DownloadUrl: "https://media.codeweavers.com/pub/crossover/cxmac/demo/crossover-25.0.1.zip",
		},
		{
			GitHubOwnerRepo:   "Gcenx/macOS_Wine_builds",
			GitHubReleaseGlob: "*.tar.xz",
		},
		{
			GitHubOwnerRepo:   "3Shain/dxmt",
			GitHubReleaseGlob: "*.tar.gz",
		},
	},
}

var OsTools = map[OperatingSystem][]Binary{
	Windows: {
		{
			Title:       "Microsoft Visual C++ Redistributable X64",
			DownloadUrl: "https://aka.ms/vs/17/release/vc_redist.x64.exe",
		},
		{
			Title:       "Microsoft Visual C++ Redistributable X86",
			DownloadUrl: "https://aka.ms/vs/17/release/vc_redist.x86.exe",
		},
	},
}
