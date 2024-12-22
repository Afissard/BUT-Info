package pgcd

/*
On peut calculer le pgcd de deux nombres par un algorithme récursif simple en
remarquant que si a > b alors pgcd(a, b) = pgcd(a - b, b), si a < b alors
pgcd(a, b) = pgcd(a, b - a) et si a = b alors pgcd(a, b) = a = b

La fonction pgcd doit implanter cet algorithme simple de manière récursive.

# Entrées
- a : un entier strictement positif
- b : un entier strictement positif

# Sorties
- c : le pgcd de a et b
*/

func pgcd(a, b uint) (c uint) {
	return c
}
