CREATE OR REPLACE PACKAGE MAJ is 
    PROCEDURE CREER_EMPLOYE (LE_NUEMPL IN NUMBER, LE_NOMEMPL IN VARCHAR2, LE_HEBDO IN NUMBER, LE_AFFECT IN NUMBER,LE_SALAIRE IN NUMBER);
    PROCEDURE MODIF_DUREE_HEBDO (NOUV_HEBDO IN NUMBER, NUEMP_CIBLE In NUMBER);
    PROCEDURE MODIF_SALAIRE (SALAIRE_NOUV IN NUMBER, EMP IN NUMBER);
    PROCEDURE MODIF_TRAVAIL (NOUV_DUREE IN NUMBER, PROJ IN NUMBER, EMP IN NUMBER);
    PROCEDURE INSERT_TRAVAIL (NOUV_EMP IN NUMBER, NOUV_NUPROJ IN NUMBER, NOUV_DUREE IN NUMBER);
    PROCEDURE INSERT_SERVICE (NOUV_NUSERV IN NUMBER, NOUV_NOMSERV IN VARCHAR, NOUV_CHEF IN NUMBER);
END;


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

END; -- Fin de "CREATE OR REPLACE PACKAGE BODY MAJ is" 
