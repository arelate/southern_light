package vangogh_integration

import (
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

func LocalImageIds() (map[string]any, error) {
	return walkFiles(Pwd.AbsDirPath(Images), filenameAsId)
}

func LocalDownloadDirs() (map[string]any, error) {
	return walkDirectories(Pwd.AbsDirPath(Downloads))
}

func AbsLocalSlugDownloads(slug string, dl DownloadsLayout) (map[string]any, error) {
	// using root product slug download dir to walk all files under it
	absSlugDownloadDir, err := AbsSlugDownloadDir(slug, Installer, dl)
	if err != nil {
		return nil, err
	}
	if _, err = os.Stat(absSlugDownloadDir); os.IsNotExist(err) {
		return map[string]any{}, nil
	}
	return walkFiles(
		absSlugDownloadDir,
		func(p string) (string, error) {
			return p, nil
		})
}

func walkFiles(dir string, transformDelegate func(string) (string, error)) (map[string]any, error) {
	fileSet := make(map[string]any)
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
				fileSet[tPath] = nil
			}
			return err
		})

	return fileSet, err
}

func walkDirectories(rootDir string) (map[string]any, error) {
	dirSet := make(map[string]any)
	err := filepath.WalkDir(
		rootDir,
		func(p string, de fs.DirEntry, err error) error {
			if de != nil && !de.IsDir() {
				return nil
			}
			if p == "" {
				return nil
			}
			dirSet[p] = nil
			return nil
		})

	return dirSet, err
}
