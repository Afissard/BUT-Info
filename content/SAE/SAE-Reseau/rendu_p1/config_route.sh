#!/bin/bash
echo "SETUP LAN : ROUTEUR"

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
echo "ajout des route VLAN"
ip r a "172.20.$idChosen.1" dev "$colChosenVLAN.$idChosen"
ip r a "172.20.$idChosen.2" dev "$colChosenVLAN.$idChosen"


echo "ajout des routes LAN (configuration du SMTP requise)"
# DNS
ip r a "192.168.0.254" dev "$colChosenLAN"
# SMTP

echo "ip SMTP :"
read ipSMTP
ip r a "$ipSMTP/24" dev "$colChosenLAN"

echo "done"
