package egs_integration

import (
	"fmt"
	"path/filepath"

	"github.com/google/uuid"
)

const manifestMagic uint32 = 0x44BEC00C

const (
	StorageUncompressed uint8 = 0x00
	StorageCompressed   uint8 = 0x01
	StorageEncrypted    uint8 = 0x02
)

type Manifest struct {
	Header       *Header
	Metadata     *Metadata
	ChunkList    *ChunkList
	FileList     *FileList
	CustomFields *CustomFields
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

func chunkDir(featureLevel uint32) string {
	if featureLevel < 3 {
		return "Chunks"
	} else if featureLevel < 6 {
		return "ChunksV2"
	} else if featureLevel < 15 {
		return "ChunksV3"
	}
	return "ChunksV4"
}

func (chk *Chunk) Path(featureLevel uint32) string {
	base := fmt.Sprintf("%02d/%016X_%X.chunk", chk.Group, chk.Hash, chk.Uuid[:])
	return filepath.Join(chunkDir(featureLevel), base)
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
