# Ma version
```Mocodo
Avis, 0N UTILISATEUR, 0N UTILISATEUR : avis
UTILISATEUR: numUtilisateur, nom, addresse, avis, dateAgrement

:
Conducteur, 1N UTILISATEUR, 11 VEHICULE: numUtilisateur
Passager, 0N UTILISATEUR, 1N TRAJET : numUtilisateur
:

VEHICULE: matricule, type, marque, energie, nbrPlace, dateMiseEnCirculation, numUtilisateur
Utilisé, 11 VEHICULE, 1N TRAJET
TRAJET: noConducteur, date, villeDep, villeAri, heureDes, passager, longueur, tarif

:
Arrivé, 11 VILLE, 11 TRAJET : nomVille
Départ, 11 VILLE, 11 TRAJET : nomVille

VILLE: nomVille
```

# Version Clement
```mocodo
:

PASSAGER:numUtilisateur, nom, adresse, avis

AvisP, 0N PASSAGER, 0N CONDUCTEUR : avis

:


SeFaitTransporter, 1N PASSAGER, 11 TRAJET

AvisC, 0N CONDUCTEUR, 0N PASSAGER : avis

CONDUCTEUR:numUtilisateur, nom, adresse, avis, date_agrement

Possède, 01 CONDUCTEUR, 11 VEHICULE


Départ, 0N VILLE, 11 TRAJET : nomVille

TRAJET: noTrajet, villeDep, villeAri, heureDes, conducteur, passager, longueur, tarif

Conduit, 1N CONDUCTEUR, 11 TRAJET 

VEHICULE: matricule, type, marque, energie, nbrPlace, dateMiseEnCirculation, numUtilisateur


VILLE: nomVille

Arrivé, 0N VILLE, 11 TRAJET : nomVille

:
```
# Version Léa
```mocodo
AVIS2, 1N PASSAGER, O1 CONDUCTEUR : avis
PASSAGER: numPassager, nom, adresse, avis, dateAgrement
SEFAITCONDUIRE, ON TRAJET, IN PASSAGER
ARRIVÉ, 11 VILLE, 11 TRAJET : nomVille

CONDUCTEUR: noConducteur, nom, addresse, avis, dateAgrement
AVIS1, 1N CONDUCTEUR, 01 PASSAGER : avis
TRAJET: noConducteur, date, villeDep, villeAri, heureDes, passager, longueur, tarif
VILLE: nomVille

POSSÈDE, IN CONDUCTEUR, ON VEHICULE
VEHICULE: matricule, type, marque, energie, nbrPlace, dateMiseEnCirculation, numUtilisateur
UTILISÉ, 11 VEHICULE, 1N TRAJET
DÉPART, 11 VILLE, 11 TRAJET : nomVille
```
# Version Final
```mocodo
DÉPART, 11 VILLE, 11 TRAJET : nomVille
SEFAITCONDUIRE, ON TRAJET, IN PASSAGER
PASSAGER: noPassager, nom, adresse, avis, dateAgrement
AVIS2, 1N PASSAGER, O1 CONDUCTEUR : avis

VILLE: nomVille
TRAJET: noConducteur, date, villeDep, villeAri, heureDes, noPassager, longueur, tarif
AVIS1, 1N CONDUCTEUR, 01 PASSAGER : avis
CONDUCTEUR: noConducteur, nom, addresse, avis, dateAgrement

ARRIVÉ, 11 VILLE, 11 TRAJET : nomVille
UTILISÉ, 11 VEHICULE, 1N TRAJET
VEHICULE: matricule, type, marque, energie, nbrPlace, dateMiseEnCirculation, noConducteur
POSSÈDE, IN CONDUCTEUR, ON VEHICULE
```
# Version Final v02
```mocodo
AVISPASSAGER: noPassager, noConducteur, avis, estPositif
AvisSurPassager, 0N AVISPASSAGER, 11 PASSAGER
PASSAGER: noPassager, nom, adresse, avis, dateAgrement
AvisPassagerConducteur, 1N PASSAGER, 1N AVISCONDUCTEUR: noConducteur, avis, estPositif

AvisConducteurPassager, 1N CONDUCTEUR, 1N AVISPASSAGER: noPassager, avis, estPositif
CONDUCTEUR: noConducteur, nom, addresse, avis, dateAgrement
:
AvisSurConducteur, 0N AVISCONDUCTEUR, 11 CONDUCTEUR
AVISCONDUCTEUR: noConducteur, noPassager, avis, estPositif

POSSÈDE, 1N CONDUCTEUR, ON VEHICULE
SEFAITCONDUIRE, ON TRAJET, 1N PASSAGER

VEHICULE: matricule, type, marque, energie, nbrPlace, dateMiseEnCirculation, noConducteur
UTILISÉ, 11 VEHICULE, 1N TRAJET
TRAJET: noConducteur, date, villeDep, villeAri, heureDes, noPassager, longueur, tarif
ARRIVÉ, 11 VILLE, 11 TRAJET : nomVille
:

:
DÉPART, 11 VILLE, 11 TRAJET : nomVille
VILLE: nomVille
```
# Encore une version final
```
UTILISATEUR: noUtilisateur, nom, adresse, dateAgrement
EstConducteur, 0N  UTILISATEUR, 11 CONDUCTEUR
CONDUCTEUR: noConducteur
Possède, 1N CONDUCTEUR, ON VEHICULE
VEHICULE: matricule, type, marque, energie, nbrPlace, dateMiseEnCirculation, noConducteur

EstNote, 01 AVIS, 11 UTILISATEUR
Conduit, 1N CONDUCTEUR, 11 TRAJET

AVIS: noUtilisateur, avis, idtrajet
Note, 11 TRAJET, 0N AVIS
TRAJET: idTrajet, date, villeDep, villeAri, heureDes, noPassager, longueur, tarif
Arrivé, 11 VILLE, 11 TRAJET : nomVille

:
Départ, 11 VILLE, 11 TRAJET : nomVille
VILLE: nomVille
```
# Yet another final version
```mocodo
VILLE: nomVille

Réside, 0N VILLE, 11 UTILISATEUR: nomVille

UTILISATEUR: idUtilisateur, nom, dateAgrement, adresse

EstConducteur, 0N  UTILISATEUR, 0N CONDUCTEUR: idUtilisateur

:

  

Arrivé, 0N VILLE, 11 TRAJET : nomVille

Départ, 0N VILLE, 11 TRAJET : nomVille

EstConduit, 1N UTILISATEUR, 1N TRAJET: idUtilisateur

EstNote, 0N AVIS, 0N UTILISATEUR

:

  

:

TRAJET: idTrajet, date, heureDes, longueur, tarif, villeDep, villeAri, idUtilisateur, idConducteur

AvisDonné, 11 TRAJET, 0N AVIS: idUtilisateur, idTrajet

AVIS: idAvis, avis, idUtilisateur, idtrajet

CONDUCTEUR: idConducteur, idUtilisateur, matricule

  

:

:

Conduit, 1N CONDUCTEUR, 11 TRAJET: idConducteur

VEHICULE: matricule, type, marque, energie, nbrPlace, dateMiseEnCirculation

Possède, 1N CONDUCTEUR, 11 VEHICULE: matricule
```
...