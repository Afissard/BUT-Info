# Diagnostiquer une configuration réseau et configurer une interface réseau


## 1. Objectifs

Pouvoir
- auditer et configurer le paramétrage réseau sur votre machine Linux
- utiliser et lire les sorties des commandes réseaux : `ip`, `ping`, `ss`, `dig`...


## 2. Connaître 

Ci-dessous, n'exécuter des commandes que lorsque vous êtes explicitement invité à le faire.

Les commandes à disposition sont : 
* `ip l[ink]` pour configurer les interfaces réseaux au niveau liaison, 
* `ip n[eigh[bo[u]r)]` pour consulter les associations (IP, MAC) présente dans le cache ARP de la machine (remplace la commande  `arp`), 
* `ip a[ddr]` pour configurer les adresses IP sur chaque interface (remplace la commande  `ifconfig`), 
* `ip r[oute]` pour configurer les routes de routage (remplace la commande  `route`), 
* `ss` pour consulter les connexions TCP/UDP au niveau transport (remplace la commande  `netstat`), 
* `dig` pour interagir avec un serveur DNS et récupérer l'adresse IP correspondant à un nom de domaine (remplace la commande  `netstat`).


### 2.1 Connaître la configuration des interfaces IP réseau

La commande `ip addr` (anciennement `ifconfig`) permet de consulter et configurer le paramétrage des interfaces réseaux. 

Pour consulter le _manuel_ d'une commande linux (savoir à quoi elle sert, quelles sont ses options, voir des exemples d'utilisation), il faut taper dans un teminal :

    man ip

`man` signifie _manuel_

La commande `ip addr` indique les _interfaces_ disponibles, leur état d'activité, leur adresse MAC, ainsi que leur configuration réseau éventuelle en termes d'adresse IP v4/v6, masque, adresse de diffusion.

Les _interfaces_ sont un moyen pour désigner les différentes portes d'entrée/sortie sur le réseau de la machine courante.

L'interface _locale (lo)_ est ce que l'on appelle la _boucle sur soi-même_ (*loopback*). Elle permet, depuis votre machine, de faire appel à une application réseau qui tourne sur votre propre machine. Une application réseau se compose en fait de deux applications : une application _cliente_ qui sollicite un service et une application _serveuse_ qui rend un service. Pour attaquer une application serveuse, il faut son adresse _IP_ et son numéro de _port_ d'écoute. L'adresse IP associée est `127.0.0.1`.

L'adresse MAC sert à identifier une machine dans un réseau local. A chaque carte réseau correspond une adresse MAC unique.
Dans les réseaux d'infrastructure Ethernet, l'adresse MAC est aussi appelée adresse Ethernet.

L'adresse IP permet d'identifier une machine à un niveau indépendant de l'infrastructure physique qui soutient le réseau. En particulier elle permet d'identifier une machine dans l'Internet (c'est-à-dire à un niveau qui interconnecte tous les réseaux). L'adresse IP compte une partie pour désigner une adresse d'un réseau et une autre partie pour identifier une machine sur ce réseau.

Un masque sert à connaître la partie de l'adresse IP qui correspond à l'identifiant du réseau IP dans lequel se trouve la machine (via un "et logique"). Par complémentarité, la partie restante désigne l'identifiant IP machine.

Une adresse de diffusion permet d'adresser dans un même temps toutes les machines d'un réseau IP.

#### Questions

Dans un terminal, tapez la commande : `ip addr`  (ou `ip a` en abrégé)

- Identifiez le nom de chaque interface. 
- Quels "petits mots" vous permet de reconnaître l'adresse MAC d'une interface, son adresse IPv4, son masque et son adresse de diffusion ? 
- Même question pour IPv6. 

Dans la description de l’activité d’une interface, `<BROADCAST,MULTICAST,UP,LOWER_UP>` le « `UP` » signifie que l’interface est « active ». 
- Quelles interfaces sont effectivement actives ?

`LOWER_UP `est relatif à l’activité au niveau liaison. Si vous êtes en filaire et que vous lisez `NO-CARRIER` c’est probablement que votre câble est mal branchée…


### 2.2 Connaître l'état du cache ARP sur la machine

ARP est un protocole qui permet de trouver l'adresse MAC correspondant à une adresse IP donnée sur le réseau local. 
La commande `ip n[eigh[bo[u]r)]` (anciennement `arp`) rend compte de l'état des associations connues à un moment donné.

Par défaut, la commande affiche les adresses numériques des hôtes. Si vous souhaitez que le système substitue les adresses numériques par les noms symboliques des hôtes qu'il connait, il faudra utiliser le paramètre -h (e.g. `ip -h n`) qui correspond à un affichage plus "humain". 

#### Questions

