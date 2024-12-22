package piecesjaunes

import "errors"

var errImpossible error = errors.New("impossible de constituer cette somme avec ces pièces")

/*
Les pièces et billets que nous utilisons dans la vie de tous les jours ont des valeurs qui ont été choisies pour qu'il soit simple de rendre la monnaie (il suffit de toujours donner la pièce ou le billet de plus grande valeur possible, jusqu'à avoir rendu la somme voulue). Dans cet exercice, on imagine que ce n'est pas le cas : on dispose d'un ensemble de valeurs de pièces quelconque, on a accès à autant de chacune de ces pièces qu'on le souhaite, et on cherche à constituer une somme avec le moins de pièces possible.

La fonction meilleurDecomposition doit trouver la meilleur décomposition d'une somme d'argent (celle qui utilise le moins de pièces) en utilisant des pièces dont les valeurs possibles sont indiquées en paramètre de la fonction.

# Entrées
- somme : la somme d'argent à constituer avec les pièces
- valeursPieces : un tableau des valeurs de pièces existantes

# Sorties
- pieces : un ensemble de pièces dont la somme vaut somme
- err : une erreur qui doit valoir errImpossible si la somme ne peut pas être réalisée à partir des valeurs de pièces considérées.

# Exemples
meilleurDecomposition(15, []int{7, 8, 9}) = [8 7]
meilleurDecomposition(15, []int{2, 4, 6}) = errImpossible
*/
func meilleurDecomposition(somme int, valeursPieces []int) (pieces []int, err error) {
	return pieces, err
}
