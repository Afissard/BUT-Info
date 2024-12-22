# Structure de donnée
Moyen de représenter des données et leurs relation en mémoire ainsi que la manière d'y accéder.
## Pourquoi plusieurs types de structures de données ?
- Opération adapté à certains algorithmes et pas à d'autres
- Compromis entre espace mémoire occupé par les données et efficacité temporel des opération temporelle des opération de base
- Décrire les relations, plus ou moins complexes entre les données
# En pratique
## Definition d'une structure de données
On définit une interface : l'ensemble des opérations qu'on doit pouvoir réaliser sur la structure
### Exemple tableau en (en Go)
- t\[i]
- len(t)
- ...
## Implantation d’une structure de données 
Il existe plusieurs façons de représenter (coder) une interface donnée, qui peuvent être plus ou moins efficaces.
# Exemple de structure de données
#Madoc voir les pdf et fichier go dans le dossier ci-présant
# Structure de données récursives
C'est une structure de donnée type "node" (je sais de à quoi je pense)
#todo recoder moi même sans doc pour voir si je suis encore compétant en go
# Benchmark
Test du temps moyen d'execution pour comparer complexité liste/node
la version avec les pointeur a un désavantage avec Go, il y a un garbage collector (ralenti à grande echelle), en C il faut s'occuper de vider. Possibilité de garder les element retiré dans une autre structure pour les réutiliser plus tard (situational)
## File/Pile
File on retire le premier, Pile on retire le dernier
# Liste
Sequence d'élément d'ont l'ordre est fixé.
# Recherche récursive dans une structure de données
#Madoc trouver fichiers sources
#TODO recoder moi même parce que je le peux