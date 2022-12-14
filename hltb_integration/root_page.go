package hltb_integration

import "golang.org/x/net/html"

type RootPage struct {
	Doc *html.Node
}

type NextBuildGetter interface {
	GetNextBuild() string
}

func (rp *RootPage) GetNextBuild() string {
	return ""
}
