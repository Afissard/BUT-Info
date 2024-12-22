#!/bin/bash
echo "Setup client"

# VLAN
modprobe 8021q
ip link add link jaune name jaune.128 type vlan id 128
ip link set up dev jaune.128
# ip a a 172.20.128.1/24 dev jaune.128 # l'ip seras donnée par le serveur DHCP

ip r a 192.168.0.0/24 via 172.20.128.3

# Une fois que le serveur DHCP est configuré:
dhclient -r
dhclient
