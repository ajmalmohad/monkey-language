package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("Parse had %d errors!", len(errors))
	for _, msg := range errors {
		t.Errorf("Parser error: %q", msg)
	}
	t.FailNow()
}

func TestLetStatements(test *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`

	lex := lexer.CreateLexer(input)
	parse := CreateParser(lex)

	program := parse.parseProgram()
	checkParserErrors(test, parse)

	if program == nil {
		test.Fatalf("parseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		test.Fatalf("program.Statements doesn't contain 3 statements, instead got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for index, testcase := range tests {
		stmt := program.Statements[index]
		if !testLetStatement(test, stmt, testcase.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let' got=%T", s)
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s' got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s' got=%s", name, letStmt.Name)
		return false
	}

	return true
}

func TestReturnStatements(test *testing.T) {
	input := `
		return 5;
		return 10;
		return 838383;
	`

	lex := lexer.CreateLexer(input)
	parse := CreateParser(lex)

	program := parse.parseProgram()
	checkParserErrors(test, parse)

	if program == nil {
		test.Fatalf("parseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		test.Fatalf("program.Statements doesn't contain 3 statements, instead got=%d",
			len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			test.Errorf("s not *ast.returnStatement got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			test.Errorf("returnStmt.TokenLiteral not 'return', got=%q", returnStmt.TokenLiteral())
			continue
		}
	}
}

func TestIdentifierExpression(test *testing.T) {
	input := `foobar;`

	lex := lexer.CreateLexer(input)
	parse := CreateParser(lex)
	program := parse.parseProgram()
	checkParserErrors(test, parse)

	if program == nil {
		test.Fatalf("parseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		test.Fatalf("program.Statements doesn't contain 1 statements, instead got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		test.Errorf("s not *ast.ExpressionStatement got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		test.Errorf("s not *ast.Identifier got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		test.Errorf("ident.Value not %s, got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		test.Errorf("ident.TokenLiteral() not %s, got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(test *testing.T) {
	input := `5;`

	lex := lexer.CreateLexer(input)
	parse := CreateParser(lex)
	program := parse.parseProgram()
	checkParserErrors(test, parse)

	if program == nil {
		test.Fatalf("parseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		test.Fatalf("program.Statements doesn't contain 1 statements, instead got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		test.Errorf("s not *ast.ExpressionStatement got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		test.Errorf("s not *ast.IntegerLiteral got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		test.Errorf("literal.Value not %d, got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		test.Errorf("ident.TokenLiteral() not %s, got=%s", "5", literal.TokenLiteral())
	}
}
