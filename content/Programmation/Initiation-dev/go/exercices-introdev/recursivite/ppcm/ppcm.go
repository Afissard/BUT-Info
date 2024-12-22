package ppcm

/*
On peut calculer le ppcm de deux nombres x0 et y0, en appliquant la méthode
récursive suivante :
- ppcm(x, y) = x si x et y sont égaux
- ppcm(x, y) = ppcm(x + x0, y) si x < y
- ppcm(x, y) = ppcm(x, y + y0) si x > y
- on commence par appeler ppcm(x0, y0)

La fonction ppcm doit être une fonction récursive (ou une fonction qui se base
sur une fonction récursive) qui calcule le ppcm de deux nombres à partir de la
technique décrite ci-dessus. Une solution qui n'est pas récursive ne rapportera
pas de points.

Vous pouvez considérer que la fonction ppcm ne sera jamais appelée avec un de ses
arguments qui soit égal à 0.

# Entrées
- x : un entier
- y : un entier

# Sortie
- z : le ppcm de x et y

# Info
2021-2022, test2, exercice 5
*/

func ppcm(x, y uint) (z uint) {
	return z
}
