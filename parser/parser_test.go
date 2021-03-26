package parser_test

import (
	"fmt"
	"testing"

	"github.com/MihaiBlebea/go-interpreter/lexer"
	"github.com/MihaiBlebea/go-interpreter/parser"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 1234;
	let foobar = 0;
	`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("program returned nil")
	}

	if len(program.Statements) < 3 {
		t.Fatalf("program should contain at least 3 statements: got %d", len(program.Statements))
	}

	cases := []struct {
		expected string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("test_%s", c.expected), func(t *testing.T) {
			p.NextToken()
		})
	}

}
