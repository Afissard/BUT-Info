-- Suppression des tables
DROP TABLE Avis;
DROP TABLE Trajet;
DROP TABLE Conducteur;
DROP TABLE Utilisateur;
DROP TABLE Vehicule;
DROP TABLE Ville;

-- Création des tables
CREATE TABLE Ville (
	nomVille VARCHAR(64) PRIMARY KEY
);

CREATE TABLE Vehicule (
    matricule VARCHAR(64) PRIMARY KEY,
    type VARCHAR(64),
    marque VARCHAR(64),
    energie VARCHAR(64),
    nbPlace INT,
    dateMiseEnCirculation DATE,
    idConducteur INT
);

CREATE TABLE Utilisateur (
	idUtilisateur INT PRIMARY KEY,
	nom VARCHAR(64),
	dateAgrement DATE,
	adresse VARCHAR(64),
	CONSTRAINT adresse FOREIGN KEY (adresse) REFERENCES Ville(nomVille)
);

CREATE TABLE Conducteur (
	idConducteur INT PRIMARY KEY,
	idUtilisateur INT,
	CONSTRAINT idUtilisateur FOREIGN KEY (idUtilisateur) REFERENCES Utilisateur(idUtilisateur),
	vehicule VARCHAR(64),
	CONSTRAINT vehicule FOREIGN KEY (vehicule) REFERENCES Vehicule(matricule)
);

CREATE TABLE Trajet (
	idTrajet INT PRIMARY KEY, 
    dateTrajet DATE,
    HeureDepart INT,
    nuPassager INT,
    longueur INT,
    tarif INT,
    idConducteurTraj int,
    CONSTRAINT idConducteurTraj FOREIGN KEY (idConducteurTraj) REFERENCES Conducteur(idConducteur),
    idUtilisateurTraj int,
    CONSTRAINT idUtilisateurTraj FOREIGN KEY (idUtilisateurTraj) REFERENCES Utilisateur(idUtilisateur),
    villeDép VARCHAR(64),
    CONSTRAINT villeDép FOREIGN KEY (villeDép) REFERENCES Ville(nomVille),
    villeArr VARCHAR(64),
    CONSTRAINT villeArr FOREIGN KEY (villeArr) REFERENCES Ville(nomVille)
);

CREATE TABLE Avis (
	idAvis INT PRIMARY KEY,
	avis NUMBER(1,0),
	idUtilisateurAvis INT,
	CONSTRAINT idUtilisateurAvis FOREIGN KEY (idUtilisateurAvis) REFERENCES Utilisateur(idUtilisateur),
    idTrajetAvis INT,
	CONSTRAINT idTrajetAvis FOREIGN KEY (idTrajetAvis) REFERENCES Trajet(idTrajet)
);

-- Insertion de donnée dans la base de donnée
-- Ville
INSERT INTO Ville (nomVille) VALUES ('Nantes');
INSERT INTO Ville (nomVille) VALUES ('Cholet');
INSERT INTO Ville (nomVille) VALUES ('Brest');
INSERT INTO Ville (nomVille) VALUES ('Paris');
INSERT INTO Ville (nomVille) VALUES ('Lyon');
INSERT INTO Ville (nomVille) VALUES ('Marseille');
INSERT INTO Ville (nomVille) VALUES ('Nice');
INSERT INTO Ville (nomVille) VALUES ('Toulouse');
INSERT INTO Ville (nomVille) VALUES ('Lille');
INSERT INTO Ville (nomVille) VALUES ('Madrid');
INSERT INTO Ville (nomVille) VALUES ('Amsterdam');
INSERT INTO Ville (nomVille) VALUES ('London');


