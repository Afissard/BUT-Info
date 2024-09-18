---
title: TP-1
draft: 
description: 
tags:
  - Réseau
---
# Td Architecture des réseaux : configuration d'un serveur web Apache

## Préambule

Dans ce Td, vous mettrez en place plusieurs sites, accessibles via plusieurs noms et/ou IPs. Ces sites seront hébergés sur une *machine virtuelle* debian que vous appellerez `serveurweb` et abrégé VM dans la suite.

* créez et lancez votre VM

* Votre VM `X` (`X` allant de 1 à 35) aura pour ip :
  `192.168.X.1/24` (qui correspond aux noms `wwwX.fai.com` et `wX.fai.com`) et `192.168.X.2/24` (associée au nom `wwX.fai.com`). Configurez ces deux adresses sur `eth0`.

* Installez apache sur votre VM:
 ```
 apt-get update
 apt install apache2
 ```
* La machine `192.168.0.254` a été configurée comme serveur de noms pour la zone `fai.com`. Pour utiliser ce serveur DNS, remplacez le contenu de votre fichier `/etc/resolv.conf`, par la ligne **unique** suivante :

  ```
  nameserver 192.168.0.254
  ```

* Pour pouvoir contacter cette machine, ajoutez la route suivante :

  ```
  ip route add 192.168.0.0/16 dev eth0
  ```

* Enfin, créez et lancez la machine correspondant au DNS : 
```
sudo vm-tp-add dns
vm-run ro-dns
```

* Testez (et validez) la résolution des trois noms :

  ```
  host wwwX.fai.com
  host wwX.fai.com
  host wX.fai.com 
  ```

* Pour pouvoir atteindre votre machine virtuelle sur les ips ajoutées, nous allons devons mettre en place, depuis un terminal sur la Silverblue, un proxy socks (dont nous ne détaillons pas le principe ici) : 
```
  ssh -fN -D 8080 root@ipSurHost0DeVotreVm
```
Pour plus de détail, consultez : [`https://www.iut-nantes.univ-nantes.fr/etu/wiki/index.php/Machines_virtuelles.html#Exemple_serveur_web_via_machine`](`https://www.iut-nantes.univ-nantes.fr/etu/wiki/index.php/Machines_virtuelles.html#Exemple_serveur_web_via_machine`)
  

* Testez `wwwX.fai.com` depuis le navigateur **Librewolf** dont vous aurez configuré manuellement le proxy : `SOCKS Host 127.0.0.1 Port 8080`. Vous devriez obtenir la page suivante :
  ![](Figs/default.png "")

La page qui s'affiche sur ce site par défaut contient informations et explications sur la configuration d'un serveur web apache sur une distribution debian. Lisez-la (attentivement).

# 1 Introduction

Un serveur Web est un serveur destiné à publier des contenus sur Internet. Il communique avec un client web en utilisant
le protocole HTTP (Hyper Text Transfert Protocol). La configuration d'un serveur web va permettre de définir en
particulier :

* les sites : quel(le)s noms, IPs, ports permettent d'accéder à quels sites

* le contenu d'un site : des fichiers dans des répertoires, éventuellement sur plusieurs partitions (ou accessibles par
  le réseau), vus comme une arborescence unique pour un client web

* les types de contenu : fichiers statiques (html, images), pages dynamiques (php, jsp,...) et leur permission :
  lecture, exécution, listage, protection par mot de passe,....

# 2 Apache

Apache est un serveur web parmi les plus utilisés. Parmi les facteurs du succès d'apache, on peut citer :

* le mode de distribution d'apache qui est fournit avec ses sources et permet (gratuitement) une utilisation non commerciale aussi bien que commerciale.
* l'architecture modulaire ; les utilisateurs d'apache peuvent facilement rajouter des fonctionnalités et adapter apache à leur propre besoin.
* la portabilité : il existe des versions d'Apache pour tous les Unix (dont Linux bien sur), mais aussi pour windows,
  ...
* enfin robustesse et sécurité

## 2.1 Module

Apache est un serveur modulaire dont seules les fonctions de base sont incluses dans le noyau du serveur. Les
fonctionnalités plus avancées sont disponibles grâce à des modules qui peuvent être chargés dans Apache. Par défaut, un
ensemble de modules de base est inclus à la compilation. Si le serveur est compilé pour utiliser des modules
dynamiquement chargés, alors des modules peuvent être compilés indépendamment et ajoutés (ou retirés) à n'importe quel
moment. Sinon, Apache doit être recompilé complètement pour ajouter ou supprimer des modules. Certains modules sont
livrés en standard, d'autre sont téléchargeables de manière séparée. Ces modules sont souvent programmés en C. Parmi les
modules apache les plus utiles, on peut citer :

