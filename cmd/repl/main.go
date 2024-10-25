package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ignoxx/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s! This is Monkey-lang\n", user.Username)
	fmt.Printf("Type your commands below\n")

	repl.Start(os.Stdin, os.Stdout)
}
