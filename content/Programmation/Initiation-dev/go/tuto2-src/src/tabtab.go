package main

import "fmt"

func main() {
	var tab [][]int = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5},
		[]int{6, 7, 8, 9, 10},
	}

	fmt.Println("tab =", tab)
	fmt.Println("tab[0] =", tab[0])
	fmt.Println("tab[0][1] =", tab[0][1])
}
