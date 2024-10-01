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
  SELECT * INTO rec_t
  FROM EMPLOYE e 
  WHERE (
    SELECT SUM(duree)
    FROM TRAVAIL t 
    WHERE e.NUEMPL = t.NUEMPL
  ) > e.HEBDO;

  RAISE_APPLICATION_ERROR(-20301, 'temps travail invalide');
  EXCEPTION 
		WHEN no_data_found THEN null; -- aucun employé ne travail plus qu'autorisé
		WHEN too_many_rows THEN RAISE_APPLICATION_ERROR(-20301, 'temps travail invalide');
END;

-- 6 tests à faire
-- UPDATE TRAVAIL SET DUREE = 40 WHERE NUEMPL = 20; -- active le trigger (temps invalide)
-- UPDATE TRAVAIL SET DUREE = 10 WHERE NUEMPL > 20; -- active le trigger  (large séléction)
-- UPDATE TRAVAIL SET DUREE = 1 WHERE NUEMPL = 20; -- active pas le trigger (-> mais ne marche pas) 

-- INSERT INTO EMPLOYE (NUEMPL, NOMEMPL, HEBDO, AFFECT, SALAIRE) VALUES (1, 'Hades', 2, 1, 1); -- employé de test
-- INSERT INTO TRAVAIL (NUEMPL, NUPROJ, DUREE) VALUES (1, 103, 3); -- test du trigger -> temps trop grand
-- INSERT INTO TRAVAIL (NUEMPL, NUPROJ, DUREE) VALUES (1, 103, 1); -- test du trigger -> temps trop faible
-- INSERT INTO TRAVAIL (NUEMPL, NUPROJ, DUREE) VALUES (1, 103, 2); -- test du trigger -> bon mais erreur
