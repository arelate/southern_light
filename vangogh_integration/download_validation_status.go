package vangogh_integration

import "github.com/boggydigital/redux"

type DownloadValidationStatus struct {
	downloadStatus   DownloadStatus
	validationResult ValidationStatus
}

func NewManualUrlDvs(manualUrl string, rdx redux.Readable) *DownloadValidationStatus {

	if err := rdx.MustHave(GogManualUrlValidationResultProperty, GogManualUrlValidationResultProperty); err != nil {
		panic(err)
	}

	dvs := &DownloadValidationStatus{
		downloadStatus:   DownloadStatusUnknown,
		validationResult: ValidationStatusUnknown,
	}

	if muss, ok := rdx.GetLastVal(GogManualUrlStatusProperty, manualUrl); ok {
		dvs.downloadStatus = ParseDownloadStatus(muss)
	}

	if vrs, ok := rdx.GetLastVal(GogManualUrlValidationResultProperty, manualUrl); ok {
		dvs.validationResult = ParseValidationStatus(vrs)
	}

	if pgs, ok := rdx.GetLastVal(GogManualUrlGeneratedChecksumProperty, manualUrl); ok && pgs != "" {
		if dvs.validationResult == ValidationStatusSuccess {
			dvs.validationResult = ValidationStatusSelfValidated
		}
	}

	return dvs
}

func NewProductDvs(id string, rdx redux.Readable) *DownloadValidationStatus {

	if err := rdx.MustHave(GogProductValidationResultProperty,
		VangoghDownloadQueuedProperty,
		VangoghDownloadStartedProperty,
		VangoghDownloadCompletedProperty); err != nil {
		panic(err)
	}

	dvs := &DownloadValidationStatus{
		downloadStatus:   DownloadStatusUnknown,
		validationResult: ValidationStatusUnknown,
	}

	if pvrs, ok := rdx.GetLastVal(GogProductValidationResultProperty, id); ok {
		dvs.validationResult = ParseValidationStatus(pvrs)
	}

	if pgs, ok := rdx.GetLastVal(GogProductGeneratedChecksumProperty, id); ok && pgs != "" {
		if dvs.validationResult == ValidationStatusSuccess {
			dvs.validationResult = ValidationStatusSelfValidated
		}
	}

	var downloadQueued, downloadStarted, downloadCompleted string

	if dq, ok := rdx.GetLastVal(VangoghDownloadQueuedProperty, id); ok {
		downloadQueued = dq
	}
	if ds, ok := rdx.GetLastVal(VangoghDownloadStartedProperty, id); ok {
		downloadStarted = ds
	}
	if dc, ok := rdx.GetLastVal(VangoghDownloadCompletedProperty, id); ok {
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
