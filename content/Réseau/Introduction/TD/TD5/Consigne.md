On passe au TP VLAN+DHCP quel que soit votre avancement sur le routage. On y reviendra si on a le temps ou durant la SAE.  

# Topo dhcp+vlan

* rappeler à quoi sert un serveur dhcp, qu'il opère dans un domaine de diffusion et qu'il fournit à minima IP et masque mais aussi IP routeur et DNS

* présenter ce que sont les VLAN. Personnellement j'ai parlé de fonctionnement de Hub et de Switch, de construction de table de commutation, des motivations pour des VLAN, les types de VLAN et notamment les VLAN tagué

  

# Test de 30 minutes

* rendre visible le test le temps du test puis le cacher à nouveau !!!

* il n'y a pas besoin de modifier le temps. Eventuellement ajouter 10 minutes pour les étudiants ayant le droit à un tiers temps

* 10 min par exercice environ

* les étudiants peuvent constater les réponses en erreur en survolant les champs de leur réponse après les avoir envoyés  

# Le TP : mise en place d'un VLAN et la configuration client et serveur DHCP  

* la partie observation d'une capture DHCP en tant que client ne marche pas systématiquement... I dont know why :-(

* les accompagner fortement pour que le TP soit terminé à la fin des 2 séances.

Le sujet est sur [https://gitlab.univ-nantes.fr/iut.info1.r2_05/r2_05.tp2](https://gitlab.univ-nantes.fr/iut.info1.r2_05/r2_05.tp2) (requiert la configuration d'un serveur DHCP sur la machine enseignante pour la partie configuration client).

Ci-dessous quelques éléments pour faire le TP  

# DHCP client  

La machine enseignant _tdreseau_ doit être allumée et le serveur DHCP (à savoir _isc-dhcp-server_) doit tourner (avec la bonne configuration). 

Depuis les machines étudiantes, il faut lancer un wireshark sur l'interface bleue que vous voulez configurer avec par exemple le filtre 

```
eth.addr == VOTRE_ADDR_MAC_SUR_BLEUE and boot
```  

bootp est une autre manière de désignée le protocole de démarrage DHCP.

Vous pourrez ensuite observer les échanges du protocoles après avoir demandé une adresse avec la commande suivante  
  

    dhclient -4 bleue

L'option -r permet de libérer le bail en cours et arrête le client DHCP en cours d'exécution.  

Par défaut, vous verrez que vos machines reçoivent des IP en 192.168... Si ce n'est pas le cas, c'est qu'il y a un problème. Vous pouvez supprimer ces IP si elles vous gênent.  

  
Vous devriez observer l'échange des paquets _discover, offer, request, et ack,_ suivants avec des ip et mac particulières :  
  

- discover : 1 client demande 1 IP (mac_dst=broadcast, ip_src=0.0.0.0 ip_dst=255.255.255.255)
- offer : 1 serveur fait une offre d'IP (mac_dst du client et ip_src du serveur ip_dst proposée) avec un id de transaction
- request : le client demande une IP offerte avec un id de transaction (mac_dst=broadcast, ip_src=0.0.0.0 ip_dst=255.255.255.255) ; tous les serveurs recoivent (y compris ceux dont l'offre n'a pas été acceptée)
- ack : le serveur confirme l'attribution au client (mac_dst du client et ip_src du serveur ip_dst proposée) avec un id de transaction

A toutes fins utiles, le support CM [hernandez-R205-reseaux1-CM_10_applications-dhcp-mail.pdf](https://madoc.univ-nantes.fr/pluginfile.php/4503598/mod_folder/content/0/hernandez-R205-reseaux1-CM_10_applications-dhcp-mail.pdf?forcedownload=1) présente le fonctionnement du protocole DCHP et fournit des éléments de configuration requis en fin de TP.

# VLAN  

RAPPELER qu'un VLAN sert à créer un LAN spécifique entre des machines ; on peut configurer un VLAN au niveau de Switch. Ici nous en configurons un au niveau de machines. Il s'agit de déclarer sur chaque machine que l'on veut regrouper en VLAN une liaison de type vlan avec un alias d'interface dédié et un identifiant de VLAN identique pour toutes les machines du VLAN.  

Le support CM hernandez-R205-reseaux1-CM_02_equipements.pdf présente ce que sont les VLAN et notamment les VLAN "taggués" mis en place dans ce TP.

  
Sur ces machines, il faut  
1. charger le module  
    modprobe 8021q

2. ajouter un lien de type vlan. Pour cela déclarer le même vlan id (e.g. 9) sur toutes les machines impliquées et indiquer l'interface sur laquelle ce lien est créée (e.g. jaune) et le nom pour désigner l'interface de ce nouveau lien (e.g. jaune.9).  
  
   ip link add link jaune name jaune.9 type vlan id 9

3. ajouter une IP à l'interface  
  

    ``ip a a 10.0.9.1/24 dev jaune.9 # sur une machine``  
    ``ip a a 10.0.9.2/24 dev jaune.9 # sur une autre machine avec l'interface vlan``  
    ``ip a a 10.0.9.3/24 dev jaune # sur une tierce machine sans l'interface vlan mais tout de même sur le même réseau que les deux précédentes``

4. activer l'interface  

    ip link set up dev jaune.9

5. lancer une communication d'une machine vers une autre (via un ping) # par exemple depuis 10.0.9.1  
    ping 10.0.9.2

6. observer la trame taguée avec wireshark. Cela se traduit par une couche intermédiaire entre la trame et le paquet IP qui porte le protocole ICMP.  

**# Contrôle d'un service**  

Prenons ici comme exemple le service DHCP connu sous le nom de _isc-dhcp-server_.  

Pour connaître l'état d'un service, faire

systemctl status isc-dhcp-server

Pour le démarrer, faire

systemctl status isc-dhcp-server

Pour l'arrêter, faire

systemctl stop isc-dhcp-server

Pour consulter le journal d'événements (log) relatifs au service, tapez la commande suivante et scrollez tout en bas avec la barre d'espace et remonter à rebours pour connaître les événements les plus récents.  

journalctl -u isc-dhcp-server

Une démarche pour debugger un service est la suivante  

1. consulter l'état du service : constater le "fail" et l"inactive" du service  
    
2. vous pouvez tenter de le démarrer pour voir si vous obtenez des informations vous permettant de debugger
3. consulter le journal pour identifier le problème  
    
4. corriger le problème de configuration
5. démarrer à nouveau et vérifier que l'état du service est désormais "active" sans problème
