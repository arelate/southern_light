package steam_vdf

import (
	"errors"
	"os"
)

type binaryParseStateFn func(parser *binaryParser) binaryParseStateFn

type binaryParser struct {
	lex   *binaryLexer
	last  *KeyValues
	stack []*KeyValues
	kv    []*KeyValues
	err   error
}

func newBinaryParser(lex *binaryLexer) *binaryParser {
	p := &binaryParser{
		lex: lex,
		kv:  make([]*KeyValues, 0),
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
		if itemKey.typ != BinaryTypeString {
			bp.err = errors.New("vdf binary key must be string")
			return nil
		}

		itemVal := bp.lex.nextItem()

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

func ParseBinary(path string) ([]*KeyValues, error) {

	vdfFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer vdfFile.Close()

	p := newBinaryParser(newBinaryLexer(vdfFile))
	if err := p.run(); err != nil {
		return nil, p.err
	}

	return p.kv, nil
}
