package token

const (
	ILLEGAL TokenType = iota // identify unknown tokens
	EOF                      // tell the parse the stop

	//Identifiers
	IDENT
	INT
	STRING
	BOOL

	//Operators
	ASSIGN
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	MODULES

	//Delimiters
	COMMA
	SEMICOLON
	DOT

	LEFT_PARENTHESIS
	RIGHT_PARENTHESIS
	LEFT_BRACKETS
	RIGHT_BRACKETS

	//keywords
	FUNCTION
	LET
)

var Tokens = map[TokenType]string{
	ILLEGAL: "ILLEGAL", //identify unknown tokens
	EOF:     "EOF",     // tell the parse the stop

	//Identifiers
	IDENT:  "IDENT",
	INT:    "INT",
	STRING: "STRING",
	BOOL:   "BOOL",

	//Operators
	ASSIGN:   "=",
	PLUS:     "+",
	MINUS:    "-",
	MULTIPLY: "*",
	DIVIDE:   "/",
	MODULES:  "%",

	//Delimiters
	COMMA:     ",",
	SEMICOLON: ";",
	DOT:       ".",

	LEFT_PARENTHESIS:  "(",
	RIGHT_PARENTHESIS: ")",
	LEFT_BRACKETS:     "{",
	RIGHT_BRACKETS:    "}",

	//keywords
	FUNCTION: "FUNCTION",
	LET:      "LET",
}

type TokenType int
type TokenLiteral string
type Token struct {
	Type    TokenType
	Literal TokenLiteral
}

//TODO: Attach filename and line number, and column number to token(using io.reader), to better track for lexing and parsing errors.
