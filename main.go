package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Mik3y-F/flipflop-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nHello %s! This is the Flipflop programming language.\n", user.Username)
	fmt.Println("Feel free to type in some code and see what happens.")
	fmt.Print("Type 'exit' to quit.\n\n")

	repl.Start(os.Stdin, os.Stdout)
}
