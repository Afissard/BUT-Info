---
title: Rapport TD-2
draft: 
description: 
tags:
  - Base-de-Données
  - SQL
---
# Exercice 1 : Triggers de type For each row et utilisation de ":NEW" et ":OLD"

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

# Exercice 2 : Trigger de type : Delete from T2 where a not in (select a from T1) / D’autres types de solutions ne sont pas acceptés pour cette exercice.

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
# Exercice 3 : Pour les triggers qui suivent, vous faites des M.A.J(insert, update) de la base de données qui déclenchent les différentes erreurs.
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
# Exercice 4

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