* `mod_dir` et `mod_autoindex` : chargement automatique d'un fichier
  (index.html par exemple) et création automatique de la liste des fichiers d'un répertoire

* `mod_alias` : association de différentes parties du système de fichier de l'hôte dans l'arborescence des documents, et
  redirection des URL.

* `mod_userdir` : permet de gérer les répertoires personnels des utilisateurs

* `mod_php` : pour faire du php...

## 2.2 Organisation/Configuration

S'il est toujours possible de (re)compiler apache pour adapter sa configuration à son environnement, les distributions de linux proposent des versions d'apache intégrant le support des modules dynamiquement chargés. La configuration est ainsi finement adaptable sans compilation. Les éléments de configuration ont été regroupés dans le répertoire `/etc/apache2`. Dans ce répertoire, on peut noter :

* `apache2.conf` ; le fichier de configuration principal qui a été réparti en plusieurs fichiers. En particulier, c'est
  depuis ce fichier qu'est inclus le contenu d'autres répertoires (comme
  `mods-enabled` et `sites-enabled`)

* les répertoires `mods-available` et `mods-enabled` qui correspondent respectivement aux modules disponibles et
  activés. Les modules disponibles se trouvent dans le répertoire `mods-available`. Les modules actifs se trouvent dans
  le répertoire `mods-enabled`. Pour activer un module, il est possible de créer un lien symbolique depuis le
  répertoire `mods-enabled` OU d'utiliser la commande `a2enmod` (_apache2 enable module_). La commande de désactivation
  est `a2dismod` (_apache2 disable module_). Bien souvent, l'utilisation du module se fait en deux temps :

    1. chargement (fichier `mod_truc.load`):
   ```
   LoadModule truc_module /usr/lib/apache2/modules/mod_truc.so
  ```

    2. configuration (fichier `mod_truc.conf`):
  ```
  <IfModule mod_truc.c>

            directives_de_configuration

  </IfModule>  
  ```
    * les répertoires `sites-available` et `sites-enabled`
      qui contiennent les sites disponibles et activés. En particulier le site "par défaut" : `000-default.conf`. De
      même il existe les commandes `a2ensite` (_apache2 enable site_) et `a2dissite` (_apache2 disable site_).


Pour activer un module, il est possible de créer un lien symbolique depuis le
  répertoire `mods-enabled` OU d'utiliser la commande `a2enmod` (_apache2 enable module_). La commande de désactivation
  est `a2dismod` (_apache2 disable module_).


La vision schématique de cette organisation est la suivante :

![](site.svg "")

En résumé, la configuration globale d'apache est dans le fichier `apache.conf` (et des fichiers inclus depuis celui-ci), celle de vos sites se met naturellement dans un fichier **séparé** et peut contredire certains éléments hérités de la configuration globale.

## 2.3 Directives

Apache se configure en plaçant des directives dans des fichiers de configuration. Les modifications dans les fichiers de
configuration ne sont prises en compte qu'au (re)démarrage d'Apache.

Un fichier de configurations contient au plus une directive par ligne. Dans les fichiers de configuration, l'écriture
des noms des directives n'est pas sensible à la casse, mais les arguments des directives le sont généralement. Les
lignes blanches et les espaces précédant une directive sont ignorés. Les commentaires ne doivent pas être inclus sur la
même ligne qu'une directive. Les lignes commençant par le caractère dièse "#" sont traitées comme des commentaires et
sont ignorées.

