package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("You must pass your name.")
	}

	name := os.Args[1]
	fmt.Printf("Hello %s\n", name)
}
