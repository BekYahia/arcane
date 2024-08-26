package ast

import (
	"arcane/token"
	"bytes"
)

type Node interface {
	TokenLiteral() token.TokenLiteral
	String() string // print AST nodes for debugging
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
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type Identifier struct {
	Token token.Token // -> token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() token.TokenLiteral {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

type IntegralLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegralLiteral) expressionNode() {}
func (il *IntegralLiteral) TokenLiteral() token.TokenLiteral {
	return il.Token.Literal
}
func (il *IntegralLiteral) String() string {
	return string(il.Token.Literal)
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
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(string(ls.TokenLiteral()) + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// :: Return Statement
type ReturnStatement struct {
	Token       token.Token // -> token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() token.TokenLiteral {
	return rs.Token.Literal
}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(string(rs.TokenLiteral()))
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// :: Expression Statement
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() token.TokenLiteral {
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
