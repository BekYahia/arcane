package lexer

import "arcane/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  //current read position in input (points to next char)
	ch           byte //current char under examination
}

func Init(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	positionOfFirstLetter := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[positionOfFirstLetter:l.position]
}

func (l *Lexer) readNumber() string {
	positionOfFirstNumber := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[positionOfFirstNumber:l.position]
}

// Note: for letters and digits lexer position is advanced within their respective function
func (l *Lexer) NextToken() token.Token {
	l.skipWhiteSpace()
	var tok token.Token

	switch {
	case isLetter(l.ch):
		tok.Literal = token.TokenLiteral(l.readIdentifier())
		tok.Type = token.LookupIdentifier(string(tok.Literal))
		return tok
	case isDigit(l.ch):
		tok.Literal = token.TokenLiteral(l.readNumber())
		tok.Type = token.INT
		return tok
	case l.ch == 0:
		tok = initToken(token.EOF, "")
	default:
		tok = initToken(token.ILLEGAL, token.TokenLiteral(l.ch))

		for key, value := range token.Tokens {
			if string(l.ch) == value {
				switch value {
				case "=": // ==
					tok = l.makeTwoCharToken(key, token.ASSIGN, token.EQUAL)
				case "!": // !=
					tok = l.makeTwoCharToken(key, token.ASSIGN, token.NOT_EQUAL)
				default:
					tok = initToken(key, token.TokenLiteral(value))
				}
				break
			}
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' || l.ch == '\v' || l.ch == '\f' {
		l.readChar()
	}
}

func initToken(tokenType token.TokenType, literal token.TokenLiteral) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: literal,
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) makeTwoCharToken(firstChar token.TokenType, secondChar token.TokenType, doubleChar token.TokenType) token.Token {
	ch := string(l.ch)
	secondCh := token.Tokens[secondChar]
	if string(l.peekChar()) == secondCh {
		l.readChar()
		return initToken(doubleChar, token.TokenLiteral(ch+secondCh))
	}
	return initToken(firstChar, token.TokenLiteral(ch))
}