-- Vehicule
INSERT INTO Vehicule (matricule, type, marque, energie, nbPlace, dateMiseEnCirculation, idConducteur) VALUES ('ABC123', 'Voiture', 'Toyota', 'Essence', 4, TO_DATE('01-01-2010', 'DD-MM-YYYY'), 1);
INSERT INTO Vehicule (matricule, type, marque, energie, nbPlace, dateMiseEnCirculation, idConducteur) VALUES ('XYZ789', 'Motorcycle', 'Harley Davidson', 'Petrol', 2, TO_DATE('01-01-2018', 'DD-MM-YYYY'), 2);
INSERT INTO Vehicule (matricule, type, marque, energie, nbPlace, dateMiseEnCirculation, idConducteur) VALUES ('DEF456', 'SUV', 'Jeep', 'Diesel', 5, TO_DATE('01-01-2015', 'DD-MM-YYYY'), 3);
INSERT INTO Vehicule (matricule, type, marque, energie, nbPlace, dateMiseEnCirculation, idConducteur) VALUES ('GHI789', 'Truck', 'Ford', 'Petrol', 3, TO_DATE('01-01-2019', 'DD-MM-YYYY'), 4);
INSERT INTO Vehicule (matricule, type, marque, energie, nbPlace, dateMiseEnCirculation, idConducteur) VALUES ('LMN123', 'Compact', 'Honda', 'Diesel', 5, TO_DATE('01-01-2016', 'DD-MM-YYYY'), 8);
INSERT INTO Vehicule (matricule, type, marque, energie, nbPlace, dateMiseEnCirculation, idConducteur) VALUES ('OPQ456', 'Sedan', 'Mercedes', 'Petrol', 4, TO_DATE('01-01-2017', 'DD-MM-YYYY'), 9);
INSERT INTO Vehicule (matricule, type, marque, energie, nbPlace, dateMiseEnCirculation, idConducteur) VALUES ('RST789', 'Hatchback', 'Volkswagen', 'Petrol', 3, TO_DATE('01-01-2018', 'DD-MM-YYYY'), 10);
INSERT INTO Vehicule (matricule, type, marque, energie, nbPlace, dateMiseEnCirculation, idConducteur) VALUES ('UVW101', 'Minivan', 'Toyota', 'Hybrid', 7, TO_DATE('01-01-2019', 'DD-MM-YYYY'), 11);
INSERT INTO Vehicule (matricule, type, marque, energie, nbPlace, dateMiseEnCirculation, idConducteur) VALUES ('XYZ202', 'SUV', 'Jeep', 'Diesel', 5, TO_DATE('01-01-2020', 'DD-MM-YYYY'), 12);
INSERT INTO Vehicule (matricule, type, marque, energie, nbPlace, dateMiseEnCirculation, idConducteur) VALUES ('ABC303', 'Electric', 'Tesla', 'Electricity', 4, TO_DATE('01-01-2021', 'DD-MM-YYYY'), 13);


-- Utilisateur
INSERT INTO Utilisateur (idUtilisateur, nom, dateAgrement, adresse) VALUES (1, 'John Doe', TO_DATE('01-01-2015', 'DD-MM-YYYY'), 'Nantes');
INSERT INTO Utilisateur (idUtilisateur, nom, dateAgrement, adresse) VALUES (2, 'Alice Wonderland', TO_DATE('01-01-2016', 'DD-MM-YYYY'), 'Paris');
INSERT INTO Utilisateur (idUtilisateur, nom, dateAgrement, adresse) VALUES (3, 'Bob Marley', TO_DATE('01-01-2014', 'DD-MM-YYYY'), 'Lyon');
INSERT INTO Utilisateur (idUtilisateur, nom, dateAgrement, adresse) VALUES (4, 'Eva Green', TO_DATE('01-01-2017', 'DD-MM-YYYY'), 'Marseille');
INSERT INTO Utilisateur (idUtilisateur, nom, dateAgrement, adresse) VALUES (8, 'Lucas Smith', TO_DATE('01-01-2015', 'DD-MM-YYYY'), 'Nice');
INSERT INTO Utilisateur (idUtilisateur, nom, dateAgrement, adresse) VALUES (9, 'Emma Johnson', TO_DATE('01-01-2016', 'DD-MM-YYYY'), 'Toulouse');
INSERT INTO Utilisateur (idUtilisateur, nom, dateAgrement, adresse) VALUES (10, 'Olivia Davis', TO_DATE('01-01-2017', 'DD-MM-YYYY'), 'Lille');
INSERT INTO Utilisateur (idUtilisateur, nom, dateAgrement, adresse) VALUES (11, 'Liam Wilson', TO_DATE('01-01-2018', 'DD-MM-YYYY'), 'Madrid');
INSERT INTO Utilisateur (idUtilisateur, nom, dateAgrement, adresse) VALUES (12, 'Noah Martinez', TO_DATE('01-01-2019', 'DD-MM-YYYY'), 'Amsterdam');
INSERT INTO Utilisateur (idUtilisateur, nom, dateAgrement, adresse) VALUES (13, 'Sophia Taylor', TO_DATE('01-01-2020', 'DD-MM-YYYY'), 'London');


