package main

import (
	"fmt"
	"os"

	"github.com/micanzhang/galerie"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%+v", os.Args)
		return
	}

	cmd, args := os.Args[1], os.Args[2:]
	command := galerie.Cmd{
		Command: cmd,
	}

	command.Run(args...)
}
