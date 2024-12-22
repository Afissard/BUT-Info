/*
  Calcul des termes de la suite de fibonacci par programmation dynamique
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var t []int

func f(n int) int {
	// On évite de recalculer des choses
	if t[n] >= 0 {
		return t[n]
	}
	// On stocke ce qu'on calcule de nouveau
	t[n] = f(n-1) + f(n-2)
	return t[n]
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Il faut exactement un argument")
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("L'argument doit être un entier")
	}
	// Initialisation du tableau t puis appel de f
	t = make([]int, n+1)
	t[0] = 0
	if n > 0 {
		t[1] = 1
	}
	for i := 2; i <= n; i++ {
		t[i] = -1
	}
	fmt.Println(f(n))
}
