package queue

/*
On considère la structure de file d'attente (queue) définie ci-dessous. Une queue contient un certain nombre de places (sa taille) et est définie par la dernière place dans cette queue.
Chaque place est occupée par une personne et est reliée à la place précédente dans la queue.
Ainsi, de place en place, on peut reconstituer toute la queue.
Quand une nouvelle personne s'ajoute à la queue, elle prend une nouvelle place à la fin de celle-ci.

La méthode ajout prend en paramètre une personne et l'ajoute à la fin de la queue sur laquelle elle s'applique.

# Entrées
- q : une queue déjà formée, cette queue sera modifiée par la méthode pour y ajouter une place contenant la personne qui souhaite s'ajouter à la queue
- p : la personne qui souhaite s'ajouter à cette queue

# Info
2022-2023, test3, exercice 8
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

func (q *queue) ajout(p personne) {
	nplace := place{
		occupant:  p,
		precedant: q.fin,
	}
	q.fin = &nplace
	q.taille ++
}
