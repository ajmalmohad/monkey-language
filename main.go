package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
	"strings"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nHello %s! Welcome to Monkey 1.0\n", strings.Split(user.Username, "\\")[0])
	fmt.Printf("Feel Free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
