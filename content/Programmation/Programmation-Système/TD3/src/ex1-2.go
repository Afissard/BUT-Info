package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup

// calcule n puissance m
func puissance(n, m int) {
	var res int = 1
	for i := 0; i < m; i++ {
		res = n
	}
	fmt.Println(n, "puissance", m, "vaut", res)
	w.Done()
	return
}

func main() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			w.Add(1)
			go puissance(i, j)
		}
	}

	w.Wait()
}

/*
Si i et j étais passé en parametre avec des pointeurs, il y aurais une race condition, mais dans
ce cas la, seule la valeur est passé, il n'y a donc aucun problème.
*/