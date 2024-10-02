CREATE OR REPLACE PACKAGE MAJ is 
    PROCEDURE CREER_EMPLOYE (
        LE_NUEMPL IN NUMBER, LE_NOMEMPL IN VARCHAR2, LE_HEBDO IN NUMBER, LE_AFFECT IN NUMBER,LE_SALAIRE IN NUMBER
        );
END;

--------------------------------------------------------------------------------------------------------------

CREATE OR REPLACE PACKAGE BODY MAJ is 
PROCEDURE CREER_EMPLOYE (
    LE_NUEMPL IN NUMBER, LE_NOMEMPL IN VARCHAR2, LE_HEBDO IN NUMBER, LE_AFFECT IN NUMBER,LE_SALAIRE IN NUMBER
    ) is
BEGIN
    SET TRANSACTION READ WRITE;
    INSERT INTO employe VALUES(LE_NUEMPL, LE_NOMEMPL, LE_HEBDO, LE_AFFECT,LE_SALAIRE);
    COMMIT;
    EXCEPTION WHEN OTHERS THEN ROLLBACK;
    IF SQLCODE=-00001 THEN 
        ROLLBACK;
        RAISE_APPLICATION_ERROR(-20401, 'Un employe avec le meme numero existe deja');
    ELSIF SQLCODE=-2291 THEN 
        ROLLBACK;
        RAISE_APPLICATION_ERROR(-20402, 'Le service auquel il est affecté n’existe pas');
    ELSIF SQLCODE=-02290 THEN 
        ROLLBACK;
        RAISE_APPLICATION_ERROR(-20403, 'la durée hebdomadaire d’un employé doit être inférieure ou égale à 35h');
    ELSIF SQLCODE=-1438 THEN 
        ROLLBACK;
        RAISE_APPLICATION_ERROR(-20404, 'Une valeur (nombre) dépasse le nombre de caractères autorisés');
    ELSIF SQLCODE=-12899 THEN 
        ROLLBACK;
        RAISE_APPLICATION_ERROR(-20405, 'Une valeur (chaine de caractère) dépasse le nombre de caractères autorisés');
    ELSIF SQLCODE=-20101 THEN 
        ROLLBACK;
        RAISE_APPLICATION_ERROR(-20405, 'Le salaire de cet employé dépasse celui de son chef de service'); --TODO
    ELSE 
        ROLLBACK;
        RAISE_APPLICATION_ERROR (-20999,'Erreur inconnue'||SQLcode);
    END IF; 
    END;
END;

--------------------------------------------------------------------------------------------------------------
-- TESTS
EXECUTE MAJ.CREER_EMPLOYE(20, 'Boby', 31, 1, 2000);
EXECUTE MAJ.CREER_EMPLOYE(100, 'Boby', 31, 66, 2000);
EXECUTE MAJ.CREER_EMPLOYE(100, 'Boby', 66, 1, 2000);
EXECUTE MAJ.CREER_EMPLOYE(100, 'Boby', 31, 1000000000000000000000000000000000000000, 2000);
EXECUTE MAJ.CREER_EMPLOYE(100, 'Jugemu Jugemu Unko Nageki Ototoi no Shin-chan no Pantsu Shinpachi no Jinsei Barumunku Fezarion Aizakku Shunaidaa Sanbun no Ichi no Junjou na Kanjou no Nokotta Sanbun no Ni wa Sakamuke ga Kininaru Kanjou Uragiri wa Boku no Namae wo Shitteiru you de Shiranai no wo Boku wa Shitteiru Rusu Surume Medaka Kazunoko Koedame Medaka... Kono Medaka wa Sakki to Chigau Yatsu Dakara Ikeno Medaka no Hou Dakara Raayu Yuuteimiyaoukimukou Pepepepepepepepepepepepe Bichiguso Maru', 31, 1, 2000);
-- EXECUTE MAJ.CREER_EMPLOYE(100, 'boby', 31, 1, 66666); -- salaire supp chef de service
-- EXECUTE MAJ.CREER_EMPLOYE(100, 'boby', 31, 1, 1); -- erreur inconue

-- Suppression des traces du tests
SELECT * FROM EMPLOYE WHERE nuempl = 100;
DELETE FROM EMPLOYE WHERE NUEMPL = 100;
COMMIT;
