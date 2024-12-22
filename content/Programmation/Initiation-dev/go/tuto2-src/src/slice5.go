package main

import "fmt"

func main() {
	var tab []int = make([]int, 5)
	fmt.Println(tab, len(tab))
	tab = append(tab, 2)
	fmt.Println(tab, len(tab))
}
