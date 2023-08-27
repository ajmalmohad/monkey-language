package lexer

import (
	"monkey/token"
)

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

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.character) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.character) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readString() string {
	position := lexer.position + 1
	for {
		lexer.readChar()
		if lexer.character == '"' || lexer.character == 0 {
			break
		}
	}
	return lexer.input[position:lexer.position]
}

func createToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func createStringToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_'
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}
