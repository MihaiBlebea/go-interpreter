package token_test

import (
	"fmt"
	"testing"

	token "github.com/MihaiBlebea/go-interpreter/token"
)

func TestIdentLookup(t *testing.T) {
	cases := []struct {
		input string
		ident token.TokenType
	}{
		{
			input: "func",
			ident: token.FUNC,
		},
		{
			input: "let",
			ident: token.LET,
		},
		{
			input: "a",
			ident: token.IDENT,
		},
		{
			input: "abc",
			ident: token.IDENT,
		},
		{
			input: "abc2",
			ident: token.IDENT,
		},
		{
			input: "abc007cba",
			ident: token.IDENT,
		},
		{
			input: "123",
			ident: token.IDENT,
		},
	}

	for _, c := range cases {
		title := fmt.Sprintf("%s_%s", c.input, c.ident)
		t.Run(title, func(t *testing.T) {
			ident := token.IdentLookup(c.input)

			if ident != c.ident {
				t.Errorf("line error: got %v want %v", ident, c.ident)
			}
		})
	}
}
