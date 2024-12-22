package main

import "fmt"

func main() {
	var reste int = 32
	var quotient int = 0
	const diviseur int = 4
	for reste >= diviseur {
		reste = reste - diviseur
		quotient++
	}
	fmt.Println(quotient)
}
