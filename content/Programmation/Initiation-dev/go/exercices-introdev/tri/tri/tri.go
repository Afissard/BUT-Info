package tri

/*
La fonction tri doit trier un tableau d'entiers du plus petit au plus grand.
Cette fonction ne doit pas modifier le tableau donné en entrée.

# Entrée
- tinit : un tableau d'entiers qui ne doit pas être modifié.

# Sortie
- tfin : un tableau contenant les mêmes entiers que tinit mais triés du plus
         petit au plus grand.

# Info
2021-2022, test2, exercice 6
*/
import "fmt"

func tri(tinit []int) (tfin []int) {
	
	for i:=0; i<len(tinit); i++{
		min := i
		for j:=i+1; j<len(tinit); j++{
			if tinit[j] < tinit[min]{
				min = j 
			}
		}
		tfin = append(tfin, min)
		fmt.Println(tinit, "->", tfin)
	}
	return tfin
}

