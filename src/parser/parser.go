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
	return nil
}