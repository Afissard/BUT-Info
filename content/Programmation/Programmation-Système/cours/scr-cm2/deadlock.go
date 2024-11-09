package main

import "sync"

var x, y int
var mx, my sync.Mutex
var wg sync.WaitGroup

func routine1() {
	mx.Lock()
	x--
	my.Lock()
	y = x
	my.Unlock()
	mx.Unlock()
	wg.Done()
}

func routine2() {
	my.Lock()
	y++
	mx.Lock()
	x = y
	mx.Unlock()
	my.Unlock()
	wg.Done()
}

func main() {
	for {
		wg.Add(2)
		go routine1()
		go routine2()
		wg.Wait()
	}
}