-- Conducteur
INSERT INTO Conducteur (idConducteur, idUtilisateur, vehicule) VALUES (1, 1, 'ABC123');
INSERT INTO Conducteur (idConducteur, idUtilisateur, vehicule) VALUES (2, 2, 'XYZ789');
INSERT INTO Conducteur (idConducteur, idUtilisateur, vehicule) VALUES (3, 3, 'DEF456');
INSERT INTO Conducteur (idConducteur, idUtilisateur, vehicule) VALUES (4, 4, 'GHI789');
INSERT INTO Conducteur (idConducteur, idUtilisateur, vehicule) VALUES (8, 8, 'LMN123');
INSERT INTO Conducteur (idConducteur, idUtilisateur, vehicule) VALUES (9, 9, 'OPQ456');
INSERT INTO Conducteur (idConducteur, idUtilisateur, vehicule) VALUES (10, 10, 'RST789');
INSERT INTO Conducteur (idConducteur, idUtilisateur, vehicule) VALUES (11, 11, 'UVW101');
INSERT INTO Conducteur (idConducteur, idUtilisateur, vehicule) VALUES (12, 12, 'XYZ202');
INSERT INTO Conducteur (idConducteur, idUtilisateur, vehicule) VALUES (13, 13, 'ABC303');


-- Trajet
INSERT INTO Trajet (idTrajet, dateTrajet, HeureDepart, nuPassager, longueur, tarif, idConducteurTraj, idUtilisateurTraj, villeDép, villeArr) VALUES (1, TO_DATE('02-01-2023', 'DD-MM-YYYY'), 8, 3, 150, 30, 1, 1, 'Nantes', 'Cholet');
INSERT INTO Trajet (idTrajet, dateTrajet, HeureDepart, nuPassager, longueur, tarif, idConducteurTraj, idUtilisateurTraj, villeDép, villeArr) VALUES (2, TO_DATE('03-01-2023', 'DD-MM-YYYY'), 10, 2, 200, 40, 2, 3, 'Paris', 'Lyon');
INSERT INTO Trajet (idTrajet, dateTrajet, HeureDepart, nuPassager, longueur, tarif, idConducteurTraj, idUtilisateurTraj, villeDép, villeArr) VALUES (3, TO_DATE('04-01-2023', 'DD-MM-YYYY'), 12, 1, 300, 50, 3, 2, 'Lyon', 'Marseille');
INSERT INTO Trajet (idTrajet, dateTrajet, HeureDepart, nuPassager, longueur, tarif, idConducteurTraj, idUtilisateurTraj, villeDép, villeArr) VALUES (4, TO_DATE('05-01-2023', 'DD-MM-YYYY'), 14, 3, 250, 45, 4, 4, 'Marseille', 'Paris');
INSERT INTO Trajet (idTrajet, dateTrajet, HeureDepart, nuPassager, longueur, tarif, idConducteurTraj, idUtilisateurTraj, villeDép, villeArr) VALUES (8, TO_DATE('09-01-2023', 'DD-MM-YYYY'), 10, 3, 220, 35, 8, 11, 'Nice', 'Madrid');
INSERT INTO Trajet (idTrajet, dateTrajet, HeureDepart, nuPassager, longueur, tarif, idConducteurTraj, idUtilisateurTraj, villeDép, villeArr) VALUES (9, TO_DATE('10-01-2023', 'DD-MM-YYYY'), 12, 2, 180, 30, 9, 10, 'Toulouse', 'Lille');
INSERT INTO Trajet (idTrajet, dateTrajet, HeureDepart, nuPassager, longueur, tarif, idConducteurTraj, idUtilisateurTraj, villeDép, villeArr) VALUES (10, TO_DATE('11-01-2023', 'DD-MM-YYYY'), 14, 1, 250, 40, 10, 12, 'Lille', 'Amsterdam');
INSERT INTO Trajet (idTrajet, dateTrajet, HeureDepart, nuPassager, longueur, tarif, idConducteurTraj, idUtilisateurTraj, villeDép, villeArr) VALUES (11, TO_DATE('12-01-2023', 'DD-MM-YYYY'), 16, 3, 300, 50, 11, 8, 'Madrid', 'Nice');
INSERT INTO Trajet (idTrajet, dateTrajet, HeureDepart, nuPassager, longueur, tarif, idConducteurTraj, idUtilisateurTraj, villeDép, villeArr) VALUES (12, TO_DATE('13-01-2023', 'DD-MM-YYYY'), 9, 2, 200, 35, 12, 13, 'Amsterdam', 'London');
INSERT INTO Trajet (idTrajet, dateTrajet, HeureDepart, nuPassager, longueur, tarif, idConducteurTraj, idUtilisateurTraj, villeDép, villeArr) VALUES (13, TO_DATE('14-01-2023', 'DD-MM-YYYY'), 11, 1, 280, 45, 13, 9, 'London', 'Toulouse');


