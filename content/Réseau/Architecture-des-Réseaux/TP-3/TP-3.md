---
title: TP-3
draft: 
description: 
tags:
  - Réseau
---
# Td Architecture des réseaux : DNS

Le but de ce tp est de vous familiariser avec le problème de la résolution de nom. Pour ce faire nous allons successivement :

* configurer un client d'un serveur de résolution de noms
* configurer un serveur de nom d'une zone

## Préambule

* Créez et lancez la machine correspondant au DNS : 
```
vm-tp-add ro-dns
vm-run ro-dns
```

* Créez et lancez une VM debian `nsX` (`ns` comme *Name Server*). Elle aura pour ip : `192.168.X.1/24` sur `eth0`.

* Installez `bind` sur votre VM :
 ```
 apt-get update
 apt install bind9
 ```

* Ajoutez la route suivante :
  ```
  ip route add 192.168.0.0/16 dev eth0
  ```
* Vérifiez que vous arrivez à contacter la machine `192.168.0.254` : 
  ```
  ping 192.168.0.254
  ```

## Introduction

Le problème de la résolution de nom sur Internet est à l'image de celui du réseau téléphonique. Il existe un service sur Internet capable de nous fournir l'adresse ip d'une machine comme www.univ-nantes.fr comme il existe un service de renseignement capable de nous donner le numéro d'une personne à partir
de son nom. De même qu'avec le téléphone où il faut connaître un seul numéro (le 118 ...), il faut connaître l'adresse ip du serveur de noms. Enfin, dernière analogie, si le 118 ... ne répond pas, même si la ligne de notre correspondant
fonctionne, nous serons incapables de le joindre.

### Principes du DNS

Le système de noms de domaine (*Domain Name System*) est la base de données distribuée du nommage hiérarchique des hôtes de l'Internet. Cette base est accessible via un mécanisme client/serveur. La partie serveur du système est assurée par des programmes appelés **serveurs de noms**. La plus populaire mise en oeuvre est BIND (Berkley Internet Name Domain), actuellement maintenue par l'ISC (Internet Software Consortium). Les utilisateurs accèdent aux serveurs de noms par des programmes appelés solveurs (*resolver*). Pour faire correspondre une adresse ip à un nom, un programme appelle le solveur et lui passe le nom de l'hôte recherché en paramètre (par exemple, `www.wikipedia.fr`). Le solveur envoie un paquet UDP au serveur DNS configuré qui peut traiter la requête de plusieurs façons :

1. le nom demandé figure dans sa base et il donne la réponse ;
2. le nom ne figure pas dans sa base. Le solveur commence par consulter un serveur de la racine. Il y a deux possibilités :
    1. le serveur de la racine fait suivre la requête, obtient la réponse (finale) et la retransmet : requête récursive.
    2. le serveur de la racine se contente d'envoyer la liste des serveurs de la `fr`. Le serveur DNS local contacte alors un des serveurs obtenus de `fr` pour obtenir la liste des serveurs de `wikipedia.fr` avant d'en contacter
       un et d'obtenir sa réponse : requête itérative
3. Le nom ne figure pas dans sa table et le serveur est configuré pour rediriger les requêtes qu'il ne peut traiter vers un redirecteur (*forwarder*) qui lui renverra la réponse.

L'adresse ip ainsi obtenue est renvoyée au solveur qui la renvoie à l'appelant. Celui-ci peut alors envoyer un paquet à l'IP obtenue.

### Domaine

La structure de la base de données est arborescente avec comme racine `.` et des domaines de haut niveau (*top level domain*). Chaque domaine est désigné par le chemin à suivre depuis la racine.

### Délégation

Pour créer un nouveau domaine, on doit avoir l'autorisation du domaine parent (ex : `univ-nantes.fr` a été autorisé par le gestionnaire de `fr`). À l'opposé, on peut créer des sous-domaines de son propre domaine sans en référer à quiconque. Par exemple : `univ-nantes.fr` peut créer `iut-nantes.univ-nantes.fr`. Ce transfert de l'autorité de la gestion d'un sous-domaine s'appelle la **délégation** de zone.

