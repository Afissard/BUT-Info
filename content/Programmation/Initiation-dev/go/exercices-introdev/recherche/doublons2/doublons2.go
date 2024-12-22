package doublons2


/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient
exactement une fois chaque nombre compris entre 1 et n. On voudrait vérifier
cela. C'est le travail de la fonction doublons.

Coder la fonction doublons de manière à ne parcourir le tableau tab qu'une seule
fois rapportera plus de points.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui vaut true si tab contient bien exactement une
fois chaque entier entre 1 et len(tab) et false sinon

# Info
2021-2022, test 1, exercice 8
*/

func doublons(tab []int) (ok bool) {
	ok = true

	// pré-test
	if ((tab == nil) || (len(tab)==0)) { 
		// si tab vide (=> initle de chercher)
		return ok
	} else if len(tab) == 1 {
		// si tab avec un seul élément (=> initle de chercher)
		if tab[0] != 1{
			ok = false
		}
		return ok
	}

	// test
	double := []int{tab[0]}
	for i:=1; i<len(tab); i++{
		if (tab[i] < 1) || (tab[i] > len(tab)) { // si hors de l'ensemble
			ok = false
			return ok
		} else { // test doublons
			for j :=0; j<len(double); j++{ 
				// On ne reparcours pas le même tab et celui-ci est plus cours (au début)
				if tab[i] == double[j] {
					ok = false
					return ok
				}
			}
			double = append(double, tab[i])
		}
	}
	return ok
}


/*
Meilleur solution (ne parcours pas dutout un second tableau) :
Créer un tab de bool, de taille len(tab), tout est initialiser à false
Si tabbool[tab[i]] == False -> tabbool[tab[i]] = True
Si tabbool[tab[i]] == True -> Doublons -> retourne ok = Flase 

On ne reparcours pas le second tableau puisque l'on sait où regarder pour le test, 
mais ne marche que parce qu'il ne peut y avoir de valeurs dans tab entre 1 et len(tab)
*/