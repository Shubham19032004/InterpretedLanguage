package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEN"
	NULL_OBJ="NULL"
)

type object interface {
	Type() ObjectType
	Inspect() string
}
type Integer struct {
	Value int64
}
type Boolean struct {
	Value bool
}
type Null struct{}

// Interger object
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// Boolean object
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

//NULL  Object
func (n *Null) Type() ObjectType {return NULL_OBJ}
func (n *Null) Inspect() string {return "null"}