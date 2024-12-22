/*
  Calcul des termes de la suite de fibonacci avec récursivité terminale
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func f(n, f1, f2 int) int {
	if n == 0 {
		return f2
	}
	if n == 1 {
		return f1
	}
	return f(n-1, f1+f2, f1)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Il faut exactement un argument")
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("L'argument doit être un entier")
	}
	fmt.Println(f(n, 1, 0))
}
