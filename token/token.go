package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Indentifiers & Literals
	IDENT = "INDENT"
	INT   = "INT"

	//operators
	ASSIGN   = "="
	PLUS     = "+"
	SUBTRACT = "-"
	MULTIPLY = "*"
	DIVISION = "/"

	//delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPARAN = "("
	RPARAN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords map[string]TokenType = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookUpIdent(identifier string) TokenType {
	if tokenType, ok := keywords[identifier]; ok {
		return tokenType
	}
	return IDENT
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
