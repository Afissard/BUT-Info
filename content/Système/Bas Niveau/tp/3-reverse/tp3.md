# TP3 : Rétro-ingénierie de code

Les objectifs de cette séance sont :

  1. Améliorer votre maîtrise de gdb.
  2. Découvrir les bases de la rétro-ingénierie de code. 
  3. Vous confronter à une série d'énigmes à résoudre. 

## La rétro-ingénierie de code

D'après [Wikipedia](https://fr.wikipedia.org/wiki/R%C3%A9tro-ing%C3%A9nierie) :

>>>
La rétro-ingénierie (aussi connue sous le nom d'ingénierie inversée ou de rétro-conception) est une méthode qui tente d'expliquer, par déduction et analyse systémique, comment un mécanisme, un dispositif, un système ou un programme existant, accomplit une tâche sans connaissance précise de la manière dont il fonctionne. La rétro-ingénierie s'applique notamment dans les domaines de l'ingénierie mécanique, informatique, électronique, chimique, biologique et dans celui du design. Le terme équivalent en anglais est reverse engineering (ou retro-engineering). 
>>>

Dans cette séance, vous allez étudier un programme qui vous est livré sous forme binaire, avec seulement une partie du code source.
Vous allez donc devoir réaliser une rétro-ingénierie du code correspondant aux parties du code source qui vous ne sont pas données.
Si vous allez ici réaliser cette activité avec la bénédiction de l'auteur du code ciblé, sachez qu'en France et plus largement en Europe, cette activité est encadrée par la loi. 
La page Wikipedia sur la rétro-ingénierie contient [une introduction à ces aspects légistlatifs](https://fr.wikipedia.org/wiki/R%C3%A9tro-ing%C3%A9nierie#Dispositions_relatives_%C3%A0_la_traduction_de_la_forme_de_code_d'un_logiciel_(directives_91/250/CEE_et_2009/24/CE)).

## Mysterio 

Dans le dossier `src`, vous trouverez le programme `mysterioso`, ainsi qu'une partie de son code source dans le fichier `main.go`.
Commencer par étudier le fichier `main.go` avant de poursuivre la lecture de ce sujet.

Comme vous avez pu le constater, vous avez 4 étapes à passer.
Pour passer une étape, vous devez soumettre une chaîne de caractères au programme.
Cette chaîne est vérifiée par une fonction.
Malheureusement, vous n'avez pas le code source des fonctions de vérification.
Vous allez donc devoir déterminer les chaînes permettant de franchir les étapes en étudiant le code machine des fonctions de vérification.

Pour chaque étape, il y a deux niveaux de réponse :
  1. être capable de produire une chaîne acceptée par la fonction de vérification ;
  2. être capable d'expliquer comment fonctionne la fonction de vérification et ainsi pouvoir caractériser l'ensemble des chaînes acceptables.

Dans tous les cas, il n'est pas nécessaire de comprendre le code ligne par ligne.
Certaines parties sont plus importantes que d'autres, à vous d'apprendre à les reconnaître.
Par ailleurs, n'hésitez pas à consulter des informations complémentaires. 
Les sujets suivants pourraient vous être utiles (liste non exhaustive) :
  - l'implantation des chaînes de caractère en go ;
  - la fonction `fmt.Sscanf` de la bibliothèque standard.

Enfin, une dernière requête avant que vous ne partiez affronter `mysterio`.
L'intérêt principal du TP est dans la résolution des énigmes que constituent chacune des étapes.
Lorsque vous avez résolu une énigme, merci de bien vouloir laisser à vos camarades la possibilité de faire de même.
Autrement dit, ne partagez pas vos solutions.
Si vous souhaitez aider un•e camarade qui est perdu•e ou qui suit une fausse piste, vous pouvez l'aider en donnant simplement un indice.

