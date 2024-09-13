SELECT * from employe;

/*
Ecrire un trigger de type for each row qui interdit la diminution 
du salaire d'un employé. Ce trigger se déclenche après la modification 
du salaire.
*/

CREATE or REPLACE TRIGGER modif_salaire 
AFTER UPDATE OF salaire on employe for EACH ROW
BEGIN
    IF :OLD.salaire > :NEW.salaire THEN
      RAISE_APPLICATION_ERROR(-20101, 'salaire non diminuable');
    END IF;
END;

-- test :
UPDATE employe set salaire = 0 WHERE nuempl = 20;