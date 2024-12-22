package main

import (
	"fmt"
)

func main() {
	var num int
	for {
		fmt.Print("Give a positive number: ")
		fmt.Scanf("%d", &num)
		if IsPrime(num) {
			fmt.Println(num, " is prime")
		} else {
			fmt.Println(num, " is not prime")
		}
	}
}
