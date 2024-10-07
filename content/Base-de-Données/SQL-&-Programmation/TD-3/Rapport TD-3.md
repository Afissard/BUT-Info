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
	PROCEDURE CREER_EMPLOYE (
		LE_NUEMPL IN NUMBER, LE_NOMEMPL IN VARCHAR2, LE_HEBDO IN NUMBER, LE_AFFECT IN NUMBER,LE_SALAIRE IN NUMBER
	) is
BEGIN
	SET TRANSACTION READ WRITE;
	INSERT INTO employe VALUES(LE_NUEMPL, LE_NOMEMPL, LE_HEBDO, LE_AFFECT,LE_SALAIRE);
	COMMIT;
	EXCEPTION WHEN OTHERS THEN ROLLBACK;
	IF SQLCODE=-00001 THEN
		ROLLBACK;
		RAISE_APPLICATION_ERROR(-20401, 'Un employe avec le meme numero existe deja');
	ELSIF SQLCODE=-2291 THEN
		ROLLBACK;
		RAISE_APPLICATION_ERROR(-20402, 'Le service auquel il est affecté n’existe pas');
	ELSIF SQLCODE=-02290 THEN
		ROLLBACK;
		RAISE_APPLICATION_ERROR(-20403, 'la durée hebdomadaire d’un employé doit être inférieure ou égale à 35h');
	ELSIF SQLCODE=-1438 THEN
		ROLLBACK;
		RAISE_APPLICATION_ERROR(-20404, 'Une valeur (nombre) dépasse le nombre de caractères autorisés');
	ELSIF SQLCODE=-12899 THEN
		ROLLBACK;
		RAISE_APPLICATION_ERROR(-20405, 'Une valeur (chaine de caractère) dépasse le nombre de caractères autorisés');
	ELSIF SQLCODE=-20101 THEN
		ROLLBACK;
		RAISE_APPLICATION_ERROR(-20405, 'Le salaire de cet employé dépasse celui de son chef de service'); --TODO
	ELSE
		ROLLBACK;
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
EXECUTE MAJ.CREER_EMPLOYE(100, 'boby', 31, 1, 66666); --TODO FIX
```
# Exercice 2
Chaque procédure doit traiter les erreurs provoquées par oracle(voir TD1) ou les erreurs qui proviennent de vos trigger (voir TD2). Vous devriez créer un nouveau tableau pour chacune des procédures.

Compléter le Package spécification(package MAJ) avec les opérations suivantes : Chaque procédure doit être rajoutée au Package MAJ et au package BODY MAJ.
- Modification de la durée hebdomadaire de la table « employe ».
- Modification du salaire d’un employé.
- Modification de la durée de la table travail correspondant à un ''nuempl'' et ''nuproj''.
- Insertion d’un enregistrement dans la table travail.

#TODO
