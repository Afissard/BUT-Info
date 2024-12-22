# Adresses particulières IPv4 
Expliquez en quoi certaines des adresses IP ci-dessous sont particulières. Précisez s'il est possible d'utiliser l'adresse pour définir un identifiant valide de machine. 
1. ``192.118.275.3`` /24 ; incorrecte : 275>255
2. ``192.168.0.1`` /24 ; valide, classe C, privé, @machine
3. ``172.17.255.0`` /16 ; valide, classe B, privé, @machine
4. ``191.100.2.255`` /16 ; valide, classe B, publique, @machine 
5. ``127.0.0.1`` ; @loopback (théoriquement classe A)
6. ``169.254.100.9`` ; valide, classe B, publique, @machine / adresse réservé par certaine application ...
7. ``0.0.0.0`` ; @initilisatation 
8. ``10.255.255.255`` /8 ; valide, classe A, privé, @broadcast
9. ``190.100.0.0`` /16 ; valide, classe B, publique, @reseau
10. ``255.255.255.255`` ; valide, @broadcast 
11. ``224.0.0.1`` /4 ; valide, classe D, publique, @machine