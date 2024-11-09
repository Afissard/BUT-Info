package main

import (
	"log"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var turn atomic.Int64 // variable pour réserver l'accès à la section critique

func swap(x, y *int, ID int64) {
	defer wg.Done()
	for {
		if turn.CompareAndSwap(ID, -1) { // demande d'accès et réservation
			//log.Print(ID)
			// début section critique
			*x, *y = *y, *x
			// fin section critique
			turn.Store(ID + 1) // libération pour le suivant
		}
	}
}

func main() {
	x := 0
	y := 1

	turn.Store(0) // initialement l'accès est ouvert à l'ID 0

	log.Print(x, y)

	var i int64
	for i = 0; i < 1000; i++ {
		wg.Add(1)
		go swap(&x, &y, i)
	}

	wg.Wait()

	log.Print(x, y)
}
