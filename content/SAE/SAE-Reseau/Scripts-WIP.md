# Setup VLAN (Client et DHCP)
```bash
#!/bin/bash
echo "SETUP VLAN : Client et DHCP"

# SETUP
echo "id machine :"
read idChosen
echo "ip machine :"
read ipChosen
echo "couleur interface :"
read colChosen
echo "
	Config :
	- id : $idChosen
	- ip : $ipChosen
	- color : $colChosen
	"

# COMMANDS
# activation prise en compte VLAN
modprobe 8021q

# lien VLAN
ip link add link "$colChosen" name "$colChosen.$idChosen" type vlan id "$idChosen"

# activation interface
ip link set up dev "$colChosen.$idChosen"

# ajout d'une ip à l'interface
ip a a "172.20.$idChosen.$ipChosen/24" dev "$colChosen.$idChosen"

echo "done"
```
# Setup LAN
## SMTP
```bash
#!/bin/bash
echo "SETUP LAN : SMTP"

# SETUP
echo "id machine :"
read idChosen
echo "couleur interface :"
read colChosen
echo "
	Config :
	- id : $idChosen
	- color : $colChosen
	"
	
# Demande au serveur DNS une ip (fourni de manière dynamique)
echo "exécuter la commande suivante (dans un autre terminal) :
	nslookup smtp$idChosen.main$idChosen.com 192.168.0.254
	"
# récupération de l'ip donnée à la fin de la commande
echo "ip donné par le serveur DNS :"
read ipGiven
echo "ip : $ipGiven"

# COMMANDS
# ajout d'une ip à l'interface
ip a a "$ipGiven" dev "$colChosen"

echo "done"
```
# Config routes
## Routeur
```bash
#!/bin/bash
echo "SETUP VLAN/LAN : ROUTEUR"

# SETUP
echo "id machine :"
read idChosen
echo "ip machine :"
read ipChosen
echo "couleur interface VLAN:"
read colChosenVLAN
echo "couleur interface LAN:"
read colChosenLAN
echo "
	Config :
	- id : $idChosen
	- ip : $ipChosen
	- color VLAN : $colChosenVLAN
	- color LAN : $colChosenLAN
	"

# COMMANDS
echo "setup interface"
# activation interface
ip link set up dev "$colChosen"
# ajout d'une ip à l'interface
ip a a "192.168.$idChosen.$ipChosen/24" dev "$colChosen"

# ROUTES
echo "forwarding"
sysctl net.ipv4.ip_forward=1

echo "ajout des route VLAN"
ip r a "172.20.$idChosen.1" dev "$colChosenVLAN.$idChosen"
ip r a "172.20.$idChosen.2" dev "$colChosenVLAN.$idChosen"

echo "ajout des routes LAN (configuration du SMTP requise)"
# DNS (optionnel ?)
#ip r a "192.168.0.254" dev "$colChosenLAN"
# SMTP
echo "ip SMTP :"
read ipSMTP
ip r a "$ipSMTP" dev "$colChosenLAN"

echo "done" 
```
## Routes via (VLAN & LAN)
```bash
#!/bin/bash

# SETUP
echo "id machine :"
read idChosen
echo "ip machine routeur :"
read ipChosen
echo "
	Config :
	- id : $idChosen
	- ip : $ipChosen
	"

# COMMANDS
ip r a 192.168.0.0/24 via "172.20.$idChosen.$ipChosen"
ip r a 172.20.0.0/24 via "192.168.0.$idChosen"

echo "done"
```
# Nota Bene
## Suppression d'une ip
```bash
# supression d'une ip
ip a d "$ipReseau/$ipMask$" dev "$colChosen.$idChosen"
```
## Suppression d'une route
```bash
ip r d "$ipReseau/$mask" dev "$ipDevice"
ip r d "$ipReseau/$mask" via "$ipGateAway"
```
## Supprimer la réinitialisation automatique des interfaces réseaux (mais n'a pas prouvé son utilité)
```bash
#!/bin/bash
echo "Supprimer la réinitialisation automatique des interfaces réseaux"

# supression et création
rm /etc/network/interfaces && touch /etc/network/interfaces
echo "
auto lo  
iface lo inet loopback  
  
allow-hotplug bleue  
iface bleue inet manual  
  
allow-hotplug jaune  
iface jaune inet manual  
  
allow-hotplug rouge  
iface rouge inet dhcp
" > /etc/network/interfaces # écriture

service networking restart
service network-manager restart

echo "done"
```