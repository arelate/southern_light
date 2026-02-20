package steam_vdf

import (
	"errors"
	"io"
	"strings"
)

type textParseStateFn func(parser *textParser) textParseStateFn

type textParser struct {
	lex   *textLexer
	last  *KeyValues
	stack ValveDataFile
	kv    ValveDataFile
	err   error
}

func newTextParser(lex *textLexer) *textParser {
	p := &textParser{
		lex: lex,
		kv:  make(ValveDataFile, 0),
	}
	return p
}

func (tp *textParser) run() error {
	for st := parseTextKey; st != nil; {
		st = st(tp)
	}
	if tp.err != nil {
		return tp.err
	}
	return nil
}

func parseTextKey(tp *textParser) textParseStateFn {
	for {
		item := tp.lex.nextItem()
		switch item.typ {
		case textItemKeyValue:
			tp.last = &KeyValues{Key: item.val}
			if len(tp.stack) == 0 {
				tp.kv = append(tp.kv, tp.last)
			} else {
				tp.stack[len(tp.stack)-1].Values = append(tp.stack[len(tp.stack)-1].Values, tp.last)
			}
			return parseTextValue
		case textItemRightCurlyBracket:
			if len(tp.stack) > 0 {
				tp.stack = tp.stack[:len(tp.stack)-1]
			}
			return parseTextKey
		case textItemEOF:
			return breakParse
		case textItemError:
			tp.err = errors.New(item.val)
			return breakParse
		default:
			tp.err = errors.New("unhandled parseTextKey item type")
			return breakParse
		}
	}
}

func breakParse(_ *textParser) textParseStateFn {
	return nil
}

func parseTextValue(tp *textParser) textParseStateFn {
	item := tp.lex.nextItem()
	switch item.typ {
	case textItemKeyValue:
		if len(tp.stack) > 0 {
			tp.last.Value = &item.val
			return parseTextKey
		}
		tp.err = errors.New("vdf cannot start with a value")
		return nil
	case textItemLeftCurlyBracket:
		tp.stack = append(tp.stack, tp.last)
		return parseTextKey
	case textItemError:
		tp.err = errors.New(item.val)
		return breakParse
	default:
		tp.err = errors.New("unhandled parseTextValue item type")
		return breakParse
	}
}

func ReadText(reader io.Reader) (ValveDataFile, error) {

	sb := &strings.Builder{}

	if _, err := io.Copy(sb, reader); err != nil {
		return nil, err
	}

	p := newTextParser(newTextLexer(sb.String()))
	if err := p.run(); err != nil {
		return nil, p.err
	}

	return p.kv, nil
}
