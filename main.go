// Created by nazarigonzalez on 16/8/17.

package main

import (
	"fmt"
	"os"
	"os/user"
	"waiig/repl"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", u.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
