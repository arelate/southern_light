package egs_integration

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"path/filepath"

	"github.com/google/uuid"
)

const magicBits uint32 = 0x44BEC00C

const (
	StorageCompressed uint8 = 0x01
	StorageEncrypted  uint8 = 0x02
)

type Manifest struct {
	Header       *Header
	Metadata     *Metadata
	ChunkList    *ChunkList
	FileList     *FileList
	CustomFields *CustomFields
}

func (mft *Manifest) Path(chk *Chunk) string {
	return filepath.Join(mft.Metadata.ChunksVersion(), chk.Base())
}

type Header struct {
	Offset           uint32
	SizeUncompressed uint32
	SizeCompressed   uint32
	ShaHash          []byte
	Storage          uint8
	FeatureLevel     uint32
}

type Metadata struct {
	Offset        uint32
	Version       uint8
	FeatureLevel  uint32
	IsFileData    bool
	AppId         uint32
	AppName       string
	BuildVersion  string
	LaunchExe     string
	LaunchCommand string
	PrereqIds     []string
	PrereqName    string
	PrereqPath    string
	PrereqArgs    string
	BuildId       string
}

func (mdt *Metadata) ChunksVersion() string {
	if mdt.Version < 3 {
		return "Chunks"
	} else if mdt.Version < 6 {
		return "ChunksV2"
	} else if mdt.Version < 15 {
		return "ChunksV3"
	}
	return "ChunksV4"
}

type ChunkList struct {
	Offset  uint32
	Version uint8
	Count   uint32
	Chunks  []*Chunk
	Lookup  map[uuid.UUID]uint32
}

type Chunk struct {
	Uuid       uuid.UUID
	Hash       uint64
	ShaHash    []byte
	Group      uint8
	WindowSize uint32
	FileSize   uint64
}

func (chk *Chunk) Base() string {
	return fmt.Sprintf("%02d/%016X_%X.chunk", chk.Group, chk.Hash, chk.Uuid[:])
}

type FileList struct {
	Offset  uint32
	Version uint8
	Count   uint32
	List    []File
}

type File struct {
	Filename      string
	SymlinkTarget string
	ShaHash       []byte
	Flags         uint8
	InstallTags   []string
	Size          uint64
	Parts         []ChunkPart
}

type ChunkPart struct {
	DataSize   uint32
	ParentUuid uuid.UUID
	Offset     uint32
	Size       uint32
	Chunk      *Chunk
}

