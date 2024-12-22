package doublons1

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient
exactement les nombres de 1 à n dans l'ordre. On voudrait vérifier cela. C'est
le travail de la fonction doublons.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui doit valoir true si tab contient exactement les entiers de
1 à len(tab) dans l'ordre et false sinon
*/

func doublons(tab []int) (ok bool) {
	ok = true

	if len(tab)<=1{ // si tab vide ou avec 1 seul élément (=> initle de chercher)
		return ok
	}
	// else implicite
	for i:=1; i<len(tab); i++{ // commence dans la cellule 2 pour le test
		if tab[i] <= tab[i-1] || tab[i]>(tab[i-1]+1) {
			ok = false
			return ok
		}
	}
	return ok
}
