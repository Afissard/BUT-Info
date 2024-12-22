#!/bin/bash
modprobe 8021q
ip link add link jaune name jaune.128 type vlan id 128
ip link set up dev jaune.128
ip a a 172.20.128.3/24 dev jaune.128

sysctl net.ipv4.ip_forward=1
sysctl -p