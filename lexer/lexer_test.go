package lexer

import (
	"monkey/token"
	"testing"
)

func TestBasicTokens(test *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := CreateLexer(input)
	for index, testcase := range tests {
		tok := lexer.NextToken()

		if tok.Type != testcase.expectedType {
			test.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				index, testcase.expectedType, tok.Type)
		}

		if tok.Literal != testcase.expectedLiteral {
			test.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				index, testcase.expectedLiteral, tok.Literal)
		}
	}
}

func TestStatementTokens(test *testing.T) {
	input := `let five = 5;
			  let ten = 10;
			  let add = fn(x, y){
				x + y;
			  };
			  let result=add(five, ten);
			  !-/*5;
			  5< 10 > 5;
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
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},

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
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	lexer := CreateLexer(input)
	for index, testcase := range tests {
		tok := lexer.NextToken()

		if tok.Type != testcase.expectedType {
			test.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				index, testcase.expectedType, tok.Type)
		}

		if tok.Literal != testcase.expectedLiteral {
			test.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				index, testcase.expectedLiteral, tok.Literal)
		}
	}
}

func TestExpressionTokens(test *testing.T) {
	input := `if(5 < 10){
				return true;
			  }else{
				return false;
			  }
			  10 == 10;
			  10 != 9;
			  10 >= 9;
			  10 <= 9;
			  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},

		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},

		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},

		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.INT, "10"},
		{token.NOTEQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},

		{token.INT, "10"},
		{token.GTE, ">="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},

		{token.INT, "10"},
		{token.LTE, "<="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	lexer := CreateLexer(input)
	for index, testcase := range tests {
		tok := lexer.NextToken()

		if tok.Type != testcase.expectedType {
			test.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				index, testcase.expectedType, tok.Type)
		}

		if tok.Literal != testcase.expectedLiteral {
			test.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				index, testcase.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	input := `let five = 5;
			let ten = 10;

			let add = fn(x, y) {
			x + y;
			};

			let result = add(five, ten);
			!-/*5;
			5 < 10 > 5;

			if (5 < 10) {
				return true;
			} else {
				return false;
			}

			10 == 10;
			10 != 9;
			"foobar"
			"foo bar"
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
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
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
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOTEQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		// {token.LBRACKET, "["},
		// {token.INT, "1"},
		// {token.COMMA, ","},
		// {token.INT, "2"},
		// {token.RBRACKET, "]"},
		// {token.SEMICOLON, ";"},
		// {token.LBRACE, "{"},
		// {token.STRING, "foo"},
		// {token.COLON, ":"},
		// {token.STRING, "bar"},
		// {token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := CreateLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
