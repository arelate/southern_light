package epic_games

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"io"
)

const magicBits uint32 = 0x44BEC00C

const (
	StorageCompressed uint8 = 0x01
	StorageEncrypted  uint8 = 0x02
)

type Manifest struct {
	Header   *Header
	Metadata *Metadata
}

type Header struct {
	HeaderSize       uint32
	SizeUncompressed uint32
	SizeCompressed   uint32
	ShaHash          []byte
	Storage          uint8
	FeatureLevel     uint32
}

type Metadata struct {
	MetadataSize    uint32
	MetadataVersion uint8
	FeatureLevel    uint32
	IsFileData      bool
	AppId           uint32
	AppName         string
	BuildVersion    string
	LaunchExe       string
	LaunchCommand   string
	PrereqIds       []string
	PrereqName      string
	PrereqPath      string
	PrereqArgs      string
	BuildId         string
}

func (h *Header) IsCompressed() bool {
	return h.Storage&StorageCompressed != 0
}

func (h *Header) IsEncrypted() bool {
	return h.Storage&StorageEncrypted != 0
}

func readUint8(r io.Reader) (val uint8, err error) {
	err = binary.Read(r, binary.LittleEndian, &val)
	return val, err
}

func readUint32(r io.Reader) (val uint32, err error) {
	err = binary.Read(r, binary.LittleEndian, &val)
	return val, err
}

func readUint64(r io.Reader) (val uint64, err error) {
	err = binary.Read(r, binary.LittleEndian, &val)
	return val, err
}

func readBool(r io.Reader) (bool, error) {
	ui8, err := readUint8(r)
	return ui8 != 0, err
}

func readBytes(r io.Reader, n int) ([]byte, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(r, buf)
	return buf, err
}

func readString(r io.Reader) (string, error) {
	if length, err := readUint32(r); err != nil || length == 0 {
		return "", err
	} else {
		var data []byte
		if data, err = readBytes(r, int(length)); err != nil {
			return "", err
		} else if readLen := len(data) - 1; data[readLen] != 0 {
			return "", errors.New("string must be null terminated")
		} else {
			return string(data[:readLen]), nil
		}
	}
}

func readMagic(r io.Reader) error {

	var magic uint32

	if err := binary.Read(r, binary.LittleEndian, &magic); err != nil {
		return err
	}

	if magic != magicBits {
		return errors.New("unsupported binary manifest format")
	}

	return nil
}

func readHeader(r io.Reader) (*Header, error) {

	var header Header
	var err error

	if header.HeaderSize, err = readUint32(r); err != nil {
		return nil, err
	}

	if header.SizeUncompressed, err = readUint32(r); err != nil {
		return nil, err
	}

	if header.SizeCompressed, err = readUint32(r); err != nil {
		return nil, err
	}

	if header.ShaHash, err = readBytes(r, 20); err != nil {
		return nil, err
	}

	if header.Storage, err = readUint8(r); err != nil {
		return nil, err
	}

	if header.FeatureLevel, err = readUint32(r); err != nil {
		return nil, err
	}

	return &header, nil
}

func readMetadata(r io.Reader) (*Metadata, error) {

	var metadata Metadata
	var err error

	if metadata.MetadataSize, err = readUint32(r); err != nil {
		return nil, err
	}

	if metadata.MetadataVersion, err = readUint8(r); err != nil {
		return nil, err
	}

	if metadata.FeatureLevel, err = readUint32(r); err != nil {
		return nil, err
	}

	if metadata.IsFileData, err = readBool(r); err != nil {
		return nil, err
	}

	if metadata.AppId, err = readUint32(r); err != nil {
		return nil, err
	}

	if metadata.AppName, err = readString(r); err != nil {
		return nil, err
	}

	return &metadata, nil
}

func ReadManifest(r io.ReadSeeker) (*Manifest, error) {

	err := readMagic(r)
	if err != nil {
		return nil, err
	}

	manifest := new(Manifest)

	if manifest.Header, err = readHeader(r); err != nil {
		return nil, err
	}

	if _, err = r.Seek(int64(manifest.Header.HeaderSize), io.SeekStart); err != nil {
		return nil, err
	}

	if manifest.Header.IsEncrypted() {
		return nil, errors.New("cannot read encrypted manifest")
	}

	var reader io.ReadSeeker

	switch manifest.Header.IsCompressed() {
	case true:
		var zr io.ReadCloser
		if zr, err = zlib.NewReader(r); err != nil {
			return nil, err
		}

		var unzipped []byte
		if unzipped, err = io.ReadAll(zr); err != nil {
			return nil, err
		}

		if len(unzipped) != int(manifest.Header.SizeUncompressed) {
			return nil, errors.New("uncompressed data size mismatch")
		}

		reader = bytes.NewReader(unzipped)
	case false:
		reader = r
	}

	if manifest.Metadata, err = readMetadata(reader); err != nil {
		return nil, err
	}

	return manifest, nil
}
