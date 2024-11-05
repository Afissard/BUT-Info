---
title: Consigne
draft: 
description: 
tags:
  - SAE
---
# qdev.tp.SAE 

**(ATTENTION document en cours de mise à jour)**

Dans le cadre de la SAE du semestre 3, il vous est demandé de développer un site web dynamique de vente en ligne. 

**Dans ce cadre, nous allons vous ~~forcer à~~ demander d'identifier des _patrons de conception_
et à les mettre en oeuvre en PHP.**


## Contexte

Dans le cadre de R3.03 | Analyse, vous allez établir différents diagrammes d'analyse, parmi lesquels un diagramme de classes d'analyse des données
métier que devra manipuler votre application :  ce diagramme de classe niveau analyse devra être mis en perspective avec votre future implémentation PHP + BDD, c.à.d qu'il vous faudra établir
un second diagramme de classes, niveau conception détaillée, comme un raffinement du diagramme de classes niveau analyse. 

A ce niveau de votre réflexion, vous identifierez différents patrons de conception, vous expliquerez comment les mettre en oeuvre au niveau de votre conception UML détaillée, puis vous devrez les implémenter au sein de votre projet d'application web.

Vous ne perdrez pas de vue qu'un projet réel va évoluer : nouveaux types d'utilisateurs, nouveaux produits, nouvelles fonctionnalités, modification d'une procédure (d'achat par exemple),... Ces réflexions alimentent donc forcément les choix de conception.


## Travail demandé

Tout au long de la ressource R3.04, nous mettons en application certains patrons de conception en utilisant le langage Kotlin ; durant cette SAE vous devrez les mettre en application en utilisant (à priori) le langage PHP.

Plus précisément, vous devrez appliquer (à minima) 3 patrons de conception au choix lors de vos développements :

- 2 patrons de conception détaillés dans les cours de la ressource R3.04,

- 1 patron de conception non détaillé dans les cours.


**Attention**, il ne s'agit pas de nous présenter les design patterns utilisés de manière hadhoc par le framework PHP que vous utiliserez, mais bien de mettre en oeuvre de 'A' à 'Z' des design patterns, sur les classes PHP métier que vous devrez développer.

**Attention 2** : MVC (et toutes ces variantes) n'est pas un design pattern. 

## Travail à rendre


Vous rendrez un mini-rapport explicitant les trois patrons de conception choisis.  Pour chaque DP, le rapport présentera les points suivants :

1.  Justifier votre choix de ce design pattern (le pourquoi ?)
	- En quoi ce DP répond au problème 
	- Quels avantages/ inconvénients par rapport à d'autres solutions
	- ... 	
	
2.  Illustrer et expliquer le DP choisi (le quoi ?)
   - Diagrammes UML niveau conception détaillée,
   - ...
   
3.  Présenter son implémentation (le comment ?)
	- Extraits de code pertinent
	- Choix/adaptation technique
   - ...


Le mini-rapport sera rédigé au format [Markdown](https://fr.wikipedia.org/wiki/Markdown) uniquement, sasn fioriture (~~.docx~~, ~~.pdf~~) Markdown est un format classique pour la documentation.

Des points seront accordés pour l'originalité des DPs choisis (= DPs non-implémentés par le framework).

Vous déposerez sur Madoc une archive .zip contenant uniquement :
- le fichier Markdown `.md` contenant vos explications 
-  un unique dossier `img/` contenant d'éventuels (mais sûrement nécessaire) diagrammes UML inclus dans votre rapport

Vous rendrez votre travail via un dépôt Madoc **à préciser** pour la fin de la semaine 2 (**date et horaires à préciser**)


Des entretiens complémentaires seront également organisés, lors de la **semaine 3**.
