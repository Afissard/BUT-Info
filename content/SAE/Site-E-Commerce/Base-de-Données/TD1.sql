/*
CREATION DU PACKAGE LECTURE
*/
DROP PACKAGE lecture;

/*INITIALISATION*/

CREATE OR REPLACE PACKAGE lecture 
    is TYPE cur_empl 
    is REF CURSOR; 
    PROCEDURE liste_employes(liste OUT cur_empl); 
    PROCEDURE li_proj_emp(idempl IN NUMBER, liste OUT cur_empl);
    PROCEDURE li_emp_proj(idproj IN NUMBER, liste OUT cur_empl);
END;

/*DÉCLARATION DES PROCÉDURES*/

-- Create or replace package BODY lecture 
--     is procedure liste_employes(liste out cur_empl) 
--     is begin 
--         open liste for select * from employe ; 
--     end ; 
-- end ;

CREATE OR REPLACE PACKAGE BODY lecture is
    /*
    Modifier la procédure liste_employes en donnant en plus en 
    sortie les noms des services auxquels sont affectés les employés .
    */
    PROCEDURE liste_employes(liste OUT cur_empl) is
    BEGIN 
        OPEN liste FOR 
            SELECT e.NUEMPL, e.NOMEMPL, e.HEBDO, e.AFFECT, e.SALAIRE, p.NOMPROJ
            FROM EMPLOYE e, PROJET p
            WHERE 
                e.AFFECT IN (SELECT c.NUSERV FROM CONCERNE c)
                AND p.NUPROJ IN (SELECT c.NUPROJ FROM CONCERNE c)
            ;
    END;

    /*
    liste des projets de la table travail pour un employé donné en paramètre. 
    Cette procédure à deux paramètres(un numéro d’employé en entrée de type 
    in et paramètre de sortie de type out). Le résultat doit contenir le nom 
    de l’employe, le nom du projet et la durée.
    */
    PROCEDURE li_proj_emp(idempl IN NUMBER, liste OUT cur_empl) is
    BEGIN
        OPEN liste FOR
        SELECT NOMEMPL, NOMPROJ, DUREE
        FROM EMPLOYE e
        JOIN TRAVAIL t ON t.NUEMPL = e.NUEMPL
        JOIN PROJET p ON p.NUPROJ = t.NUPROJ
        WHERE idempl = e.NUEMPL;
    END;

    /*
    Liste des employés pour un projet donné
    */
    PROCEDURE li_emp_proj(idproj IN NUMBER, liste OUT cur_empl) is
    BEGIN
        OPEN liste FOR
        SELECT e.NUEMPL, e.NOMEMPL
        FROM EMPLOYE e
        JOIN CONCERNE c ON c.NUSERV = e.AFFECT
        JOIN PROJET p ON p.NUPROJ = c.NUPROJ
        WHERE idproj = p.NUPROJ;
    END;

END;

/*ÉXÉCUTION*/

variable li REFCURSOR ;

execute lecture.liste_employes(:li); 
print li;

execute lecture.LI_PROJ_EMP(:li); 
print li;

execute lecture.LI_EMP_PROJ(:li); 
print li;
