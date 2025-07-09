package lexer

import (
	"fmt"
	"monkey/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextToken(t *testing.T) {
	t.Run("pre defined Literals check", func(t *testing.T) {
		input := `=+-*/(){},;`

		tests := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.ASSIGN, "="},
			{token.PLUS, "+"},
			{token.SUBTRACT, "-"},
			{token.MULTIPLY, "*"},
			{token.DIVISION, "/"},
			{token.LPARAN, "("},
			{token.RPARAN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
			{token.COMMA, ","},
			{token.SEMICOLON, ";"},
		}

		l := New(input)

		for i, test := range tests {
			tok := l.NextToken()
			assert.Equal(t, tok.Type, test.expectedType, fmt.Sprintf("tests[%d] failed; Token Type did not match; expected: %q , but got: %q", i, test.expectedType, tok.Type))
			// literal check
			assert.Equal(t, tok.Literal, test.expectedLiteral, fmt.Sprintf("tests[%d] failed; Literal did not match; expected: %q, but got : %q", i, test.expectedLiteral, tok.Literal))
		}
	})

	t.Run("testing functions and variables ", func(t *testing.T) {
		input := `let five = 5;
				  let ten = 10;

				let add = fn(x, y) {
 					 x + y;
				};

				let result = add(five, ten);
				`
		tests := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "ten"},
			{token.ASSIGN, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "add"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "fn"},
			{token.LPARAN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPARAN, ")"},
			{token.LBRACE, "{"},
			{token.IDENT, "x"},
			{token.PLUS, "+"},
			{token.IDENT, "y"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "result"},
			{token.ASSIGN, "="},
			{token.IDENT, "add"},
			{token.LPARAN, "("},
			{token.IDENT, "five"},
			{token.COMMA, ","},
			{token.IDENT, "ten"},
			{token.RPARAN, ")"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}

		l := New(input)

		for i, test := range tests {
			tok := l.NextToken()
			assert.Equal(t, test.expectedType, tok.Type, fmt.Sprintf("tests[%d] , expected : %s but got type: %s", i, test.expectedType, tok.Type))
			assert.Equal(t, test.expectedLiteral, tok.Literal, fmt.Sprintf("tests[%d] , expected : %s but got type: %s", i, test.expectedLiteral, tok.Literal))
		}

	})

}
