/*
EMPLOYE : (nuempl,nomempl,hebdo,affect)
SERVICE : (nuserv,nomserv,chef)
PROJET : (nuproj,nomproj,resp)
TRAVAIL : (nuempl,nuproj,duree)
CONCERNE : (nuserv,nuproj)
*/

-- 1 Contraintes d'intégrité

-- Reset
DROP TABLE concerne CASCADE CONSTRAINTS PURGE;
DROP TABLE travail CASCADE CONSTRAINTS PURGE;
DROP TABLE projet CASCADE CONSTRAINTS PURGE;
DROP TABLE employe CASCADE CONSTRAINTS PURGE;
DROP TABLE service CASCADE CONSTRAINTS PURGE;

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
ALTER TABLE travail add CONSTRAINT FK_employe_travail FOREIGN KEY (nuempl) REFERENCES employe(nuempl) DEFERRABLE;
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
ALTER TABLE projet add CONSTRAINT FK_resp FOREIGN KEY (resp, nuproj) REFERENCES travail (nuempl, nuproj) DEFERRABLE;

/*
La durée (temps de travail) hebdomadaire d'un employé est inférieure ou égale à 35h.
*/
ALTER TABLE employe add CONSTRAINT temps_trav_max CHECK (hebdo <=35) DEFERRABLE;

COMMIT;

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

COMMIT;
SELECT * FROM employe;

/*
La somme des durées d'un employé (de la table travail) doit être inférieur à la durée hebdomadaire `(Sum(duree) <= hebdo)`. 
Affichez les employés qui ne respectent pas cette contrainte. Vous modifiez les données de la table travail(update sur la 
durée) ou la table employé (update sur hebdo)jusqu'à ce que la requête précédente retourne un ensemble vide.
*/
SELECT e.NUEMPL, e.NOMEMPL, e.HEBDO, (SELECT SUM(t.DUREE) FROM travail t WHERE t.NUEMPL = e.NUEMPL) AS total_duree
FROM employe e
WHERE (
    SELECT SUM(t.DUREE)
    FROM travail t
    WHERE t.NUEMPL = e.NUEMPL
) > e.HEBDO;

-- à été exécuté 5 fois pour atteindre les 35h hebdomadaire
UPDATE employe
SET hebdo = LEAST(hebdo + 5, 35) -- Augmente de 5 heures sans dépasser 35 heures
WHERE nuempl IN (
    SELECT e.NUEMPL
    FROM employe e
    WHERE (
        SELECT SUM(t.DUREE)
        FROM travail t
        WHERE t.NUEMPL = e.NUEMPL
    ) > e.HEBDO
);

-- est éxécuté 1 fois pour réduire le temps de travail en dessous des 35h
UPDATE travail
SET duree = duree - 5
WHERE nuempl IN (
    SELECT e.NUEMPL
    FROM employe e
    WHERE (
        SELECT SUM(t.DUREE)
        FROM travail t
        WHERE t.NUEMPL = e.NUEMPL
    ) > e.HEBDO
);

COMMIT;

/*
Un employé ne peut être responsable de plus de 3 projets. Idem que la contrainte précédente, trouvez 
les employés qui ne respectent pas la contrainte et modifiez les données.
*/
SELECT e.NUEMPL, e.NOMEMPL, COUNT(p.NUPROJ) AS nb_projets
FROM employe e
JOIN projet p ON e.NUEMPL = p.RESP
GROUP BY e.NUEMPL, e.NOMEMPL
HAVING COUNT(p.NUPROJ) > 3;
-- Déjà bon

/*
Le chef d’un service gagne plus que les employés du service. Idem que la contrainte précédente, 
trouvez les employés qui ne respectent pas la contrainte et modifiez les données.
*/
SELECT s.CHEF, e.NOMEMPL AS chef_nom, e.SALAIRE AS salaire_chef, e2.NUEMPL AS employe_id, e2.NOMEMPL AS employe_nom, e2.SALAIRE AS salaire_employe
FROM service s
JOIN employe e ON s.CHEF = e.NUEMPL -- Chef du service
JOIN employe e2 ON e2.AFFECT = s.NUSERV -- Employés du même service
WHERE e.SALAIRE <= e2.SALAIRE AND e2.NUEMPL != e.NUEMPL; -- Ne compare pas le chef à lui-même

-- déjà bon mais voici un update au cas où
UPDATE employe
SET salaire = salaire + 1
WHERE nuempl IN (
    SELECT s.CHEF
    FROM service s
    JOIN employe e ON s.CHEF = e.NUEMPL -- Chef du service
    JOIN employe e2 ON e2.AFFECT = s.NUSERV -- Employés du même service
    WHERE e.SALAIRE <= e2.SALAIRE
    AND e2.NUEMPL != e.NUEMPL -- Ne compare pas le chef à lui-même
);

COMMIT;

/*
Un service ne peut être concerné par plus de 3 projets.
*/
SELECT s.NUSERV, s.NOMSERV, COUNT(c.NUPROJ) AS nb_projets
FROM service s
JOIN concerne c ON s.NUSERV = c.NUSERV
GROUP BY s.NUSERV, s.NOMSERV
HAVING COUNT(c.NUPROJ) > 3;

DELETE FROM concerne
WHERE (nuserv, nuproj) IN (
    SELECT nuserv, nuproj
    FROM (
        SELECT c.nuserv, c.nuproj, 
               ROW_NUMBER() OVER (PARTITION BY c.nuserv ORDER BY c.nuproj) AS rn
        FROM concerne c
    ) sub
    WHERE sub.rn > 3
);

COMMIT;