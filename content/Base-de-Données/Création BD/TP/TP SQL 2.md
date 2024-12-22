# Récupération des tables
```sql
-- Récupération des table
CREATE TABLE employe as select * FROM basetd.employe
CREATE TABLE service as select * FROM basetd.service;
CREATE TABLE travail as select * FROM basetd.travail;
CREATE TABLE projet as select * FROM basetd.projet;
CREATE TABLE concerne as select * FROM basetd.concerne;
```
# Exo 1
```sql
-- 1.1
-- Question 1 : selection des employés qui travail sur tout les projets
SELECT nomempl FROM employe, travail WHERE employe.nuempl=travail.nuempl;

-- Question 2 : Sélectionnez les noms des employés possédant les numéros suivants : 20,30 et 42
SELECT nomempl FROM employe WHERE nuempl IN (20, 30, 42);

-- Question 3 : Sélectionnez les numéros des projets dont les noms commencent par la lettre ’C’
SELECT nuproj FROM projet WHERE nomproj LIKE 'c%';

-- Question 4 : Sélectionnez les numéros des employés dont le temps hebdomadaire n’est pas renseigné.
SELECT nuempl FROM employe WHERE hebdo IS  NULL; -- table vide ?!

-- Question 5 : Sélectionnez les noms des employés avec la durée de travail sur les projets. Si un employé n’est pas impliqué dans un projet son nom doit sortir dans le résultat de la requête.
SELECT nomempl, duree FROM employe, travail WHERE employe.nuempl = travail.nuempl(+);

-- Question 6 : Sélectionnez les numéros des projets sur lesquels travaille un employé.
SELECT nuproj, nuempl FROM travail;

-- Question 7 : Sélectionnez les numéros des employés qui ne travaillent pas sur des projets
SELECT nuempl FROM employe MINUS SELECT nuempl FROM travail;

-- 1.2
-- Question 1 : Sélectionnez les numéros et noms de tous les employés dans la table Employé 
SELECT nuempl, nomempl FROM employe;

-- Question 2 : Sélectionnez le nombre d’employés
SELECT COUNT(*) from employe;

-- Question 3 : Sélectionnez le temps hebdomadaire moyen de travail des employés
SELECT AVG(hebdo) FROM employe;

-- Question 4 : Sélectionnez la somme des durées consacrées par les employés aux projets
SELECT SUM(duree) FROM travail;

-- Question 5 : Affichez les noms des employés par ordre croissant 
SELECT nomempl FROM employe ORDER BY nomempl ASC;

-- Question 6 : Affichez les numéros des employés et la durée de travail consacrée par chacun des employés à chacun des projets. Les résultats doivent être triés par numéro employé (ordre décroissant)
SELECT nuempl, duree, nuproj FROM travail ORDER BY nuempl DESC;

-- Question 7 : Affichez le nom du service numéro 1 
SELECT nomempl FROM service WHERE nuserv = 1;

-- Question 8 : Affichez les noms des autres services
SELECT nomempl FROM service WHERE nuserv != 1;

-- Question 9 : Affichez les noms des employés qui ne travaillent pas 
SELECT nomempl FROM employe WHERE NOT nuempl IN (SELECT nuempl FROM travail); -- requete imbriquée mais existe plus simple
SELECT nomempl FROM employe MINUS SELECT nomempl FROM employe, travail WHERE employe.nuempl = travail.nuempl; -- Solution "plus simple" ...

-- Question 10 : Affichez les noms des employés dont le temps de travail hebdomadaire est compris entre 20 et 30 (deux versions)
SELECT nomempl FROM employe WHERE hebdo BETWEEN 20 and 30;
SELECT nomempl FROM employe WHERE hebdo >= 20 and hebdo <=30;
```

