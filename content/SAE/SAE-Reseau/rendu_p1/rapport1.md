# Configuration réseau

* NOMS : Sacha Chauvel, Rafael Café
* GROUPE de TP : 4-2
* X : 128
* SECRET : Regulus et  Fomalhaut
* IP_CLIENT : 172.20.128.1
* IP_DHCP : 172.20.128.2
* IP_SMTP : 192.168.0.28
* IPs du ROUTEUR : 172.20.128.3 et 192.168.0.128

**Note**: Le document suivant doit rendre compte de votre plan d’adressage (i.e. la description des différents LAN, de leur interconnexion, des machines avec les IP voire @MAC que vous jugerez pertinentes), de vos tables de routage de CLIENT, ROUTEUR, SMTP et celle (supposée) de DNS, des commandes à réaliser sur CLIENT, ROUTEUR, SMTP, et tout ce qui vous semble nécessaire à la configuration de votre réseau.

## Plan d'adressage
VLAN : 172.20.128.0/24
LAN : 192.168.0.0/24

## Tables de routage
### Client

| destination    | interface | gate away    |
| -------------- | --------- | ------------ |
| 172.20.128.2   | jaune.128 | -            |
| 192.168.0.0/24 | jaune.128 | 172.20.128.3 |
### DHCP

| destination    | interface | gate away    |
| -------------- | --------- | ------------ |
| 172.20.128.1   | jaune.128 | -            |
| 192.168.0.0/24 | jaune.128 | 172.20.128.3 |
### SMTP

| destination     | interface | gate away    |
| --------------- | --------- | ------------ |
| 192.168.0.254   | bleue     | -            |
| 172.20.128.0/24 | bleue     | 192.168.0.28 |
### Routeur

| **destination** | **iface** | **gw**        |
| --------------- | --------- | ------------- |
| 172.20.128.1    | jaune.128 | 172.20.128.3  |
| 172.20.128.2    | jaune.128 | 172.20.128.3  |
| 192.168.0.28    | bleue     | 192.168.0.128 |
| 192.168.0.254   | bleue     | 192.168.0.128 |
## Commandes de configuration
### Client
```bash
#!/bin/bash
modprobe 8021q
ip link add link jaune name jaune.128 type vlan id 128
ip link set up dev jaune.128
ip a a 172.20.128.1/24 dev jaune.128

ip r a 192.168.0.0/24 via 172.20.128.3
```
### DHCP
```bash
#!/bin/bash
modprobe 8021q
ip link add link jaune name jaune.128 type vlan id 128
ip link set up dev jaune.128
ip a a 172.20.128.2/24 dev jaune.128

ip r a 192.168.0.0/24 via 172.20.128.3
```
### Routeur
```bash
#!/bin/bash
modprobe 8021q
ip link add link jaune name jaune.128 type vlan id 128
ip link set up dev jaune.128
ip a a 172.20.128.3/24 dev jaune.128

sysctl net.ipv4.ip_forward=1
```
### SMTP
```bash
#!/bin/bash
nslookup smtp128.main128.com 192.168.0.254
echo "ip donné par le serveur DNS :"
read ipGiven
ip link set up dev bleue
ip a a "$ipGiven" dev bleue

ip r a 172.20.0.0/24 via 192.168.0.128
```
# log
```
chmod 700 log
./log > nomMachine.log
```