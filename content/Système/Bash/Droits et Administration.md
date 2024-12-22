# sudo
Permet d'éxécuté une commande en temps que super utilisateur ou en temps qu'un autre utilisateur
## Exemples (tldr)
```
 - Run a command as the superuser:
   sudo {{less /var/log/syslog}}

 - Edit a file as the superuser with your default editor:
   sudo --edit {{/etc/fstab}}

 - Run a command as another user and/or group:
   sudo --user={{user}} --group={{group}} {{id -a}}

 - Repeat the last command prefixed with sudo
   sudo !!

 - Launch the default shell with superuser privileges and run login-specific files (.profile, .bash_profile, etc.):
   sudo --login

 - Launch the default shell with superuser privileges without changing the environment:
   sudo --shell

 - Launch the default shell as the specified user, loading the user's environment and reading login-specific files (.profile, .bash_profile, etc.):
   sudo --login --user={{user}}

 - List the allowed (and forbidden) commands for the invoking user:
   sudo --list
```
# chmod
Change les permission d'accès (lecture / écriture / exécution) du fichier / dossier donné
en paramètre.
## Exemples (tldr)
```
 - Give the [u]ser who owns a file the right to e[x]ecute it:
   chmod u+x {{path/to/file}}

 - Give the [u]ser rights to [r]ead and [w]rite to a file/directory:
   chmod u+rw {{path/to/file_or_directory}}

 - Remove e[x]ecutable rights from the [g]roup:
   chmod g-x {{path/to/file}}

 - Give [a]ll users rights to [r]ead and e[x]ecute:
   chmod a+rx {{path/to/file}}

 - Give [o]thers (not in the file owner's group) the same rights as the [g]roup:
   chmod o=g {{path/to/file}}

 - Remove all rights from [o]thers:
   chmod o= {{path/to/file}}

 - Change permissions recursively giving [g]roup and [o]thers the ability to [w]rite:
   chmod -R g+w,o+w {{path/to/directory}}

 - Recursively give [a]ll users [r]ead permissions to files and e[X]ecute permissions to sub-directories within a directory:
   chmod -R a+rX {{path/to/directory}}
```
