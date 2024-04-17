package parser

import (
	"fmt"
	"strconv"

	"github.com/Shubham19032004/plus/src/ast"
	"github.com/Shubham19032004/plus/src/lexer"
	"github.com/Shubham19032004/plus/src/token"
)

const (
	// 	higher preference or lower
	_           int = iota //0
	LOWEST                 //1
	EQUALS                 //2   ==
	LESSGREATER            // > or < //3
	SUM                    // +//4
	PRODUCT                // *//5
	PREFIX                 // -X or !X 6
	CALL                   // myFunction(X) 7
)

type Parser struct {
	l              *lexer.Lexer //Pointer to an instance of the lexer
	curToken       token.Token  //current Token
	errors         []string
	peekToken      token.Token //Next token
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}
type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	// Read two tokens, so curToken and PeekToken are both set
	p.nextToken()
	p.nextToken()
	p.prefixParseFns=make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT,p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	
	return p
}

// Helper function that is use to move token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParserProgram() *ast.Program {
	program := &ast.Program{}
	program.Statement = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statement = append(program.Statement, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	// calls parser for let statement
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseLetStatement()

	default:
		return p.parseExpressionStatement()

	}
}

// PARSER FOR LET STATEMENT
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// PARSER FOR RETURN STATEMENT
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// Expression Statement

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt

}
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}

// PARSER FOR ERROR STATEMENT
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// PEEK THE TOKEN
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// PEEK ERROR
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}
func (p *Parser) expectPeek(t token.TokenType) bool {
	// cheak the type of peel token
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	//convert string to integer 
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("could not pare %q as integer ", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value=value
	return lit
}

