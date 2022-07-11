package lexer

import "monkey/token"

type Lexer struct {
	input        string // contains the whole input
	position     int    // current position input
	readPosition int    // next position after current char
	ch           byte   // current char we are looking at
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// read next char into the lexer
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// peek at the next character
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// get the next token
func (l *Lexer) NextToken() token.Token {
	// initialize the token, modify it accordingly
	var tok token.Token

	// white spaces don't mean anything in this language
	l.skipWhiteSpaces()

	switch l.ch {
	case '-':
		tok = token.NewToken(token.MINUS, l.ch)
	case '!': // also check for !=
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Literal = string(ch) + string(l.ch)
			tok.Type = token.NOT_EQ
		} else {
			tok = token.NewToken(token.BANG, l.ch)
		}
	case '/':
		tok = token.NewToken(token.SLASH, l.ch)
	case '*':
		tok = token.NewToken(token.ASTERISK, l.ch)
	case '<':
		tok = token.NewToken(token.LT, l.ch)
	case '>':
		tok = token.NewToken(token.GT, l.ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case '=': // also check for ==
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Literal = string(ch) + string(l.ch)
			tok.Type = token.EQ
		} else {
			tok = token.NewToken(token.ASSIGN, l.ch)
		}
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default: // might be invalid or might be a Identifier
		// check if its a letter
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			// we do an early return because readIdentifier has already called the readChar method
			return tok
		} else if isDigit(l.ch) { // check if its a number
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar() // move forward in input
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position // starting point
	// read the full identifier
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position // starting point
	// read the full identifier
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhiteSpaces() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}
