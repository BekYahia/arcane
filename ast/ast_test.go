package ast

import (
	"arcane/token"
	"testing"
)

const testStmt = "let myVar = anotherVar;"

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != testStmt {
		t.Errorf("program.String() wrong. expected '%s' - got '%s'", testStmt, program.String())
	}
}
