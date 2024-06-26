package lexer

import (
	"arcane/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let seventy = 70;
let four = 4;

let add = fn(x, y) {
 x + y;
};
let result = add(seventy, four);
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "seventy"},
		{token.ASSIGN, "="},
		{token.INT, "70"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "four"},
		{token.ASSIGN, "="},
		{token.INT, "4"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LEFT_PARENTHESIS, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.LEFT_BRACKETS, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RIGHT_BRACKETS, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LEFT_PARENTHESIS, "("},
		{token.IDENT, "seventy"},
		{token.COMMA, ","},
		{token.IDENT, "four"},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
	}

	l := Init(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokenType wrong. expected: %d, got: %d", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != token.TokenLiteral(tt.expectedLiteral) {
			t.Fatalf("test[%d] - literal wrong. expected: %q, got: %q", i, tt.expectedLiteral, tok.Literal)
		}

	}

}
