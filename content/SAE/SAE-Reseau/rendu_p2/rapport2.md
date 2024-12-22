# Manuel utilisateur

* NOMS : Sacha Chauvel, Rafael Café
* GROUPE de TP : 4-2
* X : 128
* SECRET : Regulus et  Fomalhaut
* IP_CLIENT : 172.20.128.1
* IP_DHCP : 172.20.128.2
* IP_SMTP : 192.168.0.28
* IPs du ROUTEUR : 172.20.128.3 et 192.168.0.128 


**Note**: Ce document d’une à deux pages, doit décrire au choix : 
* soit comment comment installer un VLAN 
* soit comment configurer un serveur DHCP. 
Une personne ne sachant pas configurer un de ces services doit être capable de le faire en suivant le manuel (description globales des étapes et détails sur les différentes machines, commandes à exécuter, codes à modifier, commandes de test à réaliser, résultats à observer, etc.). Votre manuel doit être complet e.g. si votre service requiert une configuration réseau alors votre manuel doit aussi l'expliquer. Indiquer explicitement les adresses IP et éventuellement les adresses MAC des machines que vous utilisez si c’est pertinent.

# Tutoriel installation d'un serveur DHCP
## Installation du serveur
Exécuter les commande suivante pour créer un VLAN :
```bash
modprobe 8021q # activation service VLAN
ip link add link jaune name jaune.128 type vlan id 128 # ajout d'une interface pour la VLAN
ip link set up dev jaune.128 # activation de l'interface
ip a a 172.20.128.254/24 dev jaune.128 # ajout d'une adresse ip à la machine pour l'interface
```
Modifier ensuite le fichier de configuration `/etc/default/isc-dhcp-server` et modifier la ligne suivante : `INTERFACESv4=jaune.128`.
Dans le fichier `/etc/dhcp/dhcpd.conf`, commenter tout son contenu (avec des `#` en début de ligne) et copier ceci :
```dhcpd.conf
# paramètre globaux
option domain-name "FORMALHAULTREGULUS.org";
option domain-name-servers 192.168.0.254; # IP du DNS
option routers 172.20.128.254; # IP du routeur
option subnet-mask 255.255.255.0;

subnet 172.20.128.0 netmask 255.255.255.0 {
	range 172.20.128.2 172.20.128.250; # Plage d'adresses IP DHCP
	option routers 172.20.128.254; # IP du routeur
	option subnet-mask 255.255.255.0;
}

# Adresse fixe
host appareil1 {
	hardware ethernet 04:8d:38:c2:57:cd; # addresse MAC de l'interface du client
	fixed-address 172.20.128.1; # ip donnée au client
}
```
## Installation d'un client
Exécuter les commande suivante pour se connecter à la VLAN :
```bash
modprobe 8021q # activation service VLAN
ip link add link jaune name jaune.128 type vlan id 128 # ajout d'une interface pour la VLAN
ip link set up dev jaune.128 # activation de l'interface
```
Pour recevoir une nouvelle IP de la part du DHCP, assuré vous d'abord ne pas avoir de configuration DHCP cliente avec la commande suivante :
```bash
dhclient -r
```
Exécuter ensuite la commande suivante pour demander une nouvelle IP au serveur DHCP :
```bash
dhclient
```
# Analyse de trame
Appliquer le masque `bootp` pour n'afficher que les requêtes lié au DHCP.
Voici la reformulation :

## De `0.0.0.0` à `255.255.255.255` : `DHCP Discover`
Le client DHCP envoie un paquet `DHCP DISCOVER`, un message de diffusion à tous les ordinateurs du sous-réseau. Seul le serveur DHCP peut répondre à cette demande.
## De `172.20.128.254` à `172.20.128.1` : `DHCP Offer`
Tout serveur DHCP sur le sous-réseau peut répondre en envoyant un paquet `DHCP OFFER`, offrant ainsi une adresse IP potentielle au client.
## De `0.0.0.0` à `255.255.255.255` : `DHCP Request`
Après avoir reçu le paquet `DHCP OFFER`, le client peut recevoir d'autres offres d'autres serveurs DHCP. Il doit choisir l'offre du serveur qui a répondu en premier et envoyer un paquet `DHCP REQUEST`, identifiant le serveur sélectionné. Cela informe les serveurs DHCP de l'acceptation de l'offre par le client.
## De `172.20.128.254` à `172.20.128.1` : `DHCP ACK`
Les serveurs DHCP reçoivent le paquet `DHCP REQUEST` du client. Les serveurs non sélectionnés prennent cela comme un refus, tandis que le serveur choisi enregistre l'adresse IP du client et répond avec un message `DHCP ACK`. Si le serveur ne peut pas fournir l'adresse promise, il envoie un datagramme `DHCP NAK`.