package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/MihaiBlebea/go-interpreter/repl"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal()
	}

	fmt.Printf("Hello %s", usr.Username)
	repl.Start(os.Stdin, os.Stdout)
}
