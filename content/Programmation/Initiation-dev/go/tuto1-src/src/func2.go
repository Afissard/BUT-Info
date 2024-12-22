package main

import "fmt"

func g(x int, y int) bool {
	return x < y
}

func main() {
	fmt.Println(g(3, 4))
	fmt.Println(g(4, 3))
	fmt.Println(g(5, 5))
}
