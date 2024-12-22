package main

import "log"

const a int = 3

var b int

func main() {
	var c int = 1
	for i := 0; i < a; i++ {
		b += c
		c++
	}
	log.Print("a vaut ", a, ", b vaut ", b, " et c vaut ", c)
}
