package vangogh_integration

type ValidationResult int

const (
	ValidatedSuccessfully ValidationResult = iota
	ValidationResultUnknown
	ValidationQueued
	ValidationValidating
	ValidatedMissingChecksum
	ValidatedUnresolvedManualUrl
	ValidatedMissingLocalFile
	ValidationError
	ValidatedChecksumMismatch
)

var validationResultsStrings = map[ValidationResult]string{
	ValidatedSuccessfully:        "valid",
	ValidationResultUnknown:      "unknown",
	ValidationQueued:             "validation-queued",
	ValidationValidating:         "validating",
	ValidatedMissingChecksum:     "missing-checksum",
	ValidatedUnresolvedManualUrl: "unresolved-manual-url",
	ValidatedMissingLocalFile:    "missing-local-file",
	ValidationError:              "error",
	ValidatedChecksumMismatch:    "checksum-mismatch",
}

var validationResultsHumanReadableStrings = map[ValidationResult]string{
	ValidatedSuccessfully:        "Validated",
	ValidationResultUnknown:      "Unknown",
	ValidatedMissingChecksum:     "No Checksum",
	ValidatedUnresolvedManualUrl: "Unresolved",
	ValidatedMissingLocalFile:    "No File",
	ValidationError:              "Error",
	ValidatedChecksumMismatch:    "Corrupted",
}

func (vr ValidationResult) String() string {
	return validationResultsStrings[vr]
}

func (vr ValidationResult) HumanReadableString() string {
	return validationResultsHumanReadableStrings[vr]
}

func ParseValidationResult(vrs string) ValidationResult {
	for vr, str := range validationResultsStrings {
		if vrs == str {
			return vr
		}
	}
	return ValidationResultUnknown
}

func WorstValidationResult(vrs ...ValidationResult) ValidationResult {

	wvr := ValidatedSuccessfully
	if len(vrs) == 0 {
		wvr = ValidationResultUnknown
	}

	for _, vr := range vrs {
		if vr > wvr {
			wvr = vr
		}
	}

	return wvr
}
