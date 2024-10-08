/*
EMPLOYE : (nuempl,nomempl,hebdo,affect)
SERVICE : (nuserv,nomserv,chef)
PROJET : (nuproj,nomproj,resp)
TRAVAIL : (nuempl,nuproj,duree)
CONCERNE : (nuserv,nuproj)
*/

SELECT * from EMPLOYE;
SELECT * from SERVICE;
SELECT * from PROJET;
SELECT * from TRAVAIL;
SELECT * from CONCERNE;

/*
Ecrire un trigger de type for each row qui interdit la diminution 
du salaire d'un employé. Ce trigger se déclenche après la modification 
du salaire.
*/

CREATE or REPLACE TRIGGER modif_salaire 
AFTER UPDATE OF salaire on EMPLOYE for EACH ROW
BEGIN
    IF :OLD.salaire > :NEW.salaire THEN
      RAISE_APPLICATION_ERROR(-20101, 'salaire non diminuable');
    END IF;
END;

COMMIT;

-- tests :
-- UPDATE EMPLOYE SET salaire = 0 WHERE nuempl = 20; -- ne passe pas le trigger
-- UPDATE EMPLOYE Set salaire = 3501 WHERE nuempl = 23; -- passe le trigger
-- ROLLBACK;

/*
Il y a une autre contrainte qui n'est pas spécifiée "la durée hebdomadaire 
d'un employé ne peut pas augmenter", elle n'est pas descriptible. Vous écrivez 
le trigger nécessaire à la vérification de cette contrainte
*/

CREATE or REPLACE TRIGGER augmentation_hebdo
AFTER UPDATE OF hebdo on EMPLOYE FOR EACH ROW
BEGIN
  IF :OLD.hebdo < :NEW.hebdo THEN
    RAISE_APPLICATION_ERROR(-20102, 'hebdo non augmentable');
  END IF;
END;
COMMIT;

-- tests
-- UPDATE EMPLOYE SET hebdo = 30 WHERE nuempl = 23; -- active le trigger
-- UPDATE EMPLOYE SET hebdo = 20 WHERE nuempl = 23; -- active  pas le trigger
-- ROLLBACK;

/*
La spécification de l'opération supprimer_employe impose que la suppression 
d'un employé soit accompagnée de la suppression des lignes de travail correspondantes. 
Mettez en place un trigger table qui le fait. (pas de problème si on a déclaré 
"deferred" la contrainte FK_employe de la table travail vers la table employe 
"les employés de travail existent"). 
Attention : la suppression des employés de la table travail n’est possible que 
si l’employé n’est pas chef de service ou responsable de projet
*/

-- CREATE or REPLACE TRIGGER supprimer_employe BEFORE DELETE on EMPLOYE
-- BEGIN
--   DELETE FROM TRAVAIL
--   WHERE 
--     NUEMPL NOT IN (SELECT chef from SERVICE)
--     AND NUEMPL NOT IN (SELECT resp from PROJET);
-- END;

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

-- -- marche
-- DELETE FROM employe WHERE nuempl = 37; -- employé n'étant ni chef ni responsable.
-- DELETE FROM employe WHERE nuempl = 48; -- employé avec des lignes dans `travail`.
-- -- retourne une erreur (attendu)
-- DELETE FROM employe WHERE nuempl = 17; -- employé chef de service.
-- DELETE FROM employe WHERE nuempl = 20; -- employé responsable de projet.
-- ROLLBACK;

/*
La spécification de l'opération supprimer_projet impose que la suppression 
d'un projet soit accompagnée de la suppression des lignes de travail et de 
la table concerne correspondantes. Mettez en place un trigger table qui effectue 
cela. (pas de problème si on a déclaré "deferred" la contrainte FK_nuproj de 
la table travail et de la table concerne vers la table projet).
*/
CREATE OR REPLACE TRIGGER supprimer_projet
BEFORE DELETE ON projet
FOR EACH ROW
BEGIN
    -- Suppression des lignes dans la table travail associées au projet
    DELETE FROM travail WHERE nuproj = :OLD.nuproj;

    -- Suppression des lignes dans la table concerne associées au projet
    DELETE FROM concerne WHERE nuproj = :OLD.nuproj;
END;

-- DELETE FROM projet WHERE nuproj = 160; -- projet ayant des lignes dans travail et concerne.
-- DELETE FROM projet WHERE nuproj = 492; -- projet ayant uniquement des lignes dans travail.
-- DELETE FROM projet WHERE nuproj = 103; -- projet responsable par un employé spécifique.
-- ROLLBACK;

