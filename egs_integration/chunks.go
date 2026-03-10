package egs_integration

import "github.com/google/uuid"

const chunkMagic = 0xB1FE3AA2

type ChunkHeader struct {
	Magic          uint32
	Version        uint32
	Offset         uint32
	SizeCompressed uint32
	Uuid           uuid.UUID
	RollingHash    uint64
	Storage        uint8
	ShaHash        []byte
	HashType       uint32
}

func (ch *ChunkHeader) IsUncompressed() bool {
	return ch.Storage&StorageUncompressed != 0
}

func (ch *ChunkHeader) IsCompressed() bool {
	return ch.Storage&StorageCompressed != 0
}

func (ch *ChunkHeader) IsEncrypted() bool {
	return ch.Storage&StorageEncrypted != 0
}
