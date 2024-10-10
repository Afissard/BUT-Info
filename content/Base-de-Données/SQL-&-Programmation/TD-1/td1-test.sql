SELECT * from EMPLOYE;
SELECT * from SERVICE;
SELECT * from PROJET;
SELECT * from TRAVAIL;
SELECT * from CONCERNE;


-- Test de la contrainte PK_employe
INSERT INTO employe VALUES (20, 'toto', 35, 1, 1500); -- Insertion d'un employé qui existe déjà

-- Test de la contrainte PK_service
INSERT INTO service VALUES (1, 'service_test'); -- Insertion d'un service qui existe déjà

-- Test de la contrainte PK_projet
INSERT INTO projet VALUES (103, 'projet_test', 200); -- Insertion d'un projet déjà existant

-- Test de la contrainte PK_travail
INSERT INTO travail VALUES (20, 492, 10); -- Insertion d'une ligne qui existe déjà dans travail

-- Test de la contrainte FK_employe_travail
INSERT INTO travail VALUES (999, 492, 10); -- Insertion avec un employé qui n'existe pas

-- Test de la contrainte FK_projet_travail
INSERT INTO travail VALUES (20, 999, 10); -- Insertion avec un projet qui n'existe pas

-- Test de la contrainte FK_service_concerne
INSERT INTO concerne VALUES (99, 103); -- Service inexistant

-- Test de la contrainte FK_projet_concerne
INSERT INTO concerne VALUES (1, 999); -- Projet inexistant

-- Test de la contrainte UNQ_employe_affect
INSERT INTO employe VALUES (30, 'jeSuisNouveau', 35, 1, 1500); -- Affectation unique déjà utilisée

-- Test de la contrainte FK_chef
-- UPDATE service SET chef = 999 WHERE nuserv = 1; -- Chef inexistant

-- Test de la contrainte FK_affect
UPDATE employe SET affect = 99 WHERE nuempl = 20; -- Service affecté inexistant

-- Test de la contrainte FK_resp
-- UPDATE projet SET resp = 999 WHERE nuproj = 103; -- Responsable inexistant dans travail

-- Test de la contrainte temps_trav_max
INSERT INTO employe VALUES (40, 'heurs_supp', 40, 1, 1500); -- Hebdo supérieur à 35h
