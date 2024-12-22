# Tp Services Réseaux : Routage statique
## 1 Rappels

### 1.1 Adresse IP

Chaque interface réseau (i. e. carte réseau) d'un hôte du réseau 
possède (au moins) une adresse IP unique dans le réseau. Une 
interface ethernet peut être désignée par ethX où X est le numéro 
de l'interface, un nom ”prédictible” comme enp0s25. Une adresse 
IP est un nombre de 32 bits souvent noté en décimal pointé ; 
quatre entiers (compris entre 0 et 255) séparés par des points. 
Exemple : 192.168.2.200. L'adresse IP est structurée en deux 
parties : 

1. partie réseau : permet de désigner le réseau (netID)

2. partie hôte : permet de désigner l'hôte dans le réseau 
  (hostID)

Un masque de sous-réseau permet de séparer partie réseau et 
partie hôte. Ce masque ; une succession de 1 suivi d'une 
succession de 0 donne l'étendue de la partie réseau.

* mettre la partie host à 0 nous donne l'adresse du réseau. Ceci 
  se fait par un et bit à bit entre adresse IP et masque. Exemple 
  pour 192.168.2.200 avec un masque de 255.255.224.0 : 
    
```
192.168.2.200     11000000.10101000.00000010.11001000

     ET                           ET

255.255.224.0     11111111.11111111.11100000.00000000

     =                             =

192.168.0.0       11000000.10101000.00000000.00000000
```

* mettre tous les bits de la partie réseau à 0 nous donne la 
  partie hôte. 

* mettre tous les bits de la partie hôte à 1 nous donne l'adresse 
  de diffusion dans le réseau.

### 1.2 Routage

Un hôte voulant faire une transmission constitue un paquet IP qui 
contient l'adresse du destinataire et l'adresse de l'expéditeur. 
Au niveau de la couche réseau, le routage utilise une table de 
routage qui contient une ou plusieurs lignes contenant chacune 
essentiellement trois informations : 

1. une adresse de réseau

2. un masque de réseau

3. comment atteindre le réseau : soit directement par une 
  interface connectée sur ce réseau (on parle de routage direct 
  ou de remise directe), soit en passant par un routeur (on parle 
  de routage indirect) qui est identifié par son IP et 
  l'interface à utilisée pour l'atteindre.

Un routeur peut être un équipement spécialisé ou simplement un 
hôte ordinaire relié à plusieurs réseaux.

Le décision de routage se fait par la recherche d'une 
correspondance dans la table de routage en appliquant pour chaque 
ligne, le masque de réseau à l'adresse de destination. Quatre cas 
peuvent alors se présenter : 

1. le réseau de la destination est directement connecté. Il y une 
  remise directe en utilisant le réseau local sous-jacent.

2. le réseau de la destination est accessible via un routeur. Le 
  paquet est transmis au routeur sans changer les adresses IP de 
  l'émetteur et du destinataire.

3. le réseau de la destination est absent de la table de routage, 
  mais il existe une route par défaut. Le paquet est transmis au 
  routeur désigné.

4. le réseau de la destination est absent de la table de routage, 
  et il n'existe pas de route par défaut. Envoi d'un message ICMP 
  à l'émetteur : Network is unreachable

Chaque routeur recevant un paquet IP applique le même algorithme.

### 1.3 Type de routage

Les deux types connus de routage sont : 

1. le routage statique

2. le routage dynamique ou auto-adaptatif

Par routage statique, on entend que les tables de routage sont 
renseignées manuellement, par opposition au routage dynamique où 
celles-ci sont construite par application d'un algorithme 
distribué au niveau des routeurs.

### 1.4 Commandes

Les commandes réseaux sont présentes sous deux formes :

* les commandes vénérables : `ifconfig`, `arp` et `route` que nous n'allons pas utiliser

* la nouvelle commande `ip` (du paquet iproute2) : `ip link`, `ip addr`, `ip neigh` 
  et `ip route`

#### interface

consultation : `ip link`, activation d'une interface `eth1` : `ip link set eth1 up`

#### adressage

consultation : `ip addr`, configuration d'une interface `eth0` : `ip addr add 192.168.1.1/24 dev eth0`

####  voisinage

consultation  : `ip neigh`

####   routage

consultation : `ip route`, configuration : ajout avec `add`, suppression avec `del` ; route par défaut : `default`. Exemple :

```
ip route add 172.20.11.0/24 via 192.168.1.253
```

## 2 Configuration réseau

Observez votre configuration réseau initiale. 

* Combien d'interfaces réseaux sont-elles disponibles ? 
* Quelle(s) est(sont) vos ip ? 
* Qu'avez-vous comme route(s) ? 
* Quels type de liens sont disponibles ? 
* Quelles sont vos adresses MAC ? 
* Quel est l'état du cache arp ?


Après rajout d'une ip sur l'interface `bleue` (`192.168.0.Y/24` avec 
`Y` numéro de votre machine), puis `ping` d'une machine d'un voisin, 
reconsulter l'ensemble des informations.

## 3 Routage statique

Pour configurer le routage dans cette partie, vous allez utiliser 
les commandes suivantes : 

* `ip` configuration de l'adresse IPv4 et du routage

* `ping` pour générer un trafic ICMP echo request/echo reply 
  permettant de tester la connectivité

* `wireshark` pour analyser les échanges

Le `X` des adresses IP qui suivent est à remplacer par un numéro 
qui vous sera communiqué par l'enseignant.
Enfin il est rappelé que bien souvent, les stations ordinaires 
(qui ne sont pas des routeurs) se contentent d'UNE IP sur UNE 
interface avec UNE route par défaut.

### 3.1 Situation 1

