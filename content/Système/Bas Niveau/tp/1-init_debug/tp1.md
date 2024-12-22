# TP n°1 : initiation à l'utilisation de GDB

[GDB](https://sourceware.org/gdb) est l'acronyme de GNU DeBugger.
Un *debugger* (on trouve parfois la françisation débogueur) est un outil permettant d'exécuter un programme pas à pas, et d'inspecter son espace d'état après chaque pas.
GDB est un logiciel libre et multi-plateforme, qui supporte de nombreux langages de programmation (dont Go).
C'est un outil puissant, qui permet de travailler aussi bien au niveau du code source du programme exprimé dans un langage de haut niveau, qu'au niveau du code machine.
Nous allons utiliser son interface de bas niveau, qui fonctionne en mode texte et utilise un langage de commande spécifique. 
Cette interface permet d'accéder à l'intégralité des fonctionnalités de GDB, contrairement à la plupart des interfaces de plus haut niveau, qui offrent en revanche l'accès simplifié aux fonctions les plus courantes et les plus utilisées[^1].

Pour cette initiation, nous allons étudier le programme suivant :

```go
package main

import "fmt"

func main() {
	var t []int = []int{12, 34, 9, 27, 14}
	for i := 0; i < len(t)-1; i++ {
		minpos := minimumPosition(t[i:])
		t[i], t[i+minpos] = t[i+minpos], t[i]
	}
	fmt.Println(t)
}

func minimumPosition(t []int) (pos int) {
	for i := 1; i < len(t); i++ {
		if t[i] < t[pos] {
			pos = i
		}
	}
	return pos
}
```

## Construire le programme et lancer GDB

Lors de la compilation du code source et de la génération du code machine, le compilateur Go fait appel à plusieurs optimisations.
Ces optimisations améliorent les performances du programme, mais, en contrepartie, elles augmentent la distance entre le code source et le code machine, ce qui complexifie les sessions de débogage.
Pour éviter ce problème, nous allons compiler notre programme en désactivant les optimisations. 
Pour cela, nous utiliserons la commande suivante (on suppose que le module Go a été initialisé correctement) :

```bash
go build -gcflags '-N -l' 
```

On peut ensuite lancer GDB à l'aide de la commande suivante (de préférence dans un terminal passé préalablement en mode plein écran)

```bash
gdb -tui ./findmin
```
(ici, `findmin` est le nom de l'exécutable produit par la commande précédente).

Vous vous retrouvez alors avec une fenêtre scindée en deux :

  - en haut, la partie affichage du programme, pour l'instant vide ;
  - en bas, la console de commande.

Après avoir passé le message d'accueil, le code source du programme se charge dans la fenêtre haute et l'invite de commande s'affichage : `(gdb)`.

Le **focus** est actuellement sur la fenêtre de programme. Pour le passer sur la console de commande, utiliser la commande `ctrl+x o`.
Répéter cette commande permet de déplacer le focus d'une fenêtre à l'autre.

Pour ce TP, on veut afficher simultanément le code source et le code machine. 
Pour cela, circulez dans les modes de présentation de la machine en entrant deux fois la commande `layout next` dans la console. 
Vous devez normalement obtenir quelque chose qui ressemble à la capture d'écran ci-dessous.

![Capture d'écran montrant l'état souhaité de GDB](gdb-1.png)


## Lancer l'exécution et avancer pas à pas

La commande `start` permet de démarrer l'exécution du programme.
La ligne de code qui sera exécutée lors du prochain pas d'exécution est en surbrillance.
Pour avancer d'un pas d'exécution, plusieurs commandes sont possibles :

  - `stepi` : avance d'une ligne dans le code machine ;
  - `step` : avance d'une ligne dans le code source ;
  - `nexti` : comme `stepi` mais ne suit pas les appels de fonction ;
  - `next`: comme `step` mais ne suit pas les appels de fonction.

Pour faciliter les choses, chaque commande dispose d'une version raccourcie : `si`, `ni`, `s`, et `n`.
Par ailleurs, utiliser entrée dans la console avec une ligne de commande vide rappelle la dernière ligne de commande.

En testant ces commandes, on s'aperçoit qu'une ligne de code source correspond généralement à plusieurs lignes de code machine.

On s'aperçoit également qu'on se retrouve rapidement dans des fonctions internes du *runtime* de Go dont le code est automatiquement ajouté à celui du programme lors de la construction du binaire.
Pour avancer l'exécution jusqu'à la fin de la fonction courante, on utilise la commande `finish`.

---

**Exercice n°1**
Tester les différentes commandes d'exécution pas-à-pas.
À combien d'instructions du code machine correspond la ligne 9 du programme source ? 
**Réponse**
```
n 6
ni 36 // correspond à la dernière instruction machine
```
La ligne 09 du code source correspond à 36 instructions machine. 

---


## Points d'arrêts

Parfois on a besoin d'avancer rapidement dans l'exécution du programme pour se rendre à un endroit précis que l'on souhaite examiner plus en détail en avançant pas à pas (par exemple la ligne 9 du programme source).
Pour cela, les debuggers utilisent la notion de **point d'arrêt** (*breakpoint*).

Pour ajouter un point d'arrêt au programme, on utilise la commande `break <loc>` (raccourci `b`), où `<loc>` désigne un point du programme.
Plusieurs syntaxes sont possibles, par exemple :

  - le nom d'une fonction, comme dans `b main.minimumPosition`;
  - une ligne du code source actuellement affiché, comme dans `b 11` ;
  - une ligne du code source d'un fichier particulier, comme dans `b main.go:8` ;
  - une ligne du code machine, désignée par son adresse, comme dans `b *0x47e269`.

La commande `b` sans argument ajoute un point d'arrêt à la position courante.

La commande `info b` donne la liste des points d'arrêts actuellement installés.
Dans cette liste, chaque point à un numéro.
Ce numéro peut ensuite être utilisé pour supprimer le point d'arrêt grâce à la commande `delete <num>` (raccourci `d`).

Une fois les points d'arrêt installés, on peut s'en servir comme des étapes dans l'exécution du programme auxquelles on souhaite s'arrêter.
Pour cela on utilisera la commande `continue` (raccourci `c`), qui permet d'avancer l'exécution jusqu'au prochain point d'arrêt.

D'autres commandes plus avancées sont disponibles pour manipuler les points d'arrêts.
Elles ne soient pas nécessaire pour réaliser ce TP.
La personne intéressée pourra cependant en apprendre davantage en utilisant l'aide en ligne de GDB, accessible par la commande `help breakpoints`.


---

**Exercice n°2**
Tester les différentes façons de positionner des points d'arrêt.
Combien de fois le programme passe-t-il par la ligne 14 ?
**Réponse**
Le programme passe 11 fois.

---

## Afficher les variables 

Pour comprendre le comportement d'un programme et débusquer les bugs, l'exécution pas-à-pas n'est pas suffisante : il faut également pouvoir prendre connaissance du contenu des variables. 
Pour cela, on utilisera la commande `print` dont le synopsis est `print <fmt> <exp>` où `<fmt>` (optionnel) décrit le format souhaité, et `<exp>`  est une expression.

Le format est décrit sous la forme `/<f>` où `<f>` indique le format, par exemple `x` pour he**x**adécimal, `s` pour chaîne de caracère (*string*), *c* pour caractère, ...

L'expression la plus simple est le nom d'une variable, mais il est possible de construire des expressions plus complexes en utilisant des opérateurs élémentaires  (arithmétiques ou autres).
À noter que ces opérateurs ne sont pas ceux du langage Go, mais ceux du langage C.
En ce qui concerne cette ressource, deux opérateurs seront particulièrement intéressants :

  - `*<ptr>`, qui permet d'afficher le contenu de la mémoire pointée par l'adresse `<ptr>` ;
  - `&var`, qui permet d'afficher l'adresse en mémoire d'une variable.

---

**Exercice n°3**
  - Tester l'affichage de différentes expressions, dans différents formats.
  - Se rendre ligne 7, puis afficher l'adresse et le contenu de la variable `t`.
  - Se rendre ensuite ligne 15, puis afficher l'adresse et le contenu de la variable `t` (il s'agit ici du paramètre passé à la fonction `minimumPosition).
  - Dessiner l'implantation mémoire de ces deux variables lorsque l'exécution est à la ligne 15 pour la première fois.
**Réponse (problème dû à WSL)**
1. contenu de la variable t à la ligne 7 : `$1 = {array = 0x000122030, len = <optimized out>, cap = <optimized out>}`
2. contenu de la variable t à la ligne 15 : `$2 = <optimized out>`

---

La commande `print` permet d'afficher ponctuellement la valeur d'une variable.
Or, on a souvent besoin d'afficher la valeur d'une variable après chaque pas d'exécution.
Dans ce cas, on utilisera la commande `display`, qui partage la même syntaxe que `print`.
Pour oublier une des expressions enregistrée avec `display`, on utilisera la commande  `undisplay <num>`, où `<num>` est le numéro de l'expression à oublier dans la liste des expressions affichées à chaque pas d'exécution.

---


**Exercice n°4**

  - Tester l'affichage de différentes expressions avec la commande `display`.
  - Redémarrer le programme et noter les valeurs successives de la variable `minpos` à chaque tour de boucle.

---

[^1]: Il existe entre autre une extension pour l'éditeur de texte vscode qui permet de s'interfacer avec GDB.

