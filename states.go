package dson

const (
	object_state = iota
	array_state
)

type stateFn func(*lexer) stateFn

func lexState(l *lexer) stateFn {
	state = l.state.Head()
	if state == object_state {
		return lexObject
	} else if state == array_state {
		return lexArray
	}
	return nil
}

func lexObject(l *lexer) stateFn {
	for {
		if l.peekForToken(pairSeperator) {
			l.pos += len(pairSeperator)
			l.emit(tokenPairSeperator)
		} else if l.peekForToken(objectStart) {
			return lexObjStart
		} else if l.peekForToken(objectEnd) {
			return lexObjectEnd
		} else if l.peekForToken(arrayStart) {
			return lexArrayStart
		}

		switch l.next() {
		case memberSeperator1, memberSeperator2,
			memberSeperator3, memberSeperator4:
			return lexMemberSeperator
		case doubleQuote:
			return lexQuote
		}
	}
}

// state function for lexing the beginning of an object
func lexObjStart(l *lexer) stateFn {
	if l.peekForToken(objectStart) {
		l.pos += len(objectStart)
		l.emit(tokenObjectStart)
		l.state.Push(object_state)
		return lexKey
	} else {
		return l.errorf("Missing starting identifier")
	}
}

func lexObjectEnd(l *lexer) stateFn {
	if l.state.Pop() == object_state {

	} else {
		return l.errorf("expected end of object")
	}
}

func lexArray(l *lexer) stateFn {

}

func lexKey(l *lexer) stateFn {
	for {
		switch r := l.next(); {
		case isSpace(r):
			l.ignore()
		case r == doubleQuote:
			if lexQuoteRaw(l) == nil {
				return nil
			}
			return lexPairSeperator
		default:
			return l.errorf("misformatted key")
		}
	}
}

func lexValue(l *lexer) stateFn {
	if l.peekForToken(objectStart) {
		// pass
	} else if l.peekForToken(arrayStart) {
		// pass
	} else if l.peekForToken(boolTrue) {
		// pass
	} else if l.peekForToken(boolFalse) {
		// pass
	}

	switch r := l.next(); {
	case doubleQuote:
		return lexQuote
	case r == numMinus || ('0' <= r && r <= 7):
		l.backup()
		return lexNumber
	}
}

func lexMemberSeperator(l *lexer) stateFn {

}

func lexPairSeperator(l *lexer) stateFn {
	for {
		switch r := l.next(); {
		case isSpace(r):
			l.ignore()
		default:
			l.backup()
			if l.peekForToken(pairSeperator) {
				l.pos = len(pairSeperator)
				l.emit(tokenPairSeperator)
				return lexValue
			}
			return l.errorf("expected pair seperator 'is'")
		}
	}
}

// only to be called with stateCallback
func lexQuote(l *lexer) stateFn {
	if lexQuoteRaw(l) == nil { // error
		return nil
	}
	l.emit(tokenString)
	return lexState
}

func lexQuoteRaw(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("unterminated quoted string")
		case '"':
			break Loop
		}
	}
}
