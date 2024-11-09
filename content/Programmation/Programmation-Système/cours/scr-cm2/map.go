package main

import "sync"

var wg sync.WaitGroup

func writeMap(m map[int]int, k int) {
	defer wg.Done()
	m[k] = k
}

func main() {

	m := make(map[int]int)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go writeMap(m, i)
	}

	wg.Wait()

}
