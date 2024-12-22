**Tester régulièrement son code**

## Automatiser ses test
Pour ne pas avoir à sans cesse refaire des tests à la main, on peut écrire et mettre à jour un jeu de tests qui pourra être automatiquement utiliser sur notre programme.
### Pour tester un packets d'un module
- commande : *go test*
	- lit les fichier se terminant par *_test.go* du dossier courant
	- appelle une à une toutes les fonctions de test dans ces fichiers
	- affiche le résultat de ces tests (réussi ou non)
- le paquet *testing*
### Fonction de test
voir #Madoc pour exemple de la fonction de test

## Evalué l'efficacité du code
Un jeu de code de test permettra aussi de mesuré le temps de calcul de notre code sur un large ensemble de cas, ce qui permet d'évaluer son efficacité.
- *go test -bench*=
- le packet testing
voir #Madoc pour exemple de la fonction de test et benchmark avec GitHub
Avec bench, test est exécuter avec par défaut sauf si *go test -run=Bench -bench=*
