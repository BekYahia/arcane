package repl

import (
	"arcane/lexer"
	"arcane/parser"
	"bufio"
	"fmt"
	"io"
	"log"
)

func Start(in io.Reader, out io.Writer) {
	for {

		fmt.Print(">>> ")
		reader := bufio.NewReader(in)
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		l := lexer.Init(userInput)
		p := parser.Init(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			_, _ = io.WriteString(out, "Parser Errors:\n")
			for _, msg := range p.Errors() {
				_, _ = io.WriteString(out, "\t - "+msg+"\n")
			}
			continue
		}
		_, _ = io.WriteString(out, program.String())
		_, _ = io.WriteString(out, "\n")
	}
}
