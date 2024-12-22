package huitreines2

/*
Le problème des huit reines consiste à placer, sur un échiquier (un tableau de 8 cases par 8 cases), 8 reines, de telle sorte qu'aucune d'entre-elles ne soit en position d'en manger une autre (c'est à dire de telle sorte qu'il n'y ait pas deux reines sur la même ligne, la même colonne ou la même diagonale de l'échiquier).

Ce problème se généralise en problème des n reines. On cherche alors à placer de la même façon n reines sur un échiquier de n cases par n cases.

Quand on analyse le problème, on se rend rapidement compte que chaque reine doit obligatoirement être sur une ligne différente des autres : il y a exactement une reine par ligne de l'échiquier. Partant de là, on peut procéder récursivement pour résoudre le problème :
1. placer une reine sur la première case de la première ligne (l) sans reine
2. si cette reine n'est pas mangée par une autre résoudre le même problème pour n-1 reines sur un échiquier ou les lignes 0 à l sont occupées
3. s'il y a une solution, c'est gagné
4. sinon (si la reine est mangée ou si le problème à n-1 reines n'a pas de solution), placer la reine sur la case suivante de la ligne (l) et recommencer à l'étape 2.

L'objectif de cet exercice est d'écrire une fonction huitreines qui trouve toutes les solutions au problème des n reines pour n donné.

# Entrée
- n : le nombre de reines à placer sur un échiquier de n cases par n cases

# Sortie
- plateaux : un tableau de tableaux de tableaux d'entiers, chacun de taille n*n et contenant des 0 là où il n'y a pas de reine et des 1 là où il y a des reines dans une solution du problème des n reines (si elle existe), sans doublons

# Exemple
huitreines(1) = [[[1]]]
huitreines(2) = []
*/

func huitreines(n int) (plateaux [][][]int) {
	return plateaux
}
