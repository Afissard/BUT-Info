# adresse réseau, de diffusion, plage et nombre d’hôtes 
Rappelez la définition binaire d’une adresse réseau, celle de l’adresse de diffusion. Indiquez comment obtenir une adresse d’un réseau à partir d’une adresse IP d’un hôte et de son masque. Proposez une formule qui permette de calculer le nombre d’hôtes qui peuvent être adressés dans un réseau ayant un masque donc la valeur en CIDR est n. 
Puis pour chaque adresse IP ci-dessous, précisez son masque en décimal, son adresse réseau, son adresse de diffusion, la plage d'adresses IP disponibles pour désigner un hôte sur ce réseau et le nombre d'hôtes possibles. 
a 131.108.78.235/21 ; 
b 63.69.48.211/11 ; 
c 168.94.197.13/19 ; 
d 200.249.145.227/28 ; 
e 192.154.88.133/26 ; 
f 100.189.64.38/13 ; 
g 150.34.222.131/17

| IP/mask               | mask decimal      | @reseau          | @broadcast         | \#hôte             | plage d'adressage machine             |
| --------------------- | ----------------- | ---------------- | ------------------ | ------------------ | ------------------------------------- |
| ``131.108.78.235``/21 | ``255.255.248.0`` | ``131.108.72.0`` | ``131.108.79.255`` | 2<sup>11 -2 = 2046 | ``131.108.72.1`` à ``131.108.79.254`` |
| ``63.69.48.211``/11   |                   |                  |                    |                    |                                       |
| ``168.94.197.13``/19  | ``255.255.224.0`` | ``168.94.192.0`` | ``168.94.223.``    | 2<sup>13 -2 = 8192 | ``168.94.132.1`` à ``168.94.223.254`` |
