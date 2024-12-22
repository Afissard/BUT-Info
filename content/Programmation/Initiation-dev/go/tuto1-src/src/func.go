package main

import "fmt"

func f(x int) bool {
	return x < 10
}

func main() {
	fmt.Println(f(3))
	fmt.Println(f(5))
	fmt.Println(f(7))
	fmt.Println(f(11))
	fmt.Println(f(13))
}
