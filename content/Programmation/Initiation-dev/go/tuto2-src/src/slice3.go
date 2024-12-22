package main

import "fmt"

func main() {
	var tab []int = []int{2, 12, 22, 11, 3, 5, 17, 92, 7}
	var t []int = tab[3:7]
	fmt.Println(tab)
	for i := 0; i < len(t); i++ {
		t[i] = 0
	}
	fmt.Println(tab)
}
