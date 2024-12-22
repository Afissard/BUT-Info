package varswitch

/*
La fonction varswitch doit échanger les valeurs situées aux adresses ax et ay.
Ainsi, si on a *ax = x et *ay = y au moment d'un appel à varswitch, alors on
doit avoir *ax = y et *ay = x après cet appel. Cette fonction n'a donc pas de
sorties.

# Entrées
- ax : un pointeur vers un entier
- ay : un pointeur vers un entier

# Info
2021-2022, test 1, exercice 2
*/

func varswitch(ax, ay *int) {
	temp := *ax
	*ax = *ay
	*ay = temp
}
