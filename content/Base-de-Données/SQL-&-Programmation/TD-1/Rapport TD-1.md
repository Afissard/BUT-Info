---
title: Rapport TD-1
draft: 
description: 
tags:
  - SQL
  - Base-de-Données
---
## Contraintes d'intégrité

**Table des contraintes**

| nom table | nom contrainte      | type de contrainte | différer ou non |
| --------- | ------------------- | ------------------ | --------------- |
| Employe   | PK_employe          | Clé primaire       | non             |
| Service   | PK_service          | Clé primaire       | non             |
| Projet    | PK_projet           | Clé primaire       | non             |
| Employe   | FK_affect           | Clé étrangère      | non             |
| Travail   | FK_employe_travail  | Clé étrangère      | non             |
| Travail   | FK_projet_travail   | Clé étrangère      | non             |
| Concerne  | FK_service_concerne | Clé étrangère      | non             |
| Concerne  | FK_projet_concerne  | Clé étrangère      | non             |

**Table de tests des contraintes**

| nom table  | requête                                                                                | code d'erreur | message d'erreur            |
| ---------- | -------------------------------------------------------------------------------------- | ------------- | --------------------------- |
| PK_employe | `Insert into employe values(20,'toto',35,1); --insertion d'un employé qui existe déjà` | `-0001`       | Vilolation de la contrainte |
|            |                                                                                        |               |                             |
