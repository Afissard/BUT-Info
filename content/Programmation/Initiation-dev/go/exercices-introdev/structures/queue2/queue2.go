package queue

/*
On considère la structure de file d'attente (queue) définie ci-dessous. Une queue contient un certain nombre de places (sa taille) et est définie par la dernière place dans cette queue.
Chaque place est occupée par une personne et est reliée à la place précédente dans la queue.
Ainsi, de place en place, on peut reconstituer toute la queue.
Quand une nouvelle personne s'ajoute à la queue, elle prend une nouvelle place à la fin de celle-ci.

La méthode premier trouve la personne qui se trouve à la première place de la queue sur laquelle elle s'applique.

Attention : cette méthode ne doit pas modifier la queue sur laquelle elle s'applique.

# Entrée
- q : une queue de taille 1 ou plus

# Sortie
- p : la personne qui se trouve en première place de la queue q

# Info
2022-2023, test3, exercice 9
*/

type queue struct {
	fin    *place
	taille uint
}

type place struct {
	occupant  personne
	precedant *place
}

type personne struct {
	nom, prenom string
	age         uint
}

func (q *queue) premier() (p personne) {
	ptr := q.fin
	for i:=0; uint(i)<q.taille; i++{
		if ptr.precedant == nil{
			return ptr.occupant
		} else {
			ptr = ptr.precedant
		}
	}
	return
}
