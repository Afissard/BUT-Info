---
title: Rapport TD-1
draft: 
description: 
tags:
  - SQL
  - Base-de-Données
---
# Base de données

| Nom table | clefs primaire | autres colones         |
| --------- | -------------- | ---------------------- |
| employe   | nuempl         | nomempl, hebdo, affect |
| service   | nuserv         | nomserv, chef          |
| projet    | nuproj         | nomproj, resp          |
| travail   | nuempl, nuproj | duree                  |
| concerne  | nuserv, nuproj |                        |

# Création des tables
```sql
-- Reset
DROP TABLE concerne CASCADE CONSTRAINTS PURGE;
DROP TABLE travail CASCADE CONSTRAINTS PURGE;
DROP TABLE projet CASCADE CONSTRAINTS PURGE;
DROP TABLE employe CASCADE CONSTRAINTS PURGE;
DROP TABLE service CASCADE CONSTRAINTS PURGE;

-- Creation des tables
-- employe
CREATE table employe AS SELECT * FROM basetd.employe;
ALTER TABLE employe ADD CONSTRAINT PK_employe PRIMARY KEY (nuempl);
SELECT * from employe;

-- service
CREATE table service AS SELECT * FROM basetd.service;
ALTER TABLE service ADD CONSTRAINT PK_service PRIMARY KEY (nuserv);
SELECT * FROM service;

-- projet
CREATE table projet AS SELECT * FROM basetd.projet;
ALTER TABLE projet ADD CONSTRAINT PK_projet PRIMARY KEY (nuproj);
SELECT * FROM projet;

-- travail
CREATE table travail AS SELECT * FROM basetd.travail;
ALTER TABLE travail ADD CONSTRAINT PK_travail PRIMARY KEY (nuempl, nuproj);
SELECT * FROM travail;

-- concerne
CREATE table concerne AS SELECT * FROM basetd.concerne;
ALTER TABLE concerne ADD CONSTRAINT PK_concerne PRIMARY KEY (nuserv, nuproj);
SELECT * FROM concerne;
```
# Contraintes d'intégrité
```sql
-- Ajout des clefs étranères
ALTER TABLE travail add CONSTRAINT FK_employe_travail FOREIGN KEY (nuempl) REFERENCES employe(nuempl);
ALTER TABLE travail add CONSTRAINT FK_projet_travail FOREIGN KEY (nuproj) REFERENCES projet(nuproj);
ALTER TABLE concerne add CONSTRAINT FK_service_concerne FOREIGN KEY (nuserv) REFERENCES service(nuserv);
ALTER TABLE concerne add CONSTRAINT FK_projet_concerne FOREIGN KEY (nuproj) REFERENCES projet(nuproj);
```
Le chef d'un service est un employé affecté au même service. Attention les contraintes de clé étrangère `Affect` et `Chef` se croisent. Dans ce cas, vous devriez créer une contrainte unique `(nuempl, affect)` et une contrainte différée (initially deferred) de la table `service` `(chef,nuserv)` vers cette contrainte unique `(nuempl, affect)`.
```sql
ALTER TABLE employe add CONSTRAINT FK_affect FOREIGN KEY (affect) REFERENCES service (nuserv);
ALTER TABLE service add CONSTRAINT FK_chef FOREIGN KEY (chef) REFERENCES employe (nuempl) DEFERRABLE;
```
Un responsable de projet doit travailler sur le projet. Dans ce cas vous n'avez pas besoin de créer la contrainte Unique `(nuempl, nuproj)` de la table `travail`, car les deux attributs forment déjà la clé primaire. Par contre la contrainte `(resp, nuproj)` est une clé étrangère différée vers la clé primaire de la table `travail`.
```sql
ALTER TABLE projet add CONSTRAINT FK_resp FOREIGN KEY (resp, nuproj) REFERENCES travail (nuempl, nuproj) DEFERRABLE;
```
La durée (temps de travail) hebdomadaire d'un employé est inférieure ou égale à 35h.
```sql
ALTER TABLE employe add CONSTRAINT temps_trav_max CHECK (hebdo <=35) DEFERRABLE;
```

## Table des contraintes

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
| Service   | FK_chef             | Clé étrangère      | oui             |
| Projet    | FK_resp             | Clé étrangère      | oui             |
| Employe   | temps_trav_max      | ?                  | oui             |

## Tests des contraintes

| nom table  | requête                                                                                | code d'erreur | message d'erreur            |
| ---------- | -------------------------------------------------------------------------------------- | ------------- | --------------------------- |
| PK_employe | `Insert into employe values(20,'toto',35,1); --insertion d'un employé qui existe déjà` | `-0001`       | Vilolation de la contrainte |
|            |                                                                                        |               |                             |
#TODO

# Mise à jour des données
La table employé contient un attribut salaire de type number. Vous rajoutez cet attribut  
avec la requête suivante : `alter table employe add salaire number;`

Faites les modifications suivantes, à chaque fois avec une seule requête:  
- Les responsables de projet gagnent au moins 2500 Euros,  
- Les chefs de service gagnent au moins 3500 euros.  
- Les autres employés gagnent un salaire < 2000 euros.
```sql
ALTER TABLE EMPLOYE ADD SALAIRE NUMBER;

UPDATE EMPLOYE
SET SALAIRE = 2500
WHERE NUEMPL IN (SELECT RESP FROM PROJET);

UPDATE EMPLOYE
SET SALAIRE = 3500
WHERE NUEMPL IN (SELECT CHEF FROM SERVICE);

UPDATE EMPLOYE
SET SALAIRE = 1999
WHERE NUEMPL NOT IN ((SELECT CHEF FROM SERVICE) UNION (SELECT RESP FROM PROJET));
```

La somme des durées d'un employé (de la table travail) doit être inférieur à la durée hebdomadaire `(Sum(duree) <= hebdo)`. Affichez les employés qui ne respectent pas cette contrainte. Vous modifiez les données de la table travail(update sur la durée) ou la table employé (update sur hebdo)jusqu'à ce que la requête précédente retourne un ensemble vide.
```sql
SELECT * FROM employe e
WHERE (
	SELECT sum(duree) FROM travail t
	WHERE e.NUEMPL = t.NUEMPL
) > e.HEBDO ;

-- TODO
```
Un employé ne peut être responsable de plus de 3 projets. Idem que la contrainte précédente, trouvez les employés qui ne respectent pas la contrainte et modifiez les données.

Le chef d’un service gagne plus que les employés du service. Idem que la contrainte précédente, trouvez les employés qui ne respectent pas la contrainte et modifiez les données.

Un service ne peut être concerné par plus de 3 projets.
