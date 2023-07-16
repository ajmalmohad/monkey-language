package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(test *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`

	lex := lexer.CreateLexer(input)
	parse := CreateParser(lex)

	program := parse.parseProgram()
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
