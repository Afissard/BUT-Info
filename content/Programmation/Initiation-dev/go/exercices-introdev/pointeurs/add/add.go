package addptr

/*
La fonction add doit calculer la somme de deux valeurs dont les adresses sont
données en paramètres, sans modifier ces valeurs.

# Entrées
- ax : l'adresse d'un entier
- ay : l'adresse d'un entier

# Sortie
- sum : la somme des valeurs situées aux adresses ax et ay

# Info
2021-2022, test2, exercice 0
*/

func addptr(ax, ay *int) (sum int) {
	sum = *ax + *ay
	return sum
}
