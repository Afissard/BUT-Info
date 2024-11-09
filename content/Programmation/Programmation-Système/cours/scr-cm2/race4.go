package main

import (
	"log"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var open atomic.Bool // variable pour réserver l'accès à la section critique

func swap(x, y *int) {
	defer wg.Done()
	for {
		if open.CompareAndSwap(true, false) { // demande d'accès et réservation
			// début section critique
			*x, *y = *y, *x
			// fin section critique
			open.Store(true) // libération
			return
		}
	}
}

func main() {
	x := 0
	y := 1

	open.Store(true) // initialement l'accès est ouvert

	log.Print(x, y)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go swap(&x, &y)
	}

	wg.Wait()

	log.Print(x, y)
}
