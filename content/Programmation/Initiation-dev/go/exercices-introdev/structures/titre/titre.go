package titre

/*
On dispose d'une structure de données permettant 
de réprésenter des livres avec un titre, un prix 
et un nombre de page.

Étant donné un livre, on souhaite retourner 
son titre.

# Entrée
- l : un livre

# Sortie
- titre :  le titre de l

# Info
2022-2023, test3, exercice 2
*/

type livre struct {
	titre    string
	prix     float64
	numPages int
}

func obtenirTitre(l livre) (titre string) {
	return l.titre
}