# Exo 2
```sql 
-- 2.1 
-- Question 1 : Sous-requête indépendante renvoyant une seule ligne 
-- Question 1.1 : Le nom de des employés qui sont affectés au service ’achat’; 
SELECT nomempl FROM employe WHERE affect = (
	SELECT nuserv FROM service WHERE nomserv = 'achat'
	)
; 
-- Question 1.2 : Le nom du projet sur lequel travail l’employé numéro 20
SELECT nomproj FROM projet WHERE nuproj = (
	SELECT nuproj FROM travail WHERE nuempl = 20
	)
; 

-- Question 2 : Sous-requête indépendante renvoyant plusieurs lignes 
-- Question 2.1 : Les noms des employés dont la durée de travail sur chacun des projets est de 5h 
SELECT nomempl FROM employe WHERE nuempl IN (
	SELECT nuempl FROM travail WHERE duree = 5
	)
; 
-- Question 2.2 :Les noms des employés qui ne travaillent pas sur des projets. Deux versions : in et all 
-- version "IN" :
SELECT nomempl FROM employe WHERE nuempl NOT IN (SELECT nuempl FROM travail); -- version "ALL" :
SELECT nomempl FROM employe WHERE nuempl != ALL (SELECT nuempl FROM travail); 

-- Question 3 : Sous-requête dépendante de la requête principale 
-- Question 3.1 : Les employés (sans doublons) qui travaillent sur des projets dans lesquels est impliqué le responsable du projet numéro 30. 
SELECT DISTINCT nomempl FROM employe WHERE nuempl IN ( 
	SELECT nuempl FROM travail WHERE nuproj IN ( 
		SELECT nuproj FROM projet WHERE resp = 30 
		) 
	) 
; -- ne respecte pas la consigne car requête indépendante 
SELECT DISTINCT nomempl FROM employe e, travail t WHERE e.nuempl = t.nuempl AND t.nuproj IN ( 
	SELECT nuproj FROM projet WHERE resp = 30 
	)
; -- meilleur solution, plus concis mais ne respect toujours pas la consigne
SELECT DISTINCT nomempl FROM employe e, travail t WHERE e.nuempl = t.nuempl AND t.nuproj IN ( 
	SELECT nuproj FROM projet p WHERE p.resp = 30 AND p.nuproj = t.nuproj 
	)
; -- Difference entre la deuxième version est que celle ci est dépendente 
-- Question 3.2 : Les noms des employés qui travaillent sur des projets (version avec exists) 
SELECT nomempl FROM employe e WHERE EXISTS (SELECT * FROM travail t WHERE e.nuempl = t.nuempl); 
-- Question 3.3 : Les employés qui ne travaillent pas sur des projets (version avec exists) -> TODO 
<<<<<<< HEAD
=======
 
-- 2.2 
-- Question 1 : Liste des noms de projets avec le nom du responsable et le nombre d’employés qui y travaillent 
SELECT p.nomproj, r.nomempl, NBempl 
	FROM employe r, projet p, (
		SELECT nuproj, COUNT(t.nuempl) as NBempl 
			FROM travail t GROUP BY t.nuproj) tNBempl 
		WHERE 
			p.resp = r.nuempl AND 
			tNBempl.nuproj = p.nuproj 
; 
-- Correction 
SELECT nomproj, e1.nomempl, COUNT(*) 
	FROM projet p, employe e1, travail t, employe e2 
	WHERE 
		p.resp = e1.nuempl AND 
		p.nuproj = t.nuproj AND 
		t.nuempl = e2.nuempl 
		GROUP BY nomproj, e1.nomempl 
; 
-- Question 2 : Liste des noms de projets avec la totalisation du nombre d’heures passées par les employés qui y travaillent 
SELECT p.nomproj, tTotDuree.TotDuree 
	FROM projet p, (
		SELECT t.nuproj,SUM(t.duree) AS TotDuree 
			FROM travail t GROUP BY t.nuproj) tTotDuree 
		WHERE 
			tTotDuree.nuproj = p.nuproj 
; 
-- Correction 
SELECT nomproj, SUM(duree) AS totalNBh 
	FROM projet p, travail t 
	WHERE 
		p.nuproj = t.nuproj 
		GROUP BY nomproj 
; 
-- Question 3 : Liste des noms de projets avec pour chaque service concerné, le nom du service et le nombre d’employés de ce service qui travaillent sur ce projet 
SELECT nomproj, nomserv, COUNT(*) AS NBemploye 
	FROM employe e, projet p, travail t, service s, concerne c 
	WHERE 
		p.nuproj = c.nuproj AND 
		c.nuserv = s.nuserv AND 
		s.nuserv = e.affect AND 
		e.nuempl = t.nuempl 
		GROUP BY nomproj, s.nomserv 
; 
-- Question 4 : Liste des employés qui travaillent sur tous les projets 
SELECT nomempl 
	FROM employe e, travail t, projet p 
	WHERE 
		e.nuempl = t.nuempl AND 
		t.nuproj = p.nuproj 
		GROUP BY e.nuempl, e.nomempl HAVING COUNT(*) = (
			SELECT COUNT(*) FROM projet
		) 
; 
-- Question 5 : Pour le service Achat, trouvez le nom du chef de service et le nombre d’employés qui y sont affectés 
SELECT e1.nomempl, COUNT(*) 
	FROM employe e1, service s, employe e2 
	WHERE 
		e1.nuempl = s.chef AND 
		s.nomserv = 'achat' AND 
		e2.affect = s.nuserv 
		GROUP BY e1.nomempl 
; 
-- Question 6 : Liste des employés qui travaillent sur au moins un des projets sur lesquels ’Marcel’ travaille. 
SELECT e.nomempl 
	FROM employe e, projet p, travail t 
	WHERE 
		e.nuempl = t.nuempl AND 
		t.nuproj = p.nuproj AND 
		p.nuproj IN ( 
			SELECT nuproj 
				FROM travail t, employe e 
				WHERE 
					e.nuempl = t.nuempl AND 
					e.nomempl = 'marcel' 
			) 
		AND  e.nomempl != 'marcel' -- Marcel travail avec lui même 
; 
>>>>>>> 95eda11a0f81bb1dbce4b2800aff5cdf990a59f2
```