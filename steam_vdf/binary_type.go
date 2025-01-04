package steam_vdf

type BinaryType byte

// https://github.com/ValveSoftware/source-sdk-2013/blob/master/sp/src/public/tier1/kvpacker.h
const (
	BinaryTypeNone BinaryType = iota
	BinaryTypeString
	BinaryTypeInt
	BinaryTypeFloat
	BinaryTypePtr
	BinaryTypeWString
	BinaryTypeColor
	BinaryTypeUint64
	BinaryTypeNullMarker
	BinaryTypeError BinaryType = 254
	BinaryTypeEOF   BinaryType = 255
)

func IsKnownBinaryType(b BinaryType) bool {
	return b <= BinaryTypeNullMarker
}
