---
title: Rapport TD-3
draft: 
description: 
tags:
  - Base-de-Données
  - SQL
---
# Exercice 1
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
# Exercice 2
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
