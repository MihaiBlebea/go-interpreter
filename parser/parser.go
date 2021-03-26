package parser

import (
	"github.com/MihaiBlebea/go-interpreter/ast"
	"github.com/MihaiBlebea/go-interpreter/lexer"
	"github.com/MihaiBlebea/go-interpreter/token"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := Parser{l: l}

	p.NextToken()
	p.NextToken()

	return &p
}

func (p *Parser) NextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOF {
		statement := p.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		p.NextToken()
	}

	return nil
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseStatement()
	default:
		return nil
	}
}
