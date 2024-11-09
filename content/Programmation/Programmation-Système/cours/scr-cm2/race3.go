package main

import (
	"log"
	"sync"
)

var wg sync.WaitGroup
var open bool // variable pour réserver l'accès à la section critique

func swap(x, y *int) {
	defer wg.Done()
	for {
		if open { // demande d'accès
			open = false // réservation
			// début section critique
			*x, *y = *y, *x
			// fin section critique
			open = true // libération
			return
		}
	}
}

func main() {
	x := 0
	y := 1

	open = true // initialement l'accès est ouvert

	log.Print(x, y)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go swap(&x, &y)
	}

	wg.Wait()

	log.Print(x, y)
}
