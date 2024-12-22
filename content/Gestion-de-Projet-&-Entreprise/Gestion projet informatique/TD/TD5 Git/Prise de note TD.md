# TLDR (Too Long Didn't Read)
Read the f\*\*\*ing documentation : https://git-scm.com/docs/.
# Exercice 1
1. Leia clone sur sa machine (en local) le contenu du repository gitlab.univ-nantes.fr/Leia/episode4.git.
2. Leia écrit dans le fichier plan.txt, ligne 1, "Etoile Noire"
3. Leia vérifie l'état du git (montre si sa version est à jour, si des fichier on été modifié, vérifie si il y a des différence entre le serveur et la version local)
4. Leia ajout les fichiers modifier (pour un future commit -> push) : plan.txt
5. Leia vérifie l'état du git
6. Leia commit les fichier ajouté (via add -> plan.txt) avec le message "creation du plan"
7. Leia vérifie l'état du git
8. Leia écrit dans le fichier plan.txt, ligne 2, "Plan technique"
9. Leia commit les fichier ajouter (aucun) avec le message "plan". -> retourne une erreur : aucun fichier modifié ?
10. Leia ajout les fichier modifier -> plan.txt
11. Leia commit sans message -> reprend le commit actuel pour y ajouter les fichier ajouter depuis ?
12. Leia écrit dans le fichier plan.txt, ligne 3, "Diamètre : 120 kilomètres"
13. Leia ajoute plan.txt et commit avec le message "diamètre"
14. Leia supprime du repository les modification non commit du fichier plan.txt
15. Leia regarde les fichier présent dans le dossier courant
16. Leia vérifie l'état du git
17. Leia met à jour les fichiers locaux pour correspondre au contenu de la branche "head" principal et supprime les modification local (non commit ?)
18. Leia télécharge le contenu de la branche (met à jour)
19. Leia vérifie l'état du git
20. Leia pousse les modification commit sur le repo et le met à jour
21. Luke clone le repo
22. Luke regarde le log du dernier commit
23. Luke écrit dans le fichier plan.txt, ligne 4, " Equipage : 265675 hommes"
24. Luke commit "equipage" et ajoute les fichier modifier : plant.txt
25. Luke télécharge en local les modification faite sur le repo (fetch)
26. Luke regarde l'état du repo
27. Luke met à jour le repo avec les commit qu'il a réaliser
28. Leia télécharge en local les modification faite sur le repo
29. Leia regarde l'état du repo
30. Leia regarde les changement fait entre chaque commit entre la branche d'origine et la branche master
31. Leia merge les branche origine et master
32. ...