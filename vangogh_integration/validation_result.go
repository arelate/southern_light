package vangogh_integration

type ValidationResult int

const (
	ValidationResultUnknown ValidationResult = iota
	ValidatedSuccessfully
	ValidatedWithGeneratedChecksum
	ValidatedUnresolvedManualUrl
	ValidatedMissingLocalFile
	ValidatedMissingChecksum
	ValidationError
	ValidatedChecksumMismatch
)

var validationResultsStrings = map[ValidationResult]string{
	ValidationResultUnknown:        "unknown",
	ValidatedSuccessfully:          "valid",
	ValidatedWithGeneratedChecksum: "valid-with-gen-checksum",
	ValidatedUnresolvedManualUrl:   "unresolved-manual-url",
	ValidatedMissingLocalFile:      "missing-local-file",
	ValidatedMissingChecksum:       "missing-checksum",
	ValidationError:                "error",
	ValidatedChecksumMismatch:      "checksum-mismatch",
}

var validationResultsHumanReadableStrings = map[ValidationResult]string{
	ValidationResultUnknown:        "Unknown",
	ValidatedSuccessfully:          "Validated",
	ValidatedWithGeneratedChecksum: "Gen-Valid",
	ValidatedUnresolvedManualUrl:   "Unresolved",
	ValidatedMissingLocalFile:      "No File",
	ValidatedMissingChecksum:       "No Checksum",
	ValidationError:                "Error",
	ValidatedChecksumMismatch:      "Corrupted",
}

var ValidationResultsOrder = []ValidationResult{
	ValidatedSuccessfully,
	ValidatedWithGeneratedChecksum,
	ValidatedUnresolvedManualUrl,
	ValidatedMissingLocalFile,
	ValidatedMissingChecksum,
	ValidationError,
	ValidatedChecksumMismatch,
	ValidationResultUnknown,
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
