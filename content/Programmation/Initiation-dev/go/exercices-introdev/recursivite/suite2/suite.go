package suite

/*
On considère la suite définie par U(n) = 2*U(n-1) + 1 et U(0) = 2. 
La fonction terme doit permettre d'obtenir le n-ième terme de cette suite.

Pour cet exercice, les boucles for sont interdites.

# Entrée
- n, un entier

# Sortie
- un, le terme U(n) de la suite

# Info
2022-2023, test 2, exercice 0
*/

func terme(n uint) (un int) {
	if n == 0{
		return 2 // U(0) = 2
	}
	un = terme(n-1)
	return 2*un+1
}
