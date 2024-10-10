---
title: Sacha Chauvel - Compte rendu TD SQL
draft: 
description: 
tags:
  - SQL
  - Base-de-Données
---
Sacha CHAUVEL
Groupe 1
identifiant oracle : s3a02b
mot de passe oracle : K1was
# TD 3
## Exercice 1
Le tableau suivant indique les erreurs soulevées par oracle et les erreurs que vous devriez retourner par la procédure `Creer_employe`. Les codes erreurs retournés sont entre `-20000` et `-20999`.

| Test                              | Code erreur oracle -> Code d'erreur retourné | Message                                                                                                                                  |
| --------------------------------- | -------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| PK                                | -0001 -> -20401                              | Un employé avec le même numéro existe déjà                                                                                               |
| FK                                | -2291 -> -20402                              | Le service auquel il est affecté n’existe pas                                                                                            |
| CHECK                             | -02290 -> -20403                             | la durée hebdomadaire d’un employé doit être inférieure ou égale à 35h                                                                   |
| TAILLE D’UN ATTRIBUT >DESCRIPTION | -1438 -> -20404 et -12899 -> -20405          | Une valeur (nombre) dépasse le nombre de caractères autorisés Une valeur (chaine de caractère) dépasse le nombre de caractères autorisés |
| SALAIRE                           | -20406 (erreur du trigger)                   | Le salaire de cet employé dépasse celui de son chef de service                                                                           |
| AUTRES                            | -20999                                       | Erreur inattendue                                                                                                                        |
Compléter la procédure `Creer_employe` qui se trouve sur madoc, dans le fichier `Procedure_Package_MAJ.txt`. Cette procédure est dans un package. Vous devriez compléter la procédure `Creer_employe` qui se trouve dans le package `BODY`. Vous compilez la package spécifications(signature des procédures) avant le package body. Reporter ces erreurs dans le tableau(remplacer les XX).

```sql
CREATE OR REPLACE PACKAGE MAJ is
	PROCEDURE CREER_EMPLOYE (
		LE_NUEMPL IN NUMBER, LE_NOMEMPL IN VARCHAR2, LE_HEBDO IN NUMBER, LE_AFFECT IN NUMBER,LE_SALAIRE IN NUMBER
	);
END;

--------------------------------------------------------------------------------------------------------------

CREATE OR REPLACE PACKAGE BODY MAJ is

PROCEDURE CREER_EMPLOYE (LE_NUEMPL IN NUMBER, LE_NOMEMPL IN VARCHAR2, LE_HEBDO IN NUMBER, LE_AFFECT IN NUMBER,LE_SALAIRE IN NUMBER) is
BEGIN
	SET TRANSACTION READ WRITE;
	INSERT INTO employe VALUES (LE_NUEMPL, LE_NOMEMPL, LE_HEBDO, LE_AFFECT,LE_SALAIRE);
	COMMIT;
	
	EXCEPTION
	WHEN OTHERS THEN
	ROLLBACK;
	IF SQLCODE=-00001 THEN
		RAISE_APPLICATION_ERROR(-20401, 'Un employe avec le meme numero existe deja');
	ELSIF SQLCODE=-2291 THEN
		RAISE_APPLICATION_ERROR(-20402, 'Le service auquel il est affecté n’existe pas');
	ELSIF SQLCODE=-02290 THEN
		RAISE_APPLICATION_ERROR(-20403, 'la durée hebdomadaire d’un employé doit être inférieure ou égale à 35h');
	ELSIF SQLCODE=-1438 THEN
		RAISE_APPLICATION_ERROR(-20404, 'Une valeur (nombre) dépasse le nombre de caractères autorisés');
	ELSIF SQLCODE=-12899 THEN
		RAISE_APPLICATION_ERROR(-20405, 'Une valeur (chaine de caractère) dépasse le nombre de caractères autorisés');
	ELSIF SQLCODE=-20304 THEN
		RAISE_APPLICATION_ERROR(-20406, 'Le salaire de cet employé dépasse celui de son chef de service');
	ELSE
		RAISE_APPLICATION_ERROR (-20999,'Erreur inconnue'||SQLcode);
	END IF;
END;

END;
```

Test de la procedure

```sql
-- TESTS
EXECUTE MAJ.CREER_EMPLOYE(20, 'Boby', 31, 1, 2000);
EXECUTE MAJ.CREER_EMPLOYE(100, 'boby', 31, 66, 2000);
EXECUTE MAJ.CREER_EMPLOYE(100, 'Boby', 66, 1, 2000);
EXECUTE MAJ.CREER_EMPLOYE(100, 'Boby', 31, 1000000000000000000000000000000000000000, 2000);
EXECUTE MAJ.CREER_EMPLOYE(100, 'Jugemu Jugemu Unko Nageki Ototoi no Shin-chan no Pantsu Shinpachi no Jinsei Barumunku Fezarion Aizakku Shunaidaa Sanbun no Ichi no Junjou na Kanjou no Nokotta Sanbun no Ni wa Sakamuke ga Kininaru Kanjou Uragiri wa Boku no Namae wo Shitteiru you de Shiranai no wo Boku wa Shitteiru Rusu Surume Medaka Kazunoko Koedame Medaka... Kono Medaka wa Sakki to Chigau Yatsu Dakara Ikeno Medaka no Hou Dakara Raayu Yuuteimiyaoukimukou Pepepepepepepepepepepepe Bichiguso Maru', 31, 1, 2000);
EXECUTE MAJ.CREER_EMPLOYE(100, 'boby', 31, 1, 6666);
```
## Exercice 2
Chaque procédure doit traiter les erreurs provoquées par oracle(voir TD1) ou les erreurs qui proviennent de vos trigger (voir TD2). Vous devriez créer un nouveau tableau pour chacune des procédures.

Compléter le Package spécification(package MAJ) avec les opérations suivantes : Chaque procédure doit être rajoutée au Package MAJ et au package BODY MAJ.
- Modification de la durée hebdomadaire de la table "employe".
- Modification du salaire d’un employé.
- Modification de la durée de la table travail correspondant à un `nuempl` et `nuproj`.
- Insertion d’un enregistrement dans la table travail.
- Ajout d’une enregistrement dans la table service. Dans ce cas vous affectez le chef dans ce service avec un insert ou un update d’un employé qui existe déjà.

