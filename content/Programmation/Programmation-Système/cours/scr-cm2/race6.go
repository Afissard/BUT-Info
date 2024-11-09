package main

import (
	"log"
	"sync"
)

var wg sync.WaitGroup
var mux sync.Mutex // création d'un "semaphore"

func swap(x, y *int) {
	defer wg.Done()
	mux.Lock() // demande de réservation (bloquant)
	*x, *y = *y, *x
	mux.Unlock() // libération
}

func main() {
	x := 0
	y := 1

	log.Print(x, y)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go swap(&x, &y)
	}

	wg.Wait()

	log.Print(x, y)
}
