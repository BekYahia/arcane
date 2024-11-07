package parser

import (
	"arcane/ast"
	"arcane/lexer"
	"arcane/token"
	"fmt"
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
		t.Fatalf("program.Statements[0] is not a valid ast.ExpressionStatement, got=%T", program.Statements[0])
	}
	if !testIdentifierLiteral(t, stmt.Expression, "foobar") {
		return
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "7;"

	l := lexer.Init(input)
	p := Init(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program has not enough statements. got %d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Expression program.Statements[0] is not a valid ast.ExpressionStatement, got %T", program.Statements[0])
	}
	integer_literal, ok := stmt.Expression.(*ast.IntegralLiteral)
	if !ok {
		t.Fatalf("Expression is not valid ast.IntegralLiteral, got %T", program.Statements[0])
	}

	if integer_literal.TokenLiteral() != "7" {
		t.Fatalf("integer_literal.TokenLiteral is not equal to %s, got %s", "7", integer_literal.TokenLiteral())
	}
	if integer_literal.Value != 7 {
		t.Fatalf("integer_literal.Value is not equal to %d, got %d", 7, integer_literal.Value)
	}
}

func TestBooleanLiteralExpression(t *testing.T) {
	input := "true;"

	l := lexer.Init(input)
	p := Init(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program has not enough statements. got %d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Expression does not satisfy ast.ExpressionStatement- got %T", program.Statements[0])
	}
	if !testBooleanLiteral(t, stmt.Expression, true) {
		return
	}
}

func TestParsingPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input      string
		operator   string
		rightValue interface{}
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"!false", "!", false},
		{"!true", "!", true},
	}

	for _, tt := range prefixTests {
		l := lexer.Init(tt.input)
		p := Init(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statement. got %d\n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statement[0] is not an Expression Statement. got= %T", program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not a Prefix Expression. got=%T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is wrong. got=%s", exp.Operator)
		}
		if !testLiteralExpression(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integer, ok := il.(*ast.IntegralLiteral)
	if !ok {
		t.Errorf("il is not an integer. got=%T", il)
		return false
	}

	if integer.Value != value {
		t.Errorf("integer value is not %d, got=%d", value, integer.Value)
		return false
	}

	if integer.TokenLiteral() != token.TokenLiteral(fmt.Sprintf("%d", value)) {
		t.Errorf("Token Literal is not %d. got=%s", value, integer.TokenLiteral())
		return false
	}
	return true
}

func testIdentifierLiteral(t *testing.T, exp ast.Expression, value string) bool {
	identifier, ok := exp.(*ast.Identifier)

	if !ok {
		t.Errorf("exp does not satisfy ast.Identifier, got%T", exp)
		return false
	}

	if identifier.Value != value {
		t.Errorf("identifier.value is not %s, got %s", value, identifier.Value)
		return false
	}

	if identifier.TokenLiteral() != token.TokenLiteral(value) {
		t.Errorf("identifier.TokenLiteral() is not %s, got %s", value, identifier.TokenLiteral())
		return false
	}

	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	literal, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("exp does not satisfy ast.Boolean, got %T", exp)
		return false
	}
	if literal.Value != value {
		t.Errorf("boolean literal is not %t - got %t", value, literal.Value)
		return false
	}
	if literal.TokenLiteral() != token.TokenLiteral(fmt.Sprintf("%t", value)) {
		t.Errorf("literal.TokenLiteral() is not %t - got %s", value, literal.TokenLiteral())
		return false
	}

	return true
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {

	switch value := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(value))
	case int64:
		return testIntegerLiteral(t, exp, value)
	case string:
		return testIdentifierLiteral(t, exp, value)
	case bool:
		return testBooleanLiteral(t, exp, value)
	}
	t.Errorf("Unknown literal type - got %T", exp)
	return true
}
func TestParsingInfixExpression(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5+5;", 5, "+", 5},
		{"5-5;", 5, "-", 5},
		{"5*5;", 5, "*", 5},
		{"5/5;", 5, "/", 5},
		{"5>5;", 5, ">", 5},
		{"5<5;", 5, "<", 5},
		{"5==5;", 5, "==", 5},
		{"5!=5;", 5, "!=", 5},
		{"a * b;", "a", "*", "b"},
		{"false == false", false, "==", false},
		{"false != true", false, "!=", true},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
	}

	for _, tt := range infixTests {
		l := lexer.Init(tt.input)
		p := Init(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program statement does not contain %d statements, got %d\n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statement[0] dose not satisfy ast.ExpressionStatement, got %T\n", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("stmt.Expression does not satisfy ast.InfixExpression, got %T", stmt.Expression)
		}

		if !testLiteralExpression(t, exp.Left, tt.leftValue) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not %s, got %s", tt.operator, exp.Operator)
		}

		if !testLiteralExpression(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
	}

	for _, tt := range tests {
		l := lexer.Init(tt.input)
		p := Init(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		actual := program.String()

		if actual != tt.expectedValue {
			t.Errorf("expected %s, got %s", tt.expectedValue, actual)
		}
	}
}