Package MAJ
```sql
CREATE OR REPLACE PACKAGE MAJ is
	PROCEDURE CREER_EMPLOYE (LE_NUEMPL IN NUMBER, LE_NOMEMPL IN VARCHAR2, LE_HEBDO IN NUMBER, LE_AFFECT IN NUMBER,LE_SALAIRE IN NUMBER);
	PROCEDURE MODIF_DUREE_HEBDO (NOUV_HEBDO IN NUMBER, NUEMP_CIBLE In NUMBER);
	PROCEDURE MODIF_SALAIRE (SALAIRE_NOUV IN NUMBER, EMP IN NUMBER);
	PROCEDURE MODIF_TRAVAIL (NOUV_DUREE IN NUMBER, PROJ IN NUMBER, EMP IN NUMBER);
	PROCEDURE INSERT_TRAVAIL (NOUV_EMP IN NUMBER, NOUV_NUPROJ IN NUMBER, NOUV_DUREE IN NUMBER);
	PROCEDURE INSERT_SERVICE (NOUV_NUSERV IN NUMBER, NOUV_NOMSERV IN VARCHAR, NOUV_CHEF IN NUMBER);
END;
```

**Modification de la durée hebdomadaire de la table "employe".**

| Test                              | Code erreur oracle -> Code d'erreur retourné | Message                                                                                                                                  |
| --------------------------------- | -------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| Aucune ligne trouvée              | -20410 -> -20410                             | Employé inexistant                                                                                                                       |
| Durée hebdomadaire > 35h          | -20102 -> -20403                             | La durée hebdomadaire d'un employé doit être inférieure ou égale à 35h et ne peut être augmentée                                          |
| Taille du nombre (dépassement)     | -1438 -> -20404                              | Une valeur (nombre) dépasse le nombre de caractères autorisés                                                                             |
| Valeur inattendue                 | -20999 -> -20999                             | Erreur inconnue                                                                                                                          |

Procedure
```sql
PROCEDURE MODIF_DUREE_HEBDO(NOUV_HEBDO IN NUMBER, NUEMP_CIBLE In NUMBER) is
BEGIN
	UPDATE EMPLOYE SET HEBDO = NOUV_HEBDO WHERE NUEMPL = NUEMP_CIBLE;
	IF SQL%notfound = TRUE THEN ROLLBACK;
		RAISE_APPLICATION_ERROR(-20410,'aucune lignes trouvé');
	END IF;
	COMMIT;

	EXCEPTION
	WHEN OTHERS THEN
	ROLLBACK;
	IF SQLCODE =-20102 THEN
		RAISE_APPLICATION_ERROR (-20403, 'la durée hebdomadaire d"un employe doit être inférieur ou égale à 35h et ne peut être augmenté');
	ELSIF SQLCODE =- 1438 THEN
		RAISE_APPLICATION_ERROR(-20404, 'Une valeur (nombre) dépasse le nombre de caractères autorisés');
	ELSIF SQLCODE = -20410 THEN
		RAISE_APPLICATION_ERROR(-20410, 'employé inexsitant');
	ELSE
		RAISE_APPLICATION_ERROR (-20999,'Erreur inconnue'||SQLcode);
	END IF;
END;
```
Tests
```sql
EXECUTE MAJ.MODIF_DUREE_HEBDO(20, 23); -- passe
EXECUTE MAJ.MODIF_DUREE_HEBDO(30, 23); -- augmente hebdo
EXECUTE MAJ.MODIF_DUREE_HEBDO(20, 99); -- employe inexistant
EXECUTE MAJ.MODIF_DUREE_HEBDO(40, 23); -- dépasse la limite de 35h
EXECUTE MAJ.MODIF_DUREE_HEBDO(1234567890, 23); -- dépassement de la capacité du nombre
```
**Modification du salaire d’un employé.**

| Test                               | Code erreur oracle -> Code d'erreur retourné | Message                                                                                               |
| ---------------------------------- | -------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| Aucune ligne trouvée               | -20410 -> -20410                             | Employé inexistant                                                                                     |
| Salaire non diminuable             | -20101 -> -20407                             | Salaire non diminuable                                                                                 |
| Taille du nombre (dépassement)    | -1438 -> -20404                              | Une valeur (nombre) dépasse le nombre de caractères autorisés                                          |
| Valeur inattendue                  | -20999 -> -20999                             | Erreur inconnue                                                                                       |

Procedure
```sql
PROCEDURE MODIF_SALAIRE(SALAIRE_NOUV IN NUMBER,EMP IN NUMBER) is
BEGIN
	UPDATE EMPLOYE SET SALAIRE = SALAIRE_NOUV WHERE NUEMPL = EMP;
	IF SQL%notfound = TRUE THEN ROLLBACK;
		RAISE_APPLICATION_ERROR(-20410,'aucune lignes trouvé');
	END IF;
	COMMIT;

	EXCEPTION
	WHEN OTHERS THEN
	ROLLBACK;
	IF SQLCODE =-20101 THEN
		RAISE_APPLICATION_ERROR (-20407, 'salaire non diminuable');
	ELSIF SQLCODE =- 1438 THEN
		RAISE_APPLICATION_ERROR(-20404, 'Une valeur (nombre) dépasse le nombre de caractères autorisés');
	ELSIF SQLCODE = -20410 THEN
		RAISE_APPLICATION_ERROR(-20410, 'employé inexsitant');
	ELSE
		RAISE_APPLICATION_ERROR (-20999,'Erreur inconnue'||SQLcode);
	END IF;
END;
```
Test
```sql
EXECUTE MAJ.MODIF_SALAIRE(3501, 20); -- passe
EXECUTE MAJ.MODIF_SALAIRE(0, 20); -- salaire non diminuable
EXECUTE MAJ.MODIF_SALAIRE(3501, 99); -- employe inexistant
EXECUTE MAJ.MODIF_SALAIRE(1234567890, 20); -- dépassement de la capacité du nombre
```

**Modification de la durée de la table travail correspondant à un `nuempl` et `nuproj`.**

| Test                           | Code erreur oracle -> Code d'erreur retourné | Message                                                        |
| ------------------------------ | -------------------------------------------- | -------------------------------------------------------------- |
| Passage normal                 | - Aucune erreur                              | Mise à jour réussie du temps de travail.                       |
| Temps de travail invalide      | -20301 -> -20408                             | Temps de travail invalide.                                     |
| Aucune ligne trouvée (employé) | -20410 -> -20410                             | Aucune ligne trouvée.                                          |
| Aucune ligne trouvée (projet)  | -20410 -> -20410                             | Aucune ligne trouvée.                                          |
| Projet inexistant              | -2291 -> -20402                              | Le projet n’existe pas.                                        |
| Dépassement de taille          | -1438 -> -20404                              | Une valeur (nombre) dépasse le nombre de caractères autorisés. |
| Erreur inconnue                | -20999                                       | Erreur inconnue.                                               |

