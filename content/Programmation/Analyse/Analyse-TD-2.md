---
title: Analyse-TD-2
draft: 
description: 
tags:
  - Analyse
---
# Exercice 1
Une version plus élaborée est présentée ci-après. Le système est nommé TPV (Terminal de Point de Vente), les envois de messages respectent la convention de nommage, les envois de messages réflexifs sont ajoutés, les retours sont explicités, les pré- et post-conditions sont explicitées sur la ligne de vie du TPV. Notez également l’ajout d’un fragment combiné « ref ».

**Que représente ce fragment combiné « ref » intitulé « Traiter le paiement » ? Quel lien feriez- vous avec le diagramme de cas d’utilisation ?**
Le fragment "ref" sert à représenter un bout de séquence déjà utiliser dans un autre diagramme de séquence.

**Identifiez les scénarios alternatifs et d’erreurs. Identifiez où les positionner sur le Diagramme de Séquence.**
- paiement par espèces / chèque
- paiement par carte bancaire
- caisse défectueuse -> erreur
- caisse à court de monnaie -> erreur
- carte refusé -> erreur
- arrêt du scan des articles par le client
- montant insuffisant -> erreur

**Modélisez le diagramme de séquences de « Traiter le paiement » en prenant en compte les différentes possibilités de paiement. Représentez uniquement le scénario nominal. Pensez à exploiter les fragments combinés présentés en cours de Développement Objet en BUT1.**
*fait sur papier, seras réécris en UML plus tard #TODO* 
