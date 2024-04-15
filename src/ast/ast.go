// Abstract Syntax Tree

package ast

import "github.com/Shubham19032004/plus/src/token"

type Node interface{
	TokenLiteral() string
}

type Statement interface{
	Node
	statementNode()
}

type Expression interface{
	Node
	expressionNode()
}

// ROOT NODE OF EVERY AST
type Program struct{
	Statement []Statement

}
type Identifier struct{
	Token token.Token
	Value string
}
//FOR LET STATEMENT AND EXPRESSION
type LetStatement struct{
	Token token.Token //the token.LET token
	Name *Identifier
	Value Expression
}

//FOR RETURN STATEMENT
type ReturnStatement struct{
	Token token.Token //the  'return 'token
	RetrunValue Expression
}






func (p *Program) TokenLiteral() string  {
	if len(p.Statement)>0{
		return p.Statement[0].TokenLiteral()
	}else{
		return ""
	}
}

// For let statement
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}
// FOR IDENTIFIER
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// FOR RETURN STATEMENT
func (rs *ReturnStatement) statementNode(){}
func (rs  *ReturnStatement) TokenLiteral() string {return rs.Token.Literal}

