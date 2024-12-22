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