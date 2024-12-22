package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Le programme s'appelle", os.Args[0])
	fmt.Println("Il y a", len(os.Args)-1, "arguments")
	for argPos := 1; argPos < len(os.Args); argPos++ {
		fmt.Println("L'argument", argPos, "vaut", os.Args[argPos])
	}
}
