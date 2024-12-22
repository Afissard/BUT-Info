package deuxpetits

/*
La fonction lesDeuxPlusPetits retourne les deux plus petits entiers présents dans un tableau

# Entrée
- t : un tableau d'entier qui en contient au moins 2

# Sorties
- v1 : le plus petit entier dans t
- v2 : le deuxième plus petit entier dans t

# Info
2022-2023, test 4, exercice 3
*/

func lesDeuxPlusPetits(t []int) (v1, v2 int) {
	const MAXINT64 = 9223372036854775807
	v1, v2 = MAXINT64, MAXINT64 // pour ne pas initialiser les var en dessous des val du tab
	for i:=0; i<len(t); i++{
		if t[i]<=v1{
			v2 = v1
			v1 = t[i]
		} else if t[i]<v2 {
			v2 = t[i]
		}
	}
	return v1, v2
}
