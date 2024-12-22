# 1 préliminaire
## Qu'est-ce qu'une table de partition ?
Une partition de disque est le "découpage" du disque où sont répartie les différentes sections nécessaires à l'OS (fichier système (sys), root, utilisateur (usr), etc)
(Wikipedia : https://fr.wikipedia.org/wiki/Partition_(informatique))
En informatique, une **partition**, **région** ou un **disque** est une section d'un support de stockage (disque dur, SSD, carte-mémoire ...). Le partitionnement est l'opération qui consiste à diviser ce support en partitions dans lesquelles le système d'exploitation peut gérer les informations de manière séparée, généralement en y créant un système de fichiers, une manière d’organiser l’espace disponible.
## Quelles sont les différences entre des tables de partitions GPT et MBR ?
(Microsoft : https://learn.microsoft.com/fr-fr/windows-server/storage/disk-management/change-a-gpt-disk-into-an-mbr-disk)
Les disques d'enregistrement de démarrage principal (**MBR**) utilisent la table de partition BIOS standard. Les disques de table de partition GUID (**GPT**) utilisent l'interface UEFI (Unified Extensible Firmware Interface). Les disques **MBR** ne prennent pas en charge plus de quatre partitions par disque.
# 2 Installation Linux
## Qu’est ce qu’un système de fichiers ? 
(wikipedia : https://fr.wikipedia.org/wiki/Syst%C3%A8me_de_fichiers)
Soit l'organisation hiérarchique des fichiers au sein d'un système d'exploitation (on parle par exemple du file system d'une machine unix organisé à partir de sa racine (**/**))
Soit l'organisation des fichiers au sein d'un volume physique ou logique, qui peut être de différents types (par exemple NTFS, FAT, FAT32, ext2fs, ext3fs, ext4fs, zfs, btrfs, etc), et qui a également une racine mais peut en avoir plusieurs
## Qu’est ce qu’un système de fichiers journalisé ? 
(wikipedia : https://fr.wikipedia.org/wiki/Syst%C3%A8me_de_fichiers_journalis%C3%A9)
Le système de fichiers journalisés est un système de fichiers tolérant/résistant aux pannes qui permet d'assurer l'intégrité des données en cas de problème matériel, de panne de courant (ou Hot-swap / débranchement à chaud) ou d'arrêt brutal du système. Cette fonctionnalité est assurée par la tenue d'un Journal (système de fichiers) référençant les opérations d'écriture sur le support physique avant que ce dernier ne soit réellement mis à jour. Le système de fichiers doit permettre une reprise d'activité à la suite d'une coupure brutale, telle un arrêt électrique. Les métadonnées doivent alors rester cohérentes et à jour. La journalisation permet d'optimiser le contrôle d'intégrité du système de fichiers, réduisant ainsi le temps de redémarrage du système, critère important dans les environnements qui ont besoin d'une haute disponibilité.
## Quels points de montage avez vous choisi ? 
La racine : "/"
Données personnelles : "/home"
## À quoi sert le swap ?
Définition du prof (retransmit par Rafael) : Le Swap est un espace de stockage réservé
(redhat : https://access.redhat.com/documentation/fr-fr/red_hat_enterprise_linux/7/html/storage_administration_guide/ch-swapspace)
L'espace swap sur Linux est utilisé lorsque la mémoire physique (RAM) est pleine. Si le système a besoin de plus de ressources mémoire et que la mémoire RAM est pleine, les pages mémoire inactives sont alors déplacées vers l'espace swap.
# 3 Utilisateurs et groupes
1. La commande ```sudo su -``` met le répertoire courant à la racine (en super utilisateur et non l'utilisateur sae). La commande ```sudo su ``` met le répertoire courant dans le dossier personnel (en super utilisateur).
2. Pour visualisé les fichiers dans le `/etc` :
	- `cd /etc && cat passwd`
	- `cd /etc && cat group`
	- `cd /etc && sudo cat shadow`
3. Le groupe de l'utilisateur "sae" : `groups sae` retourne :
	- sae
	- adm
	- cdrom
	- sudo
	- dip
	- plugdev
	- lpadmin
	- lxd
	- sambashare
1. Création des groupes : `sudo groupadd grp1`
2. Création d'utilisateur : `sudo adduser --ingroup grp1 user1`
3. Ajout d'utilisateur à un groupe : `sudo usermod -a -G group1,group2 username` 
4. Ajout d'un répertoire accessible par un groupe : `chown user1:autre ./dossier`
5. Suppression d'un utilisateur en conservant son répertoire `/home` : `sudo deluser user2`
7. Suppression d'un utilisateur en supprimant sont répertoire `/home` `sudo deluser --remove-home user4`
# 4 Logiciels
1. mise à jour système : `sudo apt update && sudo apt upgrade`
2. installation du paquet gparted : `sudo apt install gparted`
3. suppression de la commande nslookup :
	- recherche du paquet correspondant : `dpkg -S nslookup`
	- suppression du paquet correspondant : `sudo apt remove nomPaquet`
4. installation de synaptic depuis la clé USB : `sudo dpkg -i /media/sae/UBUNTU_20_0/soft/syn/synaptic_0.84.6ubuntu5_amd64.deb` (installer avec la même commande les autres paquets (attention aux noms))
# 5 installation de virtual box
```
cd /media/sae/UBUNTU_20_0/soft/vb && sudo dpkg -i *.deb
```
# 6 Disques
1. sur gparted
2. formater la partition créer : `sudo mkfs -t ext4 /dev/sda5`
3. mount : `sudo mkdir /p6 && sudo mount /dev/sda5 /p6`
4. vérification du montage : `df -a /p6`
5. `touch unFichier.txt`
6. démonter la partition : `unmount /p6`