Procedure
```sql
PROCEDURE MODIF_TRAVAIL(NOUV_DUREE IN NUMBER, PROJ IN NUMBER, EMP IN NUMBER) is
BEGIN
	UPDATE TRAVAIL t SET DUREE = NOUV_DUREE WHERE (t.nuproj = PROJ) AND (t.nuempl = EMP) ;
	IF SQL%notfound = TRUE THEN ROLLBACK;
		RAISE_APPLICATION_ERROR(-20410,'aucune ligne trouvé');
	END IF;
	COMMIT;

	EXCEPTION
	WHEN OTHERS THEN
	ROLLBACK;
	IF SQLCODE=-20301 THEN
		RAISE_APPLICATION_ERROR (-20408, 'temps de travail invalide');
	ELSIF SQLCODE =- 1438 THEN
		RAISE_APPLICATION_ERROR(-20404, 'Une valeur (nombre) dépasse le nombre de caractères autorisés');
	ELSIF SQLCODE=-2291 THEN
		RAISE_APPLICATION_ERROR(-20402, 'Le projet n’existe pas');
	ELSIF SQLCODE = -20410 THEN
		RAISE_APPLICATION_ERROR(-20410, 'aucune ligne trouvé');
	ELSE
		RAISE_APPLICATION_ERROR (-20999,'Erreur inconnue'||SQLcode);
	END IF;
END;
```
Test
```sql
EXECUTE MAJ.MODIF_TRAVAIL(5, 237, 23); -- passe
EXECUTE MAJ.MODIF_TRAVAIL(50, 237, 23); -- temps de travail invalide
EXECUTE MAJ.MODIF_TRAVAIL(5, 999, 23); -- le projet n’existe pas
EXECUTE MAJ.MODIF_TRAVAIL(1234567890, 237, 23); -- dépassement de la capacité du nombre
EXECUTE MAJ.MODIF_TRAVAIL(5, 200, 23); -- aucune ligne trouvé
EXECUTE MAJ.MODIF_TRAVAIL(5, 200, 99); -- aucune ligne trouvé
```
**Insertion d’un enregistrement dans la table travail.**

| Test                        | Code erreur oracle -> Code d'erreur retourné | Message                                                              |
| --------------------------- | -------------------------------------------- | -------------------------------------------------------------------- |
| `Temps de travail invalide` | -20301 -> -20408                             | Temps de travail invalide.                                           |
| `Enregistrement existant`   | -00001 -> -20401                             | Un enregistrement avec ce numéro d’employé et de projet existe déjà. |
| `Projet inexistant`         | -2291 -> -20402                              | Le projet auquel il est affecté n’existe pas.                        |
| `Dépassement de taille`     | -1438 -> -20404                              | Une valeur (nombre) dépasse le nombre de caractères autorisés.       |
| `Aucune ligne trouvée`      | -20410 -> -20410                             | Aucune ligne trouvée.                                                |
| `Erreur inconnue`           | -20999                                       | Erreur inconnue.                                                     |

Procedure
```sql
PROCEDURE INSERT_TRAVAIL (NOUV_EMP IN NUMBER, NOUV_NUPROJ IN NUMBER, NOUV_DUREE IN NUMBER) is
BEGIN
	INSERT INTO TRAVAIL VALUES(NOUV_EMP, NOUV_NUPROJ, NOUV_DUREE);
	IF SQL%notfound = TRUE THEN ROLLBACK;
		RAISE_APPLICATION_ERROR(-20410,'aucune ligne trouvé');
	END IF;
	COMMIT;

	EXCEPTION
	WHEN OTHERS THEN
	ROLLBACK;
	IF SQLCODE =-20301 THEN
		RAISE_APPLICATION_ERROR (-20408, 'temps de travail invalide');
	ELSIF SQLCODE = -00001 THEN
		RAISE_APPLICATION_ERROR(-20401, 'Un enregistrement avec ce numéro d employé et de projet existe déjà');
	ELSIF SQLCODE = -02291 THEN
		RAISE_APPLICATION_ERROR(-20402, 'Le projet auquel il est affecté n’existe pas');
	ELSIF SQLCODE = -01438 THEN
		RAISE_APPLICATION_ERROR(-20404, 'Une valeur (nombre) dépasse le nombre de caractères autorisés');
	ELSIF SQLCODE = -20410 THEN
		RAISE_APPLICATION_ERROR(-20410, 'aucune ligne trouvé');
	ELSE
		RAISE_APPLICATION_ERROR (-20999,'Erreur inconnue'||SQLcode);
	END IF;
END;
```
Test
```sql
EXECUTE MAJ.INSERT_TRAVAIL(23, 103, 50); -- Insertion réussie
EXECUTE MAJ.INSERT_TRAVAIL(23, 103, 0); -- temps de travail invalide
EXECUTE MAJ.INSERT_TRAVAIL(23, 103, 50); -- Tentative d'insertion d'un enregistrement existant
EXECUTE MAJ.INSERT_TRAVAIL(23, 999, 50); -- Le projet auquel il est affecté n’existe pas
EXECUTE MAJ.INSERT_TRAVAIL(23, 103, 999999); -- Une valeur (nombre) dépasse le nombre de caractères autorisés
EXECUTE MAJ.INSERT_TRAVAIL(99, 999, 50); -- Aucune ligne trouvée (employé et projet inexistants)
```

**Ajout d’une enregistrement dans la table service. Dans ce cas vous affectez le chef dans ce
service avec un insert ou un update d’un employé qui existe déjà.**

| Test                           | Code erreur oracle -> Code d'erreur retourné | Message                                                                    |
| ------------------------------ | -------------------------------------------- | -------------------------------------------------------------------------- |
| Chef inexistant                | -2291 -> -20402                              | Le chef ou le service n’existe pas                                         |
| Dépassement de taille (nom)    | -12899 -> -20405                             | Une valeur (chaine de caractère) dépasse le nombre de caractères autorisés |
| Dépassement de taille (NUSERV) | -1438 -> -20404                              | Une valeur (nombre) dépasse le nombre de caractères autorisés              |
| Aucune ligne trouvée           | -20410                                       | Aucune ligne trouvée                                                       |
| Erreur inconnue                | -20999                                       | Erreur inconnue avec code d'erreur spécifique                              |

