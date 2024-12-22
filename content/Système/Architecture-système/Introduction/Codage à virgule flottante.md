# Rappel sur les nombre décimaux
Les ordinateurs peuvent manipuler des entier, naturels et relatif. Mais il existe d'autre ensemble de nombre tel que R (Réel), Q(Rationnel), D(Décimaux).
## Notation decimal positionnelle
Pour écrire les nombres décimaux, on peut étendre la notation positionnelle pour inclure la partie fractionnaire : on parle alors de **notation decimal fractionnelle**.
Exemple : 123/10^2 = 1,23
Voir #Madoc pour exemple des formules
On peut étendre la notation positionnelle sur la base 2.
Exemple : (base 2) : 111/10^11 = 0,111 <=> (base 10) : 7/2^3 = 7/8
# Representation machine des nombre fractionnaire
Pour les entier, le choix est fait de se limité aux nombre d'un intervalle dont la taille dépend de la "largeur" des circuit de l'UL, typiquement 64bits sur une machine moderne (int, unsigned int)
Pour les besoin plus spécifiques (ex : cryptographie), il existe des bibliothèque de "grand entiers" (BigInt) qui permettent d'atteindre un précision arbitraire, au prix de performances détériorées.

Les rationnels sont un ensemble "dense", et la representation choisie ne comptera nécessairement qu'un nombre limité d'éléments {x1, x2, ..., xi}

Choisir une representation pour les xi :
- quelle taille (en bit) pour la part entière, fractionnaire
- comment concilier précision et amplitude de la plage des nombre representable
- la représentation garentit-elle l'unicité de la representation
- la representation permet elle une réalisation materiel efficace des fonctions arithmétiques
## Representation à virgule fixe
nombre de bit allouer au entier et fractionnaire fixe.
### Avantage
- utilisation des même opérateur d'arithmétique entière
- assez simple d'estimer l'erreur d'un calcul complexe
- permet d'optimiser le format en fonction de la précision et de l'amplitude ciblé pour un cas particulier.
### Inconvenient
- #Madoc 

### Exemple de format : Qm.n
Codage qui utiliser la representation à virgule fixe : le nombre est codé sur la forme d'un entier à complément à 2 sur m + n bits, qui est ensuite diviser par 2^n.
Exemple avec #Madoc 
## Representation à virgule flottante
La position de la virgule n'est pas fixe : reprend le principe de la notation scientifique des nombre, avec un nombre de chiffre significatifs et une plage limité d'exposant.
### Avantage :
- Précision pour des petit nombre + capacité à représenter l'odre de grandeur des grand nombre
### Inconvenient
- Nécessite l'utilisation d'unité de calcul dédiées -> les Floating PU -> Graphical PU -> ?
- Performance moins bonne qu'en virgule fixe
- difficile d'évaluer les erreurs de calculs.

>rappel notation scientifique

Puisqu'on fait des puissance de 10, on peut aussi faire des puissance de 2
Exemple :  -12,375 = -1100,011 = (-1)^1\*1,100011 \* 2^3
voir #Madoc 
# Format IEEE 754
L'IEEE (Institute of Electrical and Electronics Engineers) à proposé un format standardisé pour la representation pour la representation binaire en virgule flottante, qu'on appelle communément IEEE 745.
Le format définit plusieurs precisions.
Les langages et processeurs moderns utiliser le format simple ... #Madoc 

Voir Simple precision, schema et formule 32bit sur #Madoc 
- la partie entière est toujours égale à 1 n'est pas codée.
- les 22 premier bit après la virgule son codés et le 23 est arrondi.
- L'exposant est codé avec un biais de 127 (plage de valeurs : -126, 127 car les valeurs0x00 et 0xFF sont réservé).
- La précision (=poids le plus faible) varie avec l'exposant : de 2^-149 à 2^104
Exemple : 
1. 1234,5 = 10011010010,1
2. 10011010010,1 = (-1)^0 + 1,00110100101\*2^10
3. Voir suite de l'exemple sur #Madoc 

## Cas particuliers
Pour augmenter la densité de valeurs autours de 0, le codage change quand e=0x00
Formule : voir #Madoc 
Ce codage permet de coder les valeurs 0 et -0
Par ailleurs, e = 0xFF est réservé pour les valeurs spéciales :
- NaN pour Not a Number (error maths) quand f !=0
- +l'infini quand s = 0 et f = 0
- - l'infini quand s = 1 et f = 0
## En 64 bits
Le format IEEE754 double de precision quand 64 bits sont utiliser (voir #Madoc pour les détails)

## Arithmétique
Voir #Madoc explication claire et précise avec formules
### Propriété conservé
- 0 éléments neutre pour + et absorbant pour \*
- 1 élément neutre pour \*
- Pour tout x, x-x = 0
- \* et + sont commutatives : a + b = b + a
### Propriétés non conservé
#Madoc 
- Dans FP, + et \* ne sont pas associatif