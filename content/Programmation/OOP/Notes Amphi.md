# Introduction
Langage enfant/héritier de Java
attention spécifité des boucle for avec inclusion de la dernière valeur avec *until* / *downTo*
# POO (Programmation Orienté Objet), ou OOP en anglais
C'est un paradigme (comme l'impératif ou le procedural)
Un objet est une structure avec des méthode (fonction appartenant à l'objet) avec en plus une notion d'
- Encapsulation (une partie de la structure est caché, seule les méthodes peuvent interagir avec l'objet). 
- L'utilisateur n'a besoin de connaitre que l'utilité des méthode (abstraction). 
- Autres concept : l'héritage -> on peut créer un nouvelle objet à partir d'un autre objet, tout est conservé, mais on peu réécrire une partie des méthode ou en ajouter (ainsi que des variables).
- Composition : un logiciel est un assemblage d'objet
- Polymorphisme : une même méthode peut produire des effets différents en fonction du contexte (voir héritage et polymorphisme)
L'encapsulation et l'abstraction masque la complexité du code.
## Avantage
- Pas forcement moins de code, mais plus intéressant car moins redondant
- meilleur organisation, hierachirsation
- collaboration plus simple
- maintenance et evolution plus simple
- réutilisation du code / création de librairie
## Definition d'un objet
C'est un "moule"
- attributs (propriété)
- méthode (interaction)
- éventuellement un héritage
## Manipulé un objet
- créer un objet (initialisation) => instanciation
- interagir avec l'objet avec ses méthode
un objet est une instance d'une classe
## Exemple : la classe string
le type string n'existe pas c'est un tableau de caractère, la classe string permet de masquer toute la complexité du tableau et offre des méthodes pour manipulé l'objet.
### Spécificité de Kotlin : tout les types sont des objets
