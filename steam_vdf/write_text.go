package steam_vdf

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

type VdfWriteOptions int

const (
	VdfBackupExisting VdfWriteOptions = iota
	VdfTabsIndent
	VdfSpacesIndent
)

const defaultVdfIndent = "\t"

const Ext = ".vdf"

func (kv *KeyValues) WriteString(w io.Writer, depth int, wo ...VdfWriteOptions) error {

	if slices.Contains(wo, VdfTabsIndent) && slices.Contains(wo, VdfSpacesIndent) {
		return errors.New("cannot use both tabs and spaces as VDF indent")
	}

	var indent string
	if slices.Contains(wo, VdfTabsIndent) {
		indent = "\t"
	}
	if slices.Contains(wo, VdfSpacesIndent) {
		indent = " "
	}

	if indent == "" {
		indent = defaultVdfIndent
	}

	// (depth times \t)"key"
	if _, err := io.WriteString(w, strings.Repeat(indent, depth)); err != nil {
		return err
	}
	if _, err := io.WriteString(w, fmt.Sprintf("\"%s\"", kv.Key)); err != nil {
		return err
	}

	// \t\t"value"
	if kv.Value != nil {
		if _, err := io.WriteString(w, strings.Repeat(indent, 2)); err != nil {
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
		if _, err := io.WriteString(w, strings.Repeat(indent, depth)); err != nil {
			return err
		}
		if _, err := io.WriteString(w, string(leftCurlyBracket)); err != nil {
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
		if _, err := io.WriteString(w, strings.Repeat(indent, depth)); err != nil {
			return err
		}
		if _, err := io.WriteString(w, string(rightCurlyBracket)); err != nil {
			return err
		}
		if _, err := io.WriteString(w, "\n"); err != nil {
			return err
		}
	}

	return nil
}

func WriteText(writer io.Writer, vdf ValveDataFile, wo ...VdfWriteOptions) error {

	for _, kv := range vdf {
		if err := kv.WriteString(writer, 0, wo...); err != nil {
			return err
		}
	}

	return nil
}

func CreateText(dstPath string, vdf ValveDataFile, wo ...VdfWriteOptions) error {

	if slices.Contains(wo, VdfBackupExisting) {
		if _, err := os.Stat(dstPath); err == nil {
			if err = backup(dstPath); err != nil {
				return err
			}
		}
	}

	vdfFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer vdfFile.Close()

	return WriteText(vdfFile, vdf)
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

	if _, err = io.Copy(backupFile, originalFile); err != nil {
		return err
	}

	return nil
}
