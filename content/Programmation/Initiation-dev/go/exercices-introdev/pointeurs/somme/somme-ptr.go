package somme

/*
On considère un ensemble d'adresses d'entiers (des pointeurs) 
stockées dans un tableau. On veut calculer la somme des entiers 
stockés à ces adresses et la stocker à une autre adresse donnée 
en paramètre.

# Entrées
- t, un tableau de pointeurs vers des entiers
- res, un pointeur vers un entier, cet entier doit, après 
l'appel à la fonction, contenir la somme des entiers situés 
aux adresses contenues dans t

# Info
2022-2023, test 2, exercice 2
*/

func somme(t []*int, res *int) {
	*res = 0
	for i := 0; i<len(t); i++ {
		*res = *res + *t[i]
	}
}
