---
title: Architectures Réseaux CM1
draft: 
description: 
tags:
  - Réseau
---
# Architectures des réseaux CM1
Cour (pdf) : [[cmR3.06_AR.slides.pdf]]
## Modèles OSI
*schéma du modèle OSI*
### Couche 1 : physique
fournit les moyens nécessaires aux transferts des élément binaires 
- câblage (coaxial, paire torsadées, fibres optique)
- interfaçage de connexion (prise RJ45, connecteur BNC, ...)
- codage des bits (niveau électrique)
- équipement de transmission (modems, commutateurs, ...)
- topologie du [[Réseau]]
spécifie **canal** et **signal** sans aucune sémantique de l’information -> niveau bit
#### Canal
Le canal est constitué de tout médium capable d’assurer le transfert d’une information binaire. Caractérisé par une bande passante (bande des fréquences utilisés) : `W = Fmax − FMin`
Exemple

| Type                      | Bande passante |
| ------------------------- | -------------- |
| paire torsadée            | > 100 kHz      |
| cable coaxial             | > 100 MHz      |
| fibre optique             | > 1 GHz        |
| epace entre deux antennes | variabl        |
Les données binaires sont transportés par un **signal**.

#### Débit
- les éléments transmis sur un réseaux : 0 ou 1 
- le débit mesure la rapidité d’une communication numérique 
- débit = nombre d’éléments binaires transmis par seconde
- unité : bits par seconde (bps)
- Exemple : D = 10 Mbps

#### Signal
Signal : variation d’une grandeur physique, porteuse d’informations
Exemple : 
- signal électrique : tension
- signal optique : onde lumineuse
- signal hertzien : onde électromagnétique

#### Transmission
Deux types de transmissions : 
1. transmission en bande de base 
2. transmission d’un signal modulé

##### Transmission en bande de base
- suite de bits représentant les données numériques
- changement d’états discret du signal physique
- pas de transposition en fréquence
- durée de chaque bit est constante
*figure : transmission en bande de base*

##### Transmission d'un signal modulé
- utilisation d’une onde porteuse 
- modification pour augmenter le débit, diminuer le taux d’erreur
*figure : transmission d'un signal modulé*

### Couche 2 : liaison de données
- transforme un flot binaire brut en trames 
- gère l’établissement, le maintien et la libération de la liaison entre terminaux
	- transmission
	- contrôle de flux
	- contrôle d’erreur
	- adressage des terminaux
	- accusé de réception
-> niveau trame

### Couche 1 bis : MAC *Medium Acces Control*
- sous-couche de contrôle d’accès au canal pour un réseau à diffusion
- mécanisme d’adressage des hôtes : adresses MAC
- protocole de gestion d’accès : 
	1. CSMA/CD
	2. CSMA/CA

### Couche 3 : réseau
- permet la connexion de réseaux entre systèmes ouverts :
	- fonction d’adressage des réseaux
	- fonction de routage/relayage pour l’acheminement d’un datagramme
- doit permettre l’interconnexion de réseaux hétérogènes
-> niveau paquet

### Couche 4 : transport
...