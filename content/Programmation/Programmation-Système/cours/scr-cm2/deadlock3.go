package main

import "sync"

var x int
var m sync.Mutex
var wg, w sync.WaitGroup

func routine1() {
	m.Lock()
	x++
	w.Done()
	m.Unlock()
	wg.Done()
}

func routine2() {
	m.Lock()
	x--
	w.Wait()
	m.Unlock()
	wg.Done()
}

func main() {
	for {
		wg.Add(2)
		w.Add(1)
		go routine1()
		go routine2()
		wg.Wait()
	}
}
