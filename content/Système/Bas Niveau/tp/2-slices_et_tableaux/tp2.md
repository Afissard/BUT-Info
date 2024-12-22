# TP2 : Pointeurs, tableaux, *slices*

Les objectifs de cette séance sont :

  1. Apprendre à utiliser la fonction d'exploration de la mémoire de GDB.
  2. Découvrir la représentation en mémoire des tableaux (*array*) et des tranches (*slices*).
  3. Observer l'implantation bas niveau du concept de pointeur.

# Examiner la mémoire avec GDB

Pendant une session GDB, la commande `examine` (raccourci `x`) permet d'examiner le contenu de la mémoire.
Elle s'utilise sous la forme suivante : `x /fmt exp`, où `fmt` décrit la taille de la zone mémoire à examiner et le format de présentation souhaité, et `exp` est une expression dont l'évaluation correspond à l'adresse de début de la zone mémoire à examiner.

Le **format** est composé de 3 parties :

- un entier, indiquant le nombre de bloc de la zone ;
- un caractère indiquant la taille des blocs souhaités : `b` (pour *byte*, 1 octet), `w` (pour *word*, 4 octets) ou `g` (pour *giant word*, 8 octets) ;
- un caractère indiquant le format du contenu des blocs : `a` (pour ascii), `c` (pour caractère), `d` (pour décimal), `x` (pour hexadécimal), `s` (pour *string*), ou encore `i` (pour instruction).

Ainsi, le format `/8gx` demandera l'affichage de 8 blocs de 8 octets en hexadécimal, et `/s` demandera l'affichage d'une chaîne de caractère (au sens du langage C : une séquence de caractères terminée par un octet nul).

L'expression indiquant l'**adresse de début** peut être :

- la valeur d'un registre, donné par son nom précédé du caractère `$` ;
- une adresse, donnée sous la forme d'un nombre entier ;
- une variable du programme de type pointeur ;
- une expression permettant de calculer une adresse à partir des éléments précédents et des opérateurs de l'arithmétique des pointeurs. 

Ainsi, la commande `x /2gd slice.array` affichera sous forme décimale le contenu de 2 blocs de 8 octets à partir de l'adresse pointée par la variable `slice.array` (le tableau contenant les données associés à la slice `slice`), tandis que la commande `x /20gx $rsp` affichera sous forme hexadécimale le contenu des 20 blocs de 8 octets situés au sommet de la pile.

# Exercice n°1

Soit le programme ci-dessous (fichier `src/slicesandarrays1/main.go`) : 

```go
package main

import (
	"fmt"
)

func print(t []int) {
	fmt.Print("{")
	for i := 0; i < len(t)-1; i++ {
		fmt.Print(t[i], ", ")
	}
	fmt.Printf("%d}\n", t[len(t)-1])
}

func swap(v1 *int, v2 *int) (r int) {
	tmp := *v1
	*v1 = *v2
	*v2 = tmp
	r = *v1 + *v2
	return
}

func main() {
	var w = [8]int{2, 3, 5, 7, 11, 13, 17, 19}

	x := w

	y := w[:]

	z := &x

	for cpt := 0; cpt < len(y); cpt++ {
		_ = swap(&((*z)[cpt]), &(y[(cpt+2)%len(y)]))
	}

	print(w[:])
	print(x[:])
	print(y)
	print((*z)[:])
}
```

Construire le programme en désactivant l'intégration en ligne et l'optimisation :

```bash
go build -gcflags "-l -N"
```

## Questions

1. Décrire la représentation en mémoire d'un tableau et celle d'un *slice*. Rappeler les différences, du point de vue du langage de programmation Go, entre un tableau et un slice.
   Un tableau est une *array* alors qu'un slice est une structure contenant un tableau ainsi que des méthode associé, la taille du tableau est notamment mutable.

2. Donner le type des variables `w`, `x`, `y` et `z`.
	- w est un tableau
	- x est un tableau
	- y est un slice
	- z est un pointeur vers x

3. Faire afficher la représentation en mémoire de `w` ? Même question pour `x`, `y` et `z`.

4. Mettre un point d'arrêt à la ligne 32 du fichier et avancer l'exécution jusqu'à ce point. Faire afficher la pile et identifier les différents objets qu'elle contient. 

5. Mettre un point d'arrêt à la ligne 16 du fichier. S'y rendre et faire afficher les 40 mots au sommet de la pile. Identifier les différents objets qu'elle contient.

6. Mettre un point d'arrêt ligne 19. S'y rendre puis exécuter ce code instruction par instruction, en expliquant le rôle et l'effet de chaque instruction.

7. Continuer l'exécution jusqu'à retourner dans `main.main`. Afficher à nouveau la pile et identifier les objets qu'elle contient.

8. Construire le code en réactivant les optimisations et reprendre la question 6. 

# Exercice n°2

Soit le programme ci-dessous (fichier `src/slicesandarrays2/main.go`) : 

```go
package main

import "fmt"

func print(t []int) {
	fmt.Print("{")
	for i := 0; i < len(t)-1; i++ {
		fmt.Print(t[i], ", ")
	}
	fmt.Printf("%d}\n", t[len(t)-1])
}

func add_and_append(slin []int) (slout []int) {
	l := len(slin)
	slout = append(slin, l)
	for idx := 0; idx < len(slout); idx++ {
		slout[idx] = slout[idx] + 1
	}
	return slout
}

func main() {

	var sli1, sli2, sli3, sli4, sli5 []int

	sli1 = add_and_append(sli1)
	sli2 = add_and_append(sli1)
	sli3 = add_and_append(sli2)
	sli4 = add_and_append(sli3)
	sli5 = add_and_append(sli4)

	fmt.Println("---sli1---")
	print(sli1)
	fmt.Println("----------")

	fmt.Println("---sli2---")
	print(sli2)
	fmt.Println("----------")
	fmt.Println("---sli3---")
	print(sli3)
	fmt.Println("----------")
	fmt.Println("---sli4---")
	print(sli4)
	fmt.Println("----------")
	fmt.Println("---sli5---")
	print(sli5)
	fmt.Println("----------")
}
```
Construire le programme en désactivant l'intégration en ligne et l'optimisation :

```bash
go build -gcflags "-l -N"
```

## Questions

1. Sans exécuter le programme, prédire sa sortie. Ensuite, exécuter le programme. La prédiction est-elle correcte ? Oui (je me suis spoil)

2. En cas de prédiction incorrecte, utiliser gdb pour comprendre ce qui se passe. En cas de prédiction correcte, utiliser quand même gdb pour vérifier que le raisonnement utilisé pour formuler la prédiction est conforme au comportement du programme. Décrire l'état de la mémoire à la ligne 30.

# Exercice n°3

Soit le programme ci-dessous (fichier `src/strings/main.go`) : 

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "All work and no play makes Jack a dull boy"
	str2 := str1
	str3 := str1[:]
	str4 := strings.Clone(str1)

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4)

}
```

Construire le programme en désactivant l'intégration en ligne et l'optimisation :

```bash
go build -gcflags "-l -N"
```

## Questions

1. En étudiant l'exécution de ce programme avec GDB, expliquer comment sont implantées les objets de type `string` en Go. Identifier le code d'initialisation des variables `str1`, `str2`, `str3`, et `str4`. Représenter l'implantation mémoire de ces variables sur un schéma.

2. Construire le programme en activant les optimisations puis l'étudier. Identifier comment les chaînes sont passées à la fonction d'affichage.

