---
title: Séminaire-Qdev
draft: 
description: 
tags:
  - Quebec
---
# Exemple de "poor quality code"
exemple de Toyota et son système de freinage
- ils devais suivre les normes de MISRA C (sous ensemble de C) en plus de leur propre normes
- problème du système de freinage
- analyse de la nasa, 14/37 règles non suivie -> 7 134 violation, exemple de problème : pas de cas par défaut dans les bloc de `switch`.
- autre analyse de Barr's Team -> 80k violations de MISRA C trouvé.
- Complexity Cyclomatique de McCabe (nombre d'embranchement de if)
	- 67 fonction avec une complexité > 50
	- 1300 ligne pour la fonction du frein avec ? variable
- Trop de variable global (certaine constantes (4720) mais surtout 11 253 variable global) uniquement pour le système de freinage général -> pas de variable local ou paramètre de fonctionnement
	- bug de concurrence -> lié au variable global notamment
- bug de Stack (overflow ou le stack va écrire dans le code du programme) -> si problème alors l'OS reboot = délais de boot
- mauvais codage du process de chien de garde, donc tout allais bien.

# Solution
- toujours pensé au futur dev qui va lire & exécuter le code 
- écrire du code comme si c'étais nous qui allions maintenir ce code pour l'éternité.
## le principe de qualité sans non
bien expliqué à l'oral, cf diapo
