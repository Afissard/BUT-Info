**Définition : *Séquence d'étapes de calculs qui transforment une entrée en sortie.***
Peut être plus où moins généraux et/complexes selon l'objectif du programme.

## Exemple d'algorithme : l'addition
### Problème
Prendre deux nombre entiers (positifs) x et y, exprimé en base 10, calcule et retourne la valeur de la somme *x+y*.
### Principe
Additionner les chiffres deux à deux, de droite à gauche, en propageant les éventuels retenue.
### Notations
On note : X = XnXn-1...X1X0 et Y = YmYm61...Y1Y0 et on suppose (sans perte de généralité) que n >= m.
### Ecriture de l'algorithme
Voir diapo sur #Madoc R101, page 5/10.
#TODO réécrire l'algo présenté en python.

## Autre exemple d'algorithme : le PGCD
### Problème 
Etant donnés 2 entiers positifs *a* et *b*, calculer le plus grand diviseur commun de *a* et *b*.
### Algorithme
```python
while b != 0 : 
	b = a 
	b = b % a # rappel % = modulo
return a
````
## Exemple de problème algorithmique
- Biologie : problème d'alignement de séquences d'ADN
- Internet : problème de routage, recherche de classement de données, de compression
- Industrie : problème de logistique et planification avec optimisation des ressources
## Pourquoi étudier l'algorithmie
### Evolution des performance des ordinateur
Loi de Moore (tout les 2 ans la taille des transistors est réduite) (obsolète), la mémoire augmente aussi
### Mais
- l'exécution d'algorithme a toujours un coût en énergie et en temps
- et certains problèmes sont trop complexe pour user de la force brute (mais peut évoluer avec l'arriver de nouvelle technologie, exemple ordinateur quantique)
