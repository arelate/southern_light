package vangogh_integration

import (
	"fmt"
	"github.com/boggydigital/nod"
	"strconv"
	"strings"
)

func PrintParams(
	ids []string,
	operatingSystems []OperatingSystem,
	langCodes []string,
	downloadTypes []DownloadType,
	noPatches bool) {

	ppa := nod.Begin("operating parameters:")
	defer ppa.Done()

	params := make(map[string][]string)

	for _, id := range ids {
		params[IdProperty] = append(params[IdProperty], id)
	}

	for _, os := range operatingSystems {
		params[OperatingSystemsProperty] = append(params[OperatingSystemsProperty], os.String())
	}

	for _, lc := range langCodes {
		params[LanguageCodeProperty] = append(params[LanguageCodeProperty], lc)
	}

	for _, dt := range downloadTypes {
		params[DownloadTypeProperty] = append(params[DownloadTypeProperty], dt.String())
	}

	params[NoPatchesProperty] = append(params[NoPatchesProperty], strconv.FormatBool(noPatches))

	pvs := make([]string, 0, len(params))
	for _, p := range []string{
		IdProperty,
		OperatingSystemsProperty,
		LanguageCodeProperty,
		DownloadTypeProperty} {

		if _, ok := params[p]; !ok {
			continue
		}

		pvs = append(pvs, fmt.Sprintf("%s=%s", p, strings.Join(params[p], ",")))
	}

	ppa.EndWithResult(strings.Join(pvs, "; "))
}
