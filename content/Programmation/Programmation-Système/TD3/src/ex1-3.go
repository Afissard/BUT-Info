package main

import (
	"fmt"
	"sync"
)

var res int // vas cause une race condition, res étant une var commune
var mu sync.Mutex

// calcule n puissance m
func puissance(n, m int) {
	mu.Lock()
	defer mu.Unlock() // est éxécuter au return
	res = 1
	for i := 0; i < m; i++ {
		res *= n
	}
	fmt.Println(n, "puissance", m, "vaut", res)
	return
}

func main() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			puissance(i, j)
		}
	}

}
