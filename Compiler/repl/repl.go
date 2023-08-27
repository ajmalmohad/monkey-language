package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
)

func StartREPL(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		input := scanner.Text()
		lex := lexer.CreateLexer(input)
		p := parser.CreateParser(lex)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
