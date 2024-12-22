# Linux TP 1

## 3 Lister les fichiers
**Question 1 :**
*ls -> enter* : retourne au home-directory.

**Question 2 :**
Pour obtenir un affichage de la taille d'un fichier lisible : *ls -s -h*, "-s" pour size "-h" pour human (aussi utlisable avec -l pour un affichage "long" et détaillé ou -lh pour une option combiné).

**Question 3 :**
*ls -a* : affiche tout le contenu du working-directory sans ignorer les fichier ou dossier cacher. Exemple : ".nom_fichier".

## 4.1 Chemin absolu
**Question 1 :**
*cd /etc* : Va dans le répertoire *root/etc* soit le répertoire des fichiers de configuration du système.

**Question 2 :**
Pour aller dans le répertoire "bin" ou binary : *cd /bin*. Pour ensuite retourner au répertoire personnel : *cd /home/E230824* ou *cd* ou *cd ~*.

## 4.2 Chemin relatif
**Question 1 :**
*cd ..* renvoie dans le dossier précedant soit le dossier dans lequel se trouve le répertoire où est entrée la commande. 

**Question 2 :**
Allez dans le dossier bin à partir de home avec les chemin relatif :
*/home$ cd ..
/$ cd bin*

## 5 Ne pas se perdre
**Question 1 :**
*cd ~* renvoi dans le répertoire personnel.

**Question 2 :**
La commande *pwd* renvoie le répertoire : /home/E230824.

**Question 3 :**
Liste des commande commenceant par "ma" :
- macptopbm
- mag
- make
- make-first-existing-target
- make index
- makejvf
- make-ssl-cert
- man
- mapfile
- mawk

**Question 4 :**
La commande *cd -> espace -> tab* retourne la liste de tout les répertoire diponible, dossier masquer inclus.

**Question 5 :**
*cd ../E23* ou *cd ../E230824/* retourne : "bash: cd: ../E23: no such file or directory".

**Question 6 :**
Question pas compris.

## 6 Afficher du texte
**Question 2 :**
Pour afficher plusieurs ligne de texte avec la commande echo : *echo -e hello \\\e world*

## 7 Lecure de fichiers
**Question 1 :**
Pour afficher le contenu du ficher : *cat fichierTexte.txt*.

**Question 2 :**
Pour afficher le contenue du fichier avec la numérotation des ligne : *cat -n fichierTexte.txt*.

**Question 3 :**
Avec la commande *cat fichierBinaire.pdf*, le programme retourne l'encodage pdf du fichier.

**Question 4 :**
Affichage du fichier etc/mime.types avec la commande cat : retourne la liste des media, fini en bas de la liste avec la possibilité d'entré une nouvelle commande.
Avec la commande less : retourne la liste des media en commençant par le haut.

## 8 Création de fichier
**Question 2 :**
Creation d'un repertoire "trav" : *mkdir trav*

**Question 3 :**
Pour verifier le résultat de la commande précédante : 
- Afficher le contenu du répertoire : *ls* ou *dir*.
- Se déplacer dans le répertoire créé : cd *trav*.

**Question 4 :**
Création de "monrep" dans "trav" : */trav$ mkdir monrep*.

**Question 5 :**
Création de "sousrep" depuis "trav" : */trav$ mkdir monrep/sousrep*.

**Question 6 :**
La commande *mkdir monrep/unrep/unautrerep* échoue, on ne peut créer un dossier dans un répertoire inexistant.

**Question 7 :**
Création d'un répertoire adjacent à "sousrep" en partant de "sousrep" : *mkdir ../sourep1*. 

**Question 8 :**
Création d'un répertoire similaire mais avec le chemin absolu : *mkdir /home/E230824W/reseau/Perso/Documents/cours/trav/monrep/sousrep2*.

**Question 9 :**
Création d'un répertoire "prive" dans "sousrep" avec chemin relatif : *mkdir prive*.

**Question 10 :**
Création d'un répertoire "jeu" dans "sousrep" avec le chemin absolu : *mkdir /home/E230824W/reseau/Perso/Documents/cours/trav/monrep/sousrep/jeux*.

**Question 11 :**
Vérification avec la commande *tree* :
E230824W@ubuntu:~/reseau/Perso/Documents/cours$ tree
.
├── accueil.pdf
├── fichierBinaire.pdf
├── fichierTexte.txt
├── linux_tp1.md
├── tp-1.pdf
└── trav
    └── monrep
    ├── sousrep
        │   ├── jeux
        │   └── prive
        ├── sousrep1
        └── sousrep2

7 directories, 5 files

## 9 Copie de fichier
**Question 1:**
Copiez le fichier fichierTexte.txt situ´e dans votre r´epertoire de travail dans le répertoire prive créé précédemment : *cp ../../fichierTexte.txt sousrep/prive*