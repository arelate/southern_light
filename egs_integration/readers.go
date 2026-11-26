package egs_integration

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/binary"
	"encoding/json/v2"
	"errors"
	"io"
	"unicode/utf16"
	"uuid"
)

const (
	maxStringByteLen   = 1 << 20
	chunkPartFixedSize = 28
)

var errUnsupportedBinaryManifestFormat = errors.New("unsupported binary manifest format")

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

func readBytes(r io.Reader, n uint) ([]byte, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(r, buf)
	return buf, err
}

func readString(r io.Reader) (string, error) {

	if length, err := readUint32(r); err != nil {
		return "", err
	} else if length == 0 {
		return "", nil
	} else if int32(length) < 0 {
		return readUtf16String(r, -int32(length))
	} else if length > maxStringByteLen {
		return "", errors.New("manifest string is too long")
	} else {
		data, err := readBytes(r, uint(length))
		if err != nil {
			return "", err
		}

		readLen := len(data) - 1
		if data[readLen] != 0 {
			return "", errors.New("string must be null terminated")
		}

		return string(data[:readLen]), nil
	}
}

func readUtf16String(r io.Reader, charCount int32) (string, error) {

	byteLen := int(charCount) * 2
	if byteLen > maxStringByteLen {
		return "", errors.New("manifest utf16 string is too long")
	}

	// UTF-16 characters are 2 bytes...
	if data, err := readBytes(r, uint(byteLen)); err != nil {
		return "", err
	} else {

		// ...plus a 2-byte null terminator
		readLen := len(data) - 2
		if data[readLen] != 0 || data[readLen+1] != 0 {
			return "", errors.New("utf16 string must be null terminated")
		}

		units := make([]uint16, readLen/2)
		for ii := range units {
			units[ii] = binary.LittleEndian.Uint16(data[ii*2:])
		}

		return string(utf16.Decode(units)), nil
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
		return uuid.Nil(), err
	}
	for ii, vv := range data {
		binary.LittleEndian.PutUint32(val[ii*4:(ii+1)*4], vv)
	}
	return val, err
}

func readManifestMagic(r io.Reader) error {

	var magic uint32

	if err := binary.Read(r, binary.LittleEndian, &magic); err != nil {
		return err
	}

	if magic != manifestMagic {
		return errUnsupportedBinaryManifestFormat
	}

	return nil
}

