package main

import (
	"log"
	"sync"
)

var com1, com2 chan int
var w sync.WaitGroup

func prod(n int) {
	defer w.Done()
	for i := 0; i < n; i++ {
		// com1 = append(com1, i)
		com1 <- i
	}
}

func trans() {
	defer w.Done()
	for {
		if len(com1) > 0 {
			// com2 = append(com2, com1[0])
			// com1 = com1[1:]
			com2 <- <-com1
		}
	}
}

func cons() {
	defer w.Done()
	for {
		if len(com2) > 0 {
			// log.Print(com2[0])
			// com2 = com2[1:]
			log.Println(<-com2)
		}
	}
}

func main() {
	size := 1000
	com1 = make(chan int, size)
	com2 = make(chan int, size)
	w.Add(3)
	go prod(size)
	go trans()
	go cons()
	w.Wait()
	log.Println("fini")
}
