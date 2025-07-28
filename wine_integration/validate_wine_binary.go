package wine_integration

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/arelate/southern_light/vangogh_integration"
	"github.com/boggydigital/dolo"
	"github.com/boggydigital/nod"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ValidateWineBinary(binary *vangogh_integration.WineBinaryDetails, absBinaryDir string, since time.Time, force bool) error {

	vwba := nod.NewProgress(" - %s...", binary.Filename)
	defer vwba.Done()

	absFilename := filepath.Join(absBinaryDir, binary.Filename)

	stat, err := os.Stat(absFilename)
	if err != nil {
		return err
	}

	if stat.ModTime().Before(since) && !force {
		vwba.EndWithResult("unchanged")
		return nil
	}

	if binary.Digest == "" {
		vwba.EndWithResult("digest not available")
		return nil
	}

	wbFile, err := os.Open(absFilename)
	if err != nil {
		return err
	}

	defer wbFile.Close()

	vwba.Total(uint64(stat.Size()))

	if algorithm, dstHash, ok := strings.Cut(binary.Digest, ":"); ok {

		switch algorithm {
		case "sha256":

			h := sha256.New()
			if err = dolo.CopyWithProgress(h, wbFile, vwba); err != nil {
				return err
			}

			srcHash := fmt.Sprintf("%x", h.Sum(nil))

			if srcHash == dstHash {
				vwba.EndWithResult("valid")
				return nil
			} else {
				return errors.New("digest mismatch")
			}

		default:
			return errors.New("unknown digest algorithm: " + algorithm)
		}

	} else {
		return errors.New("unknown digest format: " + binary.Digest)
	}
}