### Domaine vs zone

Pour un domaine sans sous-domaine, les notions de zone et de domaine sont identiques. Pour un domaine découpé en sous-domaines, la zone contiendra essentiellement des informations de délégation aux sous-domaines, alors que le domaine recouvre la zone “principale” et les zones déléguées

### Zone

Un serveur de noms d'une zone dispose de toutes les informations concernant cette zone. Il stocke ces informations dans un fichier local obtenu :

* soit par une saisie locale : c'est le cas d'un serveur maitre.
* soit par un téléchargement depuis un autre serveur de noms : pour un serveur esclave

### Résolution inverse

Le problème de la résolution d'un nom vers une adresse étant traité, il reste à savoir faire le contraire, c'est-à-dire passer d'une adresse à un nom. L'idée qui a été retenue est toute simple : dans un domaine particulier appelé `in-addr.arpa`, on utilise simplement des adresses comme nom des sous-domaines.

### Rôle d'un serveur de noms

Le serveur de DNS a donc un rôle double :

1. il est l'interlocuteur du solveur pour diligenter une résolution de nom
2. il détient une partie des données.

## Client : configuration

Un hôte utilisant un solveur est appelé client. C'est à la configuration de cette partie cliente – à l'image de ce que vous pouvez faire quand vous avez un abonnement chez un fournisseur d'accès internet – qu'est consacrée cette partie du tp. 

### Fichiers de configuration

#### Fichier `/etc/nsswitch.conf`

Le fichier `/etc/nsswitch.conf` permet de contrôler la manière de procéder à la résolution de nom. Cela prend la forme d'un une ou plusieurs spécifications de services. L'ordre des services sur la ligne détermine l'ordre dans lequel ces
services seront interrogés, chacun leur tour, jusqu'à ce qu'un résultat soit trouvé. Exemples :

```
hosts:          dns [!UNAVAIL=return] files
hosts:          files mdns4_minimal [NOTFOUND=return] dns myhostname
```

En parallèle, il existe un fichier `/etc/host.conf` qui historiquement jouait le rôle de `/etc/nsswitch.conf` et qui conserve un rôle pour la résolution des hôtes locaux.

### Fichier `/etc/hosts`

C'est le fichier qui contient les hôtes locaux. Il évite la mise en place d'un serveur de nom, mais la mise à jour de ce
fichier est rapidement problématique pour un réseau local évolutif.

```
127.0.0.1         localhost
```

### Fichier `/etc/resolv.conf`

Le comportement du solveur est configuré dans le fichier
`/etc/resolv.conf` à l'aide de quelques directives dont les principales sont présentées ici :

`nameserver` : cette directive fournit au solveur l'adresse ip d'un serveur de noms à interroger.

```
nameserver 10.2.3.1
nameserver 10.2.3.2
```

Si plusieurs serveurs de noms sont déclarés, le solveur commence par interroger le premier. Après un délai d'attente, si aucune réponse n'est parvenue, il interroge le deuxième, puis peut recommencer plusieurs cycles avec des délais
différents d'attente avant d'afficher son échec.

