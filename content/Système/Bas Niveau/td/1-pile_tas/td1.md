# TD n°1 : la pile et le tas
Ce sujet a été conçu par L. Jezequel.

Lorsqu'on exécute un programme sur un ordinateur, son code binaire, ainsi les données qu'il manipule (les constantes et variables), sont chargées dans la mémoire de travail de l'ordinateur.
En pratique, lorsqu'un programme est *lancé*, le système d'exploitation créé un processus pour l'exécuter.
Ce processus est associé à un **espace d'adressage**, qui peut être vu comme un tableau d'octets réservé à cette exécution du programme.
Cet espace d'adressage est découpé en différentes zones.
  - (Variable d'environnement)
  - La **pile** (*stack* en anglais), utilisée pour stocker les données locales des fonctions (paramètres réels et variables locales). Le contenu et l'état de la pile varie dynamiquement pendant l'exécution en fonction des appels et retours de fonction. (vas vers le bas)
  - Le **tas** (*heap* en anglais), utilisé pour stocker les données créées pendant l'exécution du programme dont la durée de vie n'est pas associée à celle d'une fonction, par exemple car elles continuent à être utilisées après la fin de la fonction qui les a créées. (vas ver le haut et donc attention à la rencontre entre le *stack* et le *heap*)
  - Le **segment de données**, utilisé pour stocker les données créées à la compilation du programme (variables globales et constantes). En pratique, cette zone est en fait séparée en plusieurs parties : `.data`, `.bss`, `.rodata`, etc.
  - La zone de **texte**, utilisée pour stocker le code binaire du programme.
## La pile
La pile (*stack)* est une zone mémoire qui fonctionne comme une pile au sens algorithmique du terme, c'est-à-dire une structure de donnée dynamique qui implémente une politique « dernier arrivé, premier sorti ». Quand une fonction est appelée, un espace appelé **cadre de la fonction** (*frame* en anglais) est empilé. Ce cadre sert notamment à stocker les paramètres réels de la fonction et les variables locales[^1]. Quand une fonction se termine et que le contrôle retourne vers le point d'appel, son cadre est dépilé. Les données qui s'y trouvaient sont alors inaccessibles. Une fois le retour effectué, le cadre au sommet de la pile est maintenant celui de la fonction appelante.

Lorsqu'une fonction est en cours d'exécution, elle a accès aux données de son cadre, aux données de la pile hors de son cadre dont elle connaît l'adresse, aux données du tas dont elle connaît l'addresse, et aux données globales (celles du segment de données). 
Elle n'a pas accès directement aux autres cadres contenus dans la pile.
Autrement dit, pour accéder aux variables locales de la fonction appelante (ou de la fonction appelante de la fonction appelante, ou ...), il est nécessaire d'utiliser des pointeurs[^2].

---

**Exercice n°1**
Soit le programme `pile.go`ci-dessous :
```go
package main

func main() {
	foo()
	bar()
}

func foo() {
	foobar()
}

func bar() {}

func foobar() {}

```
Dessiner l'état de la pile au début et à la fin de chaque appel de fonction.

---

## Le tas
Comme on vient de le voir, une donnée qui appartient au cadre d'une fonction n'existe plus (ou en tout cas n'est plus accessible) une fois que cette fonction a retourné. 
Or, dans certain cas, les données doivent survivre à la fonction où elles sont crées.
Ces données doivent être placées dans le tas.

---

**Exercice n°2**
Soit le programme `tas.go`:

```go
package main

func main() {
	a, b := 1, 2
	_ = additionBasique(a, b)
	_ = additionVersPointeur(a, b)
	_ = additionDepuisPointeurs(&a, &b)
}

//go:noinline
func additionBasique(x, y int) int {
	z := x + y
	return z
	/*
	x et y sont locaux à la fonction et vont dans la pile
	z est dans la pile, sa valeur retourné est copier dans la variable destiné à récupérer l'input
	*/
}

//go:noinline
func additionVersPointeur(x, y int) *int {
	z := x + y
	return &z
	/*
	x et y sont locaux à la fonction et vont dans la pile
	z est certe déclarer dans la fonction, mais son pointeur est retourné ver le programme d'appel, z est donc stocker dans le tas (attention c'est golang qui s'en occupe, en C ce n'est pas le cas est dans la pile sa valeur peut être réécrite)
	*/
}

//go:noinline
func additionDepuisPointeurs(x, y *int) int {
	z := *x + *y
	return z
	/*
	x et y sont externe à la fonction (on récupère leur pointeurs) et sont donc dans la pile de la fonction appelante.
	z est dans la pile, sa valeur retourné est copier dans la variable destiné à récupérer l'input
	*/
}

```

et le jeu de test associé `tas_test.go`

```go
package main

import (
	"math/rand"
	"testing"
)

func BenchmarkAdditionBasique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u, v := rand.Intn(1000), rand.Intn(1000)
		additionBasique(u, v)
	}
}

func BenchmarkAdditionVersPointeur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u, v := rand.Intn(1000), rand.Intn(1000)
		additionVersPointeur(u, v)
	}
}

func BenchmarkAdditionDepuisPointeurs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u, v := rand.Intn(1000), rand.Intn(1000)
		additionDepuisPointeurs(&u, &v)
	}
}
```

  - Lire ces programmes et expliquer leur fonctionnement.
  - Pour chaque variable de chaque fonction (en incluant les paramètres), expliquer si elle doit être stockée plutôt dans la pile, ou plutôt dans le tas.

---

**Remarque** 
Le commentaire `noinline` est une directive donnée au compilateur.  Elle lui interdit de faire une optimisation particulière qui consiste à remplacer les appels à une fonction par le code de cette fonction (dans le cas de ce programme, tous les appels seraient *inlinés* dans le code de la fonction main et l'ensemble des variables se retrouverait alors naturellement dans la pile). 

---

On rappelle qu'on peut exécuter les *benchmarks* du fichier `tas_test.go` par la commande

```sh
go test -bench=. tas.go tas_test.go
```

En ajoutant l'option `-benchmem` à cette commande on obtient, en plus des informations sur les temps d'exécution, des informations sur la gestion de la mémoire :

```bash
go test -bench=. -benchmem tas.go tas_test.go
```
Résultat :
```bash
BenchmarkAdditionBasique-16 53691757 22.69 ns/op 0 B/op 0 allocs/op BenchmarkAdditionVersPointeur-16 36474930 33.34 ns/op 8 B/op 1 allocs/op 
BenchmarkAdditionDepuisPointeurs-16 51104269 19.62 ns/op 0 B/op 0 allocs/op
```

---

 **Exercice n°3**
 Exécuter les *benchmarks* du fichier `tas_test.go` avec l'option `-benchmem`. 
 Expliquer les résultats obtenus (pour cela il ne faut pas hésiter à se documenter sur la signification des informations affichées par `go test`).
   - Quel(s) appel(s) de fonction(s) utilise(nt) le tas ? : `AdditionVersPointeur`
   - Qu'est-ce qui vous semble plus efficace en termes de temps de calcul entre l'utilisation du tas et l'utilisation de la pile ?

---

On peut obtenir plus d'informations sur les variables qui sont placées dans le tas en compilant le programme avec l'option `-gcflags -m` :

```bash
go build -gcflags -m tas.go
```

---

**Remarque** Quand on utilise la commande `go build` avec l'option `-gcflags` comme argument, on passe des options de compilation (des *flags*) au compilateur Go (`gc` signifie ici *Go compiler*).
En effet, la commande `go build` est quelque chose de plus général que simplement un compilateur.
Elle va ainsi vérifier les dépendances, gérer les modules, etc.
Et elle utilise bien sûr un compilateur pour produire un exécutable.
Ici, l'option `-m` permet d'afficher des informations sur les décisions d'optimisation prises par le compilateur, comme par exemple le choix de mettre une variable dans le tas ou dans la pile[^3].

---

**Exercice n°4** 
Consulter le point sur l'allocation de la mémoire dans la FAQ de Go : https://go.dev/doc/faq#stack_or_heap.
Utiliser la commande `go build -gcflags -m tas.go`.

  - Est-ce qu'une variable dont on récupère l'adresse dans un programme est obligatoirement stockée dans le tas ? : "In the current compilers, if a variable has its address taken, that variable is a candidate for allocation on the heap. However, a basic _escape analysis_ recognizes some cases when such variables will not live past the return from the function and can reside on the stack."
  - Est-ce que seules les variables dont on récupère l'adresse dans un programme sont susceptibles d'être stockées dans le tas ? : "if the compiler cannot prove that the variable is not referenced after the function returns, then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors"
  - Quelles sont les variables qui sont placées dans le tas ? Expliquer. : les grandes variables et les variable qui ne seront pas que utilisé localement.
voir documentation : https://go.dev/doc/faq#stack_or_heap 
---

**Remarque** Vous aurez peut-être noté que, quand on utilise `go build -gcflags -m tas.go` et `go test -bench=. -benchmem tas.go tas_test.go` ce n'est pas le même code qui est analysé : dans le premier cas c'est la fonction `main` qui est le point de départ du programme, alors que dans le deuxième ce sont les fonctions de `benchmark` situées dans le fichier `tas_test.go`.

---

**Remarque** Les espaces mémoires qui sont occupés dans le tas doivent être libérés quand ils ne sont plus utilisés pour éviter des fuites de mémoire (*memory leaks*) menant à une occupation trop importante de la mémoire et donc, lorsqu'il n'y a plus d'espace disponible, à des ralentissements importants des programmes, voir à un blocage complet de l'ordinateur.
Souvent, dans les langages destinés à la programmation des systèmes (comme C ou C++), le programmeur doit lui même réserver et libérer la mémoire.
Sur ce point, Go diffère donc de ces prédecesseurs de façon importante : un ramasse-miettes (*garbage collector*) se charge de temps en temps de libérer la mémoire qui n'est plus utile (les espaces du tas qui ne seront plus jamais utilisés dans le reste du programme).
Ceci simplifie la vie au programmeur mais ajoute un surcoût à l'utilisation du tas et ne permet pas au programmeur de maîtriser entièrement l'exécution.
Le fonctionnement du ramasse-miettes est assez complexe et nous ne l'aborderons pas dans ce cours.
On peut souligner que Rust, un autre langage moderne destiné à la programmation des systèmes, utilise une approche alternative : la construction du langage est telle que le compilateur est capable de déterminer par analyse à quel moment une donnée du tas devient inutile.
Il peut alors insérer un code de libération de la mémoire correspondant à celui qu'aurait pu obtenir manuellement un programmeur.
On obtient ainsi des performances et un niveau de maîtrise similaires à ce qu'on peut faire avec C ou C++, tout en profitant d'une gestion complètement automatisée de la mémoire.
Le prix à payer est un langage plus complexe que Go. 

---

## Le segment de données

Le segment de données contient les données qui peuvent être déterminées à la compilation (donc dont on connaît la taille et dont on sait qu'elles existeront dans toute exécution sans avoir a exécuter le programme). 


---

**Exercice n°5**
Soit le programme `segment_donnees.go` :

```go
package main

import "log"

const a int = 3

var b int

func main() {
	var c int = 1
	for i := 0; i < a; i++ {
		b += c
		c++
	}
	log.Print("a vaut ", a, ", b vaut ", b, " et c vaut ", c)
}

```

Tester la commande 

```bash
go build -gcflags='-S -S' segment_donnees.go
```

Attention, `-S`  apparaît bien deux fois dans cette commande (en ne le mettant qu'une seule fois, on ne verrait que le code et pas les données).

Repérer la partie données et la partie code dans le résultat produit. 
Dans le code, SP signifie *stack pointer* et SB signifie *static base pointer*.
Analyser le code. 
 
  - La variable `b` est-elle allouée statiquement ? : "b" est dans BSS, on a réservé de place dans le segment de donnée pour "b" (donc oui c'est statique avec RODATA -> read only)
  - La variable `c` est-elle allouée statiquement ? : non ce n'est pas le cas
  - Que remarque-t-on concernant la constante `a` ?
  - Que remarque-t-on concernant les chaînes de caractères utilisées dans le programme ? : c'est ce qui prend le plus de place dans le programme

 ---

## Références

1. [https://povilasv.me/go-memory-management/](https://povilasv.me/go-memory-management/) : des explications
   générales sur la gestion de la mémoire, mises en lien avec le cas
   particulier de Go.
2. [https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d](https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d) :
   des explications sur les choix faits par Go d'allouer la mémoire
   dans la pile ou dans le tas.
3. [https://go.dev/doc/asm](https://go.dev/doc/asm) : quelques informations sur le langage
   assembleur utilisé par le compilateur Go.
4. [https://go.dev/doc/faq](https://go.dev/doc/faq) : la FAQ officielle du langage Go.

---

[^1]: En fait c'est un peu plus compliqué : les variables locales et, dans certaines architectures, les paramètres, peuvent parfois être stockés dans les registres du processeur.
    
[^2]: On reparlera plus en détails des pointeurs dans la suite de ce cours.

[^3]: On peut obtenir plus d'informations en utilisant `go help build` qui nous explique que `-gcflags` passe des arguments à `go tool compile`.
En faisant `go help tool` on apprend qu'on peut obtenir des explication plus précise en utilisant `go doc cmd/compile`.
Cette dernière commande nous donne la signification exacte de `-m` et nous indique aussi toutes les autres options possibles.
On y trouve aussi des infos sur `go:noinline` dont on a parlé plus tôt.
Et bien sûr, une version en ligne de cette page de documentation existe: https://pkg.go.dev/cmd/compile.
