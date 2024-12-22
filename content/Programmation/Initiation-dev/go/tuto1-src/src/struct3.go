package main

import "fmt"

type cuple struct {
	first  int
	second int
}

func main() {
	var c cuple = cuple{first: 1, second: 2}
	fmt.Println(c, c.first, c.second)

	c.first = 3
	fmt.Println(c, c.first, c.second)
}