Procedure
```sql
PROCEDURE INSERT_SERVICE(NOUV_NUSERV IN NUMBER, NOUV_NOMSERV IN VARCHAR, NOUV_CHEF IN NUMBER) is
BEGIN
	INSERT INTO SERVICE VALUES(NOUV_NUSERV, NOUV_NOMSERV, NOUV_CHEF);
	IF SQL%notfound = TRUE THEN ROLLBACK;
		RAISE_APPLICATION_ERROR(-20410,'aucune ligne trouvé');
	END IF;
	COMMIT;

	EXCEPTION
	WHEN OTHERS THEN
	ROLLBACK;
	IF SQLCODE = -20410 THEN
		RAISE_APPLICATION_ERROR(-20410, 'aucune ligne trouvé');
	ELSIF SQLCODE=-2291 THEN
		RAISE_APPLICATION_ERROR(-20402, 'Le chef ou le service n’existe pas');
	ELSIF SQLCODE =- 1438 THEN
		RAISE_APPLICATION_ERROR(-20404, 'Une valeur (nombre) dépasse le nombre de caractères autorisés');
	ELSIF SQLCODE =- 12899 THEN
		RAISE_APPLICATION_ERROR(-20405, 'Une valeur (chaine de caractère) dépasse le nombre de caractères autorisés');
	ELSE
		RAISE_APPLICATION_ERROR (-20999,'Erreur inconnue'||SQLcode);
	END IF;
END;
```
Test
```sql
EXECUTE MAJ.INSERT_SERVICE(6, 'nouveau_service', 20); -- Insertion réussie
EXECUTE MAJ.INSERT_SERVICE(7, 'nouveau_service', 99); -- Le chef ou le service n’existe pas
EXECUTE MAJ.INSERT_SERVICE(8, 'a'*300, 20); -- Une valeur (chaine de caractère) dépasse le nombre de caractères autorisés
EXECUTE MAJ.INSERT_SERVICE(99999999999, 'nouveau_service', 20); -- Une valeur (nombre) dépasse le nombre de caractères autorisés
EXECUTE MAJ.INSERT_SERVICE(99, 'service_inexistant', 99); -- Aucune ligne trouvée (service et chef inexistants)
EXECUTE MAJ.INSERT_SERVICE(NULL, NULL, NULL); -- Cela peut provoquer une erreur inconnue selon les contraintes.
```
# Annexe
## TD1
### Base de données

| Nom table | clefs primaire | autres colones                  |
| --------- | -------------- | ------------------------------- |
| employe   | nuempl         | nomempl, hebdo, affect, salaire |
| service   | nuserv         | nomserv, chef                   |
| projet    | nuproj         | nomproj, resp                   |
| travail   | nuempl, nuproj | duree                           |
| concerne  | nuserv, nuproj |                                 |

### Création des tables
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
### Contraintes d'intégrité
```sql
-- Clefs Etrangères
ALTER TABLE travail add CONSTRAINT FK_employe_travail FOREIGN KEY (nuempl) REFERENCES employe(nuempl);
ALTER TABLE travail add CONSTRAINT FK_projet_travail FOREIGN KEY (nuproj) REFERENCES projet(nuproj);
ALTER TABLE concerne add CONSTRAINT FK_service_concerne FOREIGN KEY (nuserv) REFERENCES service(nuserv);
ALTER TABLE concerne add CONSTRAINT FK_projet_concerne FOREIGN KEY (nuproj) REFERENCES projet(nuproj);
```

Le chef d'un service est un employé affecté au même service. Attention les contraintes de clé étrangère `Affect` et `Chef` se croisent. Dans ce cas, vous devriez créer une contrainte unique `(nuempl, affect)` et une contrainte différée (initially deferred) de la table `service` `(chef,nuserv)` vers cette contrainte unique `(nuempl, affect)`.
```sql
-- Ajouter la contrainte unique sur (nuempl, affect)
ALTER TABLE employe
ADD CONSTRAINT UNQ_employe_affect UNIQUE (nuempl, affect);

-- Ajouter la contrainte étrangère pour "chef" différée
ALTER TABLE service
ADD CONSTRAINT FK_chef FOREIGN KEY (chef, nuserv)
REFERENCES employe (nuempl, affect) DEFERRABLE INITIALLY DEFERRED;

-- Ajouter la contrainte étrangère pour "affect"
ALTER TABLE employe
ADD CONSTRAINT FK_affect FOREIGN KEY (affect)
REFERENCES service (nuserv);
```

Un responsable de projet doit travailler sur le projet. Dans ce cas vous n'avez pas besoin de créer la contrainte Unique `(nuempl, nuproj)` de la table `travail`, car les deux attributs forment déjà la clé primaire. Par contre la contrainte `(resp, nuproj)` est une clé étrangère différée vers la clé primaire de la table `travail`.
```sql
ALTER TABLE projet
ADD CONSTRAINT FK_resp FOREIGN KEY (resp, nuproj)
REFERENCES travail (nuempl, nuproj) DEFERRABLE INITIALLY DEFERRED;
```

La durée (temps de travail) hebdomadaire d'un employé est inférieure ou égale à 35h.
```sql
ALTER TABLE employe ADD CONSTRAINT temps_trav_max CHECK (hebdo <= 35);
```

#### Table des contraintes

| nom table | nom contrainte      | type de contrainte | différer ou non |
| --------- | ------------------- | ------------------ | --------------- |
| Employe   | PK_employe          | Clé primaire       | non             |
| Service   | PK_service          | Clé primaire       | non             |
| Projet    | PK_projet           | Clé primaire       | non             |
| Travail   | PK_travail          | Clé primaire       | non             |
| Concerne  | PK_concerne         | Clé primaire       | non             |
| Travail   | FK_employe_travail  | Clé étrangère      | non             |
| Travail   | FK_projet_travail   | Clé étrangère      | non             |
| Concerne  | FK_service_concerne | Clé étrangère      | non             |
| Concerne  | FK_projet_concerne  | Clé étrangère      | non             |
| Employe   | UNQ_employe_affect  | Clé unique         | non             |
| Service   | FK_chef             | Clé étrangère      | oui             |
| Employe   | FK_affect           | Clé étrangère      | non             |
| Projet    | FK_resp             | Clé étrangère      | oui             |
| Employe   | temps_trav_max      | Check              | non             |

#### Tests des contraintes

| nom table           | requête                                                                                              | code d'erreur | message d'erreur                               |
| ------------------- | ---------------------------------------------------------------------------------------------------- | ------------- | ---------------------------------------------- |
| PK_employe          | `INSERT INTO employe VALUES (20, 'toto', 35, 1, 1500); -- Insertion d'un employé qui existe déjà`    | ORA-00001     | Violation de la contrainte PK_employe          |
| PK_service          | `INSERT INTO service VALUES (1, 'service_test'); -- Insertion d'un service qui existe déjà`          | ORA-00001     | Violation de la contrainte PK_service          |
| PK_projet           | `INSERT INTO projet VALUES (103, 'projet_test', 200); -- Insertion d'un projet déjà existant`        | ORA-00001     | Violation de la contrainte PK_projet           |
| PK_travail          | `INSERT INTO travail VALUES (20, 492, 10); -- Insertion d'une ligne qui existe déjà dans travail`    | ORA-00001     | Violation de la contrainte PK_travail          |
| FK_employe_travail  | `INSERT INTO travail VALUES (999, 492, 10); -- Insertion avec un employé qui n'existe pas`           | ORA-02291     | Violation de la contrainte FK_employe_travail  |
| FK_projet_travail   | `INSERT INTO travail VALUES (20, 999, 10); -- Insertion avec un projet qui n'existe pas`             | ORA-02291     | Violation de la contrainte FK_projet_travail   |
| FK_service_concerne | `INSERT INTO concerne VALUES (99, 103); -- Service inexistant`                                       | ORA-02291     | Violation de la contrainte FK_service_concerne |
| FK_projet_concerne  | `INSERT INTO concerne VALUES (1, 999); -- Projet inexistant`                                         | ORA-02291     | Violation de la contrainte FK_projet_concerne  |
| UNQ_employe_affect  | `INSERT INTO employe VALUES (30, 'jeSuisNouveau', 35, 1, 1500); -- Affectation unique déjà utilisée` | ORA-00001     | Violation de la contrainte UNQ_employe_affect  |
| FK_chef             | `UPDATE service SET chef = 999 WHERE nuserv = 1; -- Chef inexistant`                                 | ORA-02291     | Violation de la contrainte FK_chef             |
| FK_affect           | `UPDATE employe SET affect = 99 WHERE nuempl = 20; -- Service affecté inexistant`                    | ORA-02291     | Violation de la contrainte FK_affect           |
| FK_resp             | `UPDATE projet SET resp = 999 WHERE nuproj = 100; -- Responsable inexistant dans travail`            | ORA-02291     | Violation de la contrainte FK_resp             |
| temps_trav_max      | `INSERT INTO employe VALUES (40, 'heurs_supp', 40, 1, 1500); -- Hebdo supérieur à 35h`               | ORA-02290     | Violation de la contrainte temps_trav_max      |