func readHeader(r io.Reader) (*Header, error) {

	var header Header
	var err error

	if header.Size, err = readUint32(r); err != nil {
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

	if metadata.Size, err = readUint32(r); err != nil {
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
	if metadata.PrereqName, err = readString(r); err != nil {
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
	if metadata.Version >= 2 {
		if metadata.UninstallActionPath, err = readString(r); err != nil {
			return nil, err
		}
		if metadata.UninstallActionArgs, err = readString(r); err != nil {
			return nil, err
		}
	}

	return &metadata, nil
}

func readChunkList(r io.Reader) (*ChunkList, error) {

	var list ChunkList
	var err error

	if list.Size, err = readUint32(r); err != nil {
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
		if chk.Hash, err = readUint64(r); err != nil {
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

	if list.Size, err = readUint32(r); err != nil {
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
			if list.List[ii].Parts[jj].DataSize < chunkPartFixedSize {
				return nil, errors.New("manifest chunk part too small")
			}

			var parentUuid uuid.UUID
			if parentUuid, err = readUuid(r); err != nil {
				return nil, err
			}

			list.List[ii].Parts[jj].ChunkUuid = parentUuid
			if chunkId, ok := chunkList.Lookup[parentUuid]; ok {
				list.List[ii].Parts[jj].Chunk = chunkList.Chunks[chunkId]
			} else {
				return nil, errors.New("manifest parent UUID not found")
			}

			if list.List[ii].Parts[jj].Offset, err = readUint32(r); err != nil {
				return nil, err
			}
			if list.List[ii].Parts[jj].Size, err = readUint32(r); err != nil {
				return nil, err
			}

			if extra := list.List[ii].Parts[jj].DataSize - chunkPartFixedSize; extra > 0 {
				if _, err = readBytes(r, uint(extra)); err != nil {
					return nil, err
				}
			}

		}
	}

	if list.Version >= 1 {
		for ii := range list.List {
			var hasMd5 uint32
			if hasMd5, err = readUint32(r); err != nil {
				return nil, err
			}
			if hasMd5 != 0 {
				if list.List[ii].HashMd5, err = readBytes(r, 16); err != nil {
					return nil, err
				}
			}
		}
		for ii := range list.List {
			if list.List[ii].MimeType, err = readString(r); err != nil {
				return nil, err
			}
		}
	}

	if list.Version >= 2 {
		for ii := range list.List {
			if list.List[ii].HashSha256, err = readBytes(r, 32); err != nil {
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

	if fields.Size, err = readUint32(r); err != nil {
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

func readChunkHeader(r io.Reader) (*ChunkHeader, error) {

	var chunkHeader ChunkHeader
	var err error

	if chunkHeader.Magic, err = readUint32(r); err != nil {
		return nil, err
	}
	if chunkHeader.Magic != chunkMagic {
		return nil, errors.New("unsupported chunk format")
	}

	if chunkHeader.Version, err = readUint32(r); err != nil {
		return nil, err
	}
	if chunkHeader.Offset, err = readUint32(r); err != nil {
		return nil, err
	}
	if chunkHeader.SizeCompressed, err = readUint32(r); err != nil {
		return nil, err
	}
	if chunkHeader.Uuid, err = readUuid(r); err != nil {
		return nil, err
	}
	if chunkHeader.RollingHash, err = readUint64(r); err != nil {
		return nil, err
	}
	if chunkHeader.Storage, err = readUint8(r); err != nil {
		return nil, err
	}
	if chunkHeader.ShaHash, err = readBytes(r, 20); err != nil {
		return nil, err
	}
	if chunkHeader.HashType, err = readUint32(r); err != nil {
		return nil, err
	}

	return &chunkHeader, err
}

func ReadChunk(r io.ReadSeeker) (io.Reader, error) {

	var chunkHeader *ChunkHeader
	var err error

	if chunkHeader, err = readChunkHeader(r); err != nil {
		return nil, err
	}

	if chunkHeader.Version > 3 {
		return nil, errors.New("unsupported chunk version")
	}
	if _, err = r.Seek(int64(chunkHeader.Offset), io.SeekStart); err != nil {
		return nil, err
	}

	switch chunkHeader.Storage {
	case StorageUncompressed:
		return r, nil
	case StorageCompressed:
		return zlib.NewReader(r)
	case StorageEncrypted:
		return nil, errors.New("chunk is encrypted")
	default:
		return nil, errors.New("unsupported chunk storage")
	}
}

func readBinaryManifest(r io.ReadSeeker) (*Manifest, error) {

	err := readManifestMagic(r)
	if err != nil {
		return nil, err
	}

	manifest := new(Manifest)

	if manifest.Header, err = readHeader(r); err != nil {
		return nil, err
	}

	if _, err = r.Seek(int64(manifest.Header.Size), io.SeekStart); err != nil {
		return nil, err
	}

	var reader io.ReadSeeker

	switch manifest.Header.Storage {
	case StorageCompressed:
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

		if sum := sha1.Sum(unzipped); !bytes.Equal(sum[:], manifest.Header.ShaHash) {
			return nil, errors.New("manifest uncompressed data sha1 mismatch")
		}

		reader = bytes.NewReader(unzipped)
	case StorageUncompressed:
		reader = r
	case StorageEncrypted:
		return nil, errors.New("cannot read encrypted manifest")
	default:
		return nil, errors.New("unsupported manifest storage")
	}

	savedPos, err := reader.Seek(0, io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	if manifest.Metadata, err = readMetadata(reader); err != nil {
		return nil, err
	}

	if savedPos, err = reader.Seek(savedPos+int64(manifest.Metadata.Size), io.SeekStart); err != nil {
		return nil, err
	}

	if manifest.ChunkList, err = readChunkList(reader); err != nil {
		return nil, err
	}

	if savedPos, err = reader.Seek(savedPos+int64(manifest.ChunkList.Size), io.SeekStart); err != nil {
		return nil, err
	}

	if manifest.FileList, err = readFileList(reader, manifest.ChunkList); err != nil {
		return nil, err
	}

	if savedPos, err = reader.Seek(savedPos+int64(manifest.FileList.Size), io.SeekStart); err != nil {
		return nil, err
	}

	if manifest.CustomFields, err = readCustomFields(reader); err != nil {
		return nil, err
	}

	return manifest, nil
}

func readJsonManifest(r io.Reader) (*Manifest, error) {

	var jsonManifest JsonManifest

	if err := json.UnmarshalRead(r, &jsonManifest); err != nil {
		return nil, err
	}

	return jsonManifest.Manifest()
}

func ReadManifest(r io.ReadSeeker) (*Manifest, error) {

	if manifest, err := readBinaryManifest(r); err == nil {
		return manifest, nil
	} else if errors.Is(err, errUnsupportedBinaryManifestFormat) {
		if _, err = r.Seek(0, io.SeekStart); err != nil {
			return nil, err
		}
		return readJsonManifest(r)
	} else if err != nil {
		return nil, err
	}

	return nil, errors.New("unable to read manifest")
}
