package conway

/*
La suite de Conway (ou suite audioactive) est définie de la manière suivante :
- le premier terme (u(0)) est 1,
- le terme u(n) s'obtient en "prononçant" les chiffres du terme précédant et leur
  nombre d'apparitions, dans l'ordre.
Ainsi, le terme u(1) vaut 11 (prononcé "un 1"), puis u(2) vaut 21 (prononcé "deux 1"),
on a ensuite u(3) = 1211 ("un 2, un 1"), u(4) = 111221 ("un 1, un 2, deux 1"),
u(5) = 312211 ("trois 1, deux 2, un 1"), et ains de suite.

Pour simplifier, on représentera les nombres par des tableaux de chiffres, ainsi
1 sera représenté par [1], 11 par [1 1], 21 par [2 1], 1211 par [1 2 1 1], etc.

La fonction conway doit calculer les termes de la suite de Conway, représentés
par des tableaux d'entiers.

La fonction conway doit être récursive. Une fonction qui n'est pas récursive ne
rapportera pas de points.

# Entrée
- n : le numéro du terme à calculer

# Sortie
- un : un tableau d'entiers représentant les chiffres du n-ième terme de la suite
       de conway.

# Info
2021-2022, test2, exercice 8
*/

func conway(n int) (un []int) {
	return un
}
