package lexer

import (
	"arcane/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){};"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LEFT_PARENTHESIS, "("},
		{token.RIGHT_PARENTHESIS, ")"},
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
