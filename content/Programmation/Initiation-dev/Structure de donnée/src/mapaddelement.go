package main

import "fmt"

func main() {
	var m map[string]bool = map[string]bool{
		"bonjour": true,
		"salut":   false,
	}
	fmt.Println("Initialisation :", m)

	m["yo"] = false
	fmt.Println("Ajout yo :", m)

	m["bonjour"] = false
	fmt.Println("Ajout bonjour :", m)
}
