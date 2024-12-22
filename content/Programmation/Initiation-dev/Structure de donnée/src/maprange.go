package main

import "fmt"

func main() {
	var m map[string]bool = map[string]bool{
		"bonjour":   true,
		"salut":     false,
		"bye":       false,
		"au revoir": true,
		"yo":        false,
	}

	for key, value := range m {
		fmt.Println("La clé", key, "est associée à la valeur", value)
	}
}
