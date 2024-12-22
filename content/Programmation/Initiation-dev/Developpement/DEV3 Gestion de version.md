Initiation a git, schéma sur #Madoc 
# Pourquoi
Quand on travail, on veux parfois 
- revenir en arrière
- voir les modification d'un fichier au cours du temps
- savoir qui est l'auteur d'une modification donnée
# Méthode classique
Faire des copie de fichier de temps en temps (backup manuelle)
- Peut régler 1, éventuellement 2 et difficilement 3
- demande beaucoup de rigueur pour bien fonctionner
- peu efficace en terme de stockage (redondance)
# Gestion de version local
Une base de donnée qui stock les changement faits sur les fichiers au cours du temps
## Avantages
- gestion simple des version
- économie d'espace : version stocker sous forme de patchs (modification réalisé pour passé à la version suivante et non tout le code)
## Limitation
- Travail à plusieurs difficile
- risque de perte de données (code stocker uniquement sur un seul ordinateur)
# Gestion de version centralisée
La base de donnée est stockée sur un serveur distant
## Avantages
- Travail à plusieurs possible
- avec historique des action de chacun
- gestion des droits
## Limitations
- Risque de pertes de donné si problème de serveur
# Remarque importante
## Risque inévitable
A partie du moment ou on travaille à plusieurs sur les mêmes fichier, il y a des risque de *conflits*.
### Résoudre un conflit
- si les modification concernent des (partie de) fichiers différents, les conflit sont résolus automatiquement
- sinon il n'y a pas d'autre choix que de fusionné les version manuellement
# Gestion de version répartie
La base de donnée est dupliquée chez tout les utilisateurs
# Mais conflit ...
- Possibilité de conflit entre version de la BD, résolution auto
- toujours problème de conflit de fichier
## Avantages
- Tout les précédant
- resistance aux pertes de donnée -> duplication
- possibilité de travaillé avec différent groupe de devs avec des versions différente
## Inconvénients
- conflit plus complexe (entre fichier, version d'historique)
- plus laborieux à utilisé (interaction avec la base de données locale et avec le serveur)

# GIT : copie local
voir #Madoc pour commandes et description des actions exacte
