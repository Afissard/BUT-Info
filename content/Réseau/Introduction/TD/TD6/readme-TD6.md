# Tp Services Réseaux : SMTP/POP/IMAP

ce dernier TP va vous faire installer et configurer sommairement un serveur de mail.

## Préambule

* Votre machine X (allant de 1 à 35) a pour ip : `192.168.0.X/24` sur `jaune`. Vérifiez le
* Votre nom de domaine pour ce tp sera `mailX.com`
* La machine `192.168.0.254` a été configurée comme serveur de noms pour les zone `mailX.com`. Pour utiliser ce serveur DNS, vérifier que votre fichier `/etc/resolv.conf` contienne : `nameserver 192.168.0.254`
* Tester (et valider) la résolution des trois noms :

```
# host smtpX.mailX.com
# host popX.mailX.com
# host imapX.mailX.com
```
* Chaque utilisateur du système peut recevoir des mails. Vous allez donc également créer deux autres utilisateurs `foo` et `bar` à l'aide de la commande `adduser`. Ce seront les expéditeurs et destinataire des mails.

## SMTP
Le protocole SMTP (Simple Mail Transfert Protocol) est un protocole d'envoi de mail. 	Le serveur SMTP que nous mettrons en oeuvre lors de ce tp est Postfix.

### Installation
Nous procédons à une installation via le gestionnaire de paquet de debian. Après installation, le service ne se lance pas spontanément. Il faudra donc le démarrer explicitement, soit via `systemctl`, soit via `service`. En outre, lors de l'installation, les options de configuration restent plutôt basiques. `dpkg-reconfigure postfix` permet de revenir sur certaines options. En complément, nous installons également `mailutils`.

```
# apt-get update --allow-releaseinfo-change
# apt install postfix mailutils
# dpkg-reconfigure postfix

```

Configurez `postfix` :

 * Configuration type du serveur de messagerie : `Site Internet`
 * Nom de courrier : `mailX.com`
 * liste des domaines : `mailX.com localhost.localdomain, localhost`
 * Réseaux internes : ajouter `192.168.0.0/24`
 * Protocoles internet à utiliser : `ipv4`
 
Laissez les autres options à leurs valeurs par défaut. Démarrez et vérifiez le bon lancement de postfix.


### Envoyer un mail

Nous commençons, par vérifier le fonctionnement de postfix en envoyant un mail en utilisant directement le protocole SMTP.
Voici un exemple dont vous pouvez vous inspirer : 

```
$ nc 192.168.0.X smtp
ehlo 192.168.0.X
250-localhost.mailX.com
250-PIPELINING
250-SIZE 10240000
250-VRFY
250-ETRN
250-STARTTLS
250-ENHANCEDSTATUSCODES
250-8BITMIME
250-DSN
250-SMTPUTF8
250 CHUNKING
220 localhost ESMTP Postfix (Debian/GNU)
mail from: darth@starwars.univ
250 2.1.0 Ok
rcpt to: luke@starwars.univ
250 2.1.5 Ok
data
354 End data with <CR><LF>.<CR><LF>
Je suis ton père
.
250 2.0.0 Ok: queued as C8AC034E16
quit
221 2.0.0 Bye
```

### Contrôler sa réception

Vous devriez avoir un mail au nom de l'utilisateur destinataire du mail dans le répertoire `/var/mail`. Vérifiez son contenu. 

### Premier pas vers la convivialité
Utilisez la commande mail pour envoyer et recevoir un mail.

### Thunderbird
Configurez thunderbird pour envoyer des mails. Dès le premier écran, passez en configuration manuelle. Les machines à utiliser sont : `smtpX.mailX.com`, `popX.mailX.com` (le reste en auto-détection). Attention, vous devez configurer le proxy (`Preferences --> General --> Network & Disk Space --> Connection --> no proxy`). Testez.

## Serveur de réception de courrier
Il existe deux principaux protocoles pour la réception de courrier : POP et IMAP. POP a de moins en moins d'intérêt pratique mais sa simplicité permet de comprendre le principe de ces protocole. Nous allons utiliser `dovecot` qui fournit une implantation des deux protocoles.

### Installation
```
# apt install dovecot-pop3d dovecot-imapd
```

Ensuite, nous allons dégrader la sécurité en ajoutant notre réseau dans la liste des réseaux autorisés SANS authentification préalable. Il faut modifier le fichier `/etc/dovecot/dovecot.conf` (normalement cela sert plutôt à des serveurs qu'à des clients).

```
login_trusted_networks = 192.168.0.0/24
```
Avantage, ceci sera valable pour nos deux protocoles.

### POP
Après vous être assuré d'avoir démarré le service et d'avoir un mail à consulter, vous pouvez vous inspirer de ce dialogue pour lire un mail : 

```
nc 127.0.0.1 pop3
+OK Dovecot (Debian) ready.
user foo
+OK
pass foo
+OK Logged in.
list
+OK 4 messages:
1 463
.
retr 1
+OK 463 octets
Return-Path: <foo@mailX.com>
X-Original-To: foo@mailX.com
Delivered-To: foo@mailX.com
Received: by localhost (Postfix, from userid 1001)
	id 5AD42334E6; Thu,  5 May 2022 14:19:30 +0000 (UTC)
To: <foo@mailX.com>
Subject: Mail 1
X-Mailer: mail (GNU Mailutils 3.5)
Message-Id: <20220505141954.5AD42334E6@localhost>
Date: Thu,  5 May 2022 14:19:30 +0000 (UTC)
From: foo <foo@mailX.com>

Corps du mail. Avec aççénts
.
quit
+OK Logging out.
```

Puis configurez thunderbird pour la réception des mails. Attention, l'identifiant de l'utilisateur est sans `@mailX.com`.

### IMAP

Pour IMAP, le protocole commence à être trop complexe... Donc vous allez simplement configurer thunderbird pour IMAP et capturer les échanges avec wireshark.