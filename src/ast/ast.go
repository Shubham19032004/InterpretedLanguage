// Abstract Syntax Tree

package ast

import (
	"bytes"
	"github.com/Shubham19032004/plus/src/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

// All statement nodes implement this
type Statement interface {
	Node
	statementNode()
}

// All expression nodes implement this
type Expression interface {
	Node
	expressionNode()
}

// ROOT NODE OF EVERY AST
type Program struct {
	Statement []Statement
}

type Identifier struct {
	Token token.Token
	Value string
}

// FOR LET STATEMENT AND EXPRESSION
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// FOR RETURN STATEMENT
type ReturnStatement struct {
	Token       token.Token //the  'return 'token
	ReturnValue Expression
}

// FOR EXPRESSION STATEMENT
type ExpressionStatement struct {
	Token      token.Token //the first token of the expression
	Expression Expression  //hold the Expression
}

type PrefixExpression struct {
	Token    token.Token //prefix token e.g=>!
	Operator string
	Right    Expression
}

// FOR INTEGER  STATEMENT
type IntegerLiteral struct {
	Token token.Token
	Value int64 //hold the actual value of the integer not the string value
}

// FOR INFIX EXPRESSION 5 + 6
type InfixExpression struct {
	Token    token.Token //The operator token ,e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

// helper functions
func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")
	return out.String()

}

// For let statement
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// FOR IDENTIFIER
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// FOR RETURN STATEMENT
func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// FOR EXPRESSION
func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// FOR INTEGER STATEMENT
func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

func (p *Program) TokenLiteral() string {
	if len(p.Statement) > 0 {
		return p.Statement[0].TokenLiteral()
	} else {
		return ""
	}
}

// Prefix Operators
func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// Add () arround  operator and its operand
func (pe *PrefixExpression) String() string {

	var out bytes.Buffer
	out.WriteString("(")
	// Operator is going to be string that can be either ! or -
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

// creates a buffer and writes the return value of each statements String() method to it
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statement {
		out.WriteString(s.String())
	}
	return out.String()
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
func (i *Identifier) String() string { return i.Value }
