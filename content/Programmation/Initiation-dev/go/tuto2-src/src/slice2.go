package main

import "fmt"

func main() {
	var tab []int = []int{2, 12, 22, 11, 3, 5, 17, 92, 7}
	fmt.Println(tab)

	var t1 []int = tab[3:7]
	fmt.Println(t1)

	var t2 []int = tab[:7]
	fmt.Println(t2)

	var t3 []int = tab[3:]
	fmt.Println(t3)
}