type CustomFields struct {
	Offset  uint32
	Version uint8
	Count   uint32
	Fields  map[string]string
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

func readStringsSlice(r io.Reader) ([]string, error) {
	if length, err := readUint32(r); err != nil {
		return nil, err
	} else {
		strSlice := make([]string, 0, length)
		for range length {
			var str string
			if str, err = readString(r); err == nil {
				strSlice = append(strSlice, str)
			} else {
				return nil, err
			}
		}
		return strSlice, nil
	}
}

func readUuid(r io.Reader) (val uuid.UUID, err error) {
	data := make([]uint32, 4)
	err = binary.Read(r, binary.BigEndian, &data)
	if err != nil {
		return uuid.Nil, err
	}
	for ii, vv := range data {
		binary.LittleEndian.PutUint32(val[ii*4:(ii+1)*4], vv)
	}
	return val, err
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

	if header.Offset, err = readUint32(r); err != nil {
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

	if metadata.Offset, err = readUint32(r); err != nil {
		return nil, err
	}
	if metadata.Version, err = readUint8(r); err != nil {
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
	if metadata.BuildVersion, err = readString(r); err != nil {
		return nil, err
	}
	if metadata.LaunchExe, err = readString(r); err != nil {
		return nil, err
	}
	if metadata.LaunchCommand, err = readString(r); err != nil {
		return nil, err
	}
	if metadata.PrereqIds, err = readStringsSlice(r); err != nil {
		return nil, err
	}
	if metadata.PrereqPath, err = readString(r); err != nil {
		return nil, err
	}
	if metadata.PrereqArgs, err = readString(r); err != nil {
		return nil, err
	}
	if metadata.Version >= 1 {
		if metadata.BuildId, err = readString(r); err != nil {
			return nil, err
		}
	}

	return &metadata, nil
}

func readChunkList(r io.Reader) (*ChunkList, error) {

	var list ChunkList
	var err error

	if list.Offset, err = readUint32(r); err != nil {
		return nil, err
	}
	if list.Version, err = readUint8(r); err != nil {
		return nil, err
	}
	if list.Count, err = readUint32(r); err != nil {
		return nil, err
	}

	list.Chunks = make([]*Chunk, list.Count)
	for ii := range list.Count {
		list.Chunks[ii] = new(Chunk)
	}
	list.Lookup = make(map[uuid.UUID]uint32)

	for ii, chk := range list.Chunks {
		if chk.Uuid, err = readUuid(r); err == nil {
			list.Lookup[chk.Uuid] = uint32(ii)
		} else {
			return nil, err
		}
	}

	for _, chk := range list.Chunks {
		if chk.ShaHash, err = readBytes(r, 20); err != nil {
			return nil, err
		}
	}

	for _, chk := range list.Chunks {
		if chk.Group, err = readUint8(r); err != nil {
			return nil, err
		}
	}

	for _, chk := range list.Chunks {
		if chk.WindowSize, err = readUint32(r); err != nil {
			return nil, err
		}
	}

	for _, chk := range list.Chunks {
		if chk.FileSize, err = readUint64(r); err != nil {
			return nil, err
		}
	}

	return &list, nil
}

func readFileList(r io.Reader, chunkList *ChunkList) (*FileList, error) {

	var list FileList
	var err error

	if list.Offset, err = readUint32(r); err != nil {
		return nil, err
	}
	if list.Version, err = readUint8(r); err != nil {
		return nil, err
	}
	if list.Count, err = readUint32(r); err != nil {
		return nil, err
	}

	list.List = make([]File, list.Count)

	for ii := range list.List {
		if list.List[ii].Filename, err = readString(r); err != nil {
			return nil, err
		}
	}

	for ii := range list.List {
		if list.List[ii].SymlinkTarget, err = readString(r); err != nil {
			return nil, err
		}
	}

	for ii := range list.List {
		if list.List[ii].ShaHash, err = readBytes(r, 20); err != nil {
			return nil, err
		}
	}

	for ii := range list.List {
		if list.List[ii].Flags, err = readUint8(r); err != nil {
			return nil, err
		}
	}

	for ii := range list.List {
		if list.List[ii].InstallTags, err = readStringsSlice(r); err != nil {
			return nil, err
		}
	}

	for ii := range list.List {
		var partSize uint32
		if partSize, err = readUint32(r); err != nil {
			return nil, err
		}

		list.List[ii].Parts = make([]ChunkPart, partSize)

		for jj := range list.List[ii].Parts {
			if list.List[ii].Parts[jj].DataSize, err = readUint32(r); err != nil {
				return nil, err
			}
			if list.List[ii].Parts[jj].ParentUuid, err = readUuid(r); err != nil {
				return nil, err
			}
			if chunkId, ok := chunkList.Lookup[list.List[ii].Parts[jj].ParentUuid]; ok {
				list.List[ii].Parts[jj].Chunk = chunkList.Chunks[chunkId]
			} else {
				return nil, errors.New("parent UUID not found")
			}
			if list.List[ii].Parts[jj].Offset, err = readUint32(r); err != nil {
				return nil, err
			}
			if list.List[ii].Parts[jj].Size, err = readUint32(r); err != nil {
				return nil, err
			}

		}
	}

	for ii := range list.List {
		var totalSize uint64
		for jj := range list.List[ii].Parts {
			totalSize += uint64(list.List[ii].Parts[jj].Size)
		}
		list.List[ii].Size = totalSize
	}

	return &list, nil
}

func readCustomFields(r io.ReadSeeker) (*CustomFields, error) {

	var fields CustomFields
	var err error

	if fields.Offset, err = readUint32(r); err != nil {
		return nil, err
	}
	if fields.Version, err = readUint8(r); err != nil {
		return nil, err
	}
	if fields.Count, err = readUint32(r); err != nil {
		return nil, err
	}

	keys := make([]string, fields.Count)
	for ii := range keys {
		if keys[ii], err = readString(r); err != nil {
			return nil, err
		}
	}

	fields.Fields = make(map[string]string)
	for _, key := range keys {
		if fields.Fields[key], err = readString(r); err != nil {
			return nil, err
		}
	}

	return &fields, err
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

	if _, err = r.Seek(int64(manifest.Header.Offset), io.SeekStart); err != nil {
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

	savedPos, err := reader.Seek(0, io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	if manifest.Metadata, err = readMetadata(reader); err != nil {
		return nil, err
	}

	if savedPos, err = reader.Seek(savedPos+int64(manifest.Metadata.Offset), io.SeekStart); err != nil {
		return nil, err
	}

	if manifest.ChunkList, err = readChunkList(reader); err != nil {
		return nil, err
	}

	if savedPos, err = reader.Seek(savedPos+int64(manifest.ChunkList.Offset), io.SeekStart); err != nil {
		return nil, err
	}

	if manifest.FileList, err = readFileList(reader, manifest.ChunkList); err != nil {
		return nil, err
	}

	if savedPos, err = reader.Seek(savedPos+int64(manifest.FileList.Offset), io.SeekStart); err != nil {
		return nil, err
	}

	if manifest.CustomFields, err = readCustomFields(reader); err != nil {
		return nil, err
	}

	return manifest, nil
}
