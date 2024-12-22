package main

import "fmt"

func g(x int, y int) (bool, int) {
	return x < y, x + y
}

func main() {
	var b bool
	var n int
	b, n = g(3, 4)
	fmt.Println(b)
	fmt.Println(n)
}
