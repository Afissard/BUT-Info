package stevej

import (
	"errors"
)

var errPasUneDate error = errors.New("la date indiquée n'est pas valide")

/*
Steve Jobs est né le 24 février 1955. La fonction ecart devra indiquer le nombre
de jours de différence entre une date (à partir du 1er janvier 1900 jusqu'au 31
décembre 2021) et la date de naissance de Steve Jobs. L'écart en jours sera toujours
un nombre positif ou nul (on ne se préoccupera pas de savoir si la date est antérieure
ou postérieure à la naissance de steve Jobs).

# Entrées
- j : un numéro de jour
- m : un numéro de mois
- a : un numéro d'année

# Sorties
- ej : l'écart en jours entre j/m/a et le 24/2/1955
- err : une erreur qui vaudra errPasUneDate quand j/m/a n'est pas une date valide et nil sinon

# Exemple
ecart(17, 4, 1986) = 11375
*/
func ecart(j, m, a uint) (ej uint, err error) {
	// check validité
	if j>31 || m>12 || (1900<a)&& (a<2021) {
		err = errPasUneDate
		return ej, err
	}

	// calculs
	const J, M, A = 24, 02, 1955 // date de naissance de Steve Jobs
	max_j := []uint{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	// Année bisextile
	if a%4 == 0 {
		max_j[1] = 28
	}
	ans := uint(A)
	for ans != a {
		// calcule du nombre de jour de l'année "ans"
		mois := uint(M)
		for mois != m {
			// calcule du nombre de jour du mois "mois"
			/*
			Au moi du plus tard : bonne merde
			determiné nombre de jour avec max_j, faire attention si mois de départ, alors retiré des jours, pareils pour mois
			*/

			// incrémente la boucle "while"
			if mois > m {
				mois --
			} else if mois < m {
				mois ++
			}
		}

		// incrémente la boucle "while"
		if ans > a {
			ans --
		} else if ans < a {
			ans ++
		}
	}


	return ej, err
}
