package vangogh_integration

type Binary struct {
	Title           string
	Version         string
	DownloadUrl     string
	GitHubOwnerRepo string
	GitHubAssetGlob string
}

func (bin *Binary) String() string {
	switch bin.Title {
	case "":
		return bin.GitHubOwnerRepo
	default:
		return bin.Title
	}
}

var OsBinaries = map[OperatingSystem][]Binary{
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
			Title:       "Microsoft Visual C++ Redistributable X64",
			Version:     "17",
			DownloadUrl: "https://aka.ms/vs/17/release/vc_redist.x64.exe",
		},
		{
			Title:       "Microsoft Visual C++ Redistributable X86",
			Version:     "17",
			DownloadUrl: "https://aka.ms/vs/17/release/vc_redist.x86.exe",
		},
	},
}
