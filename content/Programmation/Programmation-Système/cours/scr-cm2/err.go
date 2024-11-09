package main

import (
	"log"
	"sync"
)

var wg sync.WaitGroup
var mux sync.Mutex

var x int

func routine() {
	mux.Lock()
	for x < 5 {
		log.Print(x)
		x++
	}
	mux.Unlock()
	x = 0
	wg.Done()
}

func main() {
	wg.Add(2)
	go routine()
	go routine()
	wg.Wait()
}
