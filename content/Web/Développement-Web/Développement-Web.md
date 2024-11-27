---
title: "Développement-Web"
draft: 
description: 
tags:
---
- todo
# Connections au podman

```bash
# à faire à la main potentiement
podman unshare
cd $(podman container mount web)
cd ./var/www/html/
code . --no-sandbox --user-data-dir /home/sacha/Documents/webdev-php/
```
