package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Shubham19032004/plus/src/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Plus progarming language!", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
