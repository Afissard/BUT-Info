---
title: "Untitled"
draft: 
description: 
tags:
---
1. .
2. le test passe
3. couverture de code : classe / méthode / ligne de code
	1. il y a des losange pour indiquer les "décisions", chaque prédicats soit couvert 2 fois (true / false), couvre les décisions dans le code = couverture des arcs
4. avec "nn" la couverture est plus grande mais pas complète
5. nouveau cas de test supplémentaire : ct1("nn", T), ct2("no", F)
	1. couvrir tout les arcs => couvrir tout les nœuds mais pas l'inverse en général. Néanmoins ici c'est le cas : couvrir tout les nœuds <=> tout les arcs car pour couvrir toutes les instructions (nœuds), il faut faire 2 tour de boucle, ce qui couvre aussi tout les arcs
6. 7/9 mutant tué (traquer test couvre 4-5/9 (il y a des mutant tués par plusieurs cas de test))
7. pour tuer un 8ème mutant il faut un mot non palindrome de 4 lettres ... ct3("nonn", F)
8. Attention ct4("", F) n'est pas couvert on devrait avoir ct4("", exeption) pas de vrai mot pouvant tuer le dernier mutant qui est équivalent. Changer `<` en `<=` n'a pas d'impact sur le verdict d'un test (on ne peut pas RIP).