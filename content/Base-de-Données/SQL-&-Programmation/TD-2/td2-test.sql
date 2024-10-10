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
UPDATE EMPLOYE SET salaire = 0 WHERE nuempl = 20; -- ne passe pas le trigger
UPDATE EMPLOYE Set salaire = 3501 WHERE nuempl = 23; -- passe le trigger
ROLLBACK;

/*
Il y a une autre contrainte qui n'est pas spécifiée "la durée hebdomadaire 
d'un employé ne peut pas augmenter", elle n'est pas descriptible. Vous écrivez 
le trigger nécessaire à la vérification de cette contrainte
*/

-- tests
UPDATE EMPLOYE SET hebdo = 30 WHERE nuempl = 23; -- active le trigger
UPDATE EMPLOYE SET hebdo = 20 WHERE nuempl = 23; -- active  pas le trigger
ROLLBACK;

/*
La spécification de l'opération supprimer_employe impose que la suppression 
d'un employé soit accompagnée de la suppression des lignes de travail correspondantes. 
Mettez en place un trigger table qui le fait. (pas de problème si on a déclaré 
"deferred" la contrainte FK_employe de la table travail vers la table employe 
"les employés de travail existent"). 
Attention : la suppression des employés de la table travail n’est possible que 
si l’employé n’est pas chef de service ou responsable de projet
*/
-- marche
DELETE FROM employe WHERE nuempl = 37; -- employé n'étant ni chef ni responsable.
DELETE FROM employe WHERE nuempl = 48; -- employé avec des lignes dans `travail`.
-- retourne une erreur (attendu)
DELETE FROM employe WHERE nuempl = 17; -- employé chef de service.
DELETE FROM employe WHERE nuempl = 20; -- employé responsable de projet.
ROLLBACK;

/*
La spécification de l'opération supprimer_projet impose que la suppression 
d'un projet soit accompagnée de la suppression des lignes de travail et de 
la table concerne correspondantes. Mettez en place un trigger table qui effectue 
cela. (pas de problème si on a déclaré "deferred" la contrainte FK_nuproj de 
la table travail et de la table concerne vers la table projet).
*/

DELETE FROM projet WHERE nuproj = 160; -- projet ayant des lignes dans travail et concerne.
DELETE FROM projet WHERE nuproj = 492; -- projet ayant uniquement des lignes dans travail.
DELETE FROM projet WHERE nuproj = 103; -- projet responsable par un employé spécifique.
ROLLBACK;

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
UPDATE TRAVAIL SET DUREE = 50 WHERE NUEMPL = 23 AND NUPROJ = 237;
UPDATE TRAVAIL SET DUREE = 50 WHERE NUEMPL > 23;
UPDATE TRAVAIL SET DUREE = 5 WHERE NUEMPL = 23 AND NUPROJ = 237;


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
INSERT INTO PROJET (NUPROJ, NOMPROJ, RESP) VALUES (100, 'percephone', 30); -- Créer un nouveau projet et y ajoute un employe déjà responsable 3 fois
INSERT INTO PROJET (NUPROJ, NOMPROJ, RESP) VALUES (100, 'hades', 20); -- Créer un nouveau projet et y ajoute un employe déjà responsable 1 fois
UPDATE PROJET SET RESP = 30 WHERE RESP = 20; -- assigne la responsabilité du projet de n20 à n30
UPDATE PROJET SET RESP = 57 WHERE RESP = 20; -- assigne la responsabilité du projet de n20 à n57
ROLLBACK;

/*
Ecrire un trigger qui vérifie la contrainte suivante : "un service ne peut être concerné par plus de 3 projets"
*/
UPDATE CONCERNE SET NUSERV = 1 WHERE NUPROJ = 160 AND NUSERV = 2; -- ajoute un projet à un service qui possède déjà 3 autre service
UPDATE CONCERNE SET NUSERV = 5 WHERE NUPROJ = 103 AND NUSERV = 1; -- ajoute un projet à un service qui possède moins 3 autre service 
INSERT INTO CONCERNE (NUSERV, NUPROJ) VALUES (1, 160); -- ajoute un projet à un service qui possède déjà 3 autre service
INSERT INTO CONCERNE (NUSERV, NUPROJ) VALUES (5, 103); -- ajoute un projet à un service qui possède moins 3 autre service
ROLLBACK;