![](OLD/Réseau/TD/TD4/situation1.png "Situation 1")

* configurer les interfaces des trois stations (normalement, vous 
  avez 4 interfaces à configurer)

* consulter les tables de routages des trois stations

* démarrer wireshark sur les trois stations pour observer le 
  trafic ICMP.

* avec la commande `ping`, depuis S1, déterminer quelles sont les 
  communications possibles et impossibles vers les 4 interfaces. 
  Expliquer. 

* Dans S1, rajouter une route permettant d'atteindre le réseau 
  `172.X.2.0/24`.

* Quelle(s) interface(s) ne peut-on toujours pas atteindre ?

Par défaut, le `forwarding` (la fonction de retransmission d'un 
paquet) n'est pas actif. Pour l'activer sur R1, deux possibilité existent :
 
* `echo 1 > /proc/sys/net/ipv4/ip_forward` (que fait cette commande ?)
* `sysctl net.ipv4.ip_forward=1`

`sysctl` peut permettre de pérenniser la modification. Décommenter la ligne `net.ipv4.ip_forward=1` dans le fichier `/etc/sysctl.conf`. Puis recharger la 
configuration, `sysctl -p`.

* que constate-t-on sur les captures ? 

* rajouter la route ”retour” sur S2

* tester

Il est habituel pour une station ordinaire (i.e. autre qu'un 
routeur) de ne posséder qu'une seule route, celle par défault vers son 
routeur.

* Sur S1 et S2, supprimer les routes rajoutées et mettre 
  seulement des routes par défaut.

* tester.

### 3.2 Situation 2

Dans cette situation, il va nous falloir une interface réseau de 
plus pour R2. Pour pallier le manque d'interfaces disponibles, il 
faut affecter plusieurs IP à la même interface (i.e. carte 
réseau). Exemple :

```
#ip addr add 192.168.10.42/24 dev bleue
#ip addr add 192.168.10.66/24 dev bleue
```

![](OLD/Réseau/TD/TD4/situation2.png "Situation 2")


#### Mettre en place la situation :

* configurer les interfaces 

* configurer le routage

* tester

## 4 Routage dynamique

![](situation3.png "Situation 3")


Configurer `S1`, `S2` et `S3` avec ip et route par défaut. Dans la suite, seuls les routeurs sont à modifier.

Nous allons utiliser la suite logicielle `quagga` qui fournit les protocoles de routage dynamique les plus courants : OSPF, RIP et BGP. Quagga possède une architecture modulaire : un démon (zebra) pour l'interface entre le noyau du système et les démons de routage dynamique et des démons spécifiques à chaque protocole de routage dynamique. Il est possible de configurer quagga comme on le ferait pour un routeur matériel (type CISCO) via telnet ou avec des fichiers de configuration. La consultation de l'état du routeur se fait, elle aussi, par telnet.

### 4.1 OSPF

* Pour configurer le routage ospf, trois fichiers sont à produire sur chacun des routeurs :

1. Le fichier `/etc/quagga/daemons` permet de spécifier quels démons de routage vont être activés. Pour notre cas : 

```
zebra=yes
bgpd=no
ospfd=yes
ospf6d=no
ripd=no
ripngd=no
isisd=no
```

2. Le fichier `/etc/quagga/zebra.conf`. Ce fichier est fourni avec `A` numéro du routeur en cours de configuration, `B` et `C` numéro des deux autres routeurs (à adapter !!). On y configure les adresses IP : `interface N` puis `ip address` et il est conseillé de spécifier la bande passante pour chaque interface en bits/seconde de manière à avoir un coût ”conforme” : `bandwidth 100000`

3. Le fichier `/etc/quagga/ospfd.conf`. Le routeur est un routeur ospf : `router ospf`, on déclare les réseaux sur lesquels chercher des routeurs voisins `network ... area 0`

* droits pour ces fichiers

```
chown quagga:quagga /etc/quagga/ospfd.conf && chmod 640 /etc/quagga/ospfd.conf
chown quagga:quagga /etc/quagga/zebra.conf && chmod 640 /etc/quagga/zebra.conf
```

* pour les logs 

```
mkdir -p /var/log/quagga && chown quagga:quagga /var/log/quagga 
```

* lancer quagga 

```
service zebra restart
service ospfd restart
```

* vérifier le lancement des démons (`ps ...`)

* tester (ping)

* consulter les tables de routages (`telnet localhost zebra`, puis `show ip route`)

```
R1(Zebra)> show ip route 

Codes: K - kernel route, C - connected, S - static, R - RIP, O - OSPF, I - ISIS, B - BGP, > - selected route, * - FIB route

...

O 10.1.2.0/24 [110/1] is directly connected, interface, 00:02:30
...
```

Noter le [110/1] qui indique la priorité d'une route ospf par rapport aux autres sources des routes (0 connecté, 1 static , 20 BGP, 110 ospf et 120 RIP) et le 1 le coût du lien.

Faire la même chose depuis un terminal : 

* `ip route` 

* consulter les information du routeur ospf (`telnet localhost ospf`, puis `show ip ospf` avec éventuellement `neighbor` ou `route` en complément) 

* faire un ping entre `S1` et `S2`. couper le lien entre `R1` et `R2`. 

* constater l'évolution des tables de routage

### 4.2 RIP (facultatif)

Si vous avez le temps, reprendre la situation avec RIP. Les principes restent les mêmes :

* le démon zebra (avec configuration des IPs si besoin est) ;

* un démon pour rip configuré dans `ripd.conf` ;

* la déclaration des réseaux sur lesquels chercher des routeurs voisins : `network ...`

* la nouveauté, indiquer quelles routes propager : `redistribute connected`

