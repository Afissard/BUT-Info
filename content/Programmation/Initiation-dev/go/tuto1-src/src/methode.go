package main

import "fmt"

type cuple struct {
	first  int
	second int
}

func (c cuple) add() int {
	return c.first + c.second
}

func main() {
	var c cuple = cuple{first: 1, second: 2}
	fmt.Println(c.add())
}
