package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	// END OF FILE 
	EOF     = "EOF"

	// identifiers+literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUSE  = "+"

	// DELIMIETES
	COMMA     = ","
	SEMICOLOM = ";"

	LPARAM = "("
	RPARAM = ")"
	LBRACE = "}"
	RBRACE = "{"

	// Keywords
	FUCNTION = "FUNCTION"
	LET      = "LET"
)