### Mise à jour des données
La table employé contient un attribut salaire de type number. Vous rajoutez cet attribut  
avec la requête suivante : `alter table employe add salaire number;`

Faites les modifications suivantes, à chaque fois avec une seule requête:  
- Les responsables de projet gagnent au moins 2500 Euros,  
- Les chefs de service gagnent au moins 3500 euros.  
- Les autres employés gagnent un salaire < 2000 euros.
```sql
ALTER TABLE EMPLOYE ADD
SALAIRE NUMBER;

UPDATE EMPLOYE
SET SALAIRE = 2500
WHERE NUEMPL IN (SELECT RESP FROM PROJET);

UPDATE EMPLOYE
SET SALAIRE = 3500
WHERE NUEMPL IN (SELECT CHEF FROM SERVICE);

UPDATE EMPLOYE
SET SALAIRE = 1999
WHERE NUEMPL NOT IN ((SELECT CHEF FROM SERVICE) UNION (SELECT RESP FROM PROJET));

COMMIT;
```

La somme des durées d'un employé (de la table travail) doit être inférieur à la durée hebdomadaire `(Sum(duree) <= hebdo)`. Affichez les employés qui ne respectent pas cette contrainte. Vous modifiez les données de la table travail(update sur la durée) ou la table employé (update sur hebdo)jusqu'à ce que la requête précédente retourne un ensemble vide.
```sql
SELECT e.NUEMPL, e.NOMEMPL, e.HEBDO, (SELECT SUM(t.DUREE) FROM travail t WHERE t.NUEMPL = e.NUEMPL) AS total_duree
FROM employe e
WHERE (
	SELECT SUM(t.DUREE)
	FROM travail t
	WHERE t.NUEMPL = e.NUEMPL
) > e.HEBDO;

-- à été exécuté 5 fois pour atteindre les 35h hebdomadaire
UPDATE employe
SET hebdo = LEAST(hebdo + 5, 35) -- Augmente de 5 heures sans dépasser 35 heures
WHERE nuempl IN (
	SELECT e.NUEMPL
	FROM employe e
	WHERE (
		SELECT SUM(t.DUREE)
		FROM travail t
		WHERE t.NUEMPL = e.NUEMPL
	) > e.HEBDO
);

-- est éxécuté 1 fois pour réduire le temps de travail en dessous des 35h
UPDATE travail
SET duree = duree - 5
WHERE nuempl IN (
	SELECT e.NUEMPL
	FROM employe e
	WHERE (
		SELECT SUM(t.DUREE)
		FROM travail t
		WHERE t.NUEMPL = e.NUEMPL
	) > e.HEBDO
);
```
Un employé ne peut être responsable de plus de 3 projets. Idem que la contrainte précédente, trouvez les employés qui ne respectent pas la contrainte et modifiez les données.
```sql
SELECT e.NUEMPL, e.NOMEMPL, COUNT(p.NUPROJ) AS nb_projets
FROM employe e
JOIN projet p ON e.NUEMPL = p.RESP
GROUP BY e.NUEMPL, e.NOMEMPL
HAVING COUNT(p.NUPROJ) > 3;

-- Déjà bon pour le moment
```

Le chef d’un service gagne plus que les employés du service. Idem que la contrainte précédente, trouvez les employés qui ne respectent pas la contrainte et modifiez les données.
```sql
SELECT s.CHEF, e.NOMEMPL AS chef_nom, e.SALAIRE AS salaire_chef, e2.NUEMPL AS employe_id, e2.NOMEMPL AS employe_nom, e2.SALAIRE AS salaire_employe
FROM service s
JOIN employe e ON s.CHEF = e.NUEMPL -- Chef du service
JOIN employe e2 ON e2.AFFECT = s.NUSERV -- Employés du même service
WHERE e.SALAIRE <= e2.SALAIRE AND e2.NUEMPL != e.NUEMPL; -- Ne compare pas le chef à lui-même

-- déjà bon mais voici un update au cas où
UPDATE employe
SET salaire = salaire + 1
WHERE nuempl IN (
	SELECT s.CHEF
	FROM service s
	JOIN employe e ON s.CHEF = e.NUEMPL -- Chef du service
	JOIN employe e2 ON e2.AFFECT = s.NUSERV -- Employés du même service
	WHERE e.SALAIRE <= e2.SALAIRE
		AND e2.NUEMPL != e.NUEMPL -- Ne compare pas le chef à lui-même
);
```

Un service ne peut être concerné par plus de 3 projets.
```sql
SELECT s.NUSERV, s.NOMSERV, COUNT(c.NUPROJ) AS nb_projets
FROM service s
JOIN concerne c ON s.NUSERV = c.NUSERV
GROUP BY s.NUSERV, s.NOMSERV
HAVING COUNT(c.NUPROJ) > 3;

DELETE FROM concerne
WHERE (nuserv, nuproj) IN (
	SELECT nuserv, nuproj
	FROM (
		SELECT c.nuserv, c.nuproj,
			ROW_NUMBER() OVER (PARTITION BY c.nuserv ORDER BY c.nuproj) AS rn
		FROM concerne c
	) sub
	WHERE sub.rn > 3
);
```

## TD2
### Exercice 1 : Triggers de type For each row et utilisation de ":NEW" et ":OLD"

Ecrire un trigger de type for each row qui interdit la diminution du salaire d'un employé. Ce trigger se déclenche après la modification du salaire.

