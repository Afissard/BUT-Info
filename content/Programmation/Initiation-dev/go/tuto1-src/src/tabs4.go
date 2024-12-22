package main

import "fmt"

func main() {
	var tint []int = []int{1, 2, 3, 1000, 27}
	fmt.Println(len(tint))
	tint = []int{4, 5, 6}
	fmt.Println(len(tint))
}