-- Avis
INSERT INTO Avis (idAvis, avis, idUtilisateurAvis, idTrajetAvis) VALUES (1, 1, 2, 1);
INSERT INTO Avis (idAvis, avis, idUtilisateurAvis, idTrajetAvis) VALUES (2, 1, 3, 2);
INSERT INTO Avis (idAvis, avis, idUtilisateurAvis, idTrajetAvis) VALUES (3, 0, 2, 3);
INSERT INTO Avis (idAvis, avis, idUtilisateurAvis, idTrajetAvis) VALUES (4, 1, 4, 4);
INSERT INTO Avis (idAvis, avis, idUtilisateurAvis, idTrajetAvis) VALUES (8, 1, 13, 8);
INSERT INTO Avis (idAvis, avis, idUtilisateurAvis, idTrajetAvis) VALUES (9, 1, 8, 9);
INSERT INTO Avis (idAvis, avis, idUtilisateurAvis, idTrajetAvis) VALUES (10, 0, 9, 10);
INSERT INTO Avis (idAvis, avis, idUtilisateurAvis, idTrajetAvis) VALUES (11, 1, 10, 11);
INSERT INTO Avis (idAvis, avis, idUtilisateurAvis, idTrajetAvis) VALUES (12, 1, 12, 12);
INSERT INTO Avis (idAvis, avis, idUtilisateurAvis, idTrajetAvis) VALUES (13, 0, 11, 13);



-- 1 : Liste des conducteurs qui effectuent des trajets entre Nantes et Cholet.
SELECT C.idConducteur, U.nom
FROM Conducteur C
JOIN Utilisateur U ON C.idUtilisateur = U.idUtilisateur
JOIN Trajet T ON C.idConducteur = T.idConducteurTraj
JOIN Ville V1 ON T.villeDép = V1.nomVille
JOIN Ville V2 ON T.villeArr = V2.nomVille
WHERE V1.nomVille = 'Nantes' AND V2.nomVille = 'Cholet';

