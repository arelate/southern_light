package hltb_integration

import (
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
)

const (
	buildIdMarker = "\"buildId\":"
)

type RootPage struct {
	Doc *html.Node
}

func (rp *RootPage) GetBuildId() string {
	if rp.Doc == nil {
		return ""
	}

	if ndsm := match_node.Match(rp.Doc, &nextDataScriptMatcher{}); ndsm != nil && ndsm.FirstChild != nil {
		buildString := ndsm.FirstChild.Data
		if strings.Contains(buildString, "buildId") {

			if _, version, ok := strings.Cut(buildString, buildIdMarker); ok {
				if parts := strings.Split(version, ","); len(parts) > 0 {
					return strings.Trim(parts[0], "\"")
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
