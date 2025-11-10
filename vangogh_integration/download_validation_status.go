package vangogh_integration

import "github.com/boggydigital/redux"

type DownloadValidationStatus struct {
	downloadStatus   DownloadStatus
	validationResult ValidationStatus
}

func NewManualUrlDvs(manualUrl string, rdx redux.Readable) *DownloadValidationStatus {

	if err := rdx.MustHave(ManualUrlValidationResultProperty, ManualUrlValidationResultProperty); err != nil {
		panic(err)
	}

	dvs := &DownloadValidationStatus{
		downloadStatus:   DownloadStatusUnknown,
		validationResult: ValidationStatusUnknown,
	}

	if muss, ok := rdx.GetLastVal(ManualUrlStatusProperty, manualUrl); ok {
		dvs.downloadStatus = ParseDownloadStatus(muss)
	}

	if vrs, ok := rdx.GetLastVal(ManualUrlValidationResultProperty, manualUrl); ok {
		dvs.validationResult = ParseValidationStatus(vrs)
	}

	return dvs
}

func NewProductDvs(id string, rdx redux.Readable) *DownloadValidationStatus {

	if err := rdx.MustHave(ProductValidationResultProperty,
		DownloadQueuedProperty,
		DownloadStartedProperty,
		DownloadCompletedProperty); err != nil {
		panic(err)
	}

	dvs := &DownloadValidationStatus{
		downloadStatus:   DownloadStatusUnknown,
		validationResult: ValidationStatusUnknown,
	}

	if pvrs, ok := rdx.GetLastVal(ProductValidationResultProperty, id); ok {
		dvs.validationResult = ParseValidationStatus(pvrs)
	}

	var downloadQueued, downloadStarted, downloadCompleted string

	if dq, ok := rdx.GetLastVal(DownloadQueuedProperty, id); ok {
		downloadQueued = dq
	}
	if ds, ok := rdx.GetLastVal(DownloadStartedProperty, id); ok {
		downloadStarted = ds
	}
	if dc, ok := rdx.GetLastVal(DownloadCompletedProperty, id); ok {
		downloadCompleted = dc
	}

	if downloadCompleted == "" && dvs.validationResult == ValidationStatusUnknown && downloadQueued == "" {
		dvs.downloadStatus = DownloadStatusUnknown
	} else if downloadQueued > downloadStarted &&
		downloadQueued > downloadCompleted {
		dvs.downloadStatus = DownloadStatusQueued
	} else if downloadStarted > downloadCompleted {
		dvs.downloadStatus = DownloadStatusDownloading
	} else if downloadCompleted >= downloadQueued && downloadCompleted >= downloadStarted {
		dvs.downloadStatus = DownloadStatusDownloaded
	}

	return dvs

}

func (dvs *DownloadValidationStatus) HumanReadableString() string {
	switch dvs.validationResult {
	case ValidationStatusUnknown:
		return dvs.downloadStatus.HumanReadableString()
	default:
		return dvs.validationResult.HumanReadableString()
	}
}

func (dvs *DownloadValidationStatus) ValidationStatus() ValidationStatus {
	return dvs.validationResult
}

func (dvs *DownloadValidationStatus) DownloadStatus() DownloadStatus {
	return dvs.downloadStatus
}
