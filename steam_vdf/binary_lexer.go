package steam_vdf

import (
	"encoding/binary"
	"errors"
	"io"
	"unicode/utf16"
)

type binaryLexItem struct {
	typ BinaryType
	val any
}

type binaryLexStateFn func(*binaryLexer) binaryLexStateFn

type binaryLexer struct {
	r      io.Reader
	eof    bool
	expTyp BinaryType
	items  chan binaryLexItem
}

func newBinaryLexer(r io.Reader) *binaryLexer {
	l := &binaryLexer{
		r:     r,
		items: make(chan binaryLexItem),
	}
	go l.run()
	return l
}

func (bl *binaryLexer) run() {
	for st := lexBinaryType; st != nil; {
		st = st(bl)
	}
	close(bl.items)
}

func lexBinaryType(bl *binaryLexer) binaryLexStateFn {
	t, err := bl.readByte()
	if err != nil {
		return bl.errorf(err)
	}

	bit := BinaryType(t)

	if !IsKnownBinaryType(bit) {
		bl.emit(BinaryTypeError, errors.New("unknown binary type"))
		return nil
	}

	if bit == BinaryTypeNullMarker {
		bl.emit(BinaryTypeNullMarker, nil)
		return lexBinaryType
	}

	bl.expTyp = bit
	bl.emit(bit, nil)

	return lexBinaryKey
}

func lexBinaryKey(bl *binaryLexer) binaryLexStateFn {
	str, err := bl.readString()
	if err != nil {
		return bl.errorf(err)
	}

	bl.emit(BinaryTypeString, str)

	return lexBinaryValue
}

func lexBinaryValue(bl *binaryLexer) binaryLexStateFn {

	var val any
	var err error

	switch bl.expTyp {
	case BinaryTypeNullMarker:
		// do nothing
	case BinaryTypeNone:
		// do nothing
	case BinaryTypeColor:
		fallthrough
	case BinaryTypePtr:
		fallthrough
	case BinaryTypeInt:
		val, err = bl.readInt()
	case BinaryTypeString:
		val, err = bl.readString()
	case BinaryTypeUint64:
		val, err = bl.readUint64()
	case BinaryTypeFloat:
		val, err = bl.readFloat32()
	case BinaryTypeWString:
		val, err = bl.readWString()
	default:
		err = errors.New("unsupported vdf binary item type")
	}

	if err == nil {
		bl.emit(bl.expTyp, val)
	} else {
		bl.errorf(err)
		return nil
	}

	bl.expTyp = BinaryTypeNone

	return lexBinaryType
}

func (bl *binaryLexer) readString() (string, error) {
	var bts []byte
	for {
		var b byte
		if err := binary.Read(bl.r, binary.LittleEndian, &b); err == nil {
			if b == byte(BinaryTypeNone) {
				break
			}
			bts = append(bts, b)
		} else {
			return "", err
		}
	}
	return string(bts), nil
}

func (bl *binaryLexer) readWString() (string, error) {
	var ws []uint16
	for {
		var wc uint16
		if err := binary.Read(bl.r, binary.LittleEndian, &wc); err == nil {
			if wc == uint16(BinaryTypeNone) {
				break
			}
			ws = append(ws, wc)
		} else {
			return "", err
		}
	}
	return string(utf16.Decode(ws)), nil
}

func (bl *binaryLexer) readByte() (b byte, err error) {
	err = binary.Read(bl.r, binary.LittleEndian, &b)
	return b, err
}

func (bl *binaryLexer) readInt() (i uint32, err error) {
	err = binary.Read(bl.r, binary.LittleEndian, &i)
	return i, err
}

func (bl *binaryLexer) readUint64() (ui64 uint64, err error) {
	err = binary.Read(bl.r, binary.LittleEndian, &ui64)
	return ui64, err
}

func (bl *binaryLexer) readFloat32() (f32 float32, err error) {
	err = binary.Read(bl.r, binary.LittleEndian, &f32)
	return f32, err
}

func (bl *binaryLexer) errorf(err error) binaryLexStateFn {

	if errors.Is(err, io.EOF) {
		bl.emit(BinaryTypeEOF, nil)
		return nil
	}

	bl.items <- binaryLexItem{BinaryTypeError, err}
	return nil
}

func (bl *binaryLexer) emit(t BinaryType, val any) {
	bl.items <- binaryLexItem{t, val}
}

func (bl *binaryLexer) nextItem() binaryLexItem {
	for {
		select {
		case item := <-bl.items:
			return item
		}
	}
}
