package ast

import "arcane/token"

type Node interface {
	TokenLiteral() token.TokenLiteral
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

// return the beginning of the node
func (p *Program) TokenLiteral() token.TokenLiteral {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type Identifier struct {
	Token token.Token // -> token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() token.TokenLiteral {
	return i.Token.Literal
}

// :: Let Statement
type LetStatement struct {
	Token token.Token // -> token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() token.TokenLiteral {
	return ls.Token.Literal
}

// :: Return Statement
type ReturnStatement struct {
	Token       token.Token // -> token.LET
	ReturnValue Expression
}

func (ls *ReturnStatement) statementNode() {}
func (ls *ReturnStatement) TokenLiteral() token.TokenLiteral {
	return ls.Token.Literal
}
