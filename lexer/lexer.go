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
		tok = token.New(1, l.position, token.ASSIGN, l.char)
	case '+':
		tok = token.New(1, l.position, token.PLUS, l.char)
	case '-':
		tok = token.New(1, l.position, token.MINUS, l.char)
	case '*':
		tok = token.New(1, l.position, token.MULTIPLY, l.char)
	case ':':
		tok = token.New(1, l.position, token.DIVIDE, l.char)
	case ',':
		tok = token.New(1, l.position, token.COMMA, l.char)
	case ';':
		tok = token.New(1, l.position, token.SEMICOLON, l.char)
	case '(':
		tok = token.New(1, l.position, token.LPAREN, l.char)
	case ')':
		tok = token.New(1, l.position, token.RPAREN, l.char)
	case '{':
		tok = token.New(1, l.position, token.LBRACE, l.char)
	case '}':
		tok = token.New(1, l.position, token.RBRACE, l.char)
	case 0:
		tok.Value = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.char) {
			tok.Col, tok.Value = l.readIdent()
			tok.Type = token.IdentLookup(tok.Value)

			return tok
		} else if isDigit(l.char) {
			tok.Col, tok.Value = l.readNumber()
			tok.Type = token.INT

			return tok
		} else {
			tok = token.New(0, 0, token.ILLEGAL, l.char)
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readIdent() (int, string) {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}

	return position + 1, l.input[position:l.position]
}

func (l *Lexer) readNumber() (int, string) {
	position := l.position
	for isDigit(l.char) {
		l.readChar()
	}

	return position + 1, l.input[position:l.position]
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
