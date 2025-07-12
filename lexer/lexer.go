package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, literal byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(literal),
	}
}

func (l *Lexer) readIndentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input){
		return 0
	}
	return l.input[l.readPosition]
}



func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var currentToken token.Token
	l.skipWhiteSpace()
	switch l.ch {
	case '=':
		if l.peekChar() == '='{
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			currentToken = token.Token{Type: token.EQ, Literal: literal}
		}else {
			currentToken = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		currentToken = newToken(token.SEMICOLON, l.ch)
	case '(':
		currentToken = newToken(token.LPARAN, l.ch)
	case ')':
		currentToken = newToken(token.RPARAN, l.ch)
	case ',':
		currentToken = newToken(token.COMMA, l.ch)
	case '+':
		currentToken = newToken(token.PLUS, l.ch)
	case '-':
		currentToken = newToken(token.SUBTRACT, l.ch)
	case '*':
		currentToken = newToken(token.MULTIPLY, l.ch)
	case '/':
		currentToken = newToken(token.DIVISION, l.ch)
	case '{':
		currentToken = newToken(token.LBRACE, l.ch)
	case '}':
		currentToken = newToken(token.RBRACE, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			currentToken = token.Token{Type: token.NOT_EQ, Literal: literal}
		}else{
			currentToken = newToken(token.BANG, l.ch)
		}
	case '<':
		currentToken = newToken(token.LT, l.ch)
	case '>':
		currentToken = newToken(token.GT, l.ch)
	case 0:
		currentToken.Literal = ""
		currentToken.Type = token.EOF
	default:
		if isLetter(l.ch) {
			currentToken.Literal = l.readIndentifier()
			currentToken.Type = token.LookUpIdent(currentToken.Literal)
			return currentToken
		} else if isDigit(l.ch) {
			currentToken.Type = token.INT
			currentToken.Literal = l.readNumber()
			return currentToken
		} else {
			currentToken = newToken(token.ILLEGAL, l.ch)
		}

	}
	l.readChar()
	return currentToken
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}
