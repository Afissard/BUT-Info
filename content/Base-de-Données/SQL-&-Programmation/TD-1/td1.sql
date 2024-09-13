/*
EMPLOYE : (nuempl,nomempl,hebdo,affect)
SERVICE : (nuserv,nomserv,chef)
PROJET : (nuproj,nomproj,resp)
TRAVAIL : (nuempl,nuproj,duree)
CONCERNE : (nuserv,nuproj)
*/

-- 1 Contraintes d'intégrité

-- Reset
DROP TABLE concerne;
DROP TABLE travail;
DROP TABLE projet;
DROP TABLE employe;
DROP TABLE service;

-- Creation des tables
-- employe
CREATE table employe AS SELECT * FROM basetd.employe;
ALTER TABLE employe ADD CONSTRAINT PK_employe PRIMARY KEY (nuempl);
SELECT * from employe;

-- service
CREATE table service AS SELECT * FROM basetd.service;
ALTER TABLE service ADD CONSTRAINT PK_service PRIMARY KEY (nuserv);
SELECT * FROM service;

-- projet
CREATE table projet AS SELECT * FROM basetd.projet;
ALTER TABLE projet ADD CONSTRAINT PK_projet PRIMARY KEY (nuproj);
SELECT * FROM projet;

-- travail
CREATE table travail AS SELECT * FROM basetd.travail;
ALTER TABLE travail ADD CONSTRAINT PK_travail PRIMARY KEY (nuempl, nuproj);
SELECT * FROM travail;

-- concerne
CREATE table concerne AS SELECT * FROM basetd.concerne;
ALTER TABLE concerne ADD CONSTRAINT PK_concerne PRIMARY KEY (nuserv, nuproj);
SELECT * FROM concerne;


-- Ajout des clefs étranères
/*
Contraintes implicites
*/
ALTER TABLE employe add CONSTRAINT FK_affect FOREIGN KEY (affect) REFERENCES service (nuserv);
ALTER TABLE travail add CONSTRAINT FK_employe_travail FOREIGN KEY (nuempl) REFERENCES employe(nuempl);
ALTER TABLE travail add CONSTRAINT FK_projet_travail FOREIGN KEY (nuproj) REFERENCES projet(nuproj);
ALTER TABLE concerne add CONSTRAINT FK_service_concerne FOREIGN KEY (nuserv) REFERENCES service(nuserv);
ALTER TABLE concerne add CONSTRAINT FK_projet_concerne FOREIGN KEY (nuproj) REFERENCES projet(nuproj);

/*
Le chef d'un service est un employé affecté au même service. Attention les contraintes 
de clé étrangère Affect et Chef se croisent. Dans ce cas, vous devriez créer une 
contrainte unique (nuempl, affect) et une contrainte différée(initially deferred) de la 
table service (chef,nuserv) vers cette contrainte unique(nuempl, affect).
*/
ALTER TABLE service add CONSTRAINT FK_chef FOREIGN KEY (chef) REFERENCES employe (nuempl) DEFERRABLE;

/*
Un responsable de projet doit travailler sur le projet. Dans ce cas vous n'avez pas besoin 
de créer la contrainte Unique (nuempl, nuproj) de la table travail, car les deux 
attributs forment déjà la clé primaire. Par contre la contrainte (resp, nuproj) est une clé 
étrangère différée vers la clé primaire de la table travail.
*/

/*
La durée hebdomadaire d'un employé est inférieure ou égale à 35h.
*/

-- Test des contraites
INSERT INTO EMPLOYE VALUES(20, 'toto', 35, 1);


-- 2 Mise à jour des données:

ALTER TABLE EMPLOYE ADD
SALAIRE NUMBER;

UPDATE EMPLOYE 
SET SALAIRE = 2500
WHERE NUEMPL IN (SELECT RESP FROM PROJET);

UPDATE EMPLOYE 
SET SALAIRE = 3500
WHERE NUEMPL IN (SELECT CHEF FROM SERVICE);

UPDATE EMPLOYE 
SET SALAIRE = 1999
WHERE NUEMPL NOT IN ((SELECT CHEF FROM SERVICE) UNION (SELECT RESP FROM PROJET));


-- TODO: le select doit rendre un tableau vide
SELECT * FROM employe e
WHERE (
    SELECT sum(duree) FROM travail t
    WHERE e.NUEMPL = t.NUEMPL
) > e.HEBDO ;