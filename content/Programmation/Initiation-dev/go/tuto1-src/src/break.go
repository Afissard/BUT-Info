package main

import "fmt"

func main() {
	var n int = 10
	var p int = 5
	for i := 0; i < n; i++ {
		fmt.Println("L'itération", i, "commence")
		if i == p {
			break
		}
		fmt.Println("L'itération", i, "termine")
	}
}
