---
title: Rapport TD-2
draft: 
description: 
tags:
  - Base-de-Données
  - SQL
---
## Exercice 1 : Triggers de type For each row et utilisation de ":NEW" et ":OLD"

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

## Exercice 2 : Trigger de type : Delete from T2 where a not in (select a from T1) / D’autres types de solutions ne sont pas acceptés pour cette exercice.

La spécification de l'opération supprimer_employe impose que la suppression d'un employé soit accompagnée de la suppression des lignes de travail correspondantes. Mettez en place un trigger table qui le fait. (pas de problème si on a déclaré "deferred" la contrainte FK_employe de la table travail vers la table employe "les employés de travail existent").
**Attention : la suppression des employés de la table travail n’est possible que si l’employé n’est pas chef de service ou responsable de projet**

| nom trigger | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :---------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
|             |                       |                        |           |                          |

#TODO 

La spécification de l'opération supprimer_projet impose que la suppression d'un projet soit accompagnée de la suppression des lignes de travail et de la table concerne correspondantes. Mettez en place un trigger table qui effectue cela. (pas de problème si on a déclaré "deferred" la contrainte FK_nuproj de la table travail et de la table concerne vers la table projet).

| nom trigger | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :---------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
|             |                       |                        |           |                          |

#TODO 

## Exercice 3 : Pour les triggers qui suivent, vous faites des M.A.J(insert, update) de la base de données qui déclenchent les différentes erreurs.
*Tous les trigger de cette exercice sont de type `select *` into suivi de l’instruction `raise_application_error`. Les autres solutions ne seront pas acceptées.*

**Il y a une contrainte qui n'est pas spécifiée "la somme des durées de travail d'un employé ne doit pas excéder son temps de travail hebdomadaire", elle n'est pas descriptible. Vous écrivez le (ou les) trigger nécessaire à la vérification de cette contrainte. La contrainte à mettre en place SUM(duree)<=hebdo.** 
- **Quels sont les différentes opérations(update, insert,…) sur les tables `employe` ou `travail` qui vous amènent à ne pas respecter cette contrainte.**
- **Combien de trigger allez vous mettre en place ? Indiquez dans quel cas les trigger se déclenchent. Vous utilisez la requête suivante pour construire ce(s) trigger : `Select * from employe e where (select sum(duree) from travail t where e.nuempl=t.nuempl)> hebdo`**

| nom trigger | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :---------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
|             |                       |                        |           |                          |

#TODO 

**Ecrire un trigger qui vérifie la contrainte suivante: « un employé est responsable au plus sur 3 projets ». Idem que la question précédente, vous utilisez la requête suivante pour construire votre trigger : `Select * from employe e where (select count(*) from projet p where e.nuempl=p.resp)> 3 ;`**

| nom trigger | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :---------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
|             |                       |                        |           |                          |

#TODO 

**Ecrire un trigger qui vérifie la contrainte suivante : « un service ne peut être concerné par plus de 3 projets »**

| nom trigger | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :---------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
|             |                       |                        |           |                          |

#TODO 

**Ecrire un trigger qui vérifie la contrainte suivante : « un chef de service gagne plus que les employés de son service.**

| nom trigger | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :---------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
|             |                       |                        |           |                          |

#TODO 

**Est-il possible de regrouper les deux derniers « trigger »**

| nom trigger | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :---------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
|             |                       |                        |           |                          |

#TODO 

## Exercice 4

**Lors d'augmentation de salaire ou d'embauche, l'entreprise veut enregistrer les employés (dans la table EMPLOYE_ALERTE idem que EMPLOYE) avec un salaire qui dépassent les 5000 euros. Ecrire un trigger qui permet de remplir cette table.**

| nom trigger | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :---------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
|             |                       |                        |           |                          |

#TODO 
