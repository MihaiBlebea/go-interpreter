package token

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

type TokenType string

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

func New(line int, col int, t TokenType, value byte) Token {
	return Token{line, col, t, string(value)}
}

func NewIdent(line int, col int, t TokenType, keyword string) Token {
	return Token{line, col, t, keyword}
}

func IdentLookup(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
