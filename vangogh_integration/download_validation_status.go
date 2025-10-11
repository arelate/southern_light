package vangogh_integration

import "github.com/boggydigital/redux"

type DownloadValidationStatus struct {
	manualUrlStatus  ManualUrlStatus
	validationResult ValidationResult
}

func NewDvs(manualUrl string, rdx redux.Readable) *DownloadValidationStatus {

	if err := rdx.MustHave(ManualUrlValidationResultProperty, ManualUrlValidationResultProperty); err != nil {
		panic(err)
	}

	dvs := &DownloadValidationStatus{
		manualUrlStatus:  ManualUrlStatusUnknown,
		validationResult: ValidationResultUnknown,
	}

	if muss, ok := rdx.GetLastVal(ManualUrlStatusProperty, manualUrl); ok {
		dvs.manualUrlStatus = ParseManualUrlStatus(muss)
	}

	if vrs, ok := rdx.GetLastVal(ManualUrlValidationResultProperty, manualUrl); ok {
		dvs.validationResult = ParseValidationResult(vrs)
	}

	return dvs
}

func (dvs *DownloadValidationStatus) HumanReadableString() string {
	switch dvs.validationResult {
	case ValidationResultUnknown:
		return dvs.manualUrlStatus.HumanReadableString()
	default:
		return dvs.validationResult.HumanReadableString()
	}
}

func (dvs *DownloadValidationStatus) ValidationResult() ValidationResult {
	return dvs.validationResult
}

func (dvs *DownloadValidationStatus) ManualUrlStatus() ManualUrlStatus {
	return dvs.manualUrlStatus
}
