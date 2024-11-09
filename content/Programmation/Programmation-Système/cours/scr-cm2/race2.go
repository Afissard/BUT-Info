package main

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func swap(x, y *int) {
	defer wg.Done()
	*x, *y = *y, *x
}

func main() {
	x := 0
	y := 1

	log.Print(x, y)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go swap(&x, &y)
		time.Sleep(time.Millisecond)
	}

	wg.Wait()

	log.Print(x, y)
}
