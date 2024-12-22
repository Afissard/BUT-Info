package main

import "fmt"

func main() {
	var m map[string]bool = map[string]bool{
		"bonjour": true,
		"salut":   false,
	}

	fmt.Println("Valeur associée à la clé bonjour :", m["bonjour"])
	fmt.Println("Valeur associée à la clé bye :", m["bye"])

	var b, exists bool
	var dico []string = []string{"bonjour", "bye"}
	for _, key := range dico {
		b, exists = m[key]
		if exists {
			fmt.Println("La clé", key, "est présente dans la map et correspond à la valeur", b)
		} else {
			fmt.Println("La clé", key, "n'est pas présente dans la map")
		}
	}

}
