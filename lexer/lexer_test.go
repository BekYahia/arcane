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

!-/*3;

5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

6 == 6;
15 != 9;
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
		{token.LEFT_CURLY_BRACKETS, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RIGHT_CURLY_BRACKETS, "}"},
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
		{token.NOT, "!"},
		{token.MINUS, "-"},
		{token.DIVIDE, "/"},
		{token.MULTIPLY, "*"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LEFT_PARENTHESIS, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.LEFT_CURLY_BRACKETS, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RIGHT_CURLY_BRACKETS, "}"},
		{token.ELSE, "else"},
		{token.LEFT_CURLY_BRACKETS, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RIGHT_CURLY_BRACKETS, "}"},
		{token.INT, "6"},
		{token.EQUAL, "=="},
		{token.INT, "6"},
		{token.SEMICOLON, ";"},
		{token.INT, "15"},
		{token.NOT_EQUAL, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
	}

	l := Init(input)

	for i, tt := range tests {
		tok := l.NextToken()

		// if tok.Type != tt.expectedType {
		// 	t.Fatalf("test[%d] - tokenType wrong. expected: %d, got: %d", i, tt.expectedType, tok.Type)
		// }

		if tok.Literal != token.TokenLiteral(tt.expectedLiteral) {
			t.Fatalf("test[%d] - literal wrong. expected: %q, got: %q", i, tt.expectedLiteral, tok.Literal)
		}

	}

}