-- 2 : Trouvez pour chaque conducteur habitant Nantes et conduisant un véhicule immatriculé avant 2000, le type d’énergie utilisé par le véhicule (essence, diesel etc.)
SELECT U.idUtilisateur, U.nom, V.type, V.energie
FROM Utilisateur U
JOIN Conducteur C ON U.idUtilisateur = C.idUtilisateur
JOIN Vehicule V ON C.vehicule = V.matricule
JOIN Ville V1 ON U.adresse = V1.nomVille
WHERE V1.nomVille = 'Nantes' AND EXTRACT(YEAR FROM V.dateMiseEnCirculation) < 2000;


-- 3 : Trouvez les conducteurs Nantais qui n’ont reçu que des avis positifs.
SELECT C.idConducteur, U.nom
FROM Conducteur C
JOIN Utilisateur U ON C.idUtilisateur = U.idUtilisateur
WHERE U.adresse = 'Nantes'
AND C.idConducteur NOT IN (
    SELECT CA.idConducteurTraj
    FROM Avis AV
    JOIN Trajet CA ON AV.idTrajetAvis = CA.idTrajet
    WHERE AV.avis = 0
);


-- 4 : Quels sont les conducteurs qui ont parcouru le plus de trajet
SELECT C.idConducteur, U.nom, COUNT(T.idTrajet) AS NombreDeTrajets
FROM Conducteur C
JOIN Utilisateur U ON C.idUtilisateur = U.idUtilisateur
JOIN Trajet T ON C.idConducteur = T.idConducteurTraj
GROUP BY C.idConducteur, U.nom
ORDER BY NombreDeTrajets DESC
FETCH FIRST 1 ROW ONLY;

-- 5 : Lister les trajets pour lesquels au moins un des passagers habite la ville d’arrivée
SELECT T.idTrajet, T.dateTrajet, T.villeDép, T.villeArr
FROM Trajet T
JOIN Utilisateur U ON T.idUtilisateurTraj = U.idUtilisateur
WHERE U.adresse = T.villeArr;

-- 6 : Lister les conducteurs qui ont reçu leur agrément après l’année 2014
SELECT C.idConducteur, U.nom, U.dateAgrement
FROM Conducteur C
JOIN Utilisateur U ON C.idUtilisateur = U.idUtilisateur
WHERE EXTRACT(YEAR FROM U.dateAgrement) > 2014;

-- 7 : Lister les conducteurs qui n’ont jamais transporté de passager
SELECT C.idConducteur, U.nom
FROM Conducteur C
JOIN Utilisateur U ON C.idUtilisateur = U.idUtilisateur
WHERE C.idConducteur NOT IN (
    SELECT DISTINCT idConducteurTraj
    FROM Trajet
    WHERE nuPassager > 0
);

-- 8 : Trouver le conducteur qui a transporté le plus de passagers.
SELECT C.idConducteur, U.nom, SUM(T.nuPassager) AS TotalPassagersTransportés
FROM Conducteur C
JOIN Utilisateur U ON C.idUtilisateur = U.idUtilisateur
JOIN Trajet T ON C.idConducteur = T.idConducteurTraj
GROUP BY C.idConducteur, U.nom
ORDER BY TotalPassagersTransportés DESC
FETCH FIRST 1 ROW ONLY;

-- 9 : Trouver le conducteur le plus riche (celui ayant le plus de gains cumulés)
SELECT C.idConducteur, U.nom, SUM(T.tarif) AS TotalEarnings
FROM Conducteur C
JOIN Utilisateur U ON C.idUtilisateur = U.idUtilisateur
JOIN Trajet T ON C.idConducteur = T.idConducteurTraj
GROUP BY C.idConducteur, U.nom
ORDER BY TotalEarnings DESC
FETCH FIRST 1 ROW ONLY;

-- 10 : Trouver le top 5 des utilisateurs ayant parcouru la plus grande distance cumulée
SELECT U.idUtilisateur, U.nom, SUM(T.longueur) AS TotalDistance
FROM Utilisateur U
JOIN Trajet T ON U.idUtilisateur = T.idUtilisateurTraj
GROUP BY U.idUtilisateur, U.nom
ORDER BY TotalDistance DESC
FETCH FIRST 5 ROWS ONLY;
