package main

import (
	"fmt"
	"sync"
)

var x, y int
var w sync.WaitGroup
var mu sync.Mutex // add a mutex here

func switchxy() {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		x, y = y, x
		mu.Unlock() // defer n'est pas toujours la meilleur des idée, il nécessite de return pour qu'il soit exécuté
	}
	w.Done()
}

func main() {
	x = 5
	y = 7
	w.Add(1000)
	for i := 0; i < 1000; i++ {
		go switchxy() // ont sait pas qui et quand va accèdé aux valeur et les modifie
	}
	w.Wait()
	fmt.Println("x vaut", x, "et y vaut", y)
}
