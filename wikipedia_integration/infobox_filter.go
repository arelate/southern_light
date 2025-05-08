package wikipedia_integration

import (
	"bufio"
	"io"
	"strings"
)

const (
	ListPfx    = "{{"
	ListSfx    = "}}"
	InfoboxPfx = ListPfx + "Infobox"
)

func FilterInfoboxLines(r io.Reader) ([]string, error) {

	infoboxLines := make([]string, 0)

	ts := bufio.NewScanner(r)

	collecting := false
	collected := false
	listCounter := 0

	for ts.Scan() {
		line := ts.Text()
		line = strings.TrimSpace(line)

		listCounter += strings.Count(line, ListPfx)
		listCounter -= strings.Count(line, ListSfx)

		if !collecting && strings.Contains(line, InfoboxPfx) {
			collecting = true
			continue
		}
		if collected && listCounter == 0 && strings.Contains(line, ListSfx) {
			break
		}

		if collecting {
			infoboxLines = append(infoboxLines, line)
			collected = true
		}
	}

	if err := ts.Err(); err != nil {
		return nil, err
	}

	return infoboxLines, nil
}
