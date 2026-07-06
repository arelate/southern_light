package vangogh_integration

import (
	"os"
	"path/filepath"
)

func UserDataHomeDir() (string, error) {

	currentOs := CurrentOs()

	switch currentOs {
	case Linux:
		uhd, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(uhd, ".local", "share"), nil
	case Windows:
		// TODO: verify that Windows user data home is also os.UserConfigDir
		fallthrough
	case MacOS:
		return os.UserConfigDir()
	default:
		panic(currentOs.ErrUnsupported())
	}
}
