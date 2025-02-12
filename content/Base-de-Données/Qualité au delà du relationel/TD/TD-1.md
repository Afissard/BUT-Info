---
title: TD-1
draft: 
description: 
tags:
  - Base-de-Données
---
# 1 Normalisation dans le modèle relationnel et qualité
## Exo 1

| A   | B   | C   | D   |
| --- | --- | --- | --- |
| a1  | b2  | c1  | d1  |
| a1  | b2  | c1  | d2  |
| a1  | b2  | c2  | d3  |
| a2  | b4  | c2  | d2  |
| a3  | b3  | c4  | d2  |

1. Clés candidates de la relation : AD, BD, CD
2. Dépendance fonctionnelle
	1. A->B ? : oui
	2. A->C ? : non
## Exo 2
On considère la relation R(A,B,C) suivante munie de l'ensemble F de dépendances fonctionnelles suivant: F={A->B, B->C}
1. La clé de cette relation est A (sauf si ses éléments ne sont pas atomique)
2. La forme normal de R est 2NF
## Exo 3

| A   | B   | C   |
| --- | --- | --- |
| 1   | 1   | 1   |
| 2   | 1   | 2   |
| 3   | 2   | 1   |
| 4   | 3   | 3   |
1. Clé de la relation : A, BC
2. Forme normal de la relation : 3NF
3. Couverture minimal de F : Fmin = F
## Exo 4
Soit la relation R(A,B,C,D,E) munie de l'ensemble F de dépendances fonctionnelles suivant: F={A-->B, A-->D, A-->E, E-->C}
1. Clé primaire : A (si A est atomique)
2. Forme normal : 3NF
3. F constitue-t-il une couverture minimal ? : Fmin = F
## Exo 5
Soit l'ensemble F de dépendances fonctionnelles de la relation R(A,B,C,D,E,F,G) F= {A-->B, BC --> DE, AEF -->G}
Calculer la fermeture AC+ de F par application de l’algorithme vu en cours

On peut atteindre : 
- A et C par réflexivité
- B par A
- D et E par BC
- Manque F et G
# Exo 6
Trouver la couverture minimale de F par application de l’algorithme vu en cours : 
F={AB->C, C->A, BC->D, ACD->B, BE->C, CE->FA, CF->BD, D->EF}
**Étape 1 : rendre les dépendance simple :**
F1={AB->C, C->A, BC->D, ACD->B, BE->C, CE->A, CF->B, CF->D, D->E,D->F}
**Étape 2 : éliminer les attributs superflus :**
F2={AB->C, C->A, C->D, CD->B, BE->C, CE->F, CF->B, CF->D, D->E,D->F}
**Étape 3 : éliminer les dépendance fonctionnelles redondante :**
F3={AB->C, C->A, BC->D, BE->C, CE->F, CF->B, D->E, D->F}
# Exo 7
Soit la relation suivante : R<E,R,D,P>, {E-->D, D-->R}>.

| Solution                     | R1<{E,D}, {E-->D}> R2<{D,R}, {D-->R}> | R1<{E,D}, {E-->D}> R2<{E,R,P}, {E-->R}> | R1<{D,R}, {D-->R}> R2<{E,D,P}, {E-->D}> | R1<{E,P,R}, {E --> R}> R2<{E,D,P},{E-->D} > |
| ---------------------------- | ------------------------------------- | --------------------------------------- | --------------------------------------- | ------------------------------------------- |
| Préservation des dépendances | Oui                                   | Non                                     |                                         | oui                                         |
| Préservation du contenu      | Non, on perd le P                     | Oui                                     |                                         | oui                                         |
| Normalité                    | 1NF                                   | R1 : 3NF<br>R2 : 1NF                    |                                         | R1 : 3NF<br>R2 : 3NF                        |
# Exo 8

| A   | B   | C   |
| --- | --- | --- |
| a1  | b1  | c1  |
| a1  | b2  | c3  |
| a2  | b1  | c1  |
| a2  | b4  | c3  |
| a1  | b5  | c1  |
Quelles sont les dépendances fonctionnelles incorporées dans cette relation :
- clés : AB : clé primaire de R
- forme normal : 1NF 
Quelles sont les clés candidates de la relation : 
Quelle est la forme normale de la relation :
les décompositions suivantes préservent-elles les dépendances fonctionnelles et les données? 

| Solution                     | R1(B,C), R2(A,C)     | R1(B,C), R2(A,B)     |
| ---------------------------- | -------------------- | -------------------- |
| Préservation des dépendances | oui                  | oui                  |
| Préservation du contenu      | non                  | oui                  |
| Forme Normale                | R1 : 3NF<br>R2 : 3NF | R1 : 3NF<br>R2 : 3NF |

# Exo 9

| A   | B   | C   |
| --- | --- | --- |
| a1  | b1  | c1  |
| a1  | b2  | c2  |
| a2  | b1  | c1  |
| a2  | b2  | c3  |
| a3  | b1  | c1  |
Est-ce que la relation R est en 3FNBC ? 
- F:{C->B, AB->C}
- Clés : AB, AC
- 3NF uniquement
Dans le cas où votre réponse est non, proposez une décomposition de R qui soit en 3FNBC et vérifiez la qualité de cette décomposition (sans perte de dépendances et sans perte de données).
# Exo 10

| A   | B   | C   | D   | E   |
| --- | --- | --- | --- | --- |
| 1   | 1   | 1   | 1   | 1   |
| 1   | 1   | 2   | 1   | 1   |
| 1   | 1   | 5   | 1   | 1   |
| 1   | 2   | 5   | 4   | 1   |
| 2   | 1   | 2   | 1   | 1   |
| 2   | 1   | 5   | 1   | 1   |
| 2   | 1   | 5   | 3   | 1   |
| 2   | 2   | 5   | 4   | 1   |
| 3   | 3   | 3   | 3   | 2   |
Dépendance fonctionnelles

# Exo 11

[^1]: 
