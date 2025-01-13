package vangogh_integration

import "strings"

type OperatingSystem int

const (
	AnyOperatingSystem OperatingSystem = iota
	Windows
	MacOS
	Linux
)

var operatingSystemStrings = map[OperatingSystem]string{
	AnyOperatingSystem: "any-operating-system",
	Windows:            "Windows",
	MacOS:              "macOS",
	Linux:              "Linux",
}

func AllOperatingSystems() []OperatingSystem {
	all := make([]OperatingSystem, 1, len(operatingSystemStrings))
	// order is important here given this will be used for clo default parameter
	all[0] = AnyOperatingSystem
	for os, _ := range operatingSystemStrings {
		if os == AnyOperatingSystem {
			continue
		}
		all = append(all, os)
	}
	return all
}

func (os OperatingSystem) String() string {
	str, ok := operatingSystemStrings[os]
	if ok {
		return str
	}

	return operatingSystemStrings[AnyOperatingSystem]
}

func ParseOperatingSystem(operatingSystem string) OperatingSystem {
	operatingSystem = strings.ToLower(operatingSystem)
	for os, str := range operatingSystemStrings {
		if strings.ToLower(str) == operatingSystem {
			return os
		}
	}
	return AnyOperatingSystem
}

func ParseManyOperatingSystems(osStrings []string) []OperatingSystem {
	operatingSystems := make([]OperatingSystem, 0, len(osStrings))
	for _, osStr := range osStrings {
		os := ParseOperatingSystem(osStr)
		operatingSystems = append(operatingSystems, os)
	}
	return operatingSystems
}

func IsValidOperatingSystem(os OperatingSystem) bool {
	_, ok := operatingSystemStrings[os]
	return ok
}
