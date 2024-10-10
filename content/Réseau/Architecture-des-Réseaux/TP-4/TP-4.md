---
title: TP-4
draft: 
description: 
tags:
  - Réseau
---
# Td Architecture des réseaux : NAT

NAT (Network Address Translation) permet de faire correspondre des adresses IP privées (RFC 1918) d'un intranet à des adresses externes publiques et donc routables sur internet. La mise en oeuvre sur linux passe par la table nat des iptables avec plusieurs déclinaisons du mécanisme de réécriture des adresses IP et des ports dans les paquets.
Pour rappel, les plages d'adresses IPv4 privées définies dans la RFC 1918 sont les suivantes : 
* 10/8
* 172.16/12
* 192.168/16


![Situation](fig.svg "")

Dans cette situation, figurent trois hôtes : E (Externe), I (Interne) et P (Passerelle). Le réseau `192.168.X.0/24` est **considéré** comme étant le réseau public. 

## Routage "classique"
Mettez en place la situation :
* la passerelle `P` doit autoriser le forward de paquets (à vérifier, au minimum)
* les deux stations `I` et `E` définissent une route par défaut vers `P`
* testez en produisant un échange depuis `I` vers `E` (un ping par exemple).
* capturez les échanges depuis `P`


## Masquerading simple (SNAT)

Nous allons maintenant faire les mêmes manipulations en dissimulant les adresses internes. Pour cela, nous allons utiliser la cible
`SNAT` ou `MASQUERADE` qui permet la traduction d'adresse réseau source.

* la passerelle `P` doit toujours autoriser le forward de paquets
* activez la traduction d'adresse sur `P` (la passerelle nat) :

```
iptables -t nat -A POSTROUTING -o eth1 -j MASQUERADE
```

* testez en produisant un échange depuis `I` vers `E` (un ping par exemple).
* capturez les échanges avec wireshark et observez la réécriture des adresses ip

## Miroir de port ou port forwarding (DNAT)

On souhaite rendre accessible un service localisé sur une machine de l'intranet `I` comme si le service était hébergé sur la passerelle. Nous utilisons toujours la table nat, mais avec la cible `DNAT`.

* installez apache et personnalisez la page d'accueil du serveur web de la machine `I`.
* testez depuis `E` :

```
wget --no-proxy http://192.168.X.(100+X)
```

* redirigez le trafic entrant sur le port 80 de la passerelle NAT `P` vers la machine cliente `I`

```
iptables -t nat -A PREROUTING -p tcp --dport 80 -j DNAT --to 10.0.X.X
```

* testez

## port mapping (DNAT)

* Faites de même pour du port mapping en faisant correspondre le port 8080 de la passerelle au serveur web interne.

* tester depuis `E` :

```
wget --no-proxy http://192.168.X.(100+X):8080
```

## “Load balancing”

* ajoutez une machine `J` en lui affectant une ip `10.0.X.Y/24`.
* Installez et personnalisez la page d'accueil du serveur web de cette machine.
* en utilisant le module *statistic* de iptables (`-m statistic`), redirigez alternativement le trafic sur `10.0.X.X/24` et `10.0.X.Y/24` (utilisation de `--mode nth` ou `--mode random`)
* testez depuis `E` :
```
wget --no-proxy http://192.168.X.(100+X)
```