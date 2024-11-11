package parser

import (
	"arcane/ast"
	"arcane/lexer"
	"arcane/token"
	"fmt"
	"strconv"
)

type Parser struct {
	l *lexer.Lexer

	currentToken   token.Token
	peekToken      token.Token
	errors         []string
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

type (
	prefixParseFn func() ast.Expression               // no left side, no argument
	infixParseFn  func(ast.Expression) ast.Expression // pass the left side as an argument
)

// precedence
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESS_GRATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -x or !x
	CALL        // myFunction(x)
)

func Init(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.NOT, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.LEFT_PARENTHESIS, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.DIVIDE, p.parseInfixExpression)
	p.registerInfix(token.MULTIPLY, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.EQUAL, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQUAL, p.parseInfixExpression)

	// twice so current and peek tokens are set
	p.nextToken()
	p.nextToken()

	return p
}
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Value: string(p.currentToken.Literal)}
}
func (p *Parser) parseIntegerLiteral() ast.Expression {
	defer unTrace(trace("parseIntegerLiteral"))
	literal := &ast.IntegralLiteral{Token: p.currentToken}

	value, err := strconv.ParseInt(string(p.currentToken.Literal), 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currentToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	literal.Value = value
	return literal
}
func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.currentToken, Value: p.currentTokenIs(token.TRUE)}
}
func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)

	if !p.expectedPeek(token.RIGHT_PARENTHESIS) {
		return nil
	}

	return exp
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	defer unTrace(trace("parsePrefixExpression"))
	expression := &ast.PrefixExpression{
		Token:    p.currentToken,
		Operator: string(p.currentToken.Literal),
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	defer unTrace(trace("parseInfixExpression"))
	expression := &ast.InfixExpression{
		Token:    p.currentToken,
		Operator: string(p.currentToken.Literal),
		Left:     left,
	}
	precedence := p.currentPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("Expected next token to be %s. got %s", token.Tokens[t], token.Tokens[p.peekToken.Type])
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {

	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.currentTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{
		Token: p.currentToken,
	}

	if !p.expectedPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.currentToken,
		Value: string(p.currentToken.Literal),
	}

	if !p.expectedPeek(token.ASSIGN) {
		return nil
	}

	// skip expression till the next semicolon
	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{
		Token: p.currentToken,
	}

	p.nextToken()

	// skip expression till the next semicolon
	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{
		Token: p.currentToken,
	}

	if !p.expectedPeek(token.LEFT_PARENTHESIS) {
		return nil
	}
	p.nextToken()

	expression.Condition = p.parseExpression(LOWEST)

	if !p.expectedPeek(token.RIGHT_PARENTHESIS) {
		return nil
	}

	if !p.expectedPeek(token.LEFT_CURLY_BRACKETS) {
		return nil
	}

	expression.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectedPeek(token.LEFT_CURLY_BRACKETS) {
			return nil
		}
		expression.Alternative = p.parseBlockStatement()
	}
	return expression
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{
		Token: p.currentToken,
	}
	block.Statements = []ast.Statement{}

	p.nextToken()
	for !p.currentTokenIs(token.RIGHT_CURLY_BRACKETS) && !p.currentTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectedPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	defer unTrace(trace("parseExpressionStatement"))
	stmt := &ast.ExpressionStatement{
		Token: p.currentToken,
	}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	defer unTrace(trace("parseExpression"))
	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.currentToken.Literal)
		return nil
	}
	leftExpression := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {

		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExpression
		}
		p.nextToken()
		leftExpression = infix(leftExpression)
	}

	return leftExpression
}

func (p *Parser) noPrefixParseFnError(t token.TokenLiteral) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

var precedences = map[token.TokenType]int{
	token.EQUAL:     EQUALS,
	token.NOT_EQUAL: EQUALS,
	token.LT:        LESS_GRATER,
	token.GT:        LESS_GRATER,
	token.PLUS:      SUM,
	token.MINUS:     SUM,
	token.DIVIDE:    PRODUCT,
	token.MULTIPLY:  PRODUCT,
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}
func (p *Parser) currentPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}
	return LOWEST
}
