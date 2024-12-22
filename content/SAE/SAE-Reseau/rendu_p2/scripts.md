# Client
```bash
#!/bin/bash
echo "Setup client"

# VLAN
modprobe 8021q
ip link add link jaune name jaune.128 type vlan id 128
ip link set up dev jaune.128
# ip a a 172.20.128.1/24 dev jaune.128 # l'ip seras donnée par le serveur DHCP

# Une fois que le serveur DHCP est configuré:
echo "
tester depuis la station cliente, en utilisant dhclient, avec l'interface jaune.128
"
```
# DHCP
```bash
#!/bin/bash
echo "Setup client"

# VLAN
modprobe 8021q
ip link add link jaune name jaune.128 type vlan id 128
ip link set up dev jaune.128
ip a a 172.20.128.254/24 dev jaune.128

# Serveur DHCP
# /etc/default/isc-dhcp-serveur
rm -rf /etc/default/isc-dhcp-serveur && touch /etc/default/isc-dhcp-serveur && echo "INTERFACESv4=jaune.128" > /etc/default/isc-dhcp-serveur

# /etc/dhcp/dhcpd.conf
# configurer le serveur DHCP dans le fichier `/etc/dhcp/dhcpd.conf` pour délivrer des IPs sur la plage `10.0.128.10-100/24` en fournissant route par défaut (via `10.0.128.254`), nom de domaine (`fai.com`), serveur DNS (`192.168.0.254`).

echo '
option domain-name "fai.com";
option domain-name-servers 192.168.0.254;
option routers 10.0.128.254;
option subnet-mask 255.255.255.0;
subnet 10.0.128.0 netmask 255.255.255.0 { 
	range 10.0.128.10 10.0.128.100;
}
' > /etc/dhcp/dhcpd.conf

echo "done"
```
## DHCP utilities
Pour connaître l'état d'un service : ``systemctl status isc-dhcp-server``
Pour le démarrer : ``systemctl status isc-dhcp-server``
Pour l'arrêter : ``systemctl stop isc-dhcp-server``

Pour consulter le journal d'événements (log) relatifs au service, tapez la commande suivante : ``journalctl -u isc-dhcp-server`` et scrollez tout en bas avec la barre d'espace et remonter à rebours pour connaître les événements les plus récents.  

Une démarche pour debugger un service est la suivante  
1. consulter l'état du service : constater le "``fail``" et l"``inactive``" du service  
2. vous pouvez tenter de le démarrer pour voir si vous obtenez des informations vous permettant de debugger
3. consulter le journal pour identifier le problème  
4. corriger le problème de configuration
5. démarrer à nouveau et vérifier que l'état du service est désormais "active" sans problème
