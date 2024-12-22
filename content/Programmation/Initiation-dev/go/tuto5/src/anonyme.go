package main

import "log"

func main() {

	var x int = func(x, y int) int { return x + y }(2, 3)

	var f func(int, int) int = func(x, y int) int { return x - y }

	log.Print(x, f(2, 3))

}
