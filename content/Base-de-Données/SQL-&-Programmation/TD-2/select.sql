SELECT * from EMPLOYE;
SELECT * from SERVICE;
SELECT * from PROJET;
SELECT * from TRAVAIL;
SELECT * from CONCERNE;

-- test 2.1

SELECT e.nuempl, e.NOMEMPL
FROM employe e
LEFT JOIN service s ON e.nuempl = s.chef
LEFT JOIN projet p ON e.nuempl = p.resp
WHERE s.chef IS NULL AND p.resp IS NULL;

SELECT e.nuempl, e.NOMEMPL
FROM employe e
JOIN service s ON e.nuempl = s.chef;

SELECT e.nuempl, e.NOMEMPL
FROM employe e
JOIN projet p ON e.nuempl = p.resp;

SELECT e.nuempl, e.NOMEMPL
FROM employe e
JOIN service s ON e.nuempl = s.chef
JOIN projet p ON e.nuempl = p.resp;

SELECT e.nuempl, e.NOMEMPL
FROM employe e
JOIN travail t ON e.nuempl = t.nuempl
LEFT JOIN service s ON e.nuempl = s.chef
LEFT JOIN projet p ON e.nuempl = p.resp
WHERE s.chef IS NULL AND p.resp IS NULL;

-- test 2.2

SELECT p.nuproj, p.nomproj
FROM projet p
JOIN travail t ON p.nuproj = t.nuproj
JOIN concerne c ON p.nuproj = c.nuproj;

SELECT p.nuproj, p.nomproj
FROM projet p
JOIN travail t ON p.nuproj = t.nuproj
LEFT JOIN concerne c ON p.nuproj = c.nuproj
WHERE c.nuproj IS NULL;

SELECT p.nuproj, p.nomproj
FROM projet p
JOIN employe e ON p.resp = e.nuempl;