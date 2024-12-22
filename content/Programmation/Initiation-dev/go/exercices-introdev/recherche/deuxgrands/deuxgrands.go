package deuxGrands

/*
La fonction lesDeuxPlusGrands retourne les deux plus grands entiers présents dans un tableau

# Entrée
- t : un tableau d'entiers qui en contient au moins 2

# Sorties
- v1 : le plus grand entier dans t
- v2 : le deuxième plus grand entier dans t

# Info
2022-2023, test 1, exercice 8
*/

func lesDeuxPlusGrands(t []int) (v1, v2 int) {
	for i:=0; i<len(t); i++{
		if t[i]>=v1{
			v2 = v1
			v1 = t[i]
		} else if t[i]>v2 {
			v2 = t[i]
		}
	}
	return v1, v2
}
