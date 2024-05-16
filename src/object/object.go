package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Shubham19032004/plus/src/ast"
)

type ObjectType string
type BuiltinFunction func (args ...Object)  Object
const (
	BUILTIN_OBJ="BUILTIN"
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	STRING_OBJ       = "STRING"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
)

type String struct {
	Value string
}

type Error struct {
	Message string
}
type Builtin struct{
	Fn BuiltinFunction
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

type Object interface {
	Type() ObjectType
	Inspect() string
}
type Integer struct {
	Value int64
}
type Boolean struct {
	Value bool
}
type ReturnValue struct {
	Value Object
}
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

type Null struct{}

// Integer object
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// Boolean object
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// NULL  Object
func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

// Return Object
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

// STRING OBJ
func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

//BUILDIN function object 
func (b *Builtin) Type() ObjectType {return BOOLEAN_OBJ}
func (b *Builtin) Inspect() string {return "builtin function"}

// Function Object
func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())

	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString("){\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}