`search` : cette directive permet de fixer la liste de recherche des domaines (jusqu'à 6) à parcourir. Par exemple : si vous avez une entrée

```
search dom1.com dom2.com
```

lors d'une requête de résolution pour un hôte `mach` non pleinement qualifié, on recherche d'abord `mach.dom1.com` puis `mach.dom2.com`.

### Réalisation

1. Avant configuration vérifiez que la résolution de nom ne fonctionne pas :
   ```
   host nsX.tp.com
   ```

2. La machine à l'adresse `192.168.0.254` a été configurée comme serveur de noms pour la zone `tp.com`. Configurez votre machine pour être client de ce serveur DNS (modifiez le fichier `/etc/resolv.conf`).

3. Après configuration, la résolution de nom doit fonctionner.

## Serveur : explication sur bind

### Fichier `named.conf`

Comme tout démon sous unix, bind lit un fichier de configuration (`/etc/bind/named.conf`). Nous trouvons dans ce fichier (ou dans les fichiers inclus depuis celui-ci) la déclaration des zones gérées et leur association avec les fichiers constituant la base de données.

`zone` : permet de définir le type et l'emplacement du fichier des données pour une zone.

Les trois types que vous allez rencontrer aujourd'hui sont les suivants :

1. type `master` : serveur maître (*master*) d'une zone

exemple

```
zone "unezone.com" in { //  déclaration de la zone
          type master; // déclaration type maître
          file "/etc/bind/unezone.db"; // on indique le fichier contenant les enregistrements de cette zone

};
```

Ce qui se lit “je suis serveur maître de la zone `unezone.com`, je récupère les données de cette zone dans le
fichier `unezone.db`

2. type `slave` : serveur esclave (*slave*) d'une zone
3. type `hint` : configuration de la localisation des serveurs racines

```
zone "." in {
          type hint; 
          file "/etc/bind/db.root"; 
};
```

La localisation (adresses ip) et les noms des serveurs racines sont dans le fichier indiqué par `file`. Le fichier à utiliser a été modifié pour refléter la situation particulière de notre tp avec un serveur racine interne. Le fichier original des serveurs de la racine est : `/usr/share/dns/root.hints`.

`options` : permet de définir des options globales de bind, comme le répertoire de travail, les serveurs à qui nous
relayons les requêtes, ...

exemple :

```
options {
        directory "/var/cache/bind"; // répertoire de travail par défaut
        pid-file "/var/run/named/named.pid"; 
        forwarders { 192.168.18.3; 192.168.18.4};
};
```

Sur cet exemple, tous les fichiers de zones sont à mettre dans le répertoire `/var/cache/bind` ou sous forme d'un chemin *absolu*.

#### Remarque

La différence essentielle entre un serveur maître et un esclave est la provenance de leurs données. Le maître obtient ses données à partir de fichiers alors que l'esclave les *télécharge* depuis un autre serveur de nom.

### Fichiers de zones

Le format des fichiers de configuration est défini de manière standard dans les RFC (1034 et 1035) et est donc le même quel que soit le serveur de DNS choisi. Ces fichiers contiennent des enregistrements de ressource (`RR` *Ressource
Records*). Chaque enregistrement de ressource est un quintuplet :

```
Nom_de_domaine  Durée_de_vie  Classe  Type  Valeur
```

* `Nom_de_domaine` : le domaine que concerne l'enregistrement
* `Durée_de_vie` : le champ durée de vie donne une indication sur la stabilité dans le temps de cet enregistrement (une information très stable se verra affecter une valeur comme 86400 (1 journée en secondes), une information peu stable se verra affecter une valeur comme 60 (1 minute).
* `Classe` la classe `IN` pour Internet qui peut être omise
* `Type/Valeur` indique le type d'enregistrement. Les types les plus importants sont les suivants :
    * `SOA` *Start Of Authority*
    * `A` correspondance nom-adresse : une adresse ip
    * `MX` relais de messagerie
    * `NS` serveur de nom
    * `CNAME` nom canonique (alias) : le correspondant de cet alias
    * `PTR` correspondance adresse-nom : un nom d'hôte

Chaque fichier de zone, à l'aide de ces enregistrements va définir en particulier :

* Qui a l'autorité ? L'autorité est détenue par LE serveur maître. Cette information est renseignée par un
  enregistrement SOA (un et un seul). Un enregistrement pour le serveur maître primaire.
* Quels sont les serveurs de noms de la zone : c'est le rôle des enregistrements `NS`. Il peut y avoir plusieurs
  serveurs de noms pour une zone, un maître primaire et des esclaves. Chacun va faire l'objet d'un enregistrement `NS`.
  Un enregistrement `NS` pour chaque serveur de nom de la zone.
* Enfin dans une zone de résolution directe, on va définir la correspondance nom-adresse à l'aide des enregistrements `A` ; dans une zone de résolution inverse, on va définir la correspondance adresse-nom à l'aide des enregistrements `PTR` ; Un enregistrement `A` ou `PTR` (selon la zone) pour chaque hôte de la zone.

## Serveur :  serveur maitre

Chacun va jouer le rôle d'un fai (Fournisseur d'Accès à Internet). Un fai gère généralement sa zone et doit donc administrer un serveur de nom. Chaque machine `nsX` est destinée à devenir le serveur maître d'une zone `faiX.com`, sous-domaine de com. Le serveur mis en place sur `192.168.0.254` fait office de serveur racine et gère également la zone `com`. La délégation des zones `faiX.com` y a été faite.

### Reconfiguration de `/etc/resolv.conf`

Modifiez le fichier `/etc/resolv.conf` pour indiquer que le serveur DNS par défaut c'est vous.

```
nameserver 192.168.X.1
```

### Fichier /etc/bind/named.conf

Modifiez soit le fichier `named.conf`, soit un des fichiers inclus depuis ce fichier : `named.conf.local` (conseillé).

* déclarez la zone directe `faiX.com` de type maître ;
* consultez les déclarations des zones directes et inverses de résolution locale (dans le fichier `named.conf.default-zones`) ;
* consultez le fichier `/usr/share/dns/root.hints` (pour votre culture).
* remplacez le par le fichier  [`root.hints`](https://gitlab.univ-nantes.fr/iut.info2.r3_06/r3_06.tp3/-/raw/main/root.hints) fourni, (`wget https://gitlab.univ-nantes.fr/iut.info2.r3_06/r3_06.tp3/-/raw/main/root.hints --output-document=/usr/share/dns/root.hints`): ;
* enfin, dans le fichier `named.conf.options`, passez la valeur de `dnssec-validation` de `auto` à `no`.

### Définition de la zone faiX.com

La seule machine dans votre zone c'est vous, vous en êtes le serveur maître primaire et donc le serveur de nom. Il faut mettre la correspondance nom–adresse de votre machine. Pour simplifier, télécharger et adapter le
fichier : [`fai.db`](https://gitlab.univ-nantes.fr/iut.info2.r3_06/r3_06.tp3/-/raw/main/fai.db).

### Définition des zones directe et inverse de résolution locale

Cette partie est déjà configurée sur vos VMs.

### Démarrage et tests

Le démon `named` est contrôlé de manière standard :

```
systemctl stop ou start ou restart ou status bind9 
```

À chaque (re)démarrage de `named`, les fichiers de configuration sont relus. Pour suivre l'activité du démon `named` utilisez la commande suivante :

```
journalctl -u named -f
```

Un démarrage réussi ressemble à ceci :

```
systemd[1]: Started named.service - BIND Domain Name Server.
named[20332]: starting BIND 9.XXXX
...
named[20332]:
----------------------------------------------------
named[20332]: BIND 9 is maintained by Internet Systems
Consortium,
named[20332]: Inc. (ISC), a non-profit 501(c)(3) public-benefit
named[20332]: corporation.  Support and training for BIND 9 are
named[20332]: available at https://www.isc.org/support
named[20332]:
----------------------------------------------------
...
named[20332]: listening on IPv4 interface ..., 192.168.X.1#53
...
named[20332]: command channel listening on 127.0.0.1#953
named[20332]: command channel listening on ::1#953
named[20332]: managed-keys-zone: loaded serial 339
named[20332]: zone 0.in-addr.arpa/IN: loaded serial 1
named[20332]: zone 255.in-addr.arpa/IN: loaded serial 1
named[20332]: zone faiX.com/IN: loaded serial * //---------------IMPORTANT
named[20332]: zone X.168.192.in-addr.arpa/IN: loaded serial * //---------------IMPORTANT
named[20332]: zone 127.in-addr.arpa/IN: loaded serial 1
named[20332]: zone localhost/IN: loaded serial 2
named[20332]: all zones loaded
named[20332]: running 
//---------------IMPORTANT
```

Les tests consistent à vérifier :

```
host nsX.faiX.com
```

### Rajout d'hôtes dans la zone `faiX.com`

Un serveur DNS sert à gérer toute une zone comportant plusieurs hôtes. Déclarez dans votre zone un hôte `cequevousvoulez` en `192.168.X.100`.

Redémarrez le serveur et testez.

```
host cequevousvoulez.faiX.com
```


### Alias

Votre machine `nsX` va aussi faire office de serveur web. Comme c'est l'usage, vous souhaitez que ce serveur web réponde à `www.faiX.com`. Il vous faut donc ajouter un enregistrement `CNAME` pour spécifier ceci. Faites-le, puis redémarrez le serveur et testez.

```
host www.faiX.com
```

### Serveur : serveur esclave

Configurez une deuxième VM `Y` qui va devenir un serveur secondaire de la zone `faiX.com`.

Nous commençons par rajouter une ip sur `nsY` pour que la délégation faite par le serveur racine fonctionne : `192.168.X.2/24` sur `eth0`.

### Sur nsY : modification du fichier named.conf

Il faut déclarer la zone `faiX` de type esclave.

```
zone "faiX.com" in {
   type slave;      // type esclave
   file "dbX.fai"; // nom du fichier de zone, qui sera créé sur nsY
   masterfile-format text; // format texte pour le transfert de zone
   masters { 192.168.X.1 ;}; // l'adresse du serveur maître
};
```

Ceci se lit “je suis serveur secondaire de la zone `faiY.com`, je récupère les données de cette zone au niveau du serveur primaire dont l'ip est `192.168.X.1`”

C'est tout. Le serveur esclave va télécharger depuis le (un)
serveur maître les données de la zone `faiY`. Ce qu'on appelle un
**transfert de zone**.

#### Important

Si vous indiquez un chemin absolu , il faut que le répertoire indiqué soit accessible en écriture pour `bind` (utilisateur `bind` du groupe `bind`). Si vous indiquez un chemin relatif, le répertoire de base est `/var/cache/bind`, qui a les bonnes permissions.

### Sur `nsX` : modification du fichier de zone `db.faiX`

La zone `faiX.com` comporte maintenant deux serveurs de noms. C'est cette information qu'il faut mettre à jour dans le fichier de zone :

```
IN NS nsX.faiX.com.
IN NS nsY.faiX.com.
```

L'enregistrement `NS` n'indique pas si le serveur est maître ou esclave. Les deux serveurs font autorité sur la zone.

#### Note

Si vous indiquez `nsY.faiX.com`, il faut que `nsY` soit défini dans la zone `faiX.com` par une entrée `A`.

### Démarrage et tests

Redémarrez le primaire puis l'esclave :

```
named[615]: zone faiX.com/IN: Transfer started.
named[615]: transfer of 'faiX.com/IN' from 192.168.X.1#53: connected using 192.168.X.1#53
named[615]: zone faiX.com/IN: transferred serial 1
named[615]: transfer of 'faiX.com/IN' from 192.168.X.1#53: Transfer status: success
named[615]: transfer of 'faiX.com/IN' from 192.168.X.1#53: Transfer completed: 1 messages, XX records, XX bytes, XX secs (XX bytes/sec) (serial XX)
```

Vérifiez que le fichier correspondant est créé sur `nsY` et visualisez son contenu.

## Serveur : délégation

La machine `nsX` va déléguer la gestion de la zone `sous.faiX.com` à `nsY`.

### Sur `nsX`

Dans le fichier de la zone `faiX.com`, mettre les informations de délégation :

```
sous   IN NS nsY.faiX.com.
```

### Sur `nsY`

Dans le fichier `named.conf`, rajouter une nouvelle zone primaire pour `sous.faiX.com`,

```
zone "sous.faiX.com" in {
        type master;
        file "/etc/bind/db.sous"; 
};
```

puis écrire le fichier de zone.

### Démarrage et tests.

Redémarrez les serveurs puis testez.
