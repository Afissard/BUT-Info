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