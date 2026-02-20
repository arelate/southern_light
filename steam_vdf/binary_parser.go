package steam_vdf

import (
	"errors"
	"io"
)

type binaryParseStateFn func(parser *binaryParser) binaryParseStateFn

type binaryParser struct {
	lex   *binaryLexer
	last  *KeyValues
	stack ValveDataFile
	kv    ValveDataFile
	err   error
}

func newBinaryParser(lex *binaryLexer) *binaryParser {
	p := &binaryParser{
		lex: lex,
		kv:  make(ValveDataFile, 0),
	}
	return p
}

func (bp *binaryParser) run() error {
	for st := parseNextValue; st != nil; {
		st = st(bp)
	}
	if bp.err != nil {
		return bp.err
	}
	return nil
}

func parseNextValue(bp *binaryParser) binaryParseStateFn {

	for {
		itemType := bp.lex.nextItem()

		if itemType.typ == BinaryTypeError {
			bp.err = itemType.val.(error)
			return nil
		}

		if itemType.typ == BinaryTypeEOF {
			return nil
		}

		if itemType.typ == BinaryTypeNullMarker {
			if len(bp.stack) > 0 {
				bp.stack = bp.stack[:len(bp.stack)-1]
			}
			return parseNextValue
		}

		itemKey := bp.lex.nextItem()

		if itemKey.typ == BinaryTypeEOF {
			bp.err = errors.New("unexpected vdf binary EOF while parsing key")
			return nil
		}

		if itemKey.typ == BinaryTypeError {
			bp.err = itemKey.val.(error)
			return nil
		}

		if itemKey.typ != BinaryTypeString {
			bp.err = errors.New("vdf binary key must be string")
			return nil
		}

		itemVal := bp.lex.nextItem()

		if itemKey.typ == BinaryTypeEOF {
			bp.err = errors.New("unexpected vdf binary EOF while parsing value")
			return nil
		}

		if itemType.typ == BinaryTypeError {
			bp.err = itemVal.val.(error)
			return nil
		}

		item := &KeyValues{
			Key:        itemKey.val.(string),
			Type:       itemType.typ,
			TypedValue: itemVal.val,
		}

		if itemType.typ == BinaryTypeNone {
			if len(bp.stack) == 0 {
				bp.kv = append(bp.kv, item)
			} else {
				bp.stack[len(bp.stack)-1].Values = append(bp.stack[len(bp.stack)-1].Values, item)
			}
			bp.stack = append(bp.stack, item)
			bp.last = item

		} else {
			bp.last.Values = append(bp.last.Values, item)
		}
	}
}

func ReadBinary(reader io.Reader) (ValveDataFile, error) {

	p := newBinaryParser(newBinaryLexer(reader))
	if err := p.run(); err != nil {
		return nil, p.err
	}

	return p.kv, nil
}
