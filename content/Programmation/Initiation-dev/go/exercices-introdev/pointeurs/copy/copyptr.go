package copyptr

/*
La fonction copyptr doit copier une valeur à chacune des adresses indiquées
dans un tableau.

# Entrées
- x : la valeur à copier
- tab : un tableau de pointeurs vers des entiers, à la fin de la fonction chacun
        de ces entiers doit avoir la même valeur que x

# Info
2021-2022, test2, exercice 1
*/

func copyptr(x int, tab []*int) {
        for i := 0; i < len(tab); i++ {
                *tab[i] = x
        }
}
