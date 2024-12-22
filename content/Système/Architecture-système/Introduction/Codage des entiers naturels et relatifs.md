#TODO : retéléchargé tout les documents depuis le drive
#TODO : télécharger tout les docs de #Madoc 
# Unités de mesure de l'information
## Le Bit :
Pour *Binary digit* qui est l'unité fondamentale de mesure d'une quantité d'information.
Tableau des unité :
| Nom        | Symbole | en Octets | en Bits  |
| ---------- | ------- | --------- | -------- |
| bit        | b       | -         | 1        |
| octet      | B       | 1         | 8        |
| Kibioctet  | KiB        |        4   |   16       |
| Mébioctet  |  MiB       |      8     |   32       |
| Gibioctect |  GiB       |     x      |    64      |
| Tébioctet  |   TiB      |     x      |         x |
| Kilooctet  |      KB   |        x   |        x  |
| Mégaoctet  |    MB     |     x      |  x        |
| Gigaoctet  |     GO    |  x   |  x        |
| Téraoctet  |     TO    | x   | 8\*10^12 |

Le bit est utiliser pour encoder des valeurs binaire.

## Encodage d'Entiers
### Entier naturels
Aujourd'hui, nous utilisons majoritairement la notation positionnel en base 10.
>**Positionnelle :** le poids d'un symbole (ou chiffre base 10) depend de sa position absolue dans l'écriture du nombre.

Encodage en base 10 :
Exemple : 254 = 2\*10^2 + 5\*10^1 + 4\*10^2 = 200 + 50 + 4

#### Encodage en base 2
Pour coder en binaire, même principe (on peut ignoré le 0):
Exemple : 1000 0011 (b2) = 1\*2^7 + 1\*2^2 + 1\*2^1 = 131 (b10)
Dans l'autre sens : 
487 = 256 + 231 :(2^8)
231 = 128 + 103
103 = 64 + 32 + 4 + 1
soit : 487 = 2^8 + 2^7 + 2^5 + 2^2 + 2^0 = 1 + 2^2\*(1+2^3+2^4+2^5+2^6)
résultat = 0001 1110 0111 (b2)

#### Encodage en base 16
0, 1, 2, 3, 4, 5, 6, 7, 8, 9, A, B, C, D, E, F

1CAFE (b16) = 1\*16^4 + C\*16^3 ... etc. = 117 503 (b10)

Tableau correspondence de bases
| Base 2 | Base 16 | Base 10 |
| ------ | ------- | ------- |
| 0000   | 0       | 0       |
| 0001   | 1       | 1       |
| 0010   | 2       | 2       |
| 0011   | 3       | 3       |
| 0100   | 4       | 4       |
| 0101   | 5       | 5       |
| 0110   | 6       | 6       |
| 0111   | 7       | 7       |
| 1000   | 8       | 8       |
| 1001   | 9       | 9       |
| 1010   | A       | 10      |
| 1011   | B       | 11      |
| 1100   | C       | 12      |
| 1101   | D       | 13      |
| 1110   | E       | 14      |
| 1111   | F       | 15      |

#Madoc allez récup image sur écriture en informatique des entiers et fonctionnement des calculs

## Codage d'entiers négatifs : Codage en complément à 2
Exemple du principe sur 4bits :
Encodage "naturel" : 1111 = 15
Complément à 2 : 1,111 = -1
Complément à 2 : 0,111 = +7
**Si bit de poids fort = 0, alors nombre positif, si le bit de poids fort = 1 alors négatif.** Pour calculer la valeur, penser au fait que les 0 et 1 sont inverser. Allez voir #Madoc 

Machine de complément à 2 utiliser car cela simplifie grandement les Unité Arithmétique Literal. 
