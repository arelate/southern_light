package vangogh_integration

import "github.com/arelate/southern_light/github_integration"

func OperatingSystemGitHubRepos(os OperatingSystem) []string {
	switch os {
	case Linux:
		return []string{
			github_integration.UmuProtonRepo,
			github_integration.UmuLauncherRepo,
		}
	default:
		return nil
	}
}