/*
Il y a une contrainte qui n'est pas spécifiée "la somme des durées de travail 
d'un employé ne doit pas excéder son temps de travail hebdomadaire", elle n'est 
pas descriptible. Vous écrivez le (ou les) trigger nécessaire à la vérification 
de cette contrainte. La contrainte à mettre en place SUM(duree)<=hebdo.
- Quels sont les différentes opérations(update, insert,…) sur les tables employe 
  ou travail qui vous amènent à ne pas respecter cette contrainte. 
- Combien de trigger allez vous mettre en place ? Indiquez dans quel cas les 
  trigger se déclenchent.

Vous utilisez la requête suivante pour construire ce(s) trigger : 
SELECT * FROM EMPLOYE e 
WHERE (
  SELECT SUM(duree)
  FROM TRAVAIL t where e.NUEMPL = t.NUEMPL
) > e.HEBDO;
*/

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

-- 6 tests à faire
-- UPDATE TRAVAIL SET DUREE = 40 WHERE NUEMPL = 20; -- active le trigger (temps invalide)
-- UPDATE TRAVAIL SET DUREE = 10 WHERE NUEMPL > 20; -- active le trigger  (large séléction)
-- UPDATE TRAVAIL SET DUREE = 1 WHERE NUEMPL = 20; -- active pas le trigger (-> mais ne marche pas) 

/*
Ecrire un trigger qui vérifie la contrainte suivante: "un employé est responsable au plus sur 3 projets". 
Idem que la question précédente, vous utilisez la requête suivante pour construire votre trigger : 
SELECT * FROM EMPLOYE e 
WHERE (
  SELECT COUNT(*) 
  FROM PROJET p 
  WHERE e.nuempl=p.resp
)> 3;
*/
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

/*
Ecrire un trigger qui vérifie la contrainte suivante : "un service ne peut être concerné par plus de 3 projets"
*/

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

/*
Ecrire un trigger qui vérifie la contrainte suivante : "un chef de service gagne plus que les employés de son service".
*/

CREATE OR REPLACE TRIGGER chef_plus_paye
  AFTER INSERT OR UPDATE OF SALAIRE ON EMPLOYE
  DECLARE rec SERVICE%ROWTYPE;
BEGIN
  SELECT s.NUSERV, s.NOMSERV, s.CHEF
  INTO rec
  FROM SERVICE s
  JOIN EMPLOYE e ON s.CHEF = e.NUEMPL -- Chef du service
  JOIN EMPLOYE e2 ON e2.AFFECT = s.NUSERV -- Employés du même service
  WHERE e.SALAIRE <= e2.SALAIRE AND e2.NUEMPL != e.NUEMPL; -- Ne compare pas le chef à lui-même

  RAISE_APPLICATION_ERROR(-20304, 'un chef de service gagne plus que les employés de son service');
  EXCEPTION 
		WHEN no_data_found THEN null;
		WHEN too_many_rows THEN RAISE_APPLICATION_ERROR(-20304, 'un chef de service gagne plus que les employés de son service');
END;

/*
Ecrire un trigger qui vérifie la contrainte suivante : "un chef de service gagne plus que les employés responsables de projets".
*/

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

/*
Est-il possible de regrouper les deux derniers trigger
*/

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


/*
Lors d'augmentation de salaire ou d'embauche, l'entreprise veut enregistrer les employés (dans la table 
`EMPLOYE_ALERTE` idem que `EMPLOYE`) avec un salaire qui dépassent les 5000 euros. Ecrire un trigger 
qui permet de remplir cette table.
*/

-- DROP TABLE EMPLOYE_ALERTE CASCADE CONSTRAINTS PURGE;
CREATE table EMPLOYE_ALERTE AS SELECT * FROM EMPLOYE;

CREATE OR REPLACE TRIGGER REMPLIS_EMPLOYE_ALERTE 
  AFTER INSERT OR UPDATE OF SALAIRE ON EMPLOYE
  FOR EACH ROW 
    WHEN ( NEW.SALAIRE > 5000) 
BEGIN 
	INSERT INTO EMPLOYE_ALERTE (NUEMPL, NOMEMPL, HEBDO, AFFECT, SALAIRE) 
  VALUES (:NEW.NUEMPL, :NEW.NOMEMPL, :NEW.HEBDO, :NEW.AFFECT, :NEW.SALAIRE); 

  EXCEPTION 
		WHEN no_data_found THEN null;
END;
