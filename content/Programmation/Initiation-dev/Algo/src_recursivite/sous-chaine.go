/*
Calcul de la longueur de la plus longue sous-chaine commune à deux chaînes de
caractères s1 et s2
*/
package main

import (
	"fmt"
	"log"
	"os"
)

// Tableau pour stocker les solutions déjà connues
// t[i][j] est le résultat de sousChaine(s1[:i],s2[:j])
var t [][]int

// Fonction utilitaire pour calculer le maximum de deux entiers
func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}

// Fonction récursive pour résoudre le problème, utilise la programmation
// dynamique pour éviter de faire plusieurs fois le même appel récursif
func sousChaine(s1, s2 string) int {
	// Si on ne connait pas déjà la solution on la calcule et on la stocke dans t
	if t[len(s1)][len(s2)] < 0 {
		if s1[len(s1)-1] == s2[len(s2)-1] {
			// Les derniers caractères de s1 et s2 sont identiques : ce caractère c
			// appartient à une plus longue sous-chaîne commune. La longueur de cette
			// sous-chaîne est donc 1 + la longueur d'une plus longues sous-chaîne
			// commune à s1' et s2', où s1=s1'c et s2=s2'c
			t[len(s1)][len(s2)] = 1 + sousChaine(s1[:len(s1)-1], s2[:len(s2)-1])
		} else {
			// Les derniers caractères de s1=s1'c1 et s2=s2'c2 sont différents. La
			// longueur d'une plus longue sous-chaîne commune à s1 et s2 est donc le
			// maximum de la longueur d'une plus longue sous-chaîne commune à s1' et
			// s2 et de la longueur d'une plus longue sous-chaîne commune à s1 et s2'
			t[len(s1)][len(s2)] = max(sousChaine(s1[:len(s1)], s2[:len(s2)-1]), sousChaine(s1[:len(s1)-1], s2[:len(s2)]))
		}
	}
	// On retourne la solution stockée dans t
	return t[len(s1)][len(s2)]
}

func main() {

	// Les deux chaînes à analyser sont données en arguments
	if len(os.Args) != 3 {
		log.Fatal("Il faut exactement deux arguments")
	}
	s1 := os.Args[1]
	s2 := os.Args[2]

	// Initialisation du tableau t, on sait que la plus longue sous-chaîne est de
	// longueur 0 quand l'une des chaînes est vide, c'est à dire qu'on peut
	// remplire toutes les cases de la forme t[0][.] et toutes celles de la forme
	// t[.][0] avec des 0. On ne connaît aucune autre solution, les autres cases
	// sont remplies avec des -1
	t = make([][]int, len(s1)+1)
	for i := 0; i < len(t); i++ {
		t[i] = make([]int, len(s2)+1)
		for j := 0; j < len(t[i]); j++ {
			if i == 0 {
				t[i][j] = 0
			} else if j == 0 {
				t[i][j] = 0
			} else {
				t[i][j] = -1
			}
		}
	}

	fmt.Println("La plus longue sous séquence commune à", s1, "et", s2, "a pour longueur", sousChaine(s1, s2))

}
