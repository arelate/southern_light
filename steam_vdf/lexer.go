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

type lexItem struct {
	typ itemType
	val string
}

type itemType int

const (
	itemEOF itemType = iota - 1
	itemError
	itemKeyValue
	itemLeftMeta
	itemRightMeta
)

type lexStateFn func(*lexer) lexStateFn

type lexer struct {
	input string
	start int
	pos   int
	width int
	items chan lexItem
}

func newLexer(input string) *lexer {
	l := &lexer{
		input: input,
		items: make(chan lexItem),
	}
	go l.run()
	return l
}

func (l *lexer) run() {
	for st := lexText; st != nil; {
		st = st(l)
	}
	close(l.items)
}

func (l *lexer) emit(t itemType) {
	l.items <- lexItem{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func lexText(l *lexer) lexStateFn {
	for {
		r := l.next()
		switch r {
		case quote:
			l.ignore()
			return lexKeyValue
		case leftMeta:
			l.emit(itemLeftMeta)
			return lexText
		case rightMeta:
			l.emit(itemRightMeta)
			return lexText
		}
		if unicode.IsSpace(r) {
			l.ignore()
		}
		if r == scanner.EOF {
			break
		}
	}
	l.emit(itemEOF)
	return nil
}

func lexKeyValue(l *lexer) lexStateFn {
	for {
		if l.peek() == quote {
			l.emit(itemKeyValue)
			l.next()
			return lexText
		}
		r := l.next()
		if isText(r) {
			continue
		}
		if r == scanner.EOF || r == '\n' {
			return l.errorf("unclosed key value")
		}
		return l.errorf("bad key value syntax: %q", l.input[l.start:l.pos])
	}
}

func (l *lexer) next() (rune rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return scanner.EOF
	}
	rune, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return rune
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) peek() rune {
	rn := l.next()
	l.backup()
	return rn
}

func (l *lexer) errorf(format string, args ...any) lexStateFn {
	l.items <- lexItem{
		itemError,
		fmt.Sprintf(format, args...),
	}
	return nil
}

func (l *lexer) nextItem() lexItem {
	for {
		select {
		case item := <-l.items:
			return item
		}
	}
}

func isText(r rune) bool {
	return unicode.IsNumber(r) ||
		unicode.IsLetter(r) ||
		unicode.IsPunct(r) ||
		unicode.IsSpace(r)
}
