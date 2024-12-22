package main

import "fmt"

func main() {
	var tint []int = []int{1, 2, 3, 1000, 27}
	fmt.Println(tint[0])
	fmt.Println(tint[3])
	fmt.Println(tint)
	tint[3] = 0
	fmt.Println(tint)
}
