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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	if isWhitespace(l.char) == true {
		l.readChar()
	}

	switch l.char {
	case '=':
		tok = token.New(0, 0, token.ASSIGN, l.char)
	case '+':
		tok = token.New(0, 0, token.PLUS, l.char)
	case '-':
		tok = token.New(0, 0, token.MINUS, l.char)
	case '*':
		tok = token.New(0, 0, token.MULTIPLY, l.char)
	case ':':
		tok = token.New(0, 0, token.DIVIDE, l.char)
	case ',':
		tok = token.New(0, 0, token.COMMA, l.char)
	case ';':
		tok = token.New(0, 0, token.SEMICOLON, l.char)
	case '(':
		tok = token.New(0, 0, token.LPAREN, l.char)
	case ')':
		tok = token.New(0, 0, token.RPAREN, l.char)
	case '{':
		tok = token.New(0, 0, token.LBRACE, l.char)
	case '}':
		tok = token.New(0, 0, token.RBRACE, l.char)
	case 0:
		tok.Value = ""
		tok.Type = token.EOF
	default:
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

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' ||
		'A' <= char && char <= 'Z' ||
		char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func isWhitespace(char byte) bool {
	return char == ' ' ||
		char == '\t' ||
		char == '\n' ||
		char == '\r'
}
