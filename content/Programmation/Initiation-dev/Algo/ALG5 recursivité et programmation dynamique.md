réexplicitions des fonctions recursive
voir pdf sur #Madoc

# Fibonacci récursif
exemple avec la suite (ou toute autre suite) de Fibonacci (encore avec les lapins (retour des problèmes génétiques))
## Implémentation naïve
Problème lors de grand appelle -> rallonge grandement le temps d'exécution -> "arbre d'appel de fonction" donc gourmant
Problème : on recalcule des valeurs déjà calculer
## Amélioration : Algo dynamique
Solution : utiliser un tableau pour stocker les valeur calculé, pour allez les chercher au lieu de les recalculer, on ne calcule plus qu'une seule fois chaque valeur de la suite (voir schéma arbre de l'algo modifié)
Or problème si l'on veut calculer des plus grandes valeur si pose donc le problème du stockage du tableau
## Amélioration : Algo tail-recursive
(tail-recursive : appel recursive uniquement à la fin de la fonction -> optimisation à la compilation ~= boucle)
On utilise deux var pour stocker u(n-1) et u(n-2), qui sont les seul variables necessaire pour trouver u(n). 
Plus proche d'une logique humaine qu'informatique (enfin pour moi). Voir fichier source
Cette méthode est plus situationnel que l'algo dynamique.

# Programmation dynamique
Ne sert qu'a amélioré la programmation récursive

## Exemple problème dans la recherche scientifique génétique
La plus longue sous chaine commune, voir pdf pour énoncé
algo : src_recursivité/sous-chaine.go
Utilisation de map (~= dictionnaire en python)