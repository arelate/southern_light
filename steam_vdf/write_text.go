package steam_vdf

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (kv *KeyValues) WriteString(w io.Writer, depth int) error {

	// (depth times \t)"key"
	if _, err := io.WriteString(w, strings.Repeat("\t", depth)); err != nil {
		return err
	}
	if _, err := io.WriteString(w, fmt.Sprintf("\"%s\"", kv.Key)); err != nil {
		return err
	}

	// \t\t"value"
	if kv.Value != nil {
		if _, err := io.WriteString(w, strings.Repeat("\t", 2)); err != nil {
			return err
		}
		if _, err := io.WriteString(w, fmt.Sprintf("\"%s\"", *kv.Value)); err != nil {
			return err
		}
	}

	// \n
	if _, err := io.WriteString(w, "\n"); err != nil {
		return err
	}

	// (depth times \t){\n
	if kv.Values != nil || (kv.Values == nil && kv.Value == nil) {
		if _, err := io.WriteString(w, strings.Repeat("\t", depth)); err != nil {
			return err
		}
		if _, err := io.WriteString(w, "{"); err != nil {
			return err
		}
		if _, err := io.WriteString(w, "\n"); err != nil {
			return err
		}
	}

	// write Values at depth + 1
	for _, val := range kv.Values {
		if err := val.WriteString(w, depth+1); err != nil {
			return err
		}
	}

	// (depth times \t)}\n
	if kv.Values != nil || (kv.Values == nil && kv.Value == nil) {
		if _, err := io.WriteString(w, strings.Repeat("\t", depth)); err != nil {
			return err
		}
		if _, err := io.WriteString(w, "}"); err != nil {
			return err
		}
		if _, err := io.WriteString(w, "\n"); err != nil {
			return err
		}
	}

	return nil
}

func WriteText(p string, keyValues []*KeyValues) error {

	if _, err := os.Stat(p); err == nil {
		if err := backup(p); err != nil {
			return err
		}
	}

	vdfFile, err := os.Create(p)
	if err != nil {
		return err
	}
	defer vdfFile.Close()

	for _, kv := range keyValues {
		if err := kv.WriteString(vdfFile, 0); err != nil {
			return err
		}
	}

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
