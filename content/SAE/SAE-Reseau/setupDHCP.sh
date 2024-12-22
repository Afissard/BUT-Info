#!/bin/bash
echo "Setup DHCP"

# VLAN
modprobe 8021q
ip link add link jaune name jaune.128 type vlan id 128
ip link set up dev jaune.128
ip a a 172.20.128.254/24 dev jaune.128

ip r a 192.168.0.0/24 via 172.20.128.3

# Serveur DHCP
# /etc/default/isc-dhcp-server
rm -rf /etc/default/isc-dhcp-server && touch /etc/default/isc-dhcp-server && echo "INTERFACESv4=jaune.128" > /etc/default/isc-dhcp-server

# /etc/dhcp/dhcpd.conf
rm -rf /etc/dhcp/dhcpd.conf && touch /etc/dhcp/dhcpd.conf && echo '
option domain-name "REGULUSFORMALHAULT.org";
option domain-name-servers 192.168.0.254;
option routers 172.20.128.254;
option subnet-mask 255.255.255.0;

subnet 172.20.128.0 netmask 255.255.255.0 { 
  range 172.20.128.2 172.20.128.250;  # Plage d adresses IP DHCP
  option routers 172.20.128.254;
  option subnet-mask 255.255.255.0;
}

# Adresse fixe
host appareil1 {
  hardware ethernet 04:8d:38:c2:57:cd;
  fixed-address 172.20.128.1;
}

' > /etc/dhcp/dhcpd.conf

systemctl start isc-dhcp-server
systemctl status isc-dhcp-server

echo "done"
