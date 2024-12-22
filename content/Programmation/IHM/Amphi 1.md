rappel : lié au dev objet -> langage utilisé pour le cours en Kotlin

**IHM : Interface Homme Machine**

Notion importante :
- ergonomie (strict minimum pour le moment)
- accessibilité
# Interaction avec la machine
Input: 
- clavier
- souris
- manette
- écran tactile
- casque VR
- voix
- etc
Output
- écran
- audio
# Ce que nous allons étudier
- Réaliser IHM avec client lourd (différent de client léger dans navigateur (web app)) : partie intégrante de l'application logiciel
	- interaction (souris, clavier)
	- affichage
- Placer composant graphique
- Gérer les interaction du logiciel avec l'utilisateur. Programmation événementiel (exemple : bouton cliquer -> évènement nécessitant une réaction)
- Développer des IHM en respectant le modèle d'architecture MVC (Modèle, Vue, Contrôleur)
- Introduction à la programmation réactive (binding), exemple modification de l'affichage en même temps que le remplissage d'un champ de texte ...
# Outils utilisé :
- Kotlin
- JavaFX (librairie d'interface "par défaut" de Java), le seul "problème" : la doc est en java. Attention la lib à beaucoup évolué -> faire attention à la version de la doc (v8)
# JavaFX
la lib inclue des widget (élément graphique), des conteneurs (élément de structure), des outils de programmation évènementiel.
## Schéma structure général d'une application JavaFX
#TODO récup l'image sur le diapo
## Schéma scène et graphe de scène
#TODO récup l'image sur le diapo
# Arbre d'objet
L'aobre d'objet est composer de nodes