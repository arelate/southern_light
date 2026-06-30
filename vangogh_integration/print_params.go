package vangogh_integration

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/boggydigital/nod"
)

func PrintParams(
	ids []string,
	operatingSystems []OperatingSystem,
	langCodes []string,
	noDlc, noExtras, noPatches bool) {

	ppa := nod.Begin("operating parameters:")
	defer ppa.Done()

	params := make(map[string][]string)

	for _, id := range ids {
		params[UrlIdParameter] = append(params[UrlIdParameter], id)
	}

	for _, os := range operatingSystems {
		params[UrlOperatingSystemParameter] = append(params[UrlOperatingSystemParameter], os.String())
	}

	for _, lc := range langCodes {
		params[UrlLanguageCodeParameter] = append(params[UrlLanguageCodeParameter], lc)
	}

	if noDlc {
		params[UrlNoDlcsParameter] = append(params[UrlNoDlcsParameter], strconv.FormatBool(noDlc))
	}

	if noExtras {
		params[UrlNoExtrasParameter] = append(params[UrlNoExtrasParameter], strconv.FormatBool(noExtras))
	}

	if noPatches {
		params[NoPatchesProperty] = append(params[NoPatchesProperty], strconv.FormatBool(noPatches))
	}

	pvs := make([]string, 0, len(params))
	for _, p := range []string{
		UrlIdParameter,
		UrlOperatingSystemParameter,
		UrlLanguageCodeParameter,
		UrlDownloadTypeParameter} {

		if _, ok := params[p]; !ok {
			continue
		}

		pvs = append(pvs, fmt.Sprintf("%s=%s", p, strings.Join(params[p], ",")))
	}

	ppa.EndWithResult(strings.Join(pvs, "; "))
}
