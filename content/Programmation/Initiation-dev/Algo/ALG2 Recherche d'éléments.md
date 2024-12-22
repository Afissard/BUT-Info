Presentation d'une structure de donnée : les tableaux. Maitrisé
## Problématiques :
### Recherche d'élément
Etant donnée un tableau t et un entier x, dire si x appartient à t et, si oui, donner un indice i tel que t\[i] = x
```Python
def search_tab(t, x):
	i = 0
	while i < len(t):
		if t[i] == x: return (True, i)
		else : i += 1
	return (False, -1) # -1 to prevent searching in the tab

t = [12, 32, 7, 23, 9]
print(search_tab(t, x)) # x is the target
```
#### Pourquoi se poser le nombre d'opération ?
Permet de comparer les algorithmes entre eux, savoir lequel sera le plus efficaces en temps de calcul.
#### Quelle opérations ?
Affectation de variable, test (plus très vrai sur machine moderne), opération arithmétique
#### Remarque
Pour être véritablement précis il faudrait compter les instructions en langage machine, car c'est ce qui est véritablement exécuter.
- Cela dépend du langage, du compilateur, du processeur
- donc ordre de grandeur
#### Nombre d'opération de l'algo
- 1 affectation de variable à l'initiation
- 2 test et 1 une opération arithmétique par tour de boucle.
Soit 3k+1 ou k est le nombre de tours de boucle.
#### Nombre d'opération dans le pire cas
On ne peut pas savoir à l'avance le nombre de tours de boucles effectuées, on considère ce qui va se passer au pire cas : k = len(t)
### Recherche d'un minimum
Trouver le plus petit entier du tableau t.
```Python
def search_mini(t):
	i = 1
	i_min = 0
	while i<len(t):
		if t[i]<t[i_min] : i_min = min
		i +=1
	return i_min

t = [12, 32, 7, 23, 9]
print(search_min(t))
```
#### Nombre d'opération
- 2 opération à l'initiation
- 3 opération à chaque tour de boucle
- 1 opération à chaque tour et où la condition à lieu
Soit 3\*(len(t)61+k+2)
#### Nombre d'opération dans le pire cas
La condition à lieu tout le temps : k = (len(t)-1)
### Sélection d'un éléments
Trouver tous les éléments de t qui sont plus petit que x. On veut construire un l'ensemble E tel que y appartient à E si et seulement si y appartient à t et y <= x.
```Python
def search_tab(t, x):
	i=0
	E=[]
	while i<len(t):
		if t[i]<= x :
			E.append(t[i])
		i +=1
	return E

t = [12, 32, 7, 23, 9]
print(search_tab(t, x)) # x is the target
```
#### Nombre d'opération
similaire à la recherche de minimum, à condition que les opération sur les ensemble soient simple
### Peut on faire mieux ?
#### Bilan des nombre d'opération nécessaires
- recherche d'éléments : 3\*len(t)+1
- recherche minimum et selection d'éléments : ...
#### Ordre de grandeur
Peut varié dans ces nombre d'opérations (en fonction de l'implémentation de l'algo, du compilateur, etc.) mais il restent linéaires en la taille du tableaux
#### des algo pour réduire l'ordre de grandeur
On ne peut pas faire mieux sur les tableaux quelconque, mais si on ... #Madoc pour finir

## Recherche d'éléments dans un tableau trié
Dans un tableaux trié dans un ordre voulue
### Recherche dichotomique d'éléments dans un tableaux trié
```Python
def search_tab(t, x):
	déb = 0
	fin = len(t)
	while deb < fin :
		mid = (deb + fin)/2
		if t[mid] == x : return (True, mid)
		elif t[mid]>x : fin = mid # N'a pas lieu au même tour que else
		else t[mid]<x: deb = mid+1 # N'a pas lieu au même tour que elif
	return (False, -1)

t = [7, 9, 11, 12, 23, 29, 32]
print(search_tab(t, x)) # x is the target
```
#### Nombre d'opération dans le pire cas
5\*k+2\*log(len(t))
#TODO compléter avec #Madoc 
## Remarque
Il n'est intéressant de trier un tableau que si il est grand et souvent utiliser 
