Liens MOCODO: https://mocodo.net/

## Exo télé :
A refaire si motivé

## Exo candidat :
A refaire si motivé

## Exo médiathèque :
CD: id, titre, artiste
A-FAIT, 1N ARTISTE, NN CD
ARTISTE: id, cd

CHERCHE, 11 EXEMPLAIRE, 01 CD
EXEMPLAIRE: CD, ARTISTE
EMPRUNT, 11 PRET, 0N EXEMPLAIRE
PRET: duré, CD
PREND, 0N ABONNE, 11 PRET
ABONNE: id, nom, pret

## Exo repartition IUT Info
Inscrit, 1N ETUDIANT, 11 PROMO
PROMO: annee, departement

ETUDIANT: nom, prenom, naissance, adresse, bac, annee_bac
Appartien, NN ETUDIANT, 1N MODULE
MODULE : matiere
Responsable, 11 PROF, 1N MODULE
PROF : nom

Inclus, 1N ETUDIANT, 11 GRP_TD
GRP_TD: id

## Exo OPEN TTD
COMPAGNIE : nom
Etablie, 1N COMPAGNIE, 0N VILLE
VILLE : nom

Possede, 11 COMPAGNIE,1N LIGNE
Relie, 1N LIGNE, 0N VILLE

LIGNE : id

Composée, 1N LIGNE, 2N STATION
STATION : nom
Deservie, 1N STATION, 1N BUS, 1N TRAIN
TRAIN : id

:
BUS : id

## Vente MP3
MP3 : id, marque, qualité, prix
Vend, 1N MARCHAND, 1N MP3
MARCHAND : nom; lieu, MP3
Achete, 0N ACHETEUR, 0N MARCHAND: rapport qualité/prix
ACHETEUR : nom, argent

## YAKAFAIRE Inc
### Modélisation par schéma d'association
YAKAFAIRE : argent, patron
Vend, 11 YAKAFAIRE, 1N PRODUIT, 0N CLIENT
PRODUIT : STOCK; couleur, gamme, matériaux, date fabrication, origine fabrication, certification
En-Stock, 0N PRODUIT, 0N STOCK
STOCK : PRODUIT

CLIENT : nom, adresse, téléphone, date vente
### Modélisation des donnée sur ?
flemme ?

## YAKAF
YAKAF : nom, numéro registre commerce, adresse
Facture, 11 YAKAF, 1N CLIENT : coordonnées YAKAF, date, coordonnées CLIENT, produit
CLIENT : numéro, raison sociale, adresse

## MACHINTRUC Corp
Concerné, 1N SERVICE, 11 PROJET
SERVICE : nom
Affecté, 11 EMPLOYE, 11 SERVICE

PROJET : nom
Chef, 11 EMPLOYE, 11 SERVICE
EMPLOYE : matricule, temps travail hebdo

:
Responsable, 11 EMPLOYE, 1N PROJET

:
Travail, 11 EMPLOYE, 1N PROJET : somme heures travail <= temps travail hebdo
