-- TESTS CREER_EMPLOYE
EXECUTE MAJ.CREER_EMPLOYE(20, 'Boby', 31, 1, 2000);
EXECUTE MAJ.CREER_EMPLOYE(100, 'Boby', 31, 66, 2000);
EXECUTE MAJ.CREER_EMPLOYE(100, 'Boby', 66, 1, 2000);
EXECUTE MAJ.CREER_EMPLOYE(100, 'Boby', 31, 1000000000000000000000000000000000000000, 2000);
EXECUTE MAJ.CREER_EMPLOYE(100, 'Jugemu Jugemu Unko Nageki Ototoi no Shin-chan no Pantsu Shinpachi no Jinsei Barumunku Fezarion Aizakku Shunaidaa Sanbun no Ichi no Junjou na Kanjou no Nokotta Sanbun no Ni wa Sakamuke ga Kininaru Kanjou Uragiri wa Boku no Namae wo Shitteiru you de Shiranai no wo Boku wa Shitteiru Rusu Surume Medaka Kazunoko Koedame Medaka... Kono Medaka wa Sakki to Chigau Yatsu Dakara Ikeno Medaka no Hou Dakara Raayu Yuuteimiyaoukimukou Pepepepepepepepepepepepe Bichiguso Maru', 31, 1, 2000);
EXECUTE MAJ.CREER_EMPLOYE(100, 'Boby', 31, 1, 6666); -- salaire supp chef de service
EXECUTE MAJ.CREER_EMPLOYE(100, 'Boby', 31, 1, 1); -- erreur inconue

-- Suppression des traces du tests
SELECT * FROM EMPLOYE WHERE nuempl = 100;
DELETE FROM EMPLOYE WHERE NUEMPL = 100;
COMMIT;

-- TESTS MODIF_DUREE_HEBDO
EXECUTE MAJ.MODIF_DUREE_HEBDO(20, 23); -- passe
EXECUTE MAJ.MODIF_DUREE_HEBDO(30, 23); -- augmente hebdo
EXECUTE MAJ.MODIF_DUREE_HEBDO(20, 99); -- employe inexistant
EXECUTE MAJ.MODIF_DUREE_HEBDO(40, 23); -- dépasse la limite de 35h
EXECUTE MAJ.MODIF_DUREE_HEBDO(1234567890, 23); -- dépassement de la capacité du nombre
ROLLBACK;

-- TESTS MODIF_SALAIRE
EXECUTE MAJ.MODIF_SALAIRE(3501, 20); -- passe
EXECUTE MAJ.MODIF_SALAIRE(0, 20); -- salaire non diminuable
EXECUTE MAJ.MODIF_SALAIRE(3501, 99); -- employe inexistant
EXECUTE MAJ.MODIF_SALAIRE(1234567890, 20); -- dépassement de la capacité du nombre
ROLLBACK;

-- TESTS MODIF_TRAVAIL
EXECUTE MAJ.MODIF_TRAVAIL(5, 237, 23); -- passe
EXECUTE MAJ.MODIF_TRAVAIL(50, 237, 23); -- temps de travail invalide
EXECUTE MAJ.MODIF_TRAVAIL(5, 999, 23); -- le projet n’existe pas
EXECUTE MAJ.MODIF_TRAVAIL(1234567890, 237, 23); -- dépassement de la capacité du nombre
EXECUTE MAJ.MODIF_TRAVAIL(5, 200, 23); -- aucune ligne trouvé
EXECUTE MAJ.MODIF_TRAVAIL(5, 200, 99); -- aucune ligne trouvé
ROLLBACK;

-- TESTS INSERT_TRAVAIL
EXECUTE MAJ.INSERT_TRAVAIL(23, 103, 0);
EXECUTE MAJ.INSERT_TRAVAIL(23, 103, 50);
ROLLBACK;

EXECUTE MAJ.INSERT_TRAVAIL(23, 103, 50);  -- Insertion réussie
EXECUTE MAJ.INSERT_TRAVAIL(23, 103, 0);    -- temps de travail invalide
EXECUTE MAJ.INSERT_TRAVAIL(23, 103, 50);   -- Tentative d'insertion d'un enregistrement existant
EXECUTE MAJ.INSERT_TRAVAIL(23, 999, 50);   -- Le projet auquel il est affecté n’existe pas
EXECUTE MAJ.INSERT_TRAVAIL(23, 103, 999999);  -- Une valeur (nombre) dépasse le nombre de caractères autorisés
EXECUTE MAJ.INSERT_TRAVAIL(99, 999, 50);    -- Aucune ligne trouvée (employé et projet inexistants)


-- TESTS INSERT_SERVICE
EXECUTE MAJ.INSERT_SERVICE(6, 'nouveau_service', 20)
ROLLBACK;

EXECUTE MAJ.INSERT_SERVICE(6, 'nouveau_service', 20);  -- Insertion réussie
EXECUTE MAJ.INSERT_SERVICE(7, 'nouveau_service', 99);  -- Le chef ou le service n’existe pas
EXECUTE MAJ.INSERT_SERVICE(8, 'a'*300, 20);  -- Une valeur (chaine de caractère) dépasse le nombre de caractères autorisés
EXECUTE MAJ.INSERT_SERVICE(99999999999, 'nouveau_service', 20);  -- Une valeur (nombre) dépasse le nombre de caractères autorisés
EXECUTE MAJ.INSERT_SERVICE(99, 'service_inexistant', 99);  -- Aucune ligne trouvée (service et chef inexistants)
EXECUTE MAJ.INSERT_SERVICE(NULL, NULL, NULL);  -- Cela peut provoquer une erreur inconnue selon les contraintes.
