package vangogh_integration

import (
	"github.com/boggydigital/pathways"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var excludeFiles = map[string]bool{
	".DS_Store":   true, // https://en.wikipedia.org/wiki/.DS_Store
	"desktop.ini": true, // https://en.wikipedia.org/wiki/INI_file#History

}

var excludeDirs = map[string]bool{
	"@eaDir":                    true, // https://kb.synology.com/en-us/DSM/help/FileStation/connect?version=7
	"@sharebin":                 true, // https://kb.synology.com/en-us/DSM/help/FileStation/connect?version=7
	"@tmp":                      true, // https://kb.synology.com/en-us/DSM/help/FileStation/connect?version=7
	".SynologyWorkingDirectory": true, // https://kb.synology.com/en-us/DSM/help/FileStation/connect?version=7
}

func filenameAsId(p string) (string, error) {
	_, idFile := path.Split(p)
	if !strings.HasSuffix(idFile, ".download") {
		return strings.TrimSuffix(idFile, path.Ext(idFile)), nil
	}
	return "", nil
}

func LocalImageIds() (map[string]bool, error) {
	idp, err := pathways.GetAbsDir(Images)
	if err != nil {
		return nil, err
	}
	return walkFiles(idp, filenameAsId)
}

func RecycleBinDirs() (map[string]bool, error) {
	rbdp, err := pathways.GetAbsDir(RecycleBin)
	if err != nil {
		return nil, err
	}
	return walkDirectories(rbdp)
}

func RecycleBinFiles() (map[string]bool, error) {
	rbdp, err := pathways.GetAbsDir(RecycleBin)
	if err != nil {
		return nil, err
	}
	return walkFiles(rbdp, relRecycleBinPath)
}

func LocalDownloadDirs() (map[string]bool, error) {
	ddp, err := pathways.GetAbsDir(Downloads)
	if err != nil {
		return nil, err
	}
	return walkDirectories(ddp)
}

func LocalSlugDownloads(slug string) (map[string]bool, error) {
	pDir, err := AbsProductDownloadsDir(slug)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(pDir); os.IsNotExist(err) {
		return map[string]bool{}, nil
	}
	return walkFiles(
		pDir,
		func(p string) (string, error) {
			return filepath.Rel(pDir, p)
		})
}

func walkFiles(dir string, transformDelegate func(string) (string, error)) (map[string]bool, error) {
	fileSet := make(map[string]bool)
	err := filepath.WalkDir(
		dir,
		func(p string, de fs.DirEntry, err error) error {
			if de != nil && de.IsDir() {
				return nil
			}
			dn, fn := filepath.Split(p)
			ld := filepath.Base(dn)
			if excludeDirs[ld] {
				return nil
			}
			if excludeFiles[fn] {
				return nil
			}
			tPath, err := transformDelegate(p)
			if err != nil {
				return err
			}
			if tPath != "" {
				fileSet[tPath] = true
			}
			return err
		})

	return fileSet, err
}

func walkDirectories(rootDir string) (map[string]bool, error) {
	rbdp, err := pathways.GetAbsDir(RecycleBin)
	if err != nil {
		return nil, err
	}
	dirSet := make(map[string]bool)
	err = filepath.WalkDir(
		rootDir,
		func(p string, de fs.DirEntry, err error) error {
			if de != nil && !de.IsDir() {
				return nil
			}
			if p == "" || p == rbdp {
				return nil
			}
			dirSet[p] = true
			return nil
		})

	return dirSet, err
}
