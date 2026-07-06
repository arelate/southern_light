package vangogh_integration

import (
	"runtime"
)

func CurrentOs() OperatingSystem {
	switch runtime.GOOS {
	case "windows":
		return Windows
	case "darwin":
		return MacOS
	case "linux":
		return Linux
	default:
		panic("current os is not supported")
	}
}
