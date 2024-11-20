package main

import "fmt"

func main() {

	// var c chan int = make(chan int, 3)
	var c chan int = make(chan int, 4)

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
