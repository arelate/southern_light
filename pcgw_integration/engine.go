package pcgw_integration

import "strings"

const (
	enginePagePrefix = "Engine:"
)

type Engine struct {
	Cargoquery []struct {
		Title struct {
			Engine string `json:"Engine"`
			Build  string `json:"Build"`
		} `json:"title"`
	} `json:"cargoquery"`
}

type EnginesGetter interface {
	GetEngines() []string
}

type EnginesBuildsGetter interface {
	GetEnginesBuilds() []string
}

func (e *Engine) GetEngines() []string {
	engines := make([]string, 0, len(e.Cargoquery))
	for _, cq := range e.Cargoquery {
		engines = append(engines, strings.TrimPrefix(cq.Title.Engine, enginePagePrefix))
	}
	return engines
}

func (e *Engine) GetEnginesBuilds() []string {
	enginesBuilds := make([]string, 0, len(e.Cargoquery))
	for _, cq := range e.Cargoquery {
		enginesBuilds = append(enginesBuilds, cq.Title.Build)
	}
	return enginesBuilds
}
