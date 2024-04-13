package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	// END OF FILE
	EOF = "EOF"

	// identifiers+literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// DELIMIETES
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var Keywords=map[string]TokenType{
	"fn":FUNCTION,
	"let":LET,
}

// CHECK IF GIVEN IDENTIFIER IS IN FACT A KEYWORD
func LookupIdent(ident string) TokenType{
	if tok,ok:=Keywords[ident];ok{
		return tok
	}
	return IDENT
}