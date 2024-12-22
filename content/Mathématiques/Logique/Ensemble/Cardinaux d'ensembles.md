Soit E un ensemble fini de cardinal *n*. Le cardinal de l'ensemble E<sup>p</sup> des *p-uplets* de E est n<sup>p</sup>.

Soit E un ensemble fini de cardinal *n* et *p* un entier naturel tel que *p* ≤ *n*.
- Un **p-arrengement** de E est un p-uplet d'éléments distinct de E
- Une **permutation** de E est un n-arrengement
- Une **p-combinaison** de E est une partie de E ayant p éléments

Soit E un ensemble fini de cardinal *n* et *p* un entier naturel tel que 0 ≤ *p* ≤ *n*. Le nombre de p-arrangement de E est : *A<sup>p</sup><sub>n</sub> = n(n-1) ... (n-p+1) = (n !)/((n-p) !)*
Le nombre de permutation est : *A<sup>p</sup><sub>n</sub> = n !*

Pour chaque p-combinaison il existe p! permutations possibles de ses éléments pour former des p-arrangements. A partir de la propriété précédente on peut donc en déduire la suivante : 
Le nombre de combinaison de p éléments de E est : *C<sup>p</sup><sub>n</sub> = (*A<sup>p</sup><sub>n</sub>) / (p!) = (n!) / (p! (n-p))*

Les coefficient C<sup>p</sup><sub>n</sub>, aussi noté (<sup>p</sup><sub>n</sub>), sont appelés coefficient binomiaux ou combinatoire de p éléments parmi n. Ils représentent le nombre de façons de choisir p objets parmi n. L'ordre n'importe pas dans la selection effectué et un objet ne peut être sélectionné deux fois (tirage simultané et sans remise).
Chaque p-compbinaison représente une partie de E et de l'ensemble de ces p-combinaisons (avec 0 ≤ *p* ≤ *n*) représente une partition ([[Partition d'ensemble]]) de *P(E)*. Or on a énoncé précédemment le fait que : *card(P(E)) = 2<sup>n</sup>*. On peut donc en déduire que : *∑<sup>n</sup><sub>p=0</sub> C<sup>p</sup><sub>n</sub> = 2<sup>n</sup>*