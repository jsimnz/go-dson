package dson

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/oleiade/lane"
)

type lexer struct {
	input  string
	start  int
	pos    int
	width  int
	state  *lane.Stack
	tokens chan token
}

func lex(input string) (*lexer, chan token) {
	l := &lexer{
		name:   strings.Trim(name, " "),
		input:  input,
		tokens: make(chan token),
		state:  lane.NewStack(),
	}
	go l.run()
	return l, l.tokens
}

func (l *lexer) run() {
	for state := lexObjStart; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}

func (l *lexer) emit(t tokenType) {
	l.tokens <- token{t, input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	rune, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return rune
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) peekForToken(token string) bool {
	return strings.HasPrefix(l.input[l.pos:], token)
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

func (l *lexer) acceptRun(valid string) bool {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.tokens <- token{tokenError, fmt.Sprintf(format, args...)}
	return nil
}

func isSpace(r int) bool {
	return r == ' ' || r == '\t'
}
