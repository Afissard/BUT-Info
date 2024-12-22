package combinaison

/*
Étant donné un entier n, on souhaite trouver toutes les combinaisons d'entiers (c'est-à-dire des ensemble avec répétitions) compris entre 1 et n (inclus) dont la somme vaut n. Ainsi, par exemple, pour n = 3 on doit trouver les combinaisons {1, 1, 1}, {1, 2} et {3}.

Il est possible d'écrire pour calculer les combinaisons pour un entier n une fonction récursive sur un principe proche du suivant :
* Pour chaque i entre 1 et n
- calculer les combinaisons pour n-i
- ajouter à chacune d'entre elles la valeur i
* Retourner toutes les combinaisons obtenus

En le faisant, vous remarquerez que vous construisez trop de combinaisons (il y a des doublons). Pour éviter cela, vous pouvez utiliser une borne indiquant la valeur maximum à considérer (qui sera donc le minimum entre n et cette borne), qui sera passée dans vos appels récursifs.

Le plus simple est d'écrire une fonction récursive, mais les boucles for sont autorisées.

# Entrées
- n, un entier
- res, un tableau de tableaux d'entiers, chacun de ses tableaux représente une combinaison d'entiers compris entre 1 et n et dont la somme vaut n. L'ordre des combinaisons n'a pas d'importance.

# Info
2022-2023, test 2, exercice 8
*/

func combinaisons(n int) (res [][]int) {
	return res
}
