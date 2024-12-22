package main

import "fmt"

func main() {
	var n int = 5
	var a *int = &n
	fmt.Println(n, *a)

	n = 6
	fmt.Println(n, *a)

	var m int = 7
	a = &m
	fmt.Println(n, *a, m)
}
