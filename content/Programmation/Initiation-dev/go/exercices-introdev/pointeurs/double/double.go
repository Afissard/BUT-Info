package double

/*
La fonction double doit doubler la valeur située à l'adresse donnée par son
argument. Ainsi, si x == 10 avant un appel à double(&x), alors on doit avoir
x == 20 après cet appel. Cette fonction n'a donc pas de sorties.

# Entrée
- x : un entier

# Info
2021-2022, test 1, exercice 1
*/

func double(x *int) {
	*x = *x*2
}
