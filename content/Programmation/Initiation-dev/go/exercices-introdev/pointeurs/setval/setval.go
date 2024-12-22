package setval

/*
La fonction set fixe la valeur stockée à l'adresse indiquée par un pointeur à la valeur d'un entier donné en paramètre

# Entrée
- ptr : l'adresse d'un entier, cet entier doit valoir x à l'issue de la fonction
- x : un entier

# Info
2022-2023, test 1, exercice 0
*/

func set(ptr *int, x int) {
	*ptr = x
}
