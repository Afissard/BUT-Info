Retour sur la recherche, exemple d'une fonction de recherche avec une fonction recursive
```go
// Voir #Madoc pour exemple écrit en go
```
Une fonction recursive est une fonction qui s'appelle elle même, ou via une chaine d'appel de fonction.
```go
func my_fn(i int){
	log.Print(i)
	i --
	if (i>0){
		my_func(i)
	} 
}
```
**Attention au boucle infini si un cas d'arret est oublié, la compilation ne seras pas possible ou causeras un bug à cause d'une dépassement de la mémoire alloué au programme.**
cours pas fidèle à au pdf de #Madoc 
L'appel d'une fonction est couteux en langage machine, parfois le compilateur peut de manière intelligente remplacer l'appel de la fonction par le code de la fonction (voir *Inlining*), or pour une fonction récursive ... (démonstration pour 3 appel de fonction)

La récursivité est une alternative, plus simple (mais dure à maitrisé) aux boucles.

Exemple de suite codé avec des fonction récursive
Pour la definition d'une suite par récursivité, il faut initialiser la valeurs des premier thermes
Exemple avec une boucle, comparaison avec fonction récursive (dans le pdf):
Suite : u(n-1) = u(n) = f(u(n)-1)
Algo : on appel u(i) la fonction qui retourne u(0) si i = 0 et qui retourne f(u(i-1)) sinon retourner u(n)

## Pourquoi la récursivité
Certain problème on une solution simple avec une fonction récursive, mais complexe avec une boucle : 
Exemple (code Go) Etant donné un tableau t contenant des caractères et un entier k ≤ len(t), donner toutes les chaines de k caractères que l’on peut construire en utilisant au plus une fois chaque caractère de t. 
Pour t = a b c d et k = 2 on doit trouver ab, ac, ad, ba, bc, bd, ca, cb, cd, da, db et dc. Pour t = a b c d et k = 3 on doit trouver abc, abd, acb, acd, adb, adc, bac, bad, bca, bcd, bda, bdc, cab, cad, cba, cbd, cda, cdb, dab, dac, dba, dbc, dca, dcb. 
Code dans *recursif_iterative.go*

Autres exemples plus tard Tri fusion, algorithmes sur les arbres 
#TODO récupérer les documents, mis à jours sur #Madoc ?

## Soyez prudents
Un appel de fonction est plus long qu’un tour de boucle. Les appels de fonctions non résolus occupent la mémoire. 
### Récursivité terminale 
Si l’appel récursif est la dernière chose qui se passe dans une fonction, le compilateur peut optimiser et éviter les appels de fonctions non résolus. 
### Remarque 
Une fonction récursive peut toujours être rendue récursive terminale en utilisant la notion de continuation. Cependant, c’est un peu compliqué et on n’en parlera pas dans ce cours.