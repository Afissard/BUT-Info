package main

import (
	"fmt"
	"sync"
)

var x, y int
var w sync.WaitGroup

func switchxy() {
	for i := 0; i < 1000; i++ {
		x, y = y, x
	}
	w.Done()
}

func main() {
	x = 5
	y = 7
	w.Add(1000)
	for i := 0; i < 1000; i++ {
		go switchxy()
	}
	w.Wait()
	fmt.Println("x vaut", x, "et y vaut", y)
}
