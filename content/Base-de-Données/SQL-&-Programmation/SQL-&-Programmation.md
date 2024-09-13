---
title: SQL-&-Programmation
draft: 
description: 
tags:
  - Base-de-Données
  - SQL
  - SQL-Programmation
---
# Index SQL-&-Programmation
- Cours
	- [[CM1-Trigger]]
- TD
	- [[Rapport TD-1]]
	- [[Rapport TD-2]]
# Résumé / Introduction
La différence entre SQL avec un langage de programmation et un langage de programmation, SQL permet de conservé un historique des modification des variable avec un système similaire à git.

Si une modification est faite, si celle ci ne respecte pas les contraintes dynamiques elle sera annulé lors du commit et pas avant.
Jusqu'à présent les erreurs que nous avons eu, sont des erreurs issue de contraintes statique. Associé à un langage de programmation, l'on peut donc gérer les erreurs issue de contraintes dynamiques (ou statique) en utilisant des `try & catch` et la gestion des erreurs usuel du langage utilisé.

L'on peut créer des procédures en SQL, et appelez ces procédure avec une application (codé en Kotlin par exemple), cela créer une relation client/serveur entre l'application et le serveur de la base de données. Les données envoyé / reçus sont encodé en JSON ou XML.

Répartition du cour : 
- utilisation des contraintes dynamiques.
- relation entre Kotlin et une base de données.
- utilisation d'une IHM pour interagir avec la BD.

Quand deux personnes travail sur la même BD, la concurrence d'accès est importante, tant qu'un commit n'est pas réalisé l'autre personne ne peut voir ce que la première à fait sur la BD, ce qui peut causé des conflit si chacun réalise une insertion à la même position, d'où l’intérêt de créer un verrou pour prévenir ces conflits.