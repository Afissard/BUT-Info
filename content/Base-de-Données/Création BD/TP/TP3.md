# Exo 1
## 1.1
1. cardinalité : 5 ligne
2. degré 4 colonne
3. clés de la relation : C, D sont les clés primaire de cette relation
4. Dépendance fonctionnelle A->B, A->C ? Oui pour A->B et non pour A->C.
# 1.2
1. Y a t'il des clef candidates ? : oui deux : CD, BD
2. Dépendance fonctionnel : A->B oui mais pas A->C
## 1.3
1. A est la clef de la relation R(A, B, C), on peut le prouvé avec, ici la propriété de transitivité, puis que A->B et B->C alors A->C donc A->B∪C donc A est la clef primaire.
2. Forme normale (plusieurs formes (3 Forme Normale dans le cour, mais 5 au total) : 1NF, 2NF, 3NF) de la relation : 2NF, la clef est composé d'un seul attribut alors tout le reste en dépend donc forme normale n°2 et donc valide aussi la forme normale n°1.
3. Est ce que la relation présenté est conforme à R ? C ne dépend pas de B donc inconforme. Autre analyse : analyse détaillé du tableau : celui donné est en 3NF or R est en 2NF.
## 1.4
Soit la relation R(A,B,C,D,E) munie de l’ensemble de dépendances fonctionnelles suivant: F= {A → B, A → D, A → E, E → C}
1. Clef de R : A->BDE & E->C donc par transitivité A->C donc A est clef
2. Forme normale de R : 1NF et 2NF : tout dépend de A par transitivité, pas 3NF
## 1.5
Soit la relation R(A,B,C,D,E) munie de l’ensemble de dépendances fonctionnelles suivant: F= {A → B, C → E, AE → D }
1. Pseudo transitivité avec C->E et AE->D donc AC->D
2. Forme normal de la relation : 1NF, la dépendance n'est que par augmentation (dependence indirecte)
# 1.6
1. Quelles sont les dépendances fonctionnelles incorporées dans cette relation ? :
	- D->B, 
	- A->E, C->E, D->E
	- Donc B->E
1. Quelles sont les clés candidates de la relation ? : ACD
2. Quelle la forme normale de la relation ? : 1NF car C->E, D->B
## 1.7
Quelle est l'interet des forme normales ? Donnez de exemples.
Cette forme permet de limiter les doublons
R1 (1NF)

| A   | B   | C   |
| --- | --- | --- |
| 1   | 2   | 3   |
| 1   | 3   | 1   |
| 2   | 1   | 2   |
| 3   | 2   | 3   |
| 4    |2     |3     |
R2 (2NF)

| A   | B   | C   |
| --- | --- | --- |
| 1   | 2   | 3   |
| 2   | 3   | 3   |
| 3   | 1   | 1   |
| 4   | 1   | 2   |
| 5   | 3   | 1    |
R3 (3NF)

| A   | B   | C   |
| --- | --- | --- |
| 1   | 3   | 2   |
| 2   | 2   | 3   |
| 3   | 3   | 4   |
| 4   | 2   | 5   |
| 5   | 3   | 4   | 
## 1.8
Une histoire de Dragon ...
Vérifié :
1. Non car un dragon -> Comportement Amoureux 
2. Non car pas de doublons
3. Oui
4. ?
# Exo 2
## 2.1
```
Fournisseur (#nofr, nom,ville) 
Fourniture(#nofr, #nopr, quantite) 
Produit(#nopr, nom, prix)
```
1. Clef étrangères appliqué : 
	- F-fournisseur : nofr -> nom, nofr -> ville
	- F-produit : nopr -> nom, nopr -> prix
	- F-fourniture : nofr & nopr-> quantité
	Raison application : 
	- nofr : parce qu'il faut que le fournisseur qui veut fournir soit enregistré dans la base de donnée
	- nopr : on veut que le produit soit dans le stock, un fournisseur ne peut fournir que les produit enregistré dans la base de donnée
## 2.2
```
Voiture(#noserie, modèle, marque, prix) 
Vendeur(#idvendeur, nom, tél) 
Vente(#idvendeur, #noserie, date, montant)
```
1. Clef étrangère appliqué : 
	- F-voiture : noserie -> modèle, noserie -> marque, noserie -> prix 
	- F-vendeur : idvendeur -> nom, idvendeur -> tel
	- F-vente : idvendeur & noserie -> date, idvendeur & noserie -> montant
	Raison application :
	- noserie : on veut que seules les voitures des série enregistré soit vendu
	- idvendeur : seul les vendeurs enregistré dans la base de donnée sont autorisé à vendre
2. Tuples de relation : (vendeur, voiture)
	Voiture :
	
	| n°serie | modèle | marque  | prix  |
	| ------- | ------ | ------- | ----- |
	| 1A10    | SUV    | Renault | 5000€ | 
	| 2C15    | Megane | Renault | 3000€ |
	
	Vendeur :
	
	| id  | nom  | tél            |
	| --- | ---- | -------------- |
	| 01  | Alex | 09 00 66 60 00 |
	| 02  | Bob  | 09 12 34 56 78 |

	Vente :
	| id vendeur | n°serie | date | montant  |
	| ---------- | ------- | ---- | -------- |
	| 01         | 1A10    | 23   | 5000€    |
	| 01         | 2C15    | 24   | 3000€    |
	| 02         | 2C15    | 26   | 2000.99€ |
	
3. Exemple d'insertion de tuple 
	- incorrecte : 
		- voiture essayé d'insérer une voiture qui n'existe pas
		- vendeur : essayé d'insérer un doublon d'ID
		- vente : essayé de faire une vente avec un vendeur inexistant
	- correcte :
		- voir tableau déjà fait
4. Est-il possible de supprimé un tuple dans la table vente ? : oui
5. Est-il possible de supprimé un vendeur ? : non il faudrait aussi supprimé les contrainte de clef étrangère auquel il est lié