|  nom trigger  | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :-----------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
| modif_salaire |         after         |         update         |  employé  |           oui            |

Code du trigger
```sql
CREATE or REPLACE TRIGGER modif_salaire
AFTER UPDATE OF salaire on EMPLOYE for EACH ROW
BEGIN
	IF :OLD.salaire > :NEW.salaire THEN
		RAISE_APPLICATION_ERROR(-20101, 'salaire non diminuable');
	END IF;
END;
```
Code de test
```sql
UPDATE EMPLOYE SET salaire = 0 WHERE nuempl = 20; -- ne passe pas le trigger
UPDATE EMPLOYE Set salaire = 3501 WHERE nuempl = 23; -- passe le trigger
```

**Il y a une autre contrainte qui n'est pas spécifiée "la durée hebdomadaire d'un employé ne peut pas augmenter", elle n'est pas descriptible. Vous écrivez le trigger nécessaire à la vérification de cette contrainte**

|    nom trigger     | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :----------------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
| augmentation_hebdo |         after         |         update         |  employé  |           oui            |
Code du trigger
```sql
CREATE or REPLACE TRIGGER augmentation_hebdo
AFTER UPDATE OF hebdo on EMPLOYE FOR EACH ROW
BEGIN
	IF :OLD.hebdo < :NEW.hebdo THEN
		RAISE_APPLICATION_ERROR(-20102, 'hebdo non augmentable');
	END IF;
END;
```
Code de test
```sql
UPDATE EMPLOYE SET hebdo = 30 WHERE nuempl = 23; -- active le trigger
UPDATE EMPLOYE SET hebdo = 20 WHERE nuempl = 23; -- active pas le trigger
```

### Exercice 2 : Trigger de type : Delete from T2 where a not in (select a from T1) / D’autres types de solutions ne sont pas acceptés pour cette exercice.

La spécification de l'opération supprimer_employe impose que la suppression d'un employé soit accompagnée de la suppression des lignes de travail correspondantes. Mettez en place un trigger table qui le fait. (pas de problème si on a déclaré "deferred" la contrainte FK_employe de la table travail vers la table employe "les employés de travail existent").
**Attention : la suppression des employés de la table travail n’est possible que si l’employé n’est pas chef de service ou responsable de projet**

|    nom trigger    | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :---------------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
| supprimer_employe |        before         |         delate         |  employe  |           oui            |

Code du trigger
```sql
CREATE OR REPLACE TRIGGER supprimer_employe
BEFORE DELETE ON employe
FOR EACH ROW
DECLARE
	emp_chef_count NUMBER;
	emp_resp_count NUMBER;
BEGIN
	-- Vérification si l'employé est chef de service
	SELECT COUNT(*) INTO emp_chef_count
	FROM service
	WHERE chef = :OLD.nuempl;
	-- Vérification si l'employé est responsable de projet
	SELECT COUNT(*) INTO emp_resp_count
	FROM projet
	WHERE resp = :OLD.nuempl;
	-- Si l'employé est chef de service ou responsable de projet, empêcher la suppression
	IF emp_chef_count > 0 THEN
		RAISE_APPLICATION_ERROR(-20201, 'erreur: employé est chef de service.');
	ELSIF emp_resp_count > 0 THEN
		RAISE_APPLICATION_ERROR(-20201, 'erreur: employé est responsable de projet.');
	END IF;
	-- Suppression des lignes associées à l'employé dans la table travail
	DELETE FROM travail WHERE nuempl = :OLD.nuempl;
END;
```
Code de test
```sql
-- marche
DELETE FROM employe WHERE nuempl = 37; -- employé n'étant ni chef ni responsable.
DELETE FROM employe WHERE nuempl = 48; -- employé avec des lignes dans `travail`.
-- retourne une erreur (attendu)
DELETE FROM employe WHERE nuempl = 17; -- employé chef de service.
DELETE FROM employe WHERE nuempl = 20; -- employé responsable de projet.
```

La spécification de l'opération supprimer_projet impose que la suppression d'un projet soit accompagnée de la suppression des lignes de travail et de la table concerne correspondantes. Mettez en place un trigger table qui effectue cela. (pas de problème si on a déclaré "deferred" la contrainte FK_nuproj de la table travail et de la table concerne vers la table projet).

|   nom trigger    | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :--------------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
| supprimer_projet |        before         |         delete         |  projet   |           oui            |
Code du trigger
```sql
CREATE OR REPLACE TRIGGER supprimer_projet
BEFORE DELETE ON projet
FOR EACH ROW
BEGIN
	-- Suppression des lignes dans la table travail associées au projet
	DELETE FROM travail WHERE nuproj = :OLD.nuproj;
	-- Suppression des lignes dans la table concerne associées au projet
	DELETE FROM concerne WHERE nuproj = :OLD.nuproj;
END;
```
Code de test
```sql
DELETE FROM projet WHERE nuproj = 160; -- projet ayant des lignes dans travail et concerne
DELETE FROM projet WHERE nuproj = 492; -- projet ayant uniquement des lignes dans travail.
DELETE FROM projet WHERE nuproj = 103; -- projet responsable par un employé spécifique.
```
### Exercice 3 : Pour les triggers qui suivent, vous faites des M.A.J(insert, update) de la base de données qui déclenchent les différentes erreurs.
*Tous les trigger de cette exercice sont de type `select *` into suivi de l’instruction `raise_application_error`. Les autres solutions ne seront pas acceptées.*

Il y a une contrainte qui n'est pas spécifiée "la somme des durées de travail d'un employé ne doit pas excéder son temps de travail hebdomadaire", elle n'est pas descriptible. Vous écrivez le (ou les) trigger nécessaire à la vérification de cette contrainte. La contrainte à mettre en place SUM(duree)<=hebdo. 
- Quels sont les différentes opérations(update, insert,…) sur les tables `employe` ou `travail` qui vous amènent à ne pas respecter cette contrainte.
- Combien de trigger allez vous mettre en place ? Indiquez dans quel cas les trigger se déclenchent. Vous utilisez la requête suivante pour construire ce(s) trigger : `Select * from employe e where (select sum(duree) from travail t where e.nuempl=t.nuempl)> hebdo`

| nom trigger | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :---------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
|             |                       |                        |           |                          |

Code du trigger
```sql
CREATE OR REPLACE TRIGGER duree_trav_invalide
	AFTER INSERT OR UPDATE OF duree ON TRAVAIL
	DECLARE rec_t EMPLOYE%ROWTYPE;
BEGIN
	SELECT *
	INTO rec_t
	FROM EMPLOYE e
	WHERE (
		SELECT SUM(duree)
		FROM TRAVAIL t
		WHERE e.NUEMPL = t.NUEMPL
	) > e.HEBDO;

	RAISE_APPLICATION_ERROR(-20301, 'temps travail invalide');
	EXCEPTION
		WHEN no_data_found THEN null;
		WHEN too_many_rows THEN RAISE_APPLICATION_ERROR(-20301, 'temps travail invalide');
END;
```
Tests du trigger
```sql
UPDATE TRAVAIL SET DUREE = 50 WHERE NUEMPL = 23 AND NUPROJ = 237;
UPDATE TRAVAIL SET DUREE = 50 WHERE NUEMPL > 23;
UPDATE TRAVAIL SET DUREE = 5 WHERE NUEMPL = 23 AND NUPROJ = 237;
```

