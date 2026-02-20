package steam_vdf

import (
	"encoding/binary"
	"io"
	"os"
	"slices"
)

func (kv *KeyValues) writeBinaryString(w io.Writer, s string) error {
	for _, r := range s {
		if err := binary.Write(w, binary.LittleEndian, byte(r)); err != nil {
			return err
		}
	}
	if err := binary.Write(w, binary.LittleEndian, byte(BinaryTypeNone)); err != nil {
		return err
	}
	return nil
}

func (kv *KeyValues) WriteBinary(w io.Writer) error {

	if err := binary.Write(w, binary.LittleEndian, kv.Type); err != nil {
		return err
	}

	if err := kv.writeBinaryString(w, kv.Key); err != nil {
		return err
	}

	if kv.TypedValue != nil {
		switch kv.Type {
		case BinaryTypeString:
			if err := kv.writeBinaryString(w, kv.TypedValue.(string)); err != nil {
				return err
			}
		default:
			if err := binary.Write(w, binary.LittleEndian, kv.TypedValue); err != nil {
				return err
			}
		}
	}

	for _, val := range kv.Values {
		if err := val.WriteBinary(w); err != nil {
			return err
		}
	}

	if kv.Type == BinaryTypeNone {
		if err := binary.Write(w, binary.LittleEndian, BinaryTypeNullMarker); err != nil {
			return err
		}
	}

	return nil
}

func WriteBinary(writer io.Writer, vdf ValveDataFile) error {
	for _, kv := range vdf {
		if err := kv.WriteBinary(writer); err != nil {
			return err
		}
	}

	// write top level null marker
	return binary.Write(writer, binary.LittleEndian, BinaryTypeNullMarker)
}

func CreateBinary(dstPath string, vdf ValveDataFile, wo ...VdfWriteOptions) error {

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

	return WriteBinary(vdfFile, vdf)
}
