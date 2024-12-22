package ppcm

/*
On peut calculer le ppcm de deux nombres x0 et y0, en appliquant la méthode
récursive suivante :
- ppcm(x, y) = x si x et y sont égaux
- ppcm(x, y) = ppcm(x + x0, y) si x < y
- ppcm(x, y) = ppcm(x, y + y0) si x > y
- on commence par appeler ppcm(x0, y0)
Cette méthode se généralise au calcul du ppcm de n nombres.

La fonction ppcm est une fonction récursive qui calcule le ppcm de n nombres à
partir de la technique décrite ci-dessus.

Vous pouvez considérer que la fonction ppcm ne sera jamais appelée avec un des
nombres considérés qui soit égal à 0.

# Entrées
- tab0 : un tableau d'entiers (ce tableau ne sera jamais vide)

# Sortie
- z : le ppcm des entiers contenus dans tab0
*/

func ppcm(tab0 []uint) (z uint) {
	return z
}
