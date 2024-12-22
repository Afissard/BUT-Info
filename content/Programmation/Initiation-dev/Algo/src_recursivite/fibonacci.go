/*
  Calcul des termes de la suite de fibonacci
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func f(n int) int {
	// Initialisation
	if n == 0 {
		return 0
	}
	// Initialisation
	if n == 1 {
		return 1
	}
	// Récurrence
	return f(n-1) + f(n-2)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Il faut exactement un argument")
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("L'argument doit être un entier")
	}
	fmt.Println(f(n))
}
