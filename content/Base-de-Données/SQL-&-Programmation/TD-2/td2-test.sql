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