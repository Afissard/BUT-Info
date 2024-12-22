```sql
-- Reset des tables precedement créer
DROP TABLE Concerne;
DROP TABLE Travail;
DROP TABLE Projet;
DROP TABLE Service;
DROP TABLE Employé;

-- Creation des tables
CREATE TABLE Employé 
(
    numEmpl NUMBER PRIMARY KEY,
    nomEmpl CHAR(16),
    hebdo NUMBER,
    CHECK (hebdo <= 25),
    salaire NUMBER,
    affect CHAR(16)
);

CREATE TABLE Service
(
    numServ NUMBER PRIMARY KEY,
    nomServ CHAR(16),
    chefServ NUMBER,
    FOREIGN KEY (chefServ)  REFERENCES Employé(numEmpl)
);

CREATE TABLE Projet 
(
    numProj NUMBER PRIMARY KEY,
    nomProj CHAR(16),
    respProj NUMBER,
    FOREIGN KEY (respProj)  REFERENCES Employé(numEmpl)
);

CREATE TABLE Travail 
(
    numEmpl NUMBER PRIMARY KEY,
    FOREIGN KEY (numEmpl) REFERENCES Employé(numEmpl),
    numProj NUMBER,
    FOREIGN KEY(numProj) REFERENCES Projet(numProj),
    dureeTrav DATE
);

CREATE TABLE Concerne 
(   
    numServ NUMBER,
    numProj NUMBER,
    FOREIGN KEY(numServ) REFERENCES Service(numServ),
    FOREIGN KEY(numProj) REFERENCES Projet(numProj)
);

-- Insertion de donnée dans les tables
INSERT INTO Employé(numEmpl, nomEmpl, hebdo, salaire) VALUES(01, Arthur, 1, 5000); // pas fini
```