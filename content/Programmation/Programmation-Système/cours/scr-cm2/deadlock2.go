package main

import "sync"

var c chan int
var m sync.Mutex
var wg sync.WaitGroup
var x int

func routine1() {
	m.Lock()
	x++
	c <- x
	m.Unlock()
}

func routine2() {
	m.Lock()
	x = <-c
	m.Unlock()
}

func main() {
	for {
		wg.Add(2)
		go routine1()
		go routine2()
		wg.Wait()
	}
}
