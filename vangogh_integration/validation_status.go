package vangogh_integration

type ValidationStatus int

const (
	ValidationStatusSuccess ValidationStatus = iota
	ValidationStatusUnknown
	ValidationStatusQueued
	ValidationStatusValidating
	ValidationStatusMissingChecksum
	ValidationStatusUnresolvedManualUrl
	ValidationStatusMissingLocalFile
	ValidationStatusError
	ValidationStatusChecksumMismatch
)

var validationStatusStrings = map[ValidationStatus]string{
	ValidationStatusSuccess:             "valid",
	ValidationStatusUnknown:             "unknown",
	ValidationStatusQueued:              "validation-queued",
	ValidationStatusValidating:          "validating",
	ValidationStatusMissingChecksum:     "missing-checksum",
	ValidationStatusUnresolvedManualUrl: "unresolved-manual-url",
	ValidationStatusMissingLocalFile:    "missing-local-file",
	ValidationStatusError:               "error",
	ValidationStatusChecksumMismatch:    "checksum-mismatch",
}

var validationStatusHumanReadableStrings = map[ValidationStatus]string{
	ValidationStatusSuccess:             "Validated",
	ValidationStatusUnknown:             "Unknown",
	ValidationStatusQueued:              "Validation Queued",
	ValidationStatusValidating:          "Validating",
	ValidationStatusMissingChecksum:     "No Checksum",
	ValidationStatusUnresolvedManualUrl: "Unresolved",
	ValidationStatusMissingLocalFile:    "No File",
	ValidationStatusError:               "Error",
	ValidationStatusChecksumMismatch:    "Corrupted",
}

func (vr ValidationStatus) String() string {
	return validationStatusStrings[vr]
}

func (vr ValidationStatus) HumanReadableString() string {
	return validationStatusHumanReadableStrings[vr]
}

func ParseValidationStatus(vrs string) ValidationStatus {
	for vr, str := range validationStatusStrings {
		if vrs == str {
			return vr
		}
	}
	return ValidationStatusUnknown
}

func WorstValidationStatus(vrs ...ValidationStatus) ValidationStatus {

	wvr := ValidationStatusSuccess
	if len(vrs) == 0 {
		wvr = ValidationStatusUnknown
	}

	for _, vr := range vrs {
		if vr > wvr {
			wvr = vr
		}
	}

	return wvr
}

func (vr ValidationStatus) IsValidated() bool {
	switch vr {
	case ValidationStatusUnknown:
		fallthrough
	case ValidationStatusQueued:
		fallthrough
	case ValidationStatusValidating:
		return false
	default:
		return true
	}
}
