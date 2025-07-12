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
	BANG = "!"

	//delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPARAN = "("
	RPARAN = ")"
	LBRACE = "{"
	RBRACE = "}"

	LT = "<"
	GT = ">"
	EQ = "=="
	NOT_EQ = "!="

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF		 = "IF"
	ELSE	 = "ELSE"
	RETURN	 = "RETURN"
	TRUE	 = "TRUE"
	FALSE	 = "FALSE"
)

var keywords map[string]TokenType = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
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
