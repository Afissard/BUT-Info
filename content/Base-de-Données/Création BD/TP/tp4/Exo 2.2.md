# Exo 1
Soit la structure de données suivante : 
```
Fournisseur (#nofournisseurs, nom, ville)
Mafourniture (#nofournisseur, noproduit, quantité)
Produit (#noproduit, nom-produit, couleur, origine, PVU) 
```

Exprimez chacune des questions suivantes en algèbre relationnelle : 
1 ▶ Trouvez les numéros des fournisseurs qui fournissent quelque chose.
```
R1 <- Mafourniture{nofounisseur}
``` 
2 ▶ Trouvez les numéros des fournisseurs qui ne fournissent rien.
```
R2 <- Diff(Fournisseur{nofournisseur}, MaFourniture{nofounisseur})
```
3 ▶ Trouvez les numéros des fournisseurs qui fournissent au moins le produit P6.
```
R3 <- MaFourniture[noproduit='P6']{nofournisseur}
```
4 ▶ Trouvez les numéros des fournisseurs qui fournissent quelque chose d’autre que P6.
```
R4 <- Diff(Fournisseur{nofournisseur}, MaFourniture[noproduit = 'P6'{nofournisseur})
Autre solotion plus simple :
R4 <- MaFourniture[noproduit != 'P6']{nofournisseur}
```
5 ▶ Trouvez les numéros des fournisseurs qui fournissent quelque chose mais pas P6.
``` 
On peut utilisé les résultat précédant :
R5 <- Diff(R3, R1)
```
6 ▶ Trouvez les numéros des fournisseurs qui fournissent au moins un produit originaire de leur ville.
```
A <- Join(Fournisseur, Fourniture / Founisseur.nofournisseur = Mafourniture.nofournisseur)
B <- Join(A, Produit / A.produit = produit.noproduit AND A.ville = produit.origine)
R6 <- B{nofournisseur}
```
7 ▶ Trouvez les numéros des fournisseurs qui fournissent tous les produits originaires de Nantes 
```
R7 <- Div(MaFourniture{nofournisseur, noproduit}, Produit[origine = 'Nantes']{noproduit})
```
8 ▶ Trouvez les numéros des fournisseurs dont toutes les fournitures contiennent des produits verts.
```
C <- Join(MaFourniture, Produit / MaFourniture.noproduit = Produit.noproduit AND Produi.couleur != 'Vert'){nofournisseur}
R8 <- Diff(R1, C)
```
# Exo 2
Soit la structure de données suivante : 
```
Personne (matricule, nom, sexe, âge, profession, revenu) 
Livre (no-livre, titre, auteur, thème, matricule-possesseur) 
Préférence (matricule, no-livre)
```
Exprimer chacune des questions suivantes en algèbre relationnelle : 
1 ▶ Donnez les titres des livres possédés par la personne de matricule 99.
```
R1 <- Livre[matricule-possesseur = '99']{titre}
```
2 ▶ Donnez le nom du possesseur du livre n° 36.
```
R2 <- Join(Personne, Livre / Personne.matricule = Livre.matricule-possesseur AND [Livre.no-livre = '36']){nom}
```

```
Correction :
A <- Join(Personne, Livre / Personne.matricule = Livre.matricule-possesseur)
R2 <- A[no-livre = '36']{nom}
```
3 ▶ Sélectionner les noms des personnes qui ont parmi leurs livres préférés au moins un livre du thème « science-fiction ».
```
D <- Join(Livre, Préférence / Livre.no-livre = Préférence.no-livre)
E <- Join(Personne, D / Personne.matricule = D.matricule-posssesseur)
R3 <- E[E.thème = 'SF']{nom}
```
4 ▶ Sélectionner les matricules des personnes dont tous les livres (possédé ?) sont du thème « policier ».
```
F <- A[thème = 'Policier']{Personne.matricule}
R4 <- Diff(Personne.matricule, F)
```

```
Correction
G <- Join(Livre, Personne / Personne.matricule = livre.matricule-possesseur and livre.thème = 'Policier'){matricule}
R4 <- Diff(Personne{matricule}, G)
```
5 ▶ Sélectionner les noms des personnes qui ont parmi leurs livres préférés, tous les livres de l’auteur Hergé.
```
H <- div(reference, livre[auteur = 'Hergé']{no-livre})
R5 <- Join(Personne, H / Personne.matricule = H.matricule){nom}
```