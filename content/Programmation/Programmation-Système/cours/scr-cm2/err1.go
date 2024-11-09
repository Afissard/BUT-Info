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
		log.Print("routine ", x)
		x++
	}
	x = 0
	mux.Unlock()
	wg.Done()
}

func autreRoutine() {
	for i := 0; i < 3; i++ {
		log.Print("autreRoutine ", i)
	}
	mux.Unlock()
	wg.Done()
}

func main() {
	wg.Add(3)
	go autreRoutine()
	go routine()
	go routine()
	wg.Wait()
}
