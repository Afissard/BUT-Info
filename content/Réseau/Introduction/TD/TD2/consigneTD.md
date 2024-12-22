# Td Services Réseaux : Routage statique

Une table de routage se présente sous la forme d'une liste de réseaux (ip/masque) avec le moyen de les atteindre, soit directement en utilisant une interface locale, soit via un routeur (connu par son IP), le *next hop* (prochain saut).

Figurent dans cette table trois types de routes : 

1. Route de remise directe : sans passerelle vers son réseau d'appartenance
2. Route spécifique : vers un réseau particulier
3. Route par défaut

Exemple : 

| Addresse/masque |		*Next hop*| interface|
|:----------------|:--------:		|:----------|
| `192.168.1.0/24`|	-  |   `enp7s0 `     | 
| `10.5.4.0/24`  | `192.168.1.253`    | `enp7s0`  |
| `default`       | `192.168.1.254`  | `enp7s0 `  |

Bien souvent, les stations "ordinaires" se contentent d'une route de remise directe et d'une route par défaut.
La recherche d'une correspondance dans une table de routage se fait en appliquant le masque de la ligne courante et en comparant les réseaux. Si plusieurs routes sont possible, on utilise la route **la plus spécifique**, c'est à dire celle avec le masque le plus long.
## Exercice 1
Établir les tables de routages de `A`, `R`, `S` et `T` du réseau suivant : 

![](routage1.png )

On suppose maintenant l'existence d'un accès à internet en `T` via une interface `en1`. Donner les tables de routage modifiées de `R` et `S`.

Peut-on simplifier ? : Si des destination avec même *next hope* et même *interface* alors remplaçable avec *destination = default* et *next hope* et même *interface*.

On suppose maintenant la présence d'un réseau accessible par `B`. Donner la table de routage de `B`. Peut-on la simplifier ?
### Réponses
1. Combiens de réseau présent : 3
2. Pour chaque réseau on crée un entré de la table de routage, exemple :

| destination @reseau / masque | interface | next hope / @routeur / passerelle / gate away |
| ---------------------------- | --------- | --------------------------------------------- |
|                              |           |                                               |
3. on identifie l'interface qui permet d'y accéder
4. si réseau non directement accessible, on identifie l'adresse IP du premier routeur qui permet d'y accéder.
### Table de routage exo 1

| A   | @reseau / masque | interface | next hope     |
| --- | ---------------- | --------- | ------------- |
|     | 178.42.42.0/24   | en0       | -             |
|     | 66.16.0.0/16     | en0       | 178.42.42.254 |
|     | 152.18.4.0/24    | en0       | 178.42.42.254 |

| A simplifié                        | @reseau / masque | interface | next hope     |
| ---------------------------------- | ---------------- | --------- | ------------- |
|                                    | 178.42.42.0/24   | en0       | -             |
| (correspond au 2 dernière colones) | default/0        | en0       | 178.42.42.254 |

| A internet | @reseau / masque | interface | next hope     |
| ---------- | ---------------- | --------- | ------------- |
|            | 178.42.42.0/24   | en0       | -             |
|            | default/0        | en0       | 178.42.42.254 |

| R   | @reseau / masque | interface | next hope |
| --- | ---------------- | --------- | --------- |
|     | 152.18.4.0/24    | en1       | 66.16.0.2 |
|     | 178.42.42.0/24   | en0       | -         |
|     | 66.16.0.0/16     | en1       | -         |

| R internet | @reseau / masque | interface | next hope |
| ---------- | ---------------- | --------- | --------- |
|            | 178.42.42.0/24   | en0       | -         |
|            | 66.16.0.0/16     | en1       | -         |

| S   | @reseau / masque | interface | next hope |
| --- | ---------------- | --------- | --------- |
|     | 66.16.0.0/16     | en0       | -         |
|     | 152.18.4.0/24    | en1       | -         |
|     | 178.42.42.0/24   | en0       | 66.16.0.1 |

| S internet | @reseau / masque | interface | next hope  |
| ---------- | ---------------- | --------- | ---------- |
|            | 66.16.0.0/16     | en0       | -          |
|            | 152.18.4.0/24    | en1       | -          |
|            | 178.42.42.0/24   | en0       | 66.16.0.1  |
|            | default/0        | en1       | 152.18.1.3 |

| T simplifié | @reseau / masque | interface | next hope  |
| ----------- | ---------------- | --------- | ---------- |
|             | 152.18.4.0/24    | en0       | -          |
|             | default/0        | en0       | 152.18.4.2 |

| T internet | @reseau / masque | interface | next hope  |
| ---------- | ---------------- | --------- | ---------- |
|            | 152.18.4.0/24    | en0       | -          |
|            | default/0        | en0       | 152.18.4.2 |
|            | default/0        | en1       | INTERNET   |
Pour la semaine prochaine, B à l'IP 77.77.0.0/16
## Exercice 2 
Exercice provenant de Tanenbaum, 4<sup>ème</sup> édition, exercice 43.
Un routeur possède les entrées suivantes dans sa table d’acheminement :

| Addresse/masque  | *Next hop* | interface |
| :--------------- | :--------- | :-------- |
| `135.46.56.0/22` |            | `en0`     |
| `135.46.60.0/22` |            | `en1`     |
| `192.53.40.0/23` |            | `enp30s0` |
| `default`        |            | `enp30s1` |
 
Pour chacune des adresses IP suivantes, que fera le routeur si un paquet comportant cette adresse arrive ?
1. 135.46.63.10 -> envois vers en1
2. 135.46.57.14 -> envois vers en0
3. 135.46.52.2 -> envois vers enp30s1
4. 192.53.40.7 -> envois vers enp30s0
5. 192.53.56.7 -> envois vers enp30s1
## Exercice 3

Exercice provenant de Tanenbaum, 4<sup>ème</sup> édition, exercice 41.
Un routeur vient juste de prendre connaissance des adresses IP suivantes : 
- 57.6.96.0/21, 
- 57.6.104.0/21,
- 57.6.112.0/21,
- 57.6.120.0/21.
Si elles utilisent la même interface de sortie, peuvent-elles être agrégées ? Si oui, quel est le résultat de cette agrégation ? Sinon, pourquoi ?

## Exercice 4
Établir les tables de routages de `R1`, `R2` et `R3` du réseau suivant, en supposant un accès à internet en `R4`. Simplifier les tables de routage, si possible.

![](routage2.png )

## Exercice 5

Un routeur possède les entrées suivantes dans sa table d’acheminement :

| Addresse/masque |*Next hop*| interface|
|:----------------|:--------|:----------|
| `172.16.0.0/16` |  |   `en0`     | 
| `172.16.80.0/20`|   |  `en1`     | 
| `default`       |   | `enp30s1`  |

 
Pour chacune des adresses IP suivantes, que fera le routeur si un paquet comportant cette adresse de destination arrive ?

1. 172.16.78.3
2. 172.16.85.5
3. 172.16.93.2
4. 172.16.97.1

