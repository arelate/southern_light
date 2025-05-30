package gog_integration

type GogGameInfo struct {
	BuildId    string     `json:"buildId"`
	ClientId   string     `json:"clientId"`
	GameId     string     `json:"gameId"`
	Language   string     `json:"language"`
	Languages  []string   `json:"languages"`
	Name       string     `json:"name"`
	OsBitness  []string   `json:"osBitness"`
	PlayTasks  []PlayTask `json:"playTasks"`
	RootGameId string     `json:"rootGameId"`
	Version    int        `json:"version"`
}

type PlayTask struct {
	Arguments  string   `json:"arguments,omitempty"`
	Category   string   `json:"category"`
	IsPrimary  bool     `json:"isPrimary,omitempty"`
	IsHidden   bool     `json:"isHidden,omitempty"`
	Languages  []string `json:"languages"`
	Link       string   `json:"link,omitempty"`
	Name       string   `json:"name"`
	OsBitness  []string `json:"osBitness"`
	Path       string   `json:"path,omitempty"`
	Type       string   `json:"type"`
	WorkingDir string   `json:"workingDir,omitempty"`
}

func (ggi *GogGameInfo) GetPrimaryPlayTaskPath() string {
	for _, pt := range ggi.PlayTasks {
		if pt.IsPrimary {
			return pt.Path
		}
	}
	return ""
}