Ecrire un trigger qui vérifie la contrainte suivante: "un employé est responsable au plus sur 3 projets. Idem que la question précédente, vous utilisez la requête suivante pour construire votre trigger : `SELECT * FROM EMPLOYE e WHERE (SELECT COUNT(*) FROM PROJET p WHERE e.nuempl=p.resp)> 3;`

|  nom trigger  | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :-----------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
| resp_max_proj |         after         |    insert ou update    |  projet   |           non            |

Code du trigger
```sql
CREATE OR REPLACE TRIGGER resp_max_proj
	AFTER INSERT OR UPDATE OF resp ON PROJET
	DECLARE rec EMPLOYE%ROWTYPE;
BEGIN
	SELECT *
	INTO rec
	FROM EMPLOYE e
	WHERE (
		SELECT COUNT(*)
		FROM PROJET p
		WHERE e.nuempl=p.resp
	)> 3;

	RAISE_APPLICATION_ERROR(-20302, 'employe resp sur plus de 3 projets');
	EXCEPTION
		WHEN no_data_found THEN null;
		WHEN too_many_rows THEN RAISE_APPLICATION_ERROR(-20302, 'employe resp sur plus de 3 projets');
END;
```
Tests du trigger
```sql
INSERT INTO PROJET (NUPROJ, NOMPROJ, RESP) VALUES (100, 'percephone', 30); -- Créer un nouveau projet et y ajoute un employe déjà responsable 3 fois
INSERT INTO PROJET (NUPROJ, NOMPROJ, RESP) VALUES (100, 'hades', 20); -- Créer un nouveau projet et y ajoute un employe déjà responsable 1 fois
UPDATE PROJET SET RESP = 30 WHERE RESP = 20; -- assigne la responsabilité du projet de n20 à n30
UPDATE PROJET SET RESP = 57 WHERE RESP = 20; -- assigne la responsabilité du projet de n20 à n57
```

Ecrire un trigger qui vérifie la contrainte suivante : "un service ne peut être concerné par plus de 3 projets"

|   nom trigger    | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :--------------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
| serv_max_concern |         after         |    insert ou update    | concerne  |           non            |

Code du trigger
```sql
CREATE OR REPLACE TRIGGER serv_max_concern
	AFTER INSERT OR UPDATE OF NUSERV ON CONCERNE
	DECLARE rec CONCERNE%ROWTYPE;
BEGIN
	SELECT *
	INTO rec
	FROM CONCERNE c
	WHERE (
		SELECT COUNT(*)
		FROM CONCERNE cc
		WHERE cc.NUSERV = c.NUSERV
	)>3;
	RAISE_APPLICATION_ERROR(-20303, 'un service ne peut être concerné par plus de 3 projets');
	EXCEPTION
		WHEN no_data_found THEN null;
		WHEN too_many_rows THEN RAISE_APPLICATION_ERROR(-20303, 'un service ne peut être concerné par plus de 3 projets');
END;
```
Tests du trigger
```sql
UPDATE CONCERNE SET NUSERV = 1 WHERE NUPROJ = 160 AND NUSERV = 2; -- ajoute un projet à un service qui possède déjà 3 autre service
UPDATE CONCERNE SET NUSERV = 5 WHERE NUPROJ = 103 AND NUSERV = 1; -- ajoute un projet à un service qui possède moins 3 autre service
INSERT INTO CONCERNE (NUSERV, NUPROJ) VALUES (1, 160); -- ajoute un projet à un service qui possède déjà 3 autre service
INSERT INTO CONCERNE (NUSERV, NUPROJ) VALUES (5, 103); -- ajoute un projet à un service qui possède moins 3 autre service
```

Ecrire un trigger qui vérifie la contrainte suivante : "un chef de service gagne plus que les employés de son service".

|  nom trigger   | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :------------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
| chef_plus_paye |         after         |    insert ou update    |  employe  |           non            |

Code du trigger
```sql
CREATE OR REPLACE TRIGGER chef_plus_paye
	AFTER INSERT OR UPDATE OF SALAIRE ON EMPLOYE
	DECLARE rec SERVICE%ROWTYPE;
BEGIN
	SELECT s.NUSERV, s.NOMSERV, s.CHEF
	INTO rec
	FROM service s
	JOIN employe e ON s.CHEF = e.NUEMPL -- Chef du service
	JOIN employe e2 ON e2.AFFECT = s.NUSERV -- Employés du même service
	WHERE e.SALAIRE <= e2.SALAIRE AND e2.NUEMPL != e.NUEMPL; -- Ne compare pas le chef à lui-même

	RAISE_APPLICATION_ERROR(-20304, 'un chef de service ne gagne pas plus que les employés de son service');
	EXCEPTION
		WHEN no_data_found THEN null;
		WHEN too_many_rows THEN RAISE_APPLICATION_ERROR(-20304, 'un chef de service ne gagne pas plus que les employés de son service');
END;
```
Tests du trigger
```sql
UPDATE EMPLOYE SET SALAIRE = 4000 WHERE AFFECT = 1 AND NUEMPL != 41; -- augmente le salaire d'un employé pour depassez le chef
UPDATE EMPLOYE SET SALAIRE = 2500 WHERE AFFECT = 1 AND NUEMPL = 39; -- augmente le salaire d'un employé sans depassez le chef
INSERT INTO EMPLOYE (NUEMPL, NOMEMPL, HEBDO, AFFECT, SALAIRE) VALUES (1, 'ulysse', 10, 1, 4000); -- créer un employer mieux payer que son chef
INSERT INTO EMPLOYE (NUEMPL, NOMEMPL, HEBDO, AFFECT, SALAIRE) VALUES (1, 'ulysse', 10, 1, 1500); -- créer un employer moins bien payer que son chef
```

Ecrire un trigger qui vérifie la contrainte suivante : "un chef de service gagne plus que les employés responsables de projets".

|    nom trigger     | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :----------------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
| chef_paye_sup_resp |         after         |    insert ou update    |  employe  |           non            |

