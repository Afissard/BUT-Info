package pages

/*
On dispose d'une structure de données permettant de 
réprésenter des livres avec un titre, un prix et un 
nombre de page.

On souhaite trouver dans un tableau de livres un livre 
qui a un nombre de pages donné. C'est ce que doit faire 
la fonction trouver.

# Entrée
- tab : le tableau de livres
- numPages : le nombre de pages pour lequel on souhaite 
trouver un livre

# Sortie
- l : un livre provenant de tab et qui fait numPages pages, 
s'il existe (sinon, un livre quelconque)
- trouve : un booléen qui vaut true si on a trouvé un livre 
de numPages pages dans tab et false sinon

# Info
2022-2023, test3, exercice 3
*/

type livre struct {
	titre    string
	prix     float64
	numPages int
}

func trouver(tab []livre, numPages int) (l livre, trouve bool) {
	if len(tab)!=0{
		for i:=0; i<len(tab); i++ {
			if tab[i].numPages == numPages {
				return tab[i], true
			}
		}
	}

	// livre retourné si tableau vide
	scroll_of_stupidity := livre{
		titre : "Scroll of Stupidity",
		prix : -666.99,
		numPages : 0,
	}
	return scroll_of_stupidity, false
}
