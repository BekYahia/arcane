package parser

import (
	"arcane/ast"
	"arcane/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 4;
let y = 12;
let foo = 352521;	
`
	l := lexer.Init(input)
	p := Init(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements should contain 3 statements. got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if stmt == nil {
			t.Fatalf("Statement can not be nil, %d", len(program.Statements))
			return
		}
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral expected to be 'let'. got %q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s expected to be ast.LetStatement. got %T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value expected to be %s. got %s", name, letStmt.Name.Value)
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
return 4;
return 12;
return 352521;	
`
	l := lexer.Init(input)
	p := Init(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements should contain 3 statements. got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		if stmt == nil {
			t.Fatalf("Statement can not be nil, %d", len(program.Statements))
			return
		}
		if !testReturnStatement(t, stmt) {
			return
		}
	}
}

func testReturnStatement(t *testing.T, s ast.Statement) bool {
	if s.TokenLiteral() != "return" {
		t.Errorf("s.TokenLiteral expected to be 'return'. got %q", s.TokenLiteral())
		return false
	}

	returnStmt, ok := s.(*ast.ReturnStatement)

	if !ok {
		t.Errorf("s expected to be ast.LetStatement. got %T", s)
		return false
	}
	if returnStmt.TokenLiteral() != "return" {
		t.Errorf("returnStmt.Name.Value expected to be 'return'. got %q", returnStmt.TokenLiteral())
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser has %d errors", len(errors))

	for _, err := range errors {
		t.Errorf("Parser error, %q", err)
	}
	t.FailNow()
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.Init(input)
	p := Init(l)

	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got=%d", program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("Expression hos no *ast.Identifier, got=%t", stmt.Expression)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.Token.Literal not %s. got %s", "foobar", ident.TokenLiteral())
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got %s", "foobar", ident.Value)
	}
}
