package vangogh_integration

import "slices"

type ValidationResult int

const (
	ValidatedSuccessfully ValidationResult = iota
	ValidatedUnresolvedManualUrl
	ValidationResultUnknown
	ValidatedMissingChecksum
	ValidatedMissingLocalFile
	ValidationError
	ValidatedChecksumMismatch
)

var validationResultsStrings = map[ValidationResult]string{
	ValidatedSuccessfully:        "valid",
	ValidatedUnresolvedManualUrl: "unresolved-manual-url",
	ValidationResultUnknown:      "unknown",
	ValidatedMissingChecksum:     "missing-checksum",
	ValidatedMissingLocalFile:    "missing-local-file",
	ValidationError:              "error",
	ValidatedChecksumMismatch:    "checksum-mismatch",
}

var validationResultsHumanReadableStrings = map[ValidationResult]string{
	ValidatedSuccessfully:        "Validated",
	ValidatedUnresolvedManualUrl: "Unresolved",
	ValidationResultUnknown:      "Unknown",
	ValidatedMissingChecksum:     "No Checksum",
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
	if len(vrs) == 0 {
		return ValidationResultUnknown
	}

	slices.Sort(vrs)
	return vrs[len(vrs)-1]
}
