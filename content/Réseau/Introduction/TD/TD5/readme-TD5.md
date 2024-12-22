# Tp Services Réseaux : DHCP

##  Préambule
Dans ce tp, nous allons mettre en place un intranet composé d'une station (`C` cliente) et d'un serveur DHCP (`S`).
Un serveur dhcp délivre des adresses ip à des clients dhcp. le 
client (qui est encore dépourvu d'adresse ip) utilise une 
diffusion au niveau liaison pour contacter un serveur dhcp. Pour 
isoler vos serveurs et leurs clients DHCP des autres que vous, 
sans toucher au câblage, nous allons mettre en oeuvre des VLANs.


![](plan.png "plan d'adressage")

# DHCP

DHCP (Dynamic Host Configuration Protocol) est un protocole qui 
permet de configurer les paramètres réseaux d'une station (IP, 
masque, passerelle par défaut, ...). 

## Configuration client

* consulter le manuel de dhclient

* utiliser cette commande sur `bleue` tout en capturant les 
 échanges avec wireshark (il est possible de recommencer les 
  échanges du début en libérant l'ip obtenue avec l'option -r de 
  dhclient).

* quelle est l'adresse ip du serveur ?

* quelle ip avez-vous obtenu ?

* quels protocole de transport et ports associés ont été utilisés 
  ?

* quels autres éléments de configuration avez-vous obtenu ?

## VLAN

Avec l'accroissement des débits et les coûts des équipements en 
baisse, les LAN ethernet sont de plus en plus structurés autour 
de commutateurs. Ces derniers segmentent les domaines de 
collision mais maintiennent un seul domaine de diffusion. Cet 
accroissement des domaines de diffusions pose des problèmes de 
sécurité avec la circulation de trames de diffusion (ARP, DHCP, 
...). Les VLAN (Virtual Local Area Network) sont apparus de 
manière naturelle comme le moyen d'obtenir un domaine de 
diffusion logique (où l'emplacement physique des équipements 
n'interfère pas avec la définition du domaine de diffusion). Les 
VLAN sont utilisés pour diviser les réseaux en entités plus 
petites et plus facilement administrables. L'approche 
traditionnelle de présentation des VLAN en fonction des couches 
réseaux (VLAN niveau 1 ou par port ; VLAN niveau 2 ou par adresse 
MAC et VLAN de niveau 3 ou par IP) laisse la place à une 
présentation plus fonctionnelle avec deux types de VLAN :

* VLAN basés sur les port. Le VLAN est défini par les numéros de 
  port du switch. 

* VLAN basé sur la norme 802.1q qui inscrit le VLAN 
  d'appartenance dans la trame. Ceci permet d'affecter 
  dynamiquement une trame à un VLAN. 

Activer la prise en compte des VLAN sous linux va permettre 
d'émettre des trames targuées 

* charger le module 802.1q :
```
modprobe 8021q
```

* sur les deux machines de l'intranet (C et S), on ajoute un 
  nouveau lien de type VLAN (si on omet de préciser le nom du 
  vlan, il est nommé vlan0). L'id de VLAN est la même sur C et S.

```
ip link add link jaune name jaune.X type vlan id X
```

* Ne pas oublier d'activer l'interface `jaune` et `jaune.X`

* affecter manuellement des ips à `C` et `S` sur l'interface `jaune.X` 
  prises dans le réseau `10.0.X.0/24`

* tester la connectivité entre `C` et `S`.

* ajouter une ip sur une autre station (sans configuration de vlan), toujours dans le même réseau.

* vérifier que vous n'arrivez pas à l'atteindre depuis les deux 
  autres. C'est normal, elle n'est pas dans le VLAN.

* capturer un dialogue entre les deux premières stations sur 
  l'interface jaune ; vous devriez observer des trames taguées. 


##  Configuration serveur

* Supprimer l'ip et route par défaut sur la station `C`. Nous 
  allons lui fournir via le serveur DHCP. (ip a del ...)

* (re)configurer l'ip de `S`, serveur DHCP sur `jaune.X` : `10.0.X.254`

* consulter le man de `dhcpd.conf`

* modifier `/etc/default/isc-dhcp-server` pour déclarer l'interface 
  sur laquelle le serveur va traiter les requêtes DHCP 
  (`INTERFACESv4=jaune.X`).

* configurer le serveur DHCP dans le fichier `/etc/dhcp/dhcpd.conf` 
  pour délivrer des IPs sur la plage `10.0.X.10-100/24` en 
  fournissant route par défaut (via `10.0.X.254`), nom de domaine 
  (`fai.com`), serveur DNS (`192.168.0.254`).

* tester depuis la station cliente, en utilisant dhclient, avec 
  l'interface `jaune.X`

* vérifier la configuration obtenue

* On souhaite maintenant que la station cliente se voit délivrer 
  une adresse fixe `10.0.X.1/24 en fonction de son adresse mac. 
  Modifier la configuration du serveur en conséquence. 