Dans un terminal, tapez la commande : `ip n`  (ou `ip n` en abrégé)

- Quelles IP de machines sont actuellement présentes dans votre cache ARP ? 
- Voyez-vous une différence entre `ip -h n` et `ip n` ? Qu'en déduisez-vous ?


### 2.3 Connaître l'état des routes pour accéder aux réseaux locaux interfacés et joindre le(s) routeur(s)/gateway (i.e. passerelle)

La commande `ip  r[oute] (anciennement `route`) permet de consulter et configurer la table de routage de la machine courante.

Dans une table de routage, chaque ligne spécifie une route pour joindre une destination i.e. un réseau. Une route se compose donc en général de :  
* d'un couple adresse IP réseau de destination et masque associé. Le fait qu'une machine possède une adresse réseau signifie qu'elle sait joindre toutes les machines du réseau indiqué. Une entrée supplémentaire pour indiquer comment joindre une machine appartenant à ce réseau est donc inutile. 
* d'une interface (ou *device*) par laquelle il faut sortir pour prendre la route qui conduit à la destination
* Lorsqu'une destination n'est pas directement accessible (e.g. que l'IP de la machine demandée ne se trouve pas sur un des réseaux locaux de la machine courante), on indique l'adresse IP de la *gateway/passerelle* (machine sur un réseau local directement accessible) à qui transmettre les paquets et donc *via* laquelle les paquets envoyés seront acheminés au destinataire ; on suppose que cette gateway saura quoi faire...
* On rajoute en général une ligne pour décrire une "route par défaut" (implicitement avec masque 32 en IPv4) pour traiter tous les paquets dont le destinataire ne correspondrait à aucune des routes configurées.
Si une destination n'est pas présente ou qu'il n'y a pas de route par défaut, on ne saura pas transmettre de paquets à cette destination.

Dans l'affichage suivant, on peut observer que la route par `default` passe *via* la machine ayant l'IP 192.168.1.254 et que pour atteindre cette machine il faut passer par l'interface (appelée *dev(ice)*) `wlo1`. On note aussi que la machine possède une route pour joindre le réseau 192.168.1.0 avec le masque 24. On lit sur cette ligne que la machine courante possède l'adresse IPv4 192.168.1.2 sur ce réseau.

    default via 192.168.1.254 dev wlo1 proto dhcp metric 600 
    192.168.1.0/24 dev wlo1 proto kernel scope link src 192.168.1.2 metric 600 


#### Questions

Dans un terminal, tapez la commande : `ip route`  (ou `ip r` en abrégé). 

- Identifier la liste des adresses de destination référencées. S'agit-il d'adresses machine ou réseau ?
- Pour chaque destination indiquez l'adresse de destination, le masque, l'interface de sortie et éventuellement l'adresse de la passerelle.
- Dans le cas où vous voyez le symbole « * », que signifie t’il à votre avis ?


### 2.4 Connaître les numéros de port des services

Les numéros de port servent à identifier des instances de service (serveur). Si l'on souhaite faire tourner deux serveurs identiques (par exemple deux serveurs web) sur une même machine, il faudra leur octroyer deux ports d'écoute distincts ! Les 1024 premiers numéros de port sont réservés. Pour qu'un serveur écoute derrière un port il faut l'activer.

Le fichier `/etc/services` recence les ports connus. Quelques services sélectionnés :

* `ftp` est un protocole et une application réseau pour gérer l'hébergement de fichiers à distance
* `ssh` est un protocole et une application réseau pour obtenir un shell distant sécurisé
* `smtp` est un protocole et une application réseau gérant la transmission de mail
* `dns` est un protocole et une application réseau gérant la résolution d'adresse IP et de nom de domaine
* `http` (ou Word Wide Web / www) est un protocole et une application réseau la demande et la réception de contenu distant ; réseau dit Web


La commande `ss -taupen` (et anciennement `netstat -taupen`) permet de connaître l'état des connexions de votre machine avec d'autres machines et par là de connaître les services actifs sur votre machine et les clients de votre machine impliqués dans des connexions. Une colonne désigne l'IP et le port local, une autre désigne éventuellement l'IP et le port de la machine distante. Une colonne désigne si la connexion est établie (`ESTABLISHED`) ou seulement en écoute (`LISTEN`). En "écoute", signifie que votre machine a un service (identifié par le numéro de port ou un nom symbolique) qui est disponible sur votre machine.


#### Questions

- A quoi servent les services suivants : ftp, ssh, smtp, doman (dns), http ? Quel est leur numéro de port théorique ?

Dans un terminal, tapez la commande : `ss -taupen`. 
- Les services précédemment mentionnés sont-ils actuellement actifs sur votre machine ?


### 2.5 Connaître les serveurs DNS

Pour se faire taper la commande suivante

    cat /etc/resolv.conf

(sur ubuntu et à partir de la version 12.04 `cat /etc/network/interfaces` et `/run/resolvconf/resolv.conf`)


## 3. Configuration et test d'une interface réseau

La commande `ip` (anciennement `ifconfig`) permet de configurer une interface. C'est-à-dire d'associer une adresse IP à une interface.
Avec la commande `ifconfig` quand le masque n'était pas communiqué, celui-ci était déduit de la classe d'adresse supposée de l'adresse. On va constater que ce n’est pas le cas de la commande `ip`…
L'adresse de diffusion est déduite quant à elle automatiquement du masque.
Avec la commande `ifconfig`, lorsqu'une adresse IP était associée à une interface, l'interface était automatiquement activée et une route était automatiquement ajoutée (visible avec la commande `ìp route`). Nous verrons qu’avec la commande `ip`, il faudra clairement expliciter ces différentes informations.

Chacune de vos machines ont un identifiant. Voir le câblage derrière votre poste.(11)
Par la suite vous utiliserez l'adresse IP suivante 192.168.1.VOTRE_IDENTIFIANT.
Nous travaillerons avec une interface disponible, ce devrait être le cas de l'interface 1 bleue.

Passez en root avec la commande `sudo su`

Tapez `ip a`. L'interface bleue est-elle configurée (voyez-vous une IP) ?

Ajoutez une adresse IP » à l’interface en oubliant de préciser le masque... Tapez

    ip a add 192.168.1.VOTRE_IDENTIFIANT dev bleue

puis  `ip a`. Quel est le masque associé ? (32)

Pour tenter de corriger, tapez la commande suivante en précisant un masque (ici /24)

    ip a add 192.168.1.VOTRE_IDENTIFIANT/24 dev bleue

puis `ip a`. Est-ce que cela a corrigé le problème ? (on a deux config, une avec un masque 24 et une avec un masque 32)

Pour supprimer une configuration associée à une interface, utiliser la commande suivante (attention à bien préciser le masque…).

    ip a del 192.168.1.VOTRE_IDENTIFIANT/MASQUE dev bleue

L'interface est-elle active (voyez-vous un « UP ») ? Pour activer une interface, tapez :

    sudo ip link set up dev bleue

(pour la désactiver utiliser `down` à la place de `up`)

Vérifiez que vous avez bien une route pour la remise directe (i.e. vers votre réseau local)

    ip r

Demander à votre voisin de vous indiquer son adresse MAC. (04:8d:38:cf:7d:41) Consulter votre cache ARP. Est-elle présente ? (elle ne le devrait pas) (c'est le cas, elle n'apparait pas)

Configurez votre interface pour être sur le même réseau (même masque et même identifiant réseau) que votre voisin. "Pinguez le".

    ping IP_DE_VOTRE_VOISIN
(ping 198.168.1.4)
Si l'affichage n'indique pas que le host est inconnu alors le ping est reçu et le "pong" est renvoyé.

Si le ping a fonctionné, consulter votre cache ARP, l'adresse MAC de votre voisin est-elle présente ? (oui)

Configurer votre interface pour être sur un réseau différent de votre voisin (changer seulement le masque). Est-ce que le ping marche encore ?


## 4. Pour aller plus loin

* présentation de la salle réseau (système live + baie de brassage)
* écouter via wireshark la communication ping entre deux stations et demander à une 3e station "d'emprunter" l'IP d'un de deux communiquants. Même question avec l'adresse MAC.
* simuler un client web `nc localhost 80`
* ouvrir une communication tcp entre deux stations en ouvrant un serveur sur une station `nc -l 8080` et en lançant un client depuis l'autre station `ǹc ip_station 8080`
* réaliser une capture de trames sur différents réseaux et observez ce qui a été capturé (avec/sans traffic de votre part)
* Depuis la Debian 10 (sortie en juillet 2019), les commandes `ifconfig`, `arp` et `netstat` ne sont plus présentes sur les systèmes Debian. La commande `ifconfig` a été remplacée par `ip a[ddr]`, `arp` par remplacée par `ip n[eigh[bo[u]r)]`, `route` est remplaçée par `ip r[oute]`, et `netstat` remplacée par `ss`. Les pages wikipedia de chaque commande fournissent un historique plus précis de l'obsolescence et du remplacement des anciennes commandes sur chaque OS.
* `ip -br[ief]` print only basic information in a tabular format for better readability. This option is currently only supported by `ip -br addr show`, `ip -br link show` and `ip -br neigh show` commands, not for `route`... 
* A toutes fins utiles, la commande suivante permet d'améliorer le visuel de l’affichage en transformant les espaces en tabulations… :  `ip route|column -s " " -t`