La documentation d'Apache est accessible à l'URL suivante : [`https://httpd.apache.org/docs/2.4`](`https://httpd.apache.org/docs/2.4). Vous y trouverez en
particulier le descriptif des directives, décrites selon les items suivants :

* `Description` objectif de la directive

* `Syntax` syntaxe

* `Context` indique où la directive est licite. Les valeurs possibles sont :
    * `serveur config` : pour configurer **globalement** le serveur,
    * `virtual host` : pour configurer **un** site particulier
    * `directory` : spécifique d'une **partie** d'un site
    * `.htaccess` : dans un répertoire du site. Configuration locale à ce répertoire, si ceci est autorisé (voir
      `AllowOverride`)

## 2.3 Exemple : site minimaliste

```
# site répondant sur toutes les ips, port 80
<VirtualHost *:80>
  ServerAdmin webmaster@site
		
  # définition de la racine du site
  DocumentRoot /var/www
	
  # fichier à charger automatiquement
  DirectoryIndex index.html

  # valable sur /var/www (la racine du site)
  <Directory /var/www/>
      # listing activé
      Options Indexes FollowSymLinks
      AllowOverride None
      # accès autorisé 
      Require all granted
  </Directory>
  
  # configuration des logs
  # les variables ${...} sont définies dans /etc/apache2/envvars
  ErrorLog ${APACHE_LOG_DIR}/error.log
  LogLevel warn

  CustomLog ${APACHE_LOG_DIR}/access.log combined
</VirtualHost>
```

Il n'est absolument **pas** interdit de consulter la documentation
des [directives](https://httpd.apache.org/docs/2.4/mod/quickreference.html) utilisées

## 2.4 Démarrage/arrêt/rechargement
Comme la plupart des services linux, apache doit être (re)lancé explicitement, sinon les modifications des fichiers de configurations ne seront pas prises en compte.

### Vérification :
Avant tout, il est possible de vérifier la syntaxe des fichiers de configuration :

```
apachectl configtest
```

Vous devez obtenir `syntax ok`, sinon il faut corriger...

### Démarrage  :

```
systemctl start apache2
```

### Arrêt :

```
systemctl stop apache2
```

Deux fichiers vont se révéler instructifs en cas de problème : `error.log` et `access.log`.

```
# tail -f /var/log/apache2/error.log`
```

### Rechargement de la configuration

```
# systemctl reload apache2
```

En cas de problème, 
```
# systemctl restart apache2
# systemctl status apache2
```

# 3 Réalisations

La configuration que nous allons mettre en place est la suivante :

![](sites.svg "")

Par exemple, depuis un navigateur, `http://wwX.fai.com` devra nous donner le contenu du fichier `index.html` situé dans le répertoire `web3`

## 3.1 Premier site

### Site par défaut

La configuration de base d'apache place le site par défaut dans
un [`<VirtualHost>`](https://httpd.apache.org/docs/2.4/fr/mod/core.html#virtualhost). Ce site est donc un site parmi (
éventuellement) d'autres.  Voici quelques éléments de configuration :

* [`DocumentRoot`](https://httpd.apache.org/docs/2.4/fr/mod/core.html#documentroot) permet de spécifier le répertoire
  racine du serveur.
* [`ServerAdmin`](https://httpd.apache.org/docs/2.4/fr/mod/core.html#serveradmin) définit l'adresse e-mail que le
  serveur inclut dans tout message d'erreur retourné au client.

On souhaite que `http://wwwX.fai.com` nous donne le site par défaut (le nôtre).
* Désactivez le site initial
* Créez et éditez le fichier `web.conf` dans le répertoire `/etc/apache2/sites-available`
* Récupérez et décompactez le fichier `r3_06.tp1.tgz` dans le répertoire `/var/www/html` avec la commande suivante :
```
wget https://gitlab.univ-nantes.fr/iut.info2.r3_06/r3_06.tp1/-/raw/main/r3_06.tp1.tgz --output-document=- | sudo tar xz -C /var/www/html
```
* le répertoire `/var/www/html/web` inclut un fichier `index.html` qui contient la page d'accueil du site. En vous inspirant de l'exemple plus haut, définissez votre site :
    * `DocumentRoot` correspond à  `/var/www/html/web`.
      `ServerAdmin` par `webmasterX@fai.com`.
    * Ajoutez l'entrée `<Directory>` appropriée (celle qui permet d'accorder l'accès à la racine du site)
    * Définissez `ServerName` (`wwwX.fai.com`)
* Activez ce site
* Relancez Apache
* Votre page d'accueil doit s'afficher avec l'url suivante :
  `http://wwwX.fai.com`

### Fichier par défaut / Gestion des permissions

Comme vous le constatez, la page `index.html` est chargée automatiquement. En fait, une requête sur un répertoire induit les opérations suivantes (dans cet ordre) :
1. chercher dans le répertoire un des fichiers donnés par la
   directive [`DirectoryIndex`](https://httpd.apache.org/docs/2.4/fr/mod/mod_dir.html#directoryindex) (dans l'ordre
   indiqué) et le renvoyer
2. sinon si [`Indexes`](https://httpd.apache.org/docs/2.4/fr/mod/core.html#options) est active renvoyer une page listant
   les fichiers du répertoire
3. sinon renvoyer le code `403` (interdit)

Les directives placées dans le fichier `XXX.conf` définissant un site s'appliquent à l'ensemble de celui-ci. Pour modifier la configuration pour seulement une partie du site, il faut mettre les directives correspondantes dans des sections limitant leur portée, en particulier dans une directive [`<Directory>`](https://httpd.apache.org/docs/2.4/fr/mod/core.html#directory). Un sous répertoire hérite des directives de ses parents sauf si elles sont redéfinies par des directives propres à ce sous répertoire.
La directive [`Options`](https://httpd.apache.org/docs/2.4/fr/mod/core.html#options) quant à elle contrôle quelles fonctions du serveur sont disponibles :

* `FollowSymLinks` Le serveur est autorisé à suivre les liens symboliques dans ce répertoire.
* `Indexes` : autorise le listage du contenu d'un répertoire
* ...


Dans la configuration du site, avec les directives `<Directory>`, `DirectoryIndex` et `Options`, assurez les fonctionnalités suivantes : 
* faites en sorte que le listing du contenu du répertoire `listing` soit affiché (toujours). Testez.
* Pour le répertoire `bizarre`, aucun listing ne doit s'afficher et `default.html` est prioritaire sur `index.html` 
* Pas de listing, ni de contenu par défaut pour le répertoire `rien`

## 3.2 Serveurs virtuels

Dans la terminologie apache, un site est un hôte virtuel. C'est entre des balises `<VirtualHost>` et `</VirtualHost>` qu'on regroupe les directives qui ne s'appliqueront qu'à un site particulier du serveur. La distinction entre les sites peut se faire sur : nom d'hôte, ip ou port. Apache propose les trois méthodes pour configurer (différencier) les sites :

1. par nom (_Name-based Virtual Hosts_, le navigateur doit supporter le protocole HTTP 1.1). Les adresses IPv4 se
   raréfiant, l'intérêt de cette méthode est évident. Le `Host`
   spécifié dans les entêtes HTTP, permet de préciser le site voulu.
2. par ip (IP-based Virtual Hosts). La différentiation des sites se fait par l'ip.
3. par port. Idem mais sur les numéros de port.

Il est possible d'utiliser simultanément les trois méthodes.

Pour les trois questions suivantes, il est demandé que les journaux (fichiers de log) des serveurs virtuels soient
différents.

### Serveur virtuel par nom

Nous allons faire correspondre `http://wX.fai.com` à un deuxième site. C'est le cas le plus STANDARD.

* Ajoutez un site (virtual host) répondant sur `*:80`
* Le nom utilisé (`wX.fai.com`) doit correspondre au site dont la racine est `/var/www/html/web2`.
* Testez en tapant `http://wwwX.fai.com` et `http://wX.fai.com`.

### Serveur virtuel par adresse

Nous allons utiliser la deuxième adresse IP que vous avez configuré pour faire correspondre `http://wwX.fai.com` et
`http://192.168.X.2` à un troisième site dont la racine est `/var/www/html/web3`.

* Ajoutez un site répondant sur `192.168.X.2:80`, avec
  `wwX.fai.com` comme nom ou sans nom associé
* Testez ; soit avec `http://192.168.X.1` et `http://192.168.X.2`
  soit avec `http://wwwX.fai.com`. et `http://wwX.fai.com`.

### Serveur virtuel par port (subsidiaire = à faire à la fin)

On souhaite mettre en oeuvre un site de test accessible par :
`http://wwwX.fai.com:8080` avec comme racine `/var/www/html/web4`

* Modifiez la configuration d'Apache pour écouter sur les deux ports (doubler la directive `Listen` dans le fichier
  `ports.conf`)
* Ajoutez un virtual host répondant sur `192.168.X.1:8080`
* Vérifiez que vous accédez à deux sites différents:
  `http://wwwX.fai.com` et `http://wwwX.fai.com:8080`.

## 3.3 Répertoires virtuels

### Cas général

On souhaite monter un répertoire quelconque dans l'arborescence du site.

* Consultez la documentation de la directive [`alias`](https://httpd.apache.org/docs/2.4/fr/mod/mod_alias.html#alias)
* Modifiez la configuration du site  pour que
  `http://wwwX.fai.com/rep1/` corresponde au répertoire
  `/var/www/html/rep1`.
* `http://wwwX.fai.com/rep2/` correspond naturellement au répertoire `/var/www/html/web/rep2`.
* Testez : `http://wwwX.fai.com/rep1/` et
  `http://wwwX.fai.com/rep2/`

### Répertoires de publication des utilisateurs

Sous unix, il est possible d'accéder au home d'un utilisateur `lambda` avec l'alias `~lambda`. De même sous apache, il est
possible d'associer le répertoire de publication (`public_html` par défaut) d'un utilisateur `lambda` à l'alias `~lambda` à l'aide de la directive `UserDir`.

* Créez un répertoire `public_html` dans le home de l'utilisateur
  `tdreseau`
* Activez le module `userdir`.
* Vérifiez la configuration de ce module.
* Placez un fichier `index.html` dans le répertoire `/home/tdreseau/public_html`
* Testez : `http://wwwX.fai.com/~tdreseau/`.

## 3.4 Php

Mise en place de la configuration minimale pour coder en php.

* Installer php :
  ```
  # apt update --allow-releaseinfo-change
  # apt install php7.3
   ```   
* Vérifiez l'activation du module php dans le répertoire `mods-enabled`
* Redémarrez apache
* Consultez la page `http://wwwX.fai.com/php/phpinfo.php`
* Installez des extensions php
  ```
  # apt update --allow-releaseinfo-change
  # apt install php7.3-{sqlite3,zip,bz2,mbstring,curl}
  ```   
* Consultez la page `http://wwwX.fai.com/php/phpinfo.php`. La liste des modules a du s'étoffer.
* Consultez la page `http://wwwX.fai.com/php/erreur.php` puis les logs du serveur. Quel est le code d'erreur ?
* Optionnel, gérez le cas particulier des `Userdir` 

## 3.5 `.htaccess`

Les fichiers `.htaccess` sont une spécificité d'apache. Ils permettent de placer des directive de configuration du serveur dans le site web au lieu de les mettre dans le fichier de configuration central. Ceci permet de déléguer l'administration d'un site.

Les directives d'un fichier `.htaccess` s'appliquent au répertoire (et ses sous-répertoires) où se trouve le fichier. Il remplace une balise `directory` du fichier de configuration.

il est inutile de recharger ou redémarrer apache après un changement.

Attention, une erreur dans un (seul) `.htaccess` fait dysfonctionner (tout) le site (pas le serveur). De plus Apache doit relire les fichiers `.htaccess` à chaque accès ce qui induit une baisse notable de performance ;
L'utilisation des fichiers `.htaccess` repose sur plusieurs directives. La plus importante étant : `AllowOverRide` qui définit les familles de directives qu'il est permis de mettre dans un fichier `.htaccess`. 

Dans des `.htaccess`, avec les directives `<Directory>`, `DirectoryIndex` et `Options`, assurez les fonctionnalités suivantes pour `wX.fai.com` : 
*  faites en sorte que le listing du contenu du répertoire `listing` soit affiché (toujours). Testez.
* Pour le répertoire `bizarre`, aucun listing ne doit s'afficher et `default.html` est prioritaire sur `index.html` 
* Pas de listing, ni de contenu par défaut pour le répertoire `rien`

## 3.6 Messages d'erreur personnalisés
La directive `ErrorDocument` permet de personnaliser les erreurs. En particulier, vous pouvez utiliser une page locale ou 
avec une page distante. Toujours pour `wX.fai.com`

- créez une balise `directory` pour le sous-répertoire `erreurs` avec la même configuration que le répertoire `rien` (ce qui facilite l'obtention d'erreur `403`)
- Ajoutez une directive `ErrorDocument` qui va utiliser la page `403.html` pour les erreurs `403`
- Ajoutez une directive `ErrorDocument` qui va utiliser la page `https://httpd.apache.org/docs/2.4` pour les erreurs `404`

## 3.7 Réécriture d'URL (Subsidiaire)

Apache peut modifier les requêtes entrantes pour les faire correspondre à des urls internes. La correspondance se fait en fonction d'expressions rationnelles (dites aussi expressions régulières). Le module apache qui en permet la mise en oeuvre est [`mod_rewrite`](https://httpd.apache.org/docs/2.4/rewrite/intro.html). 

- Transformez des URLs du type : 
`http://wwwX.fai.com/bd/9782723477055`
en `http://wwwX.fai.com/test.php?cat=bd&isbn=9782723477055`
(avec cat pouvant être par exemple : roman, scolaire, bd, ...)
