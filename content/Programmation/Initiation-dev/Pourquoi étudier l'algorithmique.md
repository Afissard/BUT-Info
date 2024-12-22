Evolution des performance des ordinateur (lois de Moore).
Mais exécuter un algo à quand même un coût en énergie et en temps. Il est donc necessaire de pouvoir concevoir un algo efficace.

Voir #Madoc , exemple performance d'algo, complexité d'algo
Cas A : tri par insertion bien programmer : 2n^2 opération pour un ordi puissant : 1 milliards d'opération par seconde
Cas B : rangement de n éléments par tri fusion mal programmé : 50n log2(n) opérations, sur un ordinateur peu puissant :  10 millions d'opérations par seconde.
Si n = 10^6, dans le cas A, l'éxécution nécessite 2\*(10^6)^2 opérations et dure donc  (2\*(10^6)^2)/10^9 = 2000 secondes. Alors que dans le cas B, l'éxécution nécessite 50\*10^6\*log(10^6) opérations et dure donc (50\*10^6\*log(10^6))/10^7 = 100 secondes.

## Même avec une puissance de calcul et une mémoire infinie
### Terminaison
Un algo mal conçu peut boucler sans arrêt, et donc ne jamais donnée de réponses.
### Validité
Un algo mal conçu peut donner des réponse incorrectes
### Implantation
Un algo mal compris peut être mal codé, le code executer par l'ordinateur ne correspondra alors pas à ce qui est attendu (terminaison et validité ne sont plus assurées.) Donc rester humble sur ses competences, des statistiques montre que même un programmeur expérimenter commet une erreur ~ toute les 15 min.