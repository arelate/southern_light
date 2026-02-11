package gog_integration

import (
	"encoding/json/v2"
	"errors"
	"os"
	"strings"
)

const (
	GogGameInfoFilenameTemplate = "goggame-{id}.info"
)

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

func (ggi *GogGameInfo) GetPlayTask(playTask string) (*PlayTask, error) {

	if len(ggi.PlayTasks) == 0 {
		return nil, errors.New("goggame.info has no playtasks")
	}

	if playTask == "" {
		for _, pt := range ggi.PlayTasks {
			if pt.IsPrimary {
				return &pt, nil
			}
		}
	}

	for _, pt := range ggi.PlayTasks {
		if pt.Name == playTask {
			return &pt, nil
		}
	}

	for _, pt := range ggi.PlayTasks {
		if strings.Contains(pt.Name, playTask) {
			return &pt, nil
		}
	}

	return nil, errors.New("playtask not found: " + playTask)
}

func GetGogGameInfo(path string) (*GogGameInfo, error) {

	gogGameInfoFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer gogGameInfoFile.Close()

	var gogGameInfo GogGameInfo
	if err = json.UnmarshalRead(gogGameInfoFile, &gogGameInfo); err != nil {
		return nil, err
	}

	return &gogGameInfo, nil
}
