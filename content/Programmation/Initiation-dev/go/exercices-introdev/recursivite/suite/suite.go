package suite

/*
On considère la suite (un) définie par u(n) = 3 * u(n-1) + 5 et u(0) = 0
La fonction suite doit calculer le n-ième terme de (un) de manière récursive :
une fonction qui n'est pas récursive rapportera moins de points.

# Entrée
- n : un entier indiquant le terme de la suite à calculer

# Sortie
- un : le n-ième terme de la suite

# Info
2021-2022, test 1, exercice 3
*/

func suite(n uint) (un uint) {
	if n == 0{
		return 0 // u(0) = 0
	}
	un = suite(n-1)
	return 3* un +5
}
