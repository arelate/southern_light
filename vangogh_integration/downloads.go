package vangogh_integration

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"path"
	"strconv"
	"strings"

	"github.com/arelate/southern_light/gog_integration"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/redux"
)

const (
	mbSuffix  = "MB"
	gbSuffix  = "GB"
	bytesInGB = 1024 * 1024 * 1024
	bytesInMB = 1024 * 1024
	patchStr  = "patch"
)

const (
	detailsNotFoundError = "details not found for "
	nilDetailsError      = "details are nil for "
)

type Download struct {
	ManualUrl      string
	ProductTitle   string
	Name           string
	Version        string
	Date           string
	Type           string
	Info           int
	OS             OperatingSystem
	LanguageCode   string
	DownloadType   DownloadType
	EstimatedBytes int64
}

func convertManualDownload(
	productTitle string,
	mdl *gog_integration.ManualDownload,
	dt DownloadType,
	os OperatingSystem,
	langCode string) Download {
	return Download{
		ManualUrl:      mdl.ManualUrl,
		ProductTitle:   productTitle,
		Name:           mdl.Name,
		Version:        mdl.Version,
		Date:           mdl.Date,
		Type:           mdl.Type,
		Info:           mdl.Info,
		OS:             os,
		LanguageCode:   langCode,
		DownloadType:   dt,
		EstimatedBytes: SizeToEstimatedBytes(mdl.Size),
	}
}

func convertToBytes(size string, suffix string, bytesInUnit int) int64 {
	sizeStr := strings.TrimSuffix(size, " "+suffix)
	sz, err := strconv.ParseFloat(sizeStr, 0)
	if err != nil {
		log.Printf("error parsing size: %s", size)
		return 0
	}
	return int64(sz * float64(bytesInUnit))
}

func SizeToEstimatedBytes(size string) int64 {
	if strings.HasSuffix(size, gbSuffix) {
		return convertToBytes(size, gbSuffix, bytesInGB)
	} else if strings.HasSuffix(size, mbSuffix) {
		return convertToBytes(size, mbSuffix, bytesInMB)
	} else {
		log.Printf("unknown size format: %s", size)
		return 0
	}
}

func (dl *Download) String() string {
	switch dl.DownloadType {
	case Installer:
		fallthrough
	case Movie:
		fallthrough
	case DLC:
		name := dl.Name
		if !strings.HasPrefix(dl.Name, dl.ProductTitle) {
			name = fmt.Sprintf("%s %s", dl.ProductTitle, dl.Name)
		}
		return fmt.Sprintf("%s %s (%s, %s)", name, dl.Version, dl.OS, dl.LanguageCode)
	case Extra:
		return strings.Title(dl.Name)
	default:
		return ""
	}
}

type DownloadsList []Download

func FromDetails(det *gog_integration.Details, rdx redux.Readable) (DownloadsList, error) {
	return fromGameDetails(det, rdx)
}

func fromGameDetails(det *gog_integration.Details, rdx redux.Readable) (DownloadsList, error) {
	dlList := make(DownloadsList, 0)

	if det == nil {
		return dlList, fmt.Errorf("details are nil")
	}

	installerDls, err := convertGameDetails(det, rdx, Installer)
	if err != nil {
		return dlList, err
	}
	dlList = append(dlList, installerDls...)

	for _, dlc := range det.DLCs {
		dlcDls, err := convertGameDetails(&dlc, rdx, DLC)
		if err != nil {
			return dlList, err
		}
		dlList = append(dlList, dlcDls...)
	}

	return dlList, nil
}

func convertGameDetails(det *gog_integration.Details, rdx redux.Readable, dt DownloadType) (DownloadsList, error) {

	dlList := make(DownloadsList, 0)

	downloads, err := det.GetGameDownloads()
	if err != nil {
		return dlList, err
	}

	for _, dl := range downloads {

		langCode := gog_integration.LanguageCodeByNativeName(dl.Language)

		if langCode == "" {
			return dlList, fmt.Errorf("invalid native language %s", dl.Language)
		}

		for _, winDl := range dl.Windows {
			dlList = append(dlList, convertManualDownload(det.Title, &winDl, dt, Windows, langCode))
		}
		for _, macDl := range dl.Mac {
			dlList = append(dlList, convertManualDownload(det.Title, &macDl, dt, MacOS, langCode))
		}
		for _, linDl := range dl.Linux {
			dlList = append(dlList, convertManualDownload(det.Title, &linDl, dt, Linux, langCode))
		}
	}

	for _, extraDl := range det.Extras {
		dlList = append(dlList, convertManualDownload(det.Title, &extraDl, Extra, AnyOperatingSystem, ""))
	}

	return dlList, nil
}

