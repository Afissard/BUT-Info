package main

import "fmt"

func main() {
	var n int = 10
	if n < 5 {
		fmt.Println("Moins que 5")
	} else if n <= 10 {
		fmt.Println("Entre 5 et 10")
	} else {
		fmt.Println("Plus de 10")
	}
}
