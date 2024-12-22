package main

import "fmt"

func main() {
	var n int = 5
	var a *int = &n
	var b **int = &a
	var c ***int = &b
	fmt.Println(n, *a, **b, ***c)

	***c = 7
	fmt.Println(n, *a, **b, ***c)
}
