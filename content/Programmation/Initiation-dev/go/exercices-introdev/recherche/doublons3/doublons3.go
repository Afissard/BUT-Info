package doublons3

import "fmt"

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient
exactement k fois chaque nombre compris entre 1 et n/k. On voudrait vérifier
cela. C'est le travail de la fonction doublons.

# Entrée
- tab : un tableau d'entiers
- k : un entier

# Sortie
- ok : un booléen qui doit valoir true si tab contient exactement k fois chaque
entiers de 1 à len(tab)/k et false sinon
*/

func doublons(tab []int, k int) (ok bool) {
	ok = true

	// pré-test
	if (tab == nil) || (len(tab) == 0) {
		// si tab vide (=> inutile de chercher)
		return ok
	} else if len(tab) == 1 {
		// si tab avec un seul élément (=> inutile de chercher)
		if tab[0] != 1 {
			ok = false
		}
		return ok
	}

	// test
	const NK = int(len(tab) / k) // TODO: fix this
	doublons_count := [NK]int{}  // TODO: réduire à la taille strictement nécessaire
	for i := 0; i < len(tab); i++ {
		if (tab[i] > NK) || (doublons_count[tab[i]] >= NK) {
			fmt.Println("Hors ensemble")
			return !ok
		} else {
			doublons_count[tab[i]]++
		}
	}
	for i := 0; i < len(doublons_count); i++ {
		if doublons_count[i] != k {
			fmt.Println("Pas assez ou trop")
			return !ok
		}
	}
	return ok
}
