package repl

import (
	"arcane/lexer"
	"arcane/token"
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}

	}
}
