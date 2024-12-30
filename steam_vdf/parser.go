package steam_vdf

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

type KeyValue struct {
	Key string
	Val *string
}

func (kv *KeyValue) String() string {
	if kv.Val != nil {
		return fmt.Sprintf("%s=%s", kv.Key, *kv.Val)
	} else {
		return fmt.Sprintf("%s={}", kv.Key)
	}
}

type parseStateFn func(parser *parser) parseStateFn

type parser struct {
	lex  *lexer
	path string
	kv   []*KeyValue
	err  error
}

func newParser(lex *lexer) *parser {
	p := &parser{
		lex: lex,
		kv:  make([]*KeyValue, 0),
	}
	return p
}

func (p *parser) run() error {
	for st := parseKey; st != nil; {
		st = st(p)
	}
	if p.err != nil {
		return p.err
	}
	return nil
}

func parseKey(p *parser) parseStateFn {
	for {
		item := p.lex.nextItem()
		if item.typ == itemKeyValue {
			pt := path.Join(p.path, item.val)
			p.kv = append(p.kv, &KeyValue{Key: pt})
			return parseValue
		}
		if item.typ == itemRightMeta {
			lp, _ := path.Split(p.path)
			p.path = lp
			return parseKey
		}
		if item.typ == itemEOF {
			break
		}
		if item.typ == itemError {
			p.err = errors.New(item.val)
		}
	}
	return nil
}

func parseValue(p *parser) parseStateFn {
	item := p.lex.nextItem()
	var last *KeyValue
	if len(p.kv) > 0 {
		last = p.kv[len(p.kv)-1]
	}
	if item.typ == itemKeyValue {
		if last != nil {
			last.Val = &item.val
			return parseKey
		}
		p.err = errors.New("vdf cannot start with a value")
		return nil
	}
	if item.typ == itemLeftMeta {
		if last != nil {
			p.path = p.kv[len(p.kv)-1].Key
			return parseKey
		}
		p.err = errors.New("vdf cannot start with a left meta")
		return nil
	}
	return nil
}

func Parse(path string) ([]*KeyValue, error) {

	vdfFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer vdfFile.Close()

	sb := &strings.Builder{}

	if _, err := io.Copy(sb, vdfFile); err != nil {
		return nil, err
	}

	p := newParser(newLexer(sb.String()))
	if err := p.run(); err != nil {
		return nil, p.err
	}

	return p.kv, nil
}
