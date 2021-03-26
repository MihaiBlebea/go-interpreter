package lexer_test

import (
	"fmt"
	"testing"

	lexer "github.com/MihaiBlebea/go-interpreter/lexer"
	token "github.com/MihaiBlebea/go-interpreter/token"
)

type testCase struct {
	line      int
	col       int
	tokenType token.TokenType
	value     string
}

func TestNextToken(t *testing.T) {
	input := "let a = 2 + 5"

	cases := []testCase{
		{
			line:      1,
			col:       1,
			tokenType: token.LET,
			value:     "let",
		},
		{
			line:      1,
			col:       5,
			tokenType: token.IDENT,
			value:     "a",
		},
		{
			line:      1,
			col:       7,
			tokenType: token.ASSIGN,
			value:     "=",
		},
		{
			line:      1,
			col:       9,
			tokenType: token.INT,
			value:     "2",
		},
		{
			line:      1,
			col:       11,
			tokenType: token.PLUS,
			value:     "+",
		},
		{
			line:      1,
			col:       13,
			tokenType: token.INT,
			value:     "5",
		},
	}

	l := lexer.New(input)

	for _, c := range cases {
		title := fmt.Sprintf("%s_%s", c.tokenType, c.value)
		t.Run(title, func(t *testing.T) {
			tkn := l.NextToken()

			fmt.Println(tkn)
			if tkn.Line != c.line {
				t.Errorf("line error: got %v want %v", tkn.Line, c.line)
			}

			if tkn.Col != c.col {
				t.Errorf("col error: got %v want %v", tkn.Col, c.col)
			}

			if tkn.Type != c.tokenType {
				t.Errorf("type error: got %v want %v", tkn.Type, c.tokenType)
			}

			if tkn.Value != c.value {
				t.Errorf("value error: got %v want %v", tkn.Value, c.value)
			}
		})
	}
}
