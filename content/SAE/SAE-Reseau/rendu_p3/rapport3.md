# Analyse session SMTP

* NOMS : Sacha Chauvel, Rafael Café
* GROUPE de TP : 4-2
* X : 128
* SECRET : Regulus et  Fomalhaut
* IP_CLIENT : 172.20.128.1
* IP_DHCP : 172.20.128.2
* IP_SMTP : 192.168.0.28
* IPs du ROUTEUR : 172.20.128.3 et 192.168.0.128 

# Opérations réalisées : 
- Sur la machine SMTP : 
	Setup du serveur SMTP (après la remise en configuration de la partie 1) :
	```bash
	 apt-get update --allow-releaseinfo-change
	 apt install postfix mailutils
	 dpkg-reconfigure postfix
	```
	Configuration de Postfix sur la machine SMTP :
	- Configuration type du serveur de messagerie : `Site Internet`
	- Nom de courrier : `mailX.com`
	- liste des domaines : `mailX.com localhost.localdomain, localhost`
	- Réseaux internes : ajouter `192.168.0.0/24`
	- Protocoles internet à utiliser : `ipv4`
	On ajoute le user smtp : `adduser smtp
	Lancement de WireShark pour la capture de trame sur la machine SMTP : beaucoup de chose se passent sur la liaison bleue, mais seule une partie nous intéresse. Pour l'obtenir, nous avons établi le filtre `smtp and ip.addr == 172.20.128.1`.
	Envoie du mail depuis la machine CLIENT:
	- `nc 192.168.0.X smtp`
	- `ehlo 192.168.0.28
	- `mail from: client@mail128.com`
	- `rcpt to: smtp@mail128.com`
	- `data
	- *Contenu du mail demandé*
	- ENTER, '.', ENTER
	- `quit
	Puis nous vérifions le bon reçu du mail sur la machine SMTP : 
	- `cd /var/mail
	- `cat smtp
	Le mail est là !
```smtp
From client@mail128.com  Mon Jun 10 14:43:51 2024
Return-Path: <client@mail128.com>
X-Original-To: smtp@mail128.com
Delivered-To: smtp@mail128.com
Received: from unknown (unknown [172.20.128.1])
	by localhost (Postfix) with SMTP id 2A435310C1
	for <smtp@mail128.com>; Mon, 10 Jun 2024 14:39:57 +0000 (UTC)

CAFÉ, CHAUVEL, formalhaultregulus, 128
```

# Analyse de la trame : 
Après l'application du filtre cité précédemment, on peut tout de même observer une trentaine de lignes détaillant les interactions entre le CLIENT et le SMTP lors de l'envoie du mail. Tout d'abord le `ehlo` pour confirmer le bon fonctionnement de la communication (ligne no 395). Puis, l'établissement de l'envoyeur (ligne no 500). Ensuite nous pouvons observer de la ligne 574 à 1544, deux tentatives d'établissement de receveur échouées, car à ce moment-là, le `adduser` sur la machine SMTP n'avait  pas été encore effectué. Nous avons décidé de garder cette trame contenant l'erreur car nous la jugions intéressante. Enfin,nous est présentée la ligne 2721, l'envoie du contenue du mail (40 bytes) et la ligne suivante, le '.' signant la fin du mail (2 bytes). La fin de l'échange est conclu par le 'Bye'.
