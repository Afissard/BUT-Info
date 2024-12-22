
> **Bibliothèque :** Ensemble de fonctions prêtes à être utiliser par des programmes

## Le cours d'aujourd'hui
- comment organiser son code pour un gros projet
- comment fabriquer ses propre lib
- comment gérer les dépendances de son code à des lib externes

## Organiser sont code
au sein d'un module on peut organiser son code en plusieurs fichiers et dossiers pour bien séparer les différentes fonctionnalité.
## Fabriquer ses lib
Un module est un moyen de compartimenté ... voir #Madoc 

# Qu'est ce qu'un module
## Packet (package)
Collection de fichier sources, tous situé dans le même répertoire, et qui sont compilé ensemble
- code partager entre les fichier (fonction/variable global)
- unité d'importation (en Go on importe des paquets)
## Module
Collection de paquets qui sont publiés, versionnées et distribués ensemble
- identifié par un chemin (module *path*).
- fichier *go.mod* et *go.sum* à la racine. (voir image sur madoc)

## Créer un module
> *go mod init nom-du-module*

**Attention au choix du chemin**

# Compiler un module
>*go build*

## Main
Si le module contient de quoi s'exécuter à la racine (package main + func main), un exécutable sera produit
Commande alternatives :
- Installation : *go install* : les fichiers compilés sont placés dans l'espace de travail Go (variable d'environnement)
- Exécution : *go run* : seulement si le modules contient un main, il s'éxécute directement.

# Créer des nouveaux paquets
1. créer un dossier avec le nom du packet
2. faire commencer les fichier sources Go placés dans ce paquet par *package nom-du-package*.