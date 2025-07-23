package vangogh_integration

type Binary struct {
	Title           string
	Version         string
	Digest          string
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
