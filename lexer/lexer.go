package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	character    byte
}

func CreateLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	var current token.Token
	lexer.skipWhiteSpace()

	switch lexer.character {
	case '=':
		current = lexer.makeTwoCharToken('=', token.EQ, token.ASSIGN)
	case '+':
		current = createToken(token.PLUS, lexer.character)
	case '-':
		current = createToken(token.MINUS, lexer.character)
	case '*':
		current = createToken(token.ASTERISK, lexer.character)
	case '/':
		current = createToken(token.SLASH, lexer.character)
	case '!':
		current = lexer.makeTwoCharToken('=', token.NOTEQ, token.BANG)
	case '<':
		current = lexer.makeTwoCharToken('=', token.LTE, token.LT)
	case '>':
		current = lexer.makeTwoCharToken('=', token.GTE, token.GT)
	case ';':
		current = createToken(token.SEMICOLON, lexer.character)
	case '(':
		current = createToken(token.LPAREN, lexer.character)
	case ')':
		current = createToken(token.RPAREN, lexer.character)
	case ',':
		current = createToken(token.COMMA, lexer.character)
	case '{':
		current = createToken(token.LBRACE, lexer.character)
	case '}':
		current = createToken(token.RBRACE, lexer.character)
	case '"':
		current.Type = token.STRING
		current.Literal = lexer.readString()
	case 0:
		current.Type = token.EOF
		current.Literal = ""
	default:
		if isLetter(lexer.character) {
			current.Literal = lexer.readIdentifier()
			current.Type = token.LookUpIdentity(current.Literal)
			return current
		} else if isDigit(lexer.character) {
			current.Type = token.INT
			current.Literal = lexer.readNumber()
			return current
		} else {
			current = createToken(token.ILLEGAL, lexer.character)
		}
	}

	lexer.readChar()
	return current
}
