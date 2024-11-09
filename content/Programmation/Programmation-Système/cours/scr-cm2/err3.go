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
		log.Print("boucle1 ", x)
		x++
	}
	x = 0
	for i := 0; i < 3; i++ {
		log.Print("boucle2 ", i)
	}
	mux.Unlock()
	wg.Done()
}

func main() {
	wg.Add(2)
	go routine()
	go routine()
	wg.Wait()
}
