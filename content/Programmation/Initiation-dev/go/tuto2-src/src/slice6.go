package main

import "fmt"

func main() {
	var tab []int = make([]int, 3, 5)
	fmt.Println(tab, len(tab), cap(tab))

	var tab2 []int = make([]int, 3)
	fmt.Println(tab2, len(tab2), cap(tab2))

	var tab3 []int = []int{0, 0, 0}
	fmt.Println(tab3, len(tab3), cap(tab3))
}
