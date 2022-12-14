package hltb_integration

import (
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
)

type RootPage struct {
	Doc *html.Node
}

type NextBuildGetter interface {
	GetNextBuild() string
}

func (rp *RootPage) GetNextBuild() string {
	if rp.Doc == nil {
		return ""
	}

	if ndsm := match_node.Match(rp.Doc, &nextDataScriptMatcher{}); ndsm != nil && ndsm.FirstChild != nil {
		for _, line := range strings.Split(ndsm.FirstChild.Data, "\n") {
			if strings.Contains(line, "buildId") {
				if parts := strings.Split(line, "\""); len(parts) > 3 {
					return parts[len(parts)-2]
				}
			}
		}
	}

	return ""
}

type nextDataScriptMatcher struct {
}

func (ndsm *nextDataScriptMatcher) Match(node *html.Node) bool {
	if node.DataAtom == atom.Script &&
		len(node.Attr) > 0 {

		for _, attr := range node.Attr {
			if attr.Key == "id" &&
				attr.Val == "__NEXT_DATA__" {
				return true
			}
		}

	}
	return false
}