func (list DownloadsList) Only(
	operatingSystems []OperatingSystem,
	langCodes []string,
	downloadTypes []DownloadType,
	noPatches bool) DownloadsList {
	osSet := make(map[OperatingSystem]bool)
	for _, os := range operatingSystems {
		if os == AnyOperatingSystem {
			for _, aos := range AllOperatingSystems() {
				osSet[aos] = true
			}
			break
		}
		osSet[os] = true
	}
	dtSet := make(map[DownloadType]bool)
	for _, dt := range downloadTypes {
		if dt == AnyDownloadType {
			for _, adt := range AllDownloadTypes() {
				dtSet[adt] = true
			}
			break
		}
		dtSet[dt] = true
	}
	langSet := make(map[string]bool)
	for _, lc := range langCodes {
		langSet[lc] = true
	}
	matchingList := make(DownloadsList, 0)
	for _, dl := range list {
		if dl.OS != AnyOperatingSystem &&
			!osSet[dl.OS] {
			continue
		}

		if dl.DownloadType != AnyDownloadType &&
			!dtSet[dl.DownloadType] {
			continue
		}

		if dl.LanguageCode != "" &&
			len(langSet) > 0 &&
			!langSet[dl.LanguageCode] {
			continue
		}

		if noPatches {
			if base := path.Base(dl.ManualUrl); strings.Contains(base, patchStr) {
				continue
			}
		}

		matchingList = append(matchingList, dl)
	}
	return matchingList
}

func (list DownloadsList) TotalBytesEstimate() int64 {
	var totalBytes int64
	for _, dl := range list {
		totalBytes += dl.EstimatedBytes
	}
	return totalBytes
}

func (list DownloadsList) TotalGBsEstimate() float64 {
	return float64(list.TotalBytesEstimate()) / math.Pow(1000, 3)
}

type DownloadsListProcessor interface {
	Process(id string, slug string, downloadsList DownloadsList) error
}

func MapDownloads(
	ids []string,
	rdx redux.Readable,
	operatingSystems []OperatingSystem,
	langCodes []string,
	downloadTypes []DownloadType,
	noPatches bool,
	dlProcessor DownloadsListProcessor,
	tpw nod.TotalProgressWriter) error {

	if dlProcessor == nil {
		return fmt.Errorf("vangogh_downloads: map downloads list processor is nil")
	}

	if err := rdx.MustHave(SlugProperty, ProductTypeProperty); err != nil {
		return err
	}

	detailsDir, err := AbsProductTypeDir(Details)
	if err != nil {
		return err
	}

	kvDetails, err := kevlar.New(detailsDir, kevlar.JsonExt)
	if err != nil {
		return err
	}

	packDlcProducts := make([]string, 0)
	gameProducts := make([]string, 0, len(ids))

	for _, id := range ids {
		if pt, ok := rdx.GetLastVal(ProductTypeProperty, id); ok && pt != GameProductType {
			packDlcProducts = append(packDlcProducts, id)
			continue
		}
		gameProducts = append(gameProducts, id)
	}

	if len(packDlcProducts) > 0 {
		pdpa := nod.Begin(" PACK, DLC products do not contain downloads:")
		pdpa.EndWithResult(strings.Join(packDlcProducts, ", "))
	}

	tpw.TotalInt(len(gameProducts))

	for _, id := range gameProducts {

		detSlug, ok := rdx.GetLastVal(SlugProperty, id)

		if !kvDetails.Has(id) || !ok {
			tpw.Increment()
			continue
		}

		det, err := UnmarshalDetails(id, kvDetails)
		if err != nil {
			return err
		}

		var downloads DownloadsList

		if det != nil {
			downloads, err = FromDetails(det, rdx)
			if err != nil {
				return err
			}
		}

		filteredDownloads := make([]Download, 0)

		for _, dl := range downloads.Only(operatingSystems, langCodes, downloadTypes, noPatches) {
			//some manualUrls have "0 MB" specified as size and don't seem to be used to create user clickable links.
			//resolving such manualUrls leads to an empty filename
			//given they don't contribute anything to download, size or validate commands - we're filtering them
			if dl.EstimatedBytes == 0 {
				continue
			}
			filteredDownloads = append(filteredDownloads, dl)
		}

		if det != nil && det.IsPreOrder && len(filteredDownloads) == 0 {
			nod.Log("%s is a pre-order and has no downloads", id)
			continue
		}

		// already checked for nil earlier in the function
		if err = dlProcessor.Process(
			id,
			detSlug,
			filteredDownloads); err != nil {
			return err
		}

		tpw.Increment()
	}

	return nil
}

func UnmarshalDetails(id string, kvDetails kevlar.KeyValues) (*gog_integration.Details, error) {

	rcDetails, err := kvDetails.Get(id)
	if err != nil {
		return nil, err
	}
	defer rcDetails.Close()

	var det gog_integration.Details

	if err = json.NewDecoder(rcDetails).Decode(&det); err != nil {
		if strings.Contains(err.Error(), "cannot unmarshal array into Go value of type gog_integration.Details") {
			return nil, nil
		}
		return nil, err
	}

	return &det, nil
}

func DetailsNotFoundErr(id string) error {
	return errors.New(detailsNotFoundError + id)
}

func NilDetailsErr(id string) error {
	return errors.New(nilDetailsError + id)
}

func IsDetailsNotFound(err error) bool {
	return strings.HasPrefix(err.Error(), detailsNotFoundError)
}

func IsNilDetails(err error) bool {
	return strings.HasPrefix(err.Error(), nilDetailsError)
}

func FormatBytes(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}
