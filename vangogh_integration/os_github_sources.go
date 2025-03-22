package vangogh_integration

import "github.com/arelate/southern_light/github_integration"

func OperatingSystemGitHubSources(os OperatingSystem) []*github_integration.GitHubSource {
	switch os {
	case Linux:
		return []*github_integration.GitHubSource{
			github_integration.UmuProton,
			github_integration.UmuLauncher,
		}
	default:
		return nil
	}
}
