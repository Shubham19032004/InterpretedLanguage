package parser

import (
	"testing"

	"github.com/Shubham19032004/plus/src/ast"
	"github.com/Shubham19032004/plus/src/lexer"
)

// func TestLetStatements(t *testing.T) {
// 	input := "foobar;"
// 	l := lexer.New(input)
// 	p := New(l)
// 	program := p.ParserProgram()
// 	// Error
// 	checkParserErrors(t, p)
// 	if program == nil {
// 		t.Fatalf("ParseProgram() returned nil")
// 	}
// 	if len(program.Statement)!=1{
// 		t.Fatalf("program has not enough statements. got=%d",len(program.Statement))
// 	}
// 	if len(program.Statement) != 3 {
// 		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
// 			len(program.Statement))
// 	}
// 	// FOr Error
// 	for _,stmt :=range  program.Statement{
// 		returnStmt,ok:=stmt.(*ast.ReturnStatement)
// 		if !ok{
// 			t.Errorf("stmt not *ast.returnStatement. got=%T",stmt)
// 			continue
// 		}
// 		if returnStmt.TokenLiteral() != "return" {
// 			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
// 			returnStmt.TokenLiteral())
// 			}
// 	}
// 	tests := []struct {
// 		expectedIdentifer string
// 	}{
// 		{"x"},
// 		{"y"},
// 		{"foobar"},
// 	}
// 	for i, tt := range tests {
// 		stmt := program.Statement[i]
// 		if !testLetStatement(t, stmt, tt.expectedIdentifer) {
// 			return
// 		}
// 	}

// }
func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParserProgram()
	checkParserErrors(t, p)
	if len(program.Statement) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statement))
	}
	stmt, ok := program.Statement[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statement[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.TokenLiteral())
	}
}

// FOR TESTING LET STATEMENT
func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}
	return true
}

// CHECKING PARSER FOR  ERROR
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error:%q", msg)
	}
	t.FailNow()
}
