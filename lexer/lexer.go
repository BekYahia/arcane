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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	if l.ch == 0 {
		return initToken(token.EOF, "")
	}

	for key, value := range token.Tokens {
		if string(l.ch) == value {
			tok = initToken(key, token.TokenLiteral(value))
			break
		}
	}

	l.readChar()
	return tok
}

func initToken(tokenType token.TokenType, literal token.TokenLiteral) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: literal,
	}
}
