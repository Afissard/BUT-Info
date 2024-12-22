package main

import "fmt"

func f(t []int) {
	for i := 0; i < len(t); i++ {
		t[i] = 0
	}
}

func main() {
	var tab []int = []int{1, 2, 3, 4}
	fmt.Println(tab)
	f(tab)
	fmt.Println(tab)
}
