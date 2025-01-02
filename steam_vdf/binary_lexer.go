package steam_vdf

import (
	"encoding/binary"
	"errors"
	"io"
	"unicode/utf16"
)

type binaryLexItem struct {
	typ binaryItemType
	val any
}

type binaryItemType byte

// https://github.com/ValveSoftware/source-sdk-2013/blob/master/sp/src/public/tier1/kvpacker.h
const (
	binaryItemNone binaryItemType = iota
	binaryItemString
	binaryItemInt
	binaryItemFloat
	binaryItemPtr
	binaryItemWString
	binaryItemColor
	binaryItemUint64
	binaryItemNullMarker
	binaryItemError binaryItemType = 254
	binaryItemEOF   binaryItemType = 255
)

type binaryLexStateFn func(*binaryLexer) binaryLexStateFn

type binaryLexer struct {
	r      io.Reader
	eof    bool
	expTyp binaryItemType
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

	bit := binaryItemType(t)

	if bit == binaryItemNullMarker {
		bl.emit(binaryItemNullMarker, nil)
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

	bl.emit(binaryItemString, str)

	return lexBinaryValue
}

func lexBinaryValue(bl *binaryLexer) binaryLexStateFn {

	var val any
	var err error

	switch bl.expTyp {
	case binaryItemNullMarker:
		// do nothing
	case binaryItemNone:
		// do nothing
	case binaryItemColor:
		fallthrough
	case binaryItemPtr:
		fallthrough
	case binaryItemInt:
		val, err = bl.readInt()
	case binaryItemString:
		val, err = bl.readString()
	case binaryItemUint64:
		val, err = bl.readUint64()
	case binaryItemFloat:
		val, err = bl.readFloat32()
	case binaryItemWString:
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

	bl.expTyp = binaryItemNone

	return lexBinaryType
}

func (bl *binaryLexer) readString() (string, error) {
	var bts []byte
	for {
		var b byte
		if err := binary.Read(bl.r, binary.LittleEndian, &b); err == nil {
			if b == byte(binaryItemNone) {
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
			if wc == uint16(binaryItemNone) {
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

func (bl *binaryLexer) readInt() (i int32, err error) {
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
		bl.emit(binaryItemEOF, nil)
		return nil
	}

	bl.items <- binaryLexItem{binaryItemError, err}
	return nil
}

func (bl *binaryLexer) emit(t binaryItemType, val any) {
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