/*
Ecrire un trigger qui vérifie la contrainte suivante : "un chef de service gagne plus que les employés de son service".
*/
UPDATE EMPLOYE SET SALAIRE = 4000 WHERE AFFECT = 1 AND NUEMPL != 41; -- augmente le salaire d'un employé pour depassez le chef
UPDATE EMPLOYE SET SALAIRE = 2500 WHERE AFFECT = 1 AND NUEMPL = 39; -- augmente le salaire d'un employé sans depassez le chef
INSERT INTO EMPLOYE (NUEMPL, NOMEMPL, HEBDO, AFFECT, SALAIRE) VALUES (1, 'ulysse', 10, 1, 4000); -- créer un employer mieux payer que son chef
INSERT INTO EMPLOYE (NUEMPL, NOMEMPL, HEBDO, AFFECT, SALAIRE) VALUES (1, 'ulysse', 10, 1, 1500); -- créer un employer moins bien payer que son chef
ROLLBACK;

/*
Ecrire un trigger qui vérifie la contrainte suivante : "un chef de service gagne plus que les employés responsables de projets".
*/
-- SELECT * FROM EMPLOYE WHERE AFFECT = 2 ORDER BY SALAIRE DESC;
UPDATE EMPLOYE SET SALAIRE = 4000 WHERE AFFECT = 2 AND NUEMPL = 57; -- augment le salaire d'un resp pour depasser le chef
UPDATE EMPLOYE SET SALAIRE = 3000 WHERE AFFECT = 2 AND NUEMPL = 57; -- augment le salaire d'un resp sans depasser le chef
-- pour les insert allez chercher un employé dejà resp ailleur
ROLLBACK;

/*
Est-il possible de regrouper les deux derniers trigger
*/

-- trigger 1
UPDATE EMPLOYE SET SALAIRE = 4000 WHERE AFFECT = 1 AND NUEMPL != 41; -- augmente le salaire d'un employé pour depassez le chef
UPDATE EMPLOYE SET SALAIRE = 2500 WHERE AFFECT = 1 AND NUEMPL = 39; -- augmente le salaire d'un employé sans depassez le chef
INSERT INTO EMPLOYE (NUEMPL, NOMEMPL, HEBDO, AFFECT, SALAIRE) VALUES (1, 'ulysse', 10, 1, 4000); -- créer un employer mieux payer que son chef
INSERT INTO EMPLOYE (NUEMPL, NOMEMPL, HEBDO, AFFECT, SALAIRE) VALUES (1, 'ulysse', 10, 1, 1500); -- créer un employer moins bien payer que son chef
-- trigger 2
UPDATE EMPLOYE SET SALAIRE = 4000 WHERE AFFECT = 2 AND NUEMPL = 57; -- augment le salaire d'un resp pour depasser le chef
UPDATE EMPLOYE SET SALAIRE = 3000 WHERE AFFECT = 2 AND NUEMPL = 57; -- augment le salaire d'un resp sans depasser le chef
-- pour les insert allez chercher un employé dejà resp ailleur
ROLLBACK;


/*
Lors d'augmentation de salaire ou d'embauche, l'entreprise veut enregistrer les employés (dans la table 
`EMPLOYE_ALERTE` idem que `EMPLOYE`) avec un salaire qui dépassent les 5000 euros. Ecrire un trigger 
qui permet de remplir cette table.
*/
alter trigger "CHEF_PAYE_SUP_RESP" disable;
alter trigger "CHEF_PLUS_PAYE" disable;
alter trigger "CHEF_PLUS_PAYE_EMP_ET_RESP" disable;
alter trigger "REMPLIS_EMPLOYE_ALERTE" enable;

INSERT INTO EMPLOYE (NUEMPL, NOMEMPL, HEBDO, AFFECT, SALAIRE) VALUES (1, 'cresus', 10, 1, 50000); -- créer un employer payé plus de 5000
INSERT INTO EMPLOYE (NUEMPL, NOMEMPL, HEBDO, AFFECT, SALAIRE) VALUES (2, 'diogene', 10, 1, 1); -- créer un employer payé moins de 5000
--UPDATE EMPLOYE SET SALAIRE = 50000 WHERE NUEMPL = 20; -- créer un employer payé plus de 5000
ROLLBACK;

alter trigger "CHEF_PAYE_SUP_RESP" enable;
alter trigger "CHEF_PLUS_PAYE" enable;
alter trigger "CHEF_PLUS_PAYE_EMP_ET_RESP" enable;
alter trigger "REMPLIS_EMPLOYE_ALERTE" DISABLE;

