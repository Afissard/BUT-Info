package longueur

/*
On considère la structure de chaîne définie ci-dessous, très proche de la structure
de liste chaînée qui a été vue en cours. La fonction longeur doit donner le nombre
d'éléments dans une chaîne (le dernier élément de la chaîne a le champ next qui
vaut nil, comme pour les listes).

Toutes les chaînes qui seront passées à longueur seront bien formées, c'est-à-dire
qu'en suivant les pointeurs next des éléménts on ne retombe jamais deux fois sur
le même élément (donc on est certain qu'en allant de next en next on finira par
atteindre un élément pour lequel next vaut nil).

# Entrée
- c : une chaine

# Sortie
- l : la longueur de c

# Info
2021-2022, test3, exercice 8
*/

type chaine struct {
	debut *element
}

type element struct {
	next *element
}

func longueur(c chaine) (l int) {
	return l
}
