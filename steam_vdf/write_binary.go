package steam_vdf

import (
	"encoding/binary"
	"io"
	"os"
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

func WriteBinary(p string, keyValues []*KeyValues) error {

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
		if err := kv.WriteBinary(vdfFile); err != nil {
			return err
		}
	}

	// write top level null marker
	if err := binary.Write(vdfFile, binary.LittleEndian, BinaryTypeNullMarker); err != nil {
		return err
	}

	return nil
}
