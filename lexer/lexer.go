package lexer

import (
	"monkey/token"
	"monkey/util"
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

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.character = 0
	} else {
		lexer.character = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for util.IsLetter(lexer.character) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for util.IsDigit(lexer.character) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func createToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func createStringToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func (lexer *Lexer) makeTwoCharToken(lookahead byte, foundToken token.TokenType, notFoundToken token.TokenType) token.Token {
	if lexer.peekChar() == lookahead {
		ch := lexer.character
		lexer.readChar()
		return createStringToken(foundToken, string(ch)+string(lexer.character))
	} else {
		return createToken(notFoundToken, lexer.character)
	}
}

func (lexer *Lexer) skipWhiteSpace() {
	for lexer.character == ' ' || lexer.character == '\t' || lexer.character == '\n' || lexer.character == '\r' {
		lexer.readChar()
	}
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
	case 0:
		current.Type = token.EOF
		current.Literal = ""
	default:
		if util.IsLetter(lexer.character) {
			current.Literal = lexer.readIdentifier()
			current.Type = token.LookUpIdentity(current.Literal)
			return current
		} else if util.IsDigit(lexer.character) {
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
