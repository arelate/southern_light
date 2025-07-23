package github_integration

import "time"

type GitHubPerson struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	UserViewType      string `json:"user_view_type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type GitHubAsset struct {
	Url                string       `json:"url"`
	Id                 int          `json:"id"`
	NodeId             string       `json:"node_id"`
	Name               string       `json:"name"`
	Label              interface{}  `json:"label"`
	Uploader           GitHubPerson `json:"uploader"`
	ContentType        string       `json:"content_type"`
	State              string       `json:"state"`
	Size               int          `json:"size"`
	Digest             *string      `json:"digest"`
	DownloadCount      int          `json:"download_count"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
	BrowserDownloadUrl string       `json:"browser_download_url"`
}

type GitHubReactions struct {
	Url        string `json:"url"`
	TotalCount int    `json:"total_count"`
	PlusOne    int    `json:"+1"`
	MinusOne   int    `json:"-1"`
	Laugh      int    `json:"laugh"`
	Hooray     int    `json:"hooray"`
	Confused   int    `json:"confused"`
	Heart      int    `json:"heart"`
	Rocket     int    `json:"rocket"`
	Eyes       int    `json:"eyes"`
}

type GitHubRelease struct {
	Url             string          `json:"url"`
	AssetsUrl       string          `json:"assets_url"`
	UploadUrl       string          `json:"upload_url"`
	HtmlUrl         string          `json:"html_url"`
	Id              int             `json:"id"`
	Author          GitHubPerson    `json:"author"`
	NodeId          string          `json:"node_id"`
	TagName         string          `json:"tag_name"`
	TargetCommitish string          `json:"target_commitish"`
	Name            string          `json:"name"`
	Draft           bool            `json:"draft"`
	PreRelease      bool            `json:"prerelease"`
	CreatedAt       time.Time       `json:"created_at"`
	PublishedAt     time.Time       `json:"published_at"`
	Assets          []GitHubAsset   `json:"assets"`
	TarballUrl      string          `json:"tarball_url"`
	ZipballUrl      string          `json:"zipball_url"`
	Body            string          `json:"body"`
	Reactions       GitHubReactions `json:"reactions,omitempty"`
	MentionsCount   int             `json:"mentions_count,omitempty"`
}
