package steam_vdf

import (
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
	for st := parseType; st != nil; {
		st = st(bp)
	}
	if bp.err != nil {
		return bp.err
	}
	return nil
}

func parseType(bp *binaryParser) binaryParseStateFn {
	for {
		item := bp.lex.nextItem()
		if item.typ == binaryItemEOF {
			break
		}
		if item.typ == binaryItemError {
			bp.err = item.val.(error)
		}
	}
	return nil
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
