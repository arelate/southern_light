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

	scanning := false
	previousLine := ""
	listCounter := 0

	for ts.Scan() {
		line := ts.Text()
		line = strings.TrimSpace(line)

		listCounter += strings.Count(line, ListPfx)
		listCounter -= strings.Count(line, ListSfx)

		if !scanning && strings.Contains(line, InfoboxPfx) {
			scanning = true
			continue
		}
		if listCounter == 0 && strings.HasPrefix(line, ListSfx) {
			break
		}
		if scanning && line == "" && strings.HasSuffix(previousLine, ListSfx) {
			break
		}

		if scanning {
			infoboxLines = append(infoboxLines, line)
		}

		previousLine = line
	}

	if err := ts.Err(); err != nil {
		return nil, err
	}

	return infoboxLines, nil
}
