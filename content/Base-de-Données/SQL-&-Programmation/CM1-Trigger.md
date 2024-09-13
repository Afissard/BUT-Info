---
title: SQL-&-Programmation-CM1
draft: 
description: 
tags:
  - Base-de-Données
---
# 1. Généralités
- Unicité de la description des données. Description des contraintes d'intégrité.
- Mécanisme de "vues".
- Interrogation et manipulation "non procédurale", expression d'opérations sur des ensembles.
- Indépendance traitement/stockage (logique/physique).
- Deux enjeux *forts*:
	- le maintien de l'intégrité.
	- la concurrence d'accès.
# 2. Description
```sql
entreprise(nument, noment, budgetStage); 
offre_stage(numoffre, #nument, libelle, salaire)
```
La contrainte peut être violée par la création d'une offre de stage, par la suppression d'une entreprise ou par la modification de l'affectation de l’offre de stage.

`Option on delete [cascade | set null]`
Destinée à imposer un comportement qui assure le maintien de la contrainte de référence.

La suppression d'une entreprise impliquera la suppression de toutes les offres qui sont affectées à l ’entreprise supprimée. (remplacera par "NULL" dans `offre_stage` toutes les `#nument` égaux à `nument` l’entreprise).

`Option initially [immediate | deferred]`
Indique si la contrainte doit être vérifiée :
- **immédiatement** après l'instruction qui peut la violer,
- ou si la vérification doit être **différée** à la confirmation (au moment du `commit`).
Le défaut d'ORACLE est `initially immediate`.

`Option deferrable`
Permet de choisir le moment de la vérification pour les transactions. Le défaut d'ORACLE est `not deferrable``.

`set constraint [{nom_contrainte,} … | all} [immediate | deferred]`
# 3. Déclencheurs (triggers) 
#TODO
