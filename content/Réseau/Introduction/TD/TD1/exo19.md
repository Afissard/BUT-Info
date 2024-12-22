# Découpage en sous-réseaux 
Dans chacun des cas ci-dessous, exprimez les plages d'adresses IP des sous-réseaux lorsque cela est possible. 
Calculez le masque CIDR, le masque décimal. Identifiez l'incrément puis les identificateurs de sous-réseau. Finalement, exprimez les plages de chaque sous-réseau ainsi que l'adresse de diffusion associée pour le sous-réseau. 
## a
220.100.80/24 avec 4 réseaux logiques et 10 hôtes par réseau 
````
220.100.80.0 (dec)
DC.64.50.00 (hex)
1101 1100.0110 0100.0101 0000.|0000 0000 (bin)

4 sous réseaux : 
1101 1100.0110 0100.0101 0000.|00/00| 0000
1101 1100.0110 0100.0101 0000.|00/01| 0000
1101 1100.0110 0100.0101 0000.|00/10| 0000
1101 1100.0110 0100.0101 0000.|00/11| 0000

Broadcast des sous réseau
1101 1100.0110 0100.0101 0000.|00/00| 1111
...
1101 1100.0110 0100.0101 0000.|00/11| 1111
````
## b 
172.18/16 avec 10 réseaux logiques et 500 hôtes par réseau
```
172.18.0.0 (dec)
AC.12.0.0 (hex)
1010 1100.0001 0010.|0000 0000.0000 0000 (bin)

10 sous réseaux (4 bits) pour 500 hôtes (10 bits)
1010 1100.0001 0010.|00/00 00|00.0000 0000
...
1010 1100.0001 0010.|00/10 10|00.0000 0000

Broadcast des sous réseaux :
1010 1100.0001 0010.|00/00 00|11.1111 1111
...
1010 1100.0001 0010.|00/10 10|11.1111 1111
```
## c 
10/8 avec 20 réseaux logiques et 1000 hôtes par réseau 
## d 
10.160/13 avec 60 réseaux logiques et 500 hôtes par réseau 
## e 
10.163.128/19 avec 6 réseaux logiques et 200 hôtes par réseau 
## f
20/9 avec 15 000 réseaux logiques et 500 hôtes par réseau
````
20 ...
````
## g 
120/8 avec 100 000 réseaux logiques et 100 hôtes par réseau