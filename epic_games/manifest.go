package epic_games

import (
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
	Header *Header
}

type Header struct {
	HeaderSize       uint32
	SizeUncompressed uint32
	SizeCompressed   uint32
	ShaHash          []byte
	Storage          uint8
	Version          uint32
}

func (h *Header) IsCompressed() bool {
	return h.Storage&StorageCompressed != 0
}

func (h *Header) IsEncrypted() bool {
	return h.Storage&StorageEncrypted != 0
}

func readUint8(r io.Reader) (uint8, error) {
	var v uint8
	err := binary.Read(r, binary.LittleEndian, &v)
	return v, err
}

func readUint32(r io.Reader) (uint32, error) {
	var v uint32
	err := binary.Read(r, binary.LittleEndian, &v)
	return v, err
}

func readUint64(r io.Reader) (uint64, error) {
	var v uint64
	err := binary.Read(r, binary.LittleEndian, &v)
	return v, err
}

func readBool(r io.Reader) (bool, error) {
	var b uint8
	err := binary.Read(r, binary.LittleEndian, &b)
	return b != 0, err
}

func readBytes(r io.Reader, n int) ([]byte, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(r, buf)
	return buf, err
}

func readString(r io.Reader) (string, error) {
	length, err := readUint32(r)
	if err != nil {
		return "", err
	}
	data, err := readBytes(r, int(length))
	if err != nil {
		return "", err
	}
	return string(data), nil
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

	if header.Version, err = readUint32(r); err != nil {
		return nil, err
	}

	return &header, nil
}

func ReadManifest(r io.ReadSeeker) (*Manifest, error) {

	err := readMagic(r)
	if err != nil {
		return nil, err
	}

	manifest := new(Manifest)

	manifest.Header, err = readHeader(r)
	if err != nil {
		return nil, err
	}

	if _, err = r.Seek(int64(manifest.Header.HeaderSize), io.SeekStart); err != nil {
		return nil, err
	}

	return manifest, nil
}
