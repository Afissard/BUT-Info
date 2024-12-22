package produit

/*
On considère un ensemble d'adresses d'entiers (des pointeurs) stockées dans un tableau. 
On veut calculer le produit des entiers stockés à ces adresses et le stocker à une autre 
adresse donnée en paramètre.

# Entrées
- t, un tableau de pointeurs vers des entiers
- res, un pointeur vers un entier, cet entier doit, après l'appel à la fonction, 
contenir le produit des entiers situés aux adresses contenues dans t

# Info
2022-2023, test 4, exercice 1
*/

func produit(t []*int, res *int) {
	*res = 1
	for i := 0; i<len(t); i++ {
		*res = *res * *t[i]
	}
}
