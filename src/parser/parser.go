package parser

import (
	"github.com/Shubham19032004/plus/src/ast"
	"github.com/Shubham19032004/plus/src/lexer"
	"github.com/Shubham19032004/plus/src/token"
)

type Parser struct {
	l         *lexer.Lexer //Pointer to an instance of the lexer
	curToken  token.Token  //current Token
	peekToken token.Token  //Next token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// Read two tokens, so curToken and PeekToken are both set
	p.nextToken()
	p.nextToken()
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
	default:
		return nil

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

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
	p.nextToken()
	return true
	} else {
	return false
	}
	}