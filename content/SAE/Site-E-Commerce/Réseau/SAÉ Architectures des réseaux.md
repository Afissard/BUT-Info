# R3.06 SAÉ Architectures des réseaux


## Présentation du sujet
Cette SAÉ aura pour objectif de : 
- configurer un réseau comportant un VLAN ;
- d'installer, configurer et déployer plusieurs services.

## Modalités
- vous travaillerez en groupe. 
- vous rendrez, sur madoc, une archive `Y_NOM1[-NOM?]*.tar.gz` d'un répertoire où `Y` désigne votre groupe de TD et `NOM1` et `NOM?` sont les noms de famille des membres du goupe. Ce répertoire contiendra vos rendus. Exemple de production d'une archive :
	```
	tar -czvf 3_MARTIN_MEYER.tar.gz 3_MARTIN_MEYER/ 
	```

## Travail à réaliser
Ci-dessous un schéma de l’organisation minimale du réseau que vous aurez à mettre en place. Vous choisirez vos noms de machines qui remplaceront ROUTEUR, INTERNE et EXTERNE. 

![](SAE/Site-E-Commerce/Réseau/fig.svg "")

### DNS
Choisissez un nom de domaine dans `.com`. Votre serveur DNS gérera a minima les noms suivants : ROUTEUR, www qui sera un alias d'une autre machine. Vous fournirez également la procédure et le(s) fichier(s) pour faire une délégation depuis le serveur de `.com`

### VLAN
Mettez en place un VLAN sur la partie privée de votre réseau.

### DHCP
Installez et configurez un serveur DHCP. Celui-ci fournira une route par défaut vers la machine ROUTEUR, le nom de domaine, et l’IP du serveur DNS. L'hôte INTERNE doit se voir délivrer une adresse fixe en fonction de son adresse MAC.

### HTTP
Installez et configurez un serveur HTTP sur INTERNE. Ce serveur doit répondre sur le port 80 et être accessible via http://ROUTEUR:8080 depuis EXTERNE.

### Autre
Vous pouvez compléter votre configuration par les éléments de votre choix (filtrage, services supplémentaires, mise en place d'une DMZ, plus de réseaux/machine avec routage dynamique, ...) 

## Rendu
Rédiger un document (`rapport.md`) expliquant vos choix. Vous fournirez également : 
	- des captures de trames démontrant le bon fonctionnement de votre réseau. En particulier, une capture de trame faite sur le ROUTEUR d'une requête HTTP émise depuis l'extérieur et une capture montrant que INTERNE peut accéder au LAN.
	- le résultat des commandes  `./log > MACHINE.log`, où MACHINE correspond à la machine sur laquelle doit être exécutée la commande avec les droits `root`, et doit être remplacée successivement par INTERNE et ROUTEUR
