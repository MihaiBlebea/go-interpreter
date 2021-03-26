package ast

import "github.com/MihaiBlebea/go-interpreter/token"

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) TokenValue() string    { return ls.Token.Value }
func (ls *LetStatement) statementNode() string { return "" }

type Identifier struct {
	Token token.Token
	Value Expression
}

func (i *Identifier) TokenValue() string     { return i.Token.Value }
func (i *Identifier) expressionNode() string { return "" }
