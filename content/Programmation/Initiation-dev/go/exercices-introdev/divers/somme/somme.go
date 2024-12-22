package somme

/*
La fonction somme doit calculer la somme des nombres contenus dans un tableau
d'entiers puis retourner cette somme. Il faudra bien penser à se demander ce
qu'est la somme d'un tableau vide.

# Entrée
- tab : un tableau d'entiers

# Sortie
- res : un entier correspondant à la somme des éléments de tab

# Info
2021-2022, test 1, exercice 4
*/

func somme(tab []int) (res int) {
	for i:=0; i<len(tab); i++{
		res += tab[i]
	}
	return res
}
