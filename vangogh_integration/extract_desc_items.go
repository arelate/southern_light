package vangogh_integration

import "regexp"

var (
	descItems     = regexp.MustCompile(`https://items\.gog\.com/([-\pL0-9()!@:%_,’'™×\ \+\[\]\x{2013}/u.~#?&\/\/=]*)`)
	descGameLinks = regexp.MustCompile(`https://www\.gog\.com/(en/)?game/([-\pL0-9()!@:%_,’'™×\ \+\[\]\x{2013}/u.~#?&\/\/=]*)`)
)

func ExtractDescItems(desc string) []string {
	return descItems.FindAllString(desc, -1)
}

func ExtractGameLinks(desc string) []string {
	return descGameLinks.FindAllString(desc, -1)
}
