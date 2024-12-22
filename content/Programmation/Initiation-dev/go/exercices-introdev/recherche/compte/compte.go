package compte

/*
La fonction compte doit indiquer combien de fois une valeur v apparaît dans un
tableau tab.

# Entrées
- tab : un tableau d'entiers
- v : la valeur à chercher

# Sortie
- num : le nombre de fois que v apparaît dans tab

# Info
2021-2022, test 1, exercice 6
*/

func compte(tab []int, v int) (num int) {
	for i:=0; i<len(tab); i++{
		if tab[i]==v{
			num++
		}
	}
	return num
}
