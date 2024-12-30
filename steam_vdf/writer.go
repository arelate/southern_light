package steam_vdf

import (
	"io"
	"os"
	"path/filepath"
	"time"
)

func Write(p string, keyValues []*KeyValue) error {

	//if _, err := os.Stat(p); err == nil {
	//	if err := backup(p); err != nil {
	//		return err
	//	}
	//}
	//
	//
	//
	//vdfFile, err := os.Create(p)
	//if err != nil {
	//	return err
	//}
	//defer vdfFile.Close()

	return nil
}

func backup(path string) error {

	timestamp := time.Now().Format("2006-01-02-15-04-05")

	originalFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer originalFile.Close()

	backupPath := filepath.Join(path + "." + timestamp)

	backupFile, err := os.Create(backupPath)
	if err != nil {
		return err
	}
	defer backupFile.Close()

	if _, err := io.Copy(backupFile, originalFile); err != nil {
		return err
	}

	return nil
}
