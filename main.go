package main

import (
	"arcane/repl"
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf(`
=========================================
Hello %s!

This is Arcane Programming Language!

Fee free to type in command	
=========================================

`, user.Username)
	fmt.Println()
	repl.Start(os.Stdin, os.Stdout)
}
