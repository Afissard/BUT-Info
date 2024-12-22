package main

import "fmt"

func main() {
	var hello = "Hello, BUT"
	var helloAsRuneArray = []rune(hello)
	for i := 0; i < len(helloAsRuneArray); i++ {
		fmt.Print(string(helloAsRuneArray[i]), ".")
	}
	fmt.Println()
}
