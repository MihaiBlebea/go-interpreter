package lexer

import (
	token "github.com/MihaiBlebea/go-interpreter/token"
)

type Lexer struct {
	input    string
	position int
	readPos  int
	char     byte
}

func New(input string) *Lexer {
	l := Lexer{input: input}
	l.readChar()

	return &l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPos]
	}

	l.position = l.readPos
	l.readPos += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPos]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	if isWhitespace(l.char) == true {
		l.readChar()
	}

	col := l.position + 1

	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()

			tok = token.NewWithString(1, col, token.EQ, "==")
		} else {
			tok = token.New(1, col, token.ASSIGN, l.char)
		}
	case '+':
		tok = token.New(1, col, token.PLUS, l.char)
	case '-':
		tok = token.New(1, col, token.MINUS, l.char)
	case '*':
		tok = token.New(1, col, token.MULTIPLY, l.char)
	case ':':
		tok = token.New(1, col, token.DIVIDE, l.char)
	case ',':
		tok = token.New(1, col, token.COMMA, l.char)
	case ';':
		tok = token.New(1, col, token.SEMICOLON, l.char)
	case '(':
		tok = token.New(1, col, token.LPAREN, l.char)
	case ')':
		tok = token.New(1, col, token.RPAREN, l.char)
	case '{':
		tok = token.New(1, col, token.LBRACE, l.char)
	case '}':
		tok = token.New(1, col, token.RBRACE, l.char)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.NewWithString(1, col, token.BANG, "!=")
		} else {
			tok = token.New(1, col, token.BANG, l.char)
		}
	case '>':
		tok = token.New(1, col, token.GT, l.char)
	case '<':
		tok = token.New(1, col, token.LT, l.char)
	case 0:
		tok.Value = ""
		tok.Type = token.EOF
	default:
		tok.Line = 1
		tok.Col = col

		if isLetter(l.char) {
			tok.Value = l.readIdent()
			tok.Type = token.IdentLookup(tok.Value)

			return tok
		} else if isDigit(l.char) {
			tok.Value = l.readNumber()
			tok.Type = token.INT

			return tok
		} else {
			tok = token.New(0, 0, token.ILLEGAL, l.char)
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readIdent() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.char) {
		l.readChar()
	}

	return l.input[position:l.position]
}
