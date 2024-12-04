package github_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"strings"
)

func ReleasesUrl(owner, repo string) *url.URL {

	// https://docs.github.com/en/rest/releases/releases
	releasesPath := strings.Replace(releasesPathTemplate, "{owner}", owner, 1)
	releasesPath = strings.Replace(releasesPath, "{repo}", repo, 1)

	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   githubApiHost,
		Path:   releasesPath,
	}
}
