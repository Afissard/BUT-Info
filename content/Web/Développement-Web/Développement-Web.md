---
title: "Développement-Web"
draft: 
description: 
tags:
---
# Développement-Web
- todo
## Connections au podman

```bash
# à faire à la main potentiement
podman unshare
cd $(podman container mount web)
cd ./var/www/html/
code . --no-sandbox --user-data-dir /home/sacha/Documents/webdev-php/
```
## Connection VM
```bash
ssh debian@172.21.44.93
# http://172.21.44.93/
```
[websiteVM](http://172.21.44.93/)

*dpkg est cassé :'(*
# FF j'ai setup Apache
la doc : 
- [doc fedora fr](https://doc.fedora-fr.org/wiki/Installation_et_configuration_d%27Apache)
- [doc fedora](https://docs.fedoraproject.org/en-US/quick-docs/getting-started-with-apache-http-server/)
- M.Berdjugin
```bash
sudo systemctl start httpd.service
sudo systemctl restart httpd.service
sudo systemctl stop httpd.service
```
site ici : http://localhost/webdev/ ou juste localhost
