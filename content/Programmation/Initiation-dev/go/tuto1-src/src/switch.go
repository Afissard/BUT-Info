package main

import "fmt"

func main() {
	var n int = 10
	switch {
	case n < 5:
		fmt.Println("Moins que 5")
	case n <= 10:
		fmt.Println("Entre 5 et 10")
	default:
		fmt.Println("Plus de 10")
	}
}
