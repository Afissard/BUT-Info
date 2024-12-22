#!/bin/bash
nslookup smtp128.main128.com 192.168.0.254
echo "ip donné par le serveur DNS :"
read ipGiven
ip link set up dev bleue
ip a a "$ipGiven" dev bleue

ip r a 172.20.0.0/24 via 192.168.0.128