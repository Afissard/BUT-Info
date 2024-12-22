package filtre

import "errors"

/*
Étant donné un nom de fichier, on souhaite obtenir une copie de ce fichier dont les lignes de longueur impaire (sans compter le caractère de fin de ligne \n) ont été retirées.

# Entrée
- fName : le nom du fichier à traiter

# Sorties
- pairs : une chaîne de caractère qui est une copie du contenu du fichier dont les lignes de longueur impaire ont été retirées
- err : nil si le fichier fName existe, errImpossible sinon

# Info
2022-2023, test3, exercice 6
*/

var errImpossible error = errors.New("Le fichier n'existe pas")

func filtrePairs(fName string) (pairs string, err error) {
	return
}
