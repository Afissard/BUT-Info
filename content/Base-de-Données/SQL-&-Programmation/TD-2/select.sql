SELECT * from EMPLOYE;
SELECT * from SERVICE;
SELECT * from PROJET;
SELECT * from TRAVAIL;
SELECT * from CONCERNE;

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
