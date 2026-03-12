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
