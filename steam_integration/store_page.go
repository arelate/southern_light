package steam_integration

import (
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
)

const (
	tagsContainerClassVariant1 = "glance_tags popular_tags"
	tagsContainerClassVariant2 = "popular_tags glance_tags"
)

type StorePage struct {
	Doc *html.Node
}

type SteamTagsGetter interface {
	GetSteamTags() []string
}

func (sp *StorePage) GetSteamTags() []string {
	if sp.Doc == nil {
		return nil
	}

	if tagsContainer := match_node.Match(sp.Doc, &appTagMatcher{}); tagsContainer != nil {
		return extractTagsFromContainer(tagsContainer)
	}

	return nil
}

type appTagMatcher struct{}

func (atm *appTagMatcher) Match(node *html.Node) bool {
	if node.Type != html.ElementNode ||
		node.DataAtom != atom.Div ||
		len(node.Attr) == 0 {
		return false
	}

	for _, a := range node.Attr {
		if a.Key == "class" &&
			(strings.Contains(a.Val, tagsContainerClassVariant1) ||
				strings.Contains(a.Val, tagsContainerClassVariant2)) {
			return true
		}
	}

	return false
}

func extractTagsFromContainer(node *html.Node) []string {
	tags := make([]string, 0)
	if node == nil ||
		node.FirstChild == nil {
		return tags
	}

	for ch := node.FirstChild; ch != nil; ch = ch.NextSibling {
		if ch.DataAtom != atom.A {
			continue
		}

		tags = append(tags, strings.Trim(ch.FirstChild.Data, "\n\t"))
	}

	return tags
}
