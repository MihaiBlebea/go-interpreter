package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = ":"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNC = "FUNC"
	LET  = "LET"
)

var keywords = map[string]TokenType{
	"let":  LET,
	"func": FUNC,
}

type Token struct {
	Line  int
	Col   int
	Type  TokenType
	Value string
}

// TODO: handle error cases when line = 0 or col = 0
func New(line int, col int, t TokenType, value byte) Token {
	return Token{line, col, t, string(value)}
}

// TODO: handle error cases when line = 0 or col = 0
func NewIdent(line int, col int, t TokenType, keyword string) Token {
	return Token{line, col, t, keyword}
}

func IdentLookup(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
