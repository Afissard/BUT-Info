package main

import (
	"fmt"
	"sync"
)

var x, y int
var w sync.WaitGroup
var mx, my sync.Mutex

func switchxy() {
	// il existe un certain pourcentage de chance que la fonction ne récupère qu'un des lock
	mx.Lock() 
	// my.Lock()
	for i := 0; i < 1000; i++ {
		x, y = y, x
	}
	mx.Unlock()
	// my.Unlock()
	w.Done()
}

func switchyx() {
	// my.Lock()
	mx.Lock()
	for i := 0; i < 1000; i++ {
		x, y = y, x
	}
	// my.Unlock()
	mx.Unlock()
	w.Done()
}

func main() {
	x = 5
	y = 7
	w.Add(2)
	go switchxy()
	go switchyx()
	w.Wait()
	fmt.Println("x vaut", x, "et y vaut", y)

	// pas de deadlock ?
}