Code du trigger
```sql
CREATE OR REPLACE TRIGGER chef_paye_sup_resp
	AFTER INSERT OR UPDATE OF SALAIRE ON EMPLOYE
	DECLARE rec SERVICE%ROWTYPE;
BEGIN
	SELECT s.NUSERV, s.NOMSERV, s.CHEF
	INTO rec
	FROM SERVICE s
	JOIN EMPLOYE e ON s.CHEF = e.NUEMPL -- Chef du service
	JOIN EMPLOYE e2 ON e2.AFFECT = s.NUSERV -- Employés du même service
	JOIN PROJET p ON p.RESP = e2.NUEMPL -- Employé responsable de projet
WHERE e.SALAIRE <= e2.SALAIRE AND e2.NUEMPL != e.NUEMPL; -- Ne compare pas le chef à lui-même

	RAISE_APPLICATION_ERROR(-20305, 'un chef de service gagne plus que les employés responsables de projets');
	EXCEPTION
		WHEN no_data_found THEN null;
		WHEN too_many_rows THEN RAISE_APPLICATION_ERROR(-20305, 'un chef de service gagne plus que les employés responsables de projets');
END;
```
Tests du triggers
```sql
-- SELECT * FROM EMPLOYE WHERE AFFECT = 2 ORDER BY SALAIRE DESC;
UPDATE EMPLOYE SET SALAIRE = 4000 WHERE AFFECT = 2 AND NUEMPL = 57; -- augment le salaire d'un resp pour depasser le chef
UPDATE EMPLOYE SET SALAIRE = 3000 WHERE AFFECT = 2 AND NUEMPL = 57; -- augment le salaire d'un resp sans depasser le chef
-- pour les insert allez chercher un employé dejà resp ailleur
```


Est-il possible de regrouper les deux derniers « trigger »

|        nom trigger         | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :------------------------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
| chef_plus_paye_emp_et_resp |         after         |    insert ou update    |  employe  |           non            |

Code du trigger
```sql
CREATE OR REPLACE TRIGGER chef_plus_paye_emp_et_resp
	AFTER INSERT OR UPDATE OF SALAIRE ON EMPLOYE
	DECLARE rec SERVICE%ROWTYPE;
BEGIN
	-- trigger 1
	SELECT s.NUSERV, s.NOMSERV, s.CHEF
	INTO rec
	FROM SERVICE s
	JOIN EMPLOYE e ON s.CHEF = e.NUEMPL -- Chef du service
	JOIN EMPLOYE e2 ON e2.AFFECT = s.NUSERV -- Employés du même service
	WHERE e.SALAIRE <= e2.SALAIRE AND e2.NUEMPL != e.NUEMPL; -- Ne compare pas le chef à lui-même
	-- trigger 2
	SELECT s.NUSERV, s.NOMSERV, s.CHEF
	INTO rec
	FROM SERVICE s
	JOIN EMPLOYE e ON s.CHEF = e.NUEMPL -- Chef du service
	JOIN EMPLOYE e2 ON e2.AFFECT = s.NUSERV -- Employés du même service
	JOIN PROJET p ON p.RESP = e2.NUEMPL -- -- Employé responsable de projet
	WHERE e.SALAIRE <= e2.SALAIRE AND e2.NUEMPL != e.NUEMPL; -- Ne compare pas le chef à lui-même
	
	RAISE_APPLICATION_ERROR(-20306, 'un chef de service gagne plus que les employés de son service');
	EXCEPTION
		WHEN no_data_found THEN null;
		WHEN too_many_rows THEN RAISE_APPLICATION_ERROR(-20306, 'un chef de service gagne plus que les employés de son service');
END;
```
Tests du trigger
```sql
-- trigger 1
UPDATE EMPLOYE SET SALAIRE = 4000 WHERE AFFECT = 1 AND NUEMPL != 41; -- augmente le salaire d'un employé pour depassez le chef
UPDATE EMPLOYE SET SALAIRE = 2500 WHERE AFFECT = 1 AND NUEMPL = 39; -- augmente le salaire d'un employé sans depassez le chef
INSERT INTO EMPLOYE (NUEMPL, NOMEMPL, HEBDO, AFFECT, SALAIRE) VALUES (1, 'ulysse', 10, 1, 4000); -- créer un employer mieux payer que son chef
INSERT INTO EMPLOYE (NUEMPL, NOMEMPL, HEBDO, AFFECT, SALAIRE) VALUES (1, 'ulysse', 10, 1, 1500); -- créer un employer moins bien payer que son chef
-- trigger 2
UPDATE EMPLOYE SET SALAIRE = 4000 WHERE AFFECT = 2 AND NUEMPL = 57; -- augment le salaire d'un resp pour depasser le chef
UPDATE EMPLOYE SET SALAIRE = 3000 WHERE AFFECT = 2 AND NUEMPL = 57; -- augment le salaire d'un resp sans depasser le chef
-- pour les insert allez chercher un employé dejà resp ailleur
```
### Exercice 4

Lors d'augmentation de salaire ou d'embauche, l'entreprise veut enregistrer les employés (dans la table `EMPLOYE_ALERTE` idem que `EMPLOYE`) avec un salaire qui dépassent les 5000 euros. Ecrire un trigger qui permet de remplir cette table.

|      nom trigger       | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :--------------------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
| remplis_employe_alerte |         after         |    insert ou update    |  employe  |           oui            |

Code du trigger
```sql
CREATE table EMPLOYE_ALERTE AS SELECT * FROM EMPLOYE;
ALTER TABLE EMPLOYE_ALERTE ADD CONSTRAINT PK_employe_alerte PRIMARY KEY (nuempl);

CREATE OR REPLACE TRIGGER remplis_employe_alerte
	AFTER INSERT OR UPDATE OF SALAIRE ON EMPLOYE
	FOR EACH ROW WHEN ( NEW.SALAIRE > 5000)
BEGIN
	INSERT INTO EMPLOYE_ALERTE (NUEMPL, NOMEMPL, HEBDO, AFFECT, SALAIRE)
	VALUES (:NEW.NUEMPL, :NEW.NOMEMPL, :NEW.HEBDO, :NEW.AFFECT, :NEW.SALAIRE);

	EXCEPTION
		WHEN no_data_found THEN null;
END;
```
Tests du trigger
```sql
alter trigger "CHEF_PAYE_SUP_RESP" disable;
alter trigger "CHEF_PLUS_PAYE" disable;
alter trigger "CHEF_PLUS_PAYE_EMP_ET_RESP" disable;

INSERT INTO EMPLOYE (NUEMPL, NOMEMPL, HEBDO, AFFECT, SALAIRE) VALUES (1, 'cresus', 10, 1, 50000); -- créer un employer payé plus de 5000
INSERT INTO EMPLOYE (NUEMPL, NOMEMPL, HEBDO, AFFECT, SALAIRE) VALUES (2, 'diogene', 10, 1, 1); -- créer un employer payé moins de 5000
--UPDATE EMPLOYE SET SALAIRE = 50000 WHERE NUEMPL = 41; -- créer un employer payé plus de 5000
```
