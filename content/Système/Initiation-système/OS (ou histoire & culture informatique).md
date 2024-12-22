Operating System
- premier programme chargé au démarrage
- interface entre l'utilisateur et les resources matérielles
Sert à faire abstraction et presenter à l'utilisateur une machine virtuel plus convivial. et fourni des service : interface de plus haut niveau, machine virtuelle (niveau machine, l'os permet l'abstration entre la machine et les programmes logiciel, pas emulation d'une machine dans la machine (ex avec virtual box)). Le role de l'OS est la gestion des ressources informatique, le partage de la machine physique entre les applications.
# Couches d'abstration
Utilisateurs <-> Applications <-> Système d'exploitation <-> Matériel
## Abstraction du matériel vers logicielle 
schéma sur #Madoc 
matériel -> pilotes -> noyau -> shell -> abstraction graphique
# Gestion de ressources
schéma sur #Madoc 
- Processeur
- Memoire
- Périphérique dont le système de fichiers
# Caractéristique des OS
Elements de differentiation
- mono/multi utilisateur
- mono-tâche / multicache
- structure monolithique / en couche / micronoyaux / modulaire
- ...
Type d'OS
- **Traitement par lots (batch) (année 50 à 60)** : tache soumises jusqu'à terminaison
- **Système multicache multiprogrammés** : plusieurs programmes en mémoire pour utiliser le processeur en cas d'entrée/sortie -> maximise utilisation processeur
- **Temps partagé** : partage de temps de processeur entre les usagers (tranche de durée fixe) -> Unix,-> minimise temps de réponse
- **système répartis** (80 ->) abstraction de la localisation physique des resource partagée au sein d'un réseaux
- **système embarquer** (90 ->) adapté à des contraintes : énergie, temps-réel, ressource limité, coût
# Historique
- Première génération (ENIAC, Colossus) : pas de système d'exploitation
- Unix
- suite sur #Madoc 
# Constitution d'un OS
# Noyaux
Composant s'éxécutant en mode superviseur, chargé en mémoire centrale au demurrage du système. Contient des les fonctions fondamentales qui doivent être exécuté en mode superviseur
- tout ce qui est requis pour implémenté les service
- tout ce qui est requis pour sécuriser le système
Au minimum
- gestion des espaces d'adressage
- des contextes d'éxécution
- prise en compte basique des événements materiel et des exceptions logicielle
Services, programme de base
- #Madoc 
# Mode d'éxécutions / commutation de contexte
#Madoc pour les schémas
Un programme utilisateur s'éxécute dans un mode non privilégié : un mode utilisateur
le système s'éxécute dans un mode privilégié : un mode noyau
## Qu'est ce qui provoque le changement de contexte ?
- Interruption materiel : asynchrone
	- fin timer
	- frappe clavier
- appel système
- faute synchrone
	- division par zero
	- default de protection
	- ...
# Gestion des processus
Processus = l'ensembles des activité résultant de l'éxécution d'un programme au sein d'un système informatique.
Un processus encapsule :
- flot de controle logique : donne l'illusion d'un accès permanant et exclusif à UC
- un espace d'adressage virtuel : illusion de mémoire ...
#TODO #Madoc compléter ici
schémas sur #Madoc : explication de la gestion asynchrone