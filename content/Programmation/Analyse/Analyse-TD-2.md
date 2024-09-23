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
# Exercice 2
Nous reprenons l’énoncé de l’exercice 2 du TD 1 : Dans le cadre de l’amélioration qu’elle veut apporter à son système d’information, une entreprise souhaite modéliser, dans un premier temps, le processus de formation de ses employés afin que quelques-unes de leurs tâches soient informatisées.
1. Le processus de formation est initialisé lorsque le responsable formation reçoit une demande de formation de la part d’un employé. Cette demande est instruite par le responsable qui la qualifie et transmet son accord ou son désaccord à l’intéressé.
2. En cas d’accord, le responsable recherche dans le catalogue des formations agréées un stage qui correspond à la demande. Il informe l’employé du contenu de la formation et lui propose une liste des prochaines sessions. Lorsque l’employé a fait son choix, le responsable formation inscrit le participant à la session auprès de l’organisme de formation concerné.
3. En cas d’empêchement, l’employé doit informer le responsable de formation au plus tôt pour annuler l’inscription ou la demande.
4. À la fin de sa formation, l’employé doit remettre au responsable formation une appréciation sur le stage qu’il a effectué, ainsi qu’un document justifiant de sa présence.
5. Le responsable formation contrôle par la suite la facture que l’organisme de formation lui a envoyée avant de la transmettre au comptable achats.

**Décrivez la dynamique du processus de formation au moyen d’un diagramme d’activité. Utilisez des partitions verticales pour affecter les responsabilités aux acteurs.**
