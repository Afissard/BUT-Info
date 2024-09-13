---
title: TP-0
draft: false
description: 
tags:
  - Réseau
---
# architecture des réseaux TP0
*prise de note*
# Tp Services Réseaux : Routage

Le `X` des adresses IP qui suivent est à remplacer par un numéro 
qui vous sera communiqué par l'enseignant.

## 0 Préambule

Nous allons travailler sur des machines virtuelles pour les TPs réseaux de cette année. Afin de se familiariser avec cette noouveauté, nous allons reprendre une activité de l'année passée.

Vous disposez de la possibilité d'ajouter des machines virtuelles depuis la Silverblue via la commande `vm-add`.

* consulter le manuel de `vm-add`
* ajouter trois machines virtuelles **debian** nommées respectivement `S1`, `S2` et `R1`.
* consulter le manuel de `vm-ls`
* lister les machines virtuelles disponibles
* consulter le manuel de `vm-run`
* démarrer `S1` et `S2`

```bash
exit # quitte ubuntu pour silverblue
sudo vm-add -d S1
sudo vm-add -d S2
sudo vm-add -d R1
# dans des terminaux séparés
vm-run S1
vm-run S2
vm-run R1
```

## 1 Rappels

### 1.1 Adresse IP

Chaque interface réseau (i. e. carte réseau) d'un hôte du réseau possède (au moins) une adresse IP unique dans le réseau. Une interface ethernet peut être désignée par un nom ”traditionnel” comme ethX où X est le numéro de l'interface ou un nom ”prédictible” comme enp0s25. Une adresse IP est un nombre de 32 bits souvent noté en décimal pointé ; quatre entiers (compris entre 0 et 255) séparés par des points. 

Exemple : 192.168.2.200. L'adresse IP est structurée en deux parties : 
1. partie réseau : permet de désigner le réseau (netID)
2. partie hôte : permet de désigner l'hôte dans le réseau 
  (hostID)

Un masque de sous-réseau permet de séparer partie réseau et partie hôte. Ce masque ; une succession de 1 suivi d'une succession de 0 donne l'étendue de la partie réseau.
* mettre la partie host à 0 nous donne l'adresse du réseau. Ceci se fait par un et bit à bit entre adresse IP et masque. Exemple pour 192.168.2.200 avec un masque de 255.255.224.0 : 
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

Un hôte voulant faire une transmission constitue un paquet IP qui contient l'adresse du destinataire et l'adresse de l'expéditeur. Au niveau de la couche réseau, le routage utilise une table de routage qui contient une ou plusieurs lignes contenant chacune essentiellement trois informations : 
1. une adresse de réseau
2. un masque de réseau
3. comment atteindre le réseau : soit directement par une interface connectée sur ce réseau (on parle de routage direct ou de remise directe), soit en passant par un routeur (on parle de routage indirect) qui est identifié par son IP et l'interface à utilisée pour l'atteindre.

Un routeur peut être un équipement spécialisé ou simplement un hôte ordinaire relié à plusieurs réseaux.

Le décision de routage se fait par la recherche d'une correspondance dans la table de routage en appliquant pour chaque ligne, le masque de réseau à l'adresse de destination. Quatre cas peuvent alors se présenter : 
1. le réseau de la destination est directement connecté. Il y une remise directe en utilisant le réseau local sous-jacent.
2. le réseau de la destination est accessible via un routeur. Le paquet est transmis au routeur sans changer les adresses IP de l'émetteur et du destinataire.
3. le réseau de la destination est absent de la table de routage, mais il existe une route par défaut. Le paquet est transmis au routeur désigné.
4. le réseau de la destination est absent de la table de routage, et il n'existe pas de route par défaut. Envoi d'un message ICMP à l'émetteur : Network is unreachable

Chaque routeur recevant un paquet IP applique le même algorithme.

### 1.3 Type de routage
Les deux types connus de routage sont : 
1. le routage statique
2. le routage dynamique ou auto-adaptatif
Par routage statique, on entend que les tables de routage sont renseignées manuellement, par opposition au routage dynamique où celles-ci sont construite par application d'un algorithme distribué au niveau des routeurs.

### 1.4 Commandes
Les commandes réseaux sont présentes sous deux formes :
* les commandes vénérables : `ifconfig`, `arp` et `route` que nous n'allons pas utiliser
* la nouvelle commande `ip` (du paquet iproute2) : `ip link`, `ip addr`, `ip neigh` et `ip route`

#### interface
consultation : `ip link`, activation d'une interface `eth1` : `ip link set eth1 up`

#### adressage
consultation : `ip addr`, configuration d'une interface `eth0` : `ip addr add 192.168.1.1/24 dev eth0`

#### voisinage
consultation  : `ip neigh`

#### routage
consultation : `ip route`, configuration : ajout avec `add`, suppression avec `del` ; route par défaut : `default`. Exemple :

```
ip route add 172.20.11.0/24 via 192.168.1.253
```

## 2 Configuration réseau
Observez la configuration réseau initiale de `S1`. 
* Combien d'interfaces réseaux sont-elles disponibles ? -> 3 interfaces
* Quelle(s) est(sont) vos ip ? -> `172.21.180.173`
* Qu'avez-vous comme route(s) ? -> 7
* Quels type de liens sont disponibles ? -> loopback, host, eth0 eth1, eth2, eth3
* Quelles sont vos adresses MAC ? -> *c'est long à recopier*
* Quel est l'état du cache arp ? -> *???*
* ajouter les ip `192.168.X.101/24` sur `s1` et `192.168.X.102/24` sur `s2` en utilisant dans les deux cas l'interface `eth3`
* lancer un `ping` (qui génére un trafic ICMP `echo request/echo reply
  permettant de tester la connectivité) entre `S1`et `S2`
* re-consulter l'ensemble des informations.

## 3 Routage statique
Il est rappelé que bien souvent, les stations ordinaires (qui ne sont pas des routeurs) se contentent d'UNE IP sur UNE interface avec UNE route par défaut.

### 3.1 Situation 1

![[situation1.png]]

- configurer les interfaces des trois stations (normalement, vous avez 4 interfaces à configurer). Il faut impérativement utiliser des interfaces homologues (i.e. de même nom) si vous voulez les faire communiquer.
- consulter les tables de routages des trois stations
- avec la commande `ping`, depuis `S1`, déterminer quelles sont les communications possibles et impossibles vers les 4 interfaces. Expliquer. 
- Dans `S1`, ajouter une route permettant d'atteindre le réseau `172.X.2.0/24`.
- Attention le `forwarding` (la fonction de retransmission d'un paquet) peut ne pas être actif. Pour en consulter la valeur, deux possibilités existent :
	- `cat /proc/sys/net/ipv4/ip_forward`
	- `sysctl net.ipv4.ip_forward`

pour modifier cette valeur, deux possibilités existent :
* `echo 1 > /proc/sys/net/ipv4/ip_forward` (que fait cette commande ?)
* `sysctl net.ipv4.ip_forward=1`

`sysctl` peut permettre de pérenniser la modification. (Dé)commenter la ligne `net.ipv4.ip_forward=1` dans le fichier `/etc/sysctl.conf`. Puis recharger la configuration, `sysctl -p`.
* ajouter la route ”retour” sur `S2`
* tester
* Sur `S1` et `S2`, supprimer les routes ajoutées et mettre seulement des routes par défaut.
* tester.

### 3.2 Situation 2

![[situation2.png]]
#### Mettre en place la situation :
* configurer les interfaces 
* configurer le routage
* tester
