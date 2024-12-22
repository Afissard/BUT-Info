package main

import "fmt"

type cuple struct {
	first  int
	second int
}

func (c *cuple) exchange() {
	c.first, c.second = c.second, c.first
}

func main() {
	var c cuple = cuple{first: 1, second: 2}
	fmt.Println(c)
	c.exchange()
	fmt.Println(c)
}
