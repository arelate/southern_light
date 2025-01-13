package vangogh_integration

import (
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func MoveToRecycleBin(typeRootDir, absPath string) error {
	if absPath == "" {
		nod.Log("move to recycle bin: empty filename")
		return nil
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		nod.Log("move to recycle bin: file not found: %s", absPath)
		return nil
	}

	root, _ := filepath.Split(typeRootDir)

	relPath, err := filepath.Rel(root, absPath)
	if err != nil {
		nod.Log("move to recycle bin: relPath error: root=%s target=%s", root, absPath)
		return err
	}

	rbdp, err := pathways.GetAbsDir(RecycleBin)
	if err != nil {
		return err
	}

	rbFilepath := filepath.Join(rbdp, relPath)
	rbDir, _ := filepath.Split(rbFilepath)
	if _, err := os.Stat(rbDir); os.IsNotExist(err) {
		if err := os.MkdirAll(rbDir, 0755); err != nil {
			return err
		}
	}

	if err := os.Rename(absPath, rbFilepath); err != nil {
		// inspired by https://github.com/golang/go/issues/41487
		if strings.Contains(err.Error(), "invalid cross-device link") {
			if err := copyDelete(absPath, rbFilepath); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func copyDelete(src, dst string) error {
	srcf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcf.Close()

	dstf, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstf.Close()

	_, err = io.Copy(dstf, srcf)
	if err != nil {
		return err
	}

	return os.Remove(src)
}
