Ensemble des activité au sein d'un système informatique
encapsule un flot de contrôle logique ... rappel du précédant amphi + tp
#TODO récupérer le pdf sur #madoc pour les schémas
fonctionnement d'un processus, shéma

Ordonnanceur : logiciel du noyau chargé de déterminé l'ordre d'éxécution du processeur
- efficacité : utiliser le temps processeur le plus vite possible
- équité : s'assuré que chaque processus reçoit une part du temps processeur
- temps de réponse : minimiser le temps de réponse pour le mode interactif
- temps d'éxécution : minimiser l'attent des utilisateur en mode non-interactif
- rendement : maximiser le nombre de travaux effectué en une unité de temps
Impossible de tout satisfaire -> différente politique de gestion (schéma)
- Premier arrivé, premier servi
- Plus court d'abord
- file de priorité
- système de tourniquet, à chaque processus est alloué une tranche temps, à boucle soit repasse en file d'attente, soit sort si fini
Ce qui est utilisé aujourd'hui : un compromis entre file de priorité et système de tourniquet

Bloc de contrôle du processus : quand sortie du processus, sauvegarde de l'état avec PCB  (Process Control Block)
Démonstration avec exécution et vu depuis assembleur et registre