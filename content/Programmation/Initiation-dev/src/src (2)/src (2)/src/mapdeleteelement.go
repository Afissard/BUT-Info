package main

import "fmt"

func main() {
	var m map[string]bool = map[string]bool{
		"bonjour": true,
		"salut":   false,
	}

	fmt.Println("Avant suppression de la clé bonjour :", m)
	delete(m, "bonjour")
	fmt.Println("Après suppression de la clé bonjour :", m)

	var dico []string = []string{"salut", "bye"}
	for _, key := range dico {
		delete(m, key)
	}

	fmt.Println("Après suppression des clés de dico :", m)
}
