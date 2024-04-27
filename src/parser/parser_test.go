package parser

import (
	"fmt"
	"testing"

	"github.com/Shubham19032004/plus/src/ast"
	"github.com/Shubham19032004/plus/src/lexer"
)

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
	input string
	leftValue int64
	operator string
	rightValue int64
	}{
	{"5 + 5;", 5, "+", 5},
	{"5 - 5;", 5, "-", 5},
	{"5 * 5;", 5, "*", 5},
	{"5 / 5;", 5, "/", 5},
	{"5 > 5;", 5, ">", 5},
	{"5 < 5;", 5, "<", 5},
	{"5 == 5;", 5, "==", 5},
	{"5 != 5;", 5, "!=", 5},
	}
	for _, tt := range infixTests {
	l := lexer.New(tt.input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statement) != 1 {
	t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
	1, len(program.Statement))
	}
	stmt, ok := program.Statement[0].(*ast.ExpressionStatement)
	if !ok {
	t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
	program.Statement[0])
	}
	exp, ok := stmt.Expression.(*ast.InfixExpression)
	if !ok {
	t.Fatalf("exp is not ast.InfixExpression. got=%T", stmt.Expression)
	}
	if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
	return
	}
	if exp.Operator != tt.operator {
	t.Fatalf("exp.Operator is not '%s'. got=%s",
	tt.operator, exp.Operator)
	}
	if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
	return
	}
	}
	}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
	t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
	return false
	}
	if integ.Value != value {
	t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
	return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
	t.Errorf("integ.TokenLiteral not %d. got=%s", value,
	integ.TokenLiteral())
	return false
	}
	return true
	}
	func checkParserErrors(t *testing.T, p *Parser) {
		errors := p.Errors()
		if len(errors) == 0 {
		return
		}
		t.Errorf("parser has %d errors", len(errors))
		for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
		}
		t.FailNow()
		}