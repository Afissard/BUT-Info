package monnaie2

import (
	"errors"
)

var errPasAssezPaye error = errors.New("le montant payé est inférieur au montant de l'achat, impossible de rendre la monnaie")

/*
Étant donnés un montant d'achat en euros et un montant payé en euros,
la fonction rendreMonnaie retourne une liste de pièces et billets qui
permet de rendre au client la somme qu'il faut.

# Entrées
- eurosAchat : la partie entière du montant d'achat en euros
- centimesAchat : la partie décimale du montant d'achat (donc les centimes)
- eurosPayes : la partie entière du montant payé en euros
- centimesPayes : la partie décimale du montant payé (donc les centimes)

# Sorties
- eurosRendus : un tableau contenant les valeurs des pièces et billets en euros à rendre, ces valeurs doivent être prises parmi 1, 2, 5, 10, 20, 50, 100, 200 et 500
- centimesRendus : un tableau contenant les valeurs des pièces en centimes à rendre, ces valeurs doivent être prises parmi 1, 2, 5, 10, 20 et 50
- err : une erreur qui doit être errPasAssezPaye si le montant payé est inférieur au montant de l'achat et nil sinon

# Exemple
rendreMonnaie(12, 20, 15, 0) = [1, 1], [50, 20, 10] (ce n'est pas la seule possibilité pour ce rendu)
*/
func rendreMonnaie(eurosAchat, centimesAchat, eurosPayes, centimesPayes int) (eurosRendus, centimesRendus []int, err error) {

	return eurosRendus, centimesRendus, err
}
