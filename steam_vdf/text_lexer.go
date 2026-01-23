package steam_vdf

import (
	"fmt"
	"text/scanner"
	"unicode"
	"unicode/utf8"
)

const (
	leftMeta  = '{'
	rightMeta = '}'
	quote     = '"'
)

type textLexItem struct {
	typ textItemType
	val string
}

type textItemType int

const (
	textItemEOF textItemType = iota - 1
	textItemError
	textItemKeyValue
	textItemLeftMeta
	textItemRightMeta
)

type textLexStateFn func(*textLexer) textLexStateFn

type textLexer struct {
	input string
	start int
	pos   int
	width int
	items chan textLexItem
}

func newTextLexer(input string) *textLexer {
	l := &textLexer{
		input: input,
		items: make(chan textLexItem),
	}
	go l.run()
	return l
}

func (tl *textLexer) run() {
	for st := lexText; st != nil; {
		st = st(tl)
	}
	close(tl.items)
}

func (tl *textLexer) emit(t textItemType) {
	tl.items <- textLexItem{t, tl.input[tl.start:tl.pos]}
	tl.start = tl.pos
}

func lexText(tl *textLexer) textLexStateFn {
	for {
		r := tl.next()
		switch r {
		case quote:
			tl.ignore()
			return lexTextKeyValue
		case leftMeta:
			tl.emit(textItemLeftMeta)
			return lexText
		case rightMeta:
			tl.emit(textItemRightMeta)
			return lexText
		}
		if unicode.IsSpace(r) {
			tl.ignore()
		}
		if r == scanner.EOF {
			break
		}
	}
	tl.emit(textItemEOF)
	return nil
}

func lexTextKeyValue(tl *textLexer) textLexStateFn {
	for {
		if tl.peek() == quote {
			tl.emit(textItemKeyValue)
			tl.next()
			return lexText
		}
		r := tl.next()
		if isText(r) {
			continue
		}
		if r == scanner.EOF || r == '\n' {
			return tl.errorf("unclosed key value")
		}
		return tl.errorf("bad key value syntax: %q", tl.input[tl.start:tl.pos])
	}
}

func (tl *textLexer) next() (rune rune) {
	if tl.pos >= len(tl.input) {
		tl.width = 0
		return scanner.EOF
	}
	rune, tl.width = utf8.DecodeRuneInString(tl.input[tl.pos:])
	tl.pos += tl.width
	return rune
}

func (tl *textLexer) ignore() {
	tl.start = tl.pos
}

func (tl *textLexer) backup() {
	tl.pos -= tl.width
}

func (tl *textLexer) peek() rune {
	rn := tl.next()
	tl.backup()
	return rn
}

func (tl *textLexer) errorf(format string, args ...any) textLexStateFn {
	tl.items <- textLexItem{
		textItemError,
		fmt.Sprintf(format, args...),
	}
	return nil
}

func (tl *textLexer) nextItem() textLexItem {
	for {
		select {
		case item := <-tl.items:
			return item
		}
	}
}

func isText(r rune) bool {
	return unicode.IsNumber(r) ||
		unicode.IsLetter(r) ||
		unicode.IsPunct(r) ||
		unicode.IsSpace(r) ||
		unicode.IsSymbol(r)
}
