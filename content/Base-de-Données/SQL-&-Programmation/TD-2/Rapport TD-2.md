---
title: Rapport TD-2
draft: 
description: 
tags:
  - Base-de-Données
  - SQL
---
## Exercice 1 : Triggers de type For each row et utilisation de « :NEW » et « :OLD »

**Ecrire un trigger de type for each row qui interdit la diminution du salaire d'un employé. Ce trigger se déclenche après la modification du salaire.**

|  nom trigger  | type : before / after | Insert, delete, update | nom table | for each row : oui / non |
| :-----------: | :-------------------: | :--------------------: | :-------: | :----------------------: |
| modif_salaire |         after         |         update         |  employé  |           oui            |

```sql
CREATE or REPLACE TRIGGER modif_salaire
AFTER UPDATE OF salaire on employe for EACH ROW
BEGIN
	IF :OLD.salaire > :NEW.salaire THEN
		RAISE_APPLICATION_ERROR(-20101, 'salaire non diminuable');
	END IF;
END;

-- test :
UPDATE employe set salaire = 0 WHERE nuempl = 20;
```

**Il y a une autre contrainte qui n'est pas spécifiée "la durée hebdomadaire d'un employé ne peut pas augmenter", elle n'est pas descriptible. Vous écrivez le trigger nécessaire à la vérification de cette contrainte**
