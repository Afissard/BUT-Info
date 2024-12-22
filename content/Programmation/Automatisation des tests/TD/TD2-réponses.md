# Division
```kotlin
/** 
* Diviser deux entiers naturels 
* @param dividende : entier naturel 
* @param diviseur : entier naturel 
* @return quotient : flottant 
* @throws ArithmeticException */ 
fun diviserNaturel(dividende: Int, diviseur: Int): Float { … }
```
## Identifiez les variables qui forment chaque Donnée de Test, ainsi que les variables qui seront contrôlées par l’Oracle. 
```
DIVISION NATUREL
dt:array<Int> = [dividende:Int, diviseur:Int]
oracle:Float = res
```
## Réalisez pour chacune des variables formant les Données de Test une analyse partitionnelle, afin d’en déduire des classes d’équivalences. 

```
DIVISION NATUREL
1 Types des variables
dividende : entier naturel : Int
diviseur : entier naturel : Int

2 Plage de la variable
	A Intervalles nominaux 
		dividende :
			type Int -> [Int.min, Int.max]
			plage (nominal) -> [0, Int.max]
		diviseur :
			type Int -> [Int.min, Int.max]
			plage (nominal) -> ]0, Int.max] 
	B Plage fonctionnement exeptionnel
		dividende -> [Int.min, 0[-> [Int.min, 0[
		diviseur -> [Int.min, 0[ & []

3 Partition fonctionnel
-> [0]
```
## Etablissez une table de décision décrivant le comportement attendu. 

|            |   Variable    |        Contenu        | Test 1 | Test 2 | Test 3 | Test 4 |
| :--------: | :-----------: | :-------------------: | :----: | :----: | :----: | :----: |
| **Entré**  | **Dividende** |    \[Int.min, 0\[     |   X    |   I    |   I    |        |
|            |               |     \[0, Int.max]     |        |   I    |   I    |   X    |
|            | **Diviseur**  |    \[Int.min, 0\[     |   I    |   X    |        |        |
|            |               |         \[0]          |   I    |        |   X    |        |
|            |               |     \[0, Int.max]     |   I    |        |        |   X    |
| **Sortie** | **Quotient**  |    \[0, Float.max]    |        |        |        |   X    |
|            |               | Arithmétique exeption |   X    |   X    |   X    |        |

## Déduisez un ensemble fini de Cas de Test.

|     | Dividende | Diviseur | Quotient               |
| --- | --------- | -------- | ---------------------- |
| 1   | -4        | 2        | Arithmétique exception |
| 2   | 4         | -2       | Arithmétique exception |
| 3   | 4         | 0        | Arithmétique exception |
| 4   | 1         | 2        | 0.5                    |
# Multiplication
```kotlin
/** 
* Multiplier deux entiers 
* @param op1, op2 : int Opérandes à multiplier 
* @return produit : int 
* @throws ArithmeticException : out of Int bounds */ 
* fun multiplier(op1: Int, op2: Int): Int { … }
```
## Identifiez les variables qui forment chaque Donnée de Test, ainsi que les variables qui seront contrôlées par l’Oracle. 
```
op1: Int
op2: Int
res: Int
```
## Réalisez pour chacune des variables formant les Données de Test une analyse partitionnelle, afin d’en déduire des classes d’équivalences. 
```
MULTIPLICATION
1 Types des variables
op1 : Int
op2 : Int

2 Plage de la variable
	A Intervalles nominaux 
		op :
			type Int -> [Int.min, Int.max]
			plage (nominal) -> [Int.min, Int.max]
	B Plage fonctionnement exeptionnel
		?

3 Partition fonctionnel
TODO?
```

## Etablissez une table de décision décrivant le comportement attendu. 

|        | Variable | Contenu                       | Test 1 | Test 2 | Test 3 |
| ------ | -------- | ----------------------------- | ------ | ------ | ------ |
| Entré  | op1+op2  | \[Int.min + Int.min, Int.max] | X      |        |        |
|        | op1+op2  | \[Int.min, Int.max]           |        |        | X      |
|        | op1+op2  | ]Int.max, Int.max + Int.max\[ |        | X      |        |
| Sortie | retour   | op1 + op2                     |        |        | X      |
|        |          | Arithmétique exception        | X      | X      |        |

## Déduisez un ensemble fini de Cas de Test.
