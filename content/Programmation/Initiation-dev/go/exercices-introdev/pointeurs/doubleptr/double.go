package double

import "errors"

/*
La fonction double (pour double pointeur) doit mettre l'adresse x à l'adresse y.
En d'autre terme, si on appelle cette fonction avec pour arguments &n et &ptr
alors ptr devrait ensuite contenir l'adresse de n. Dans le cas où ceci n'est pas
possible, la fonction devrait retourner l'erreur errImpossibleAssignment définie
ci-dessous.

# Entrées
- x : l'adresse d'un entier
- y : l'adresse de l'adresse d'un entier

# Sortie
- err : une erreur indiquant que la fonction n'a pas pu fonctionner

# Info
2021-2022, test3, exercice 7
*/

var errImpossibleAssignment error = errors.New("Impossible assignment")

func double(x *int, y **int) (err error) {
	if y == nil {
		err = errImpossibleAssignment
	} else {
		*y = *&x
	}
	return err
}
