# cd 
Change Directory
```bash
cd # va à la racine de l'environement utilsateur
cd / # va à la racine du système
cd - # retourne au précédant dossier

# Se déplace, en avant, dans un dossier
cd Documents
cd ./Documents/golang

# Se déplace, en arrière, vers un dossier précédant
cd ..
cd ../..

# Combinaison des deux
cd ../Images
cd ../Document/golang
```
# ls
Liste le contenu du dossier
## Exemple (tldr)
```
 - List files one per line:
   ls -1

 - List all files, including hidden files:
   ls -a

 - List all files, with trailing "/" added to directory names:
   ls -F

 - Long format list (permissions, ownership, size, and modification date) of all files:
   ls -la

 - Long format list with size displayed using human-readable units (KiB, MiB, GiB):
   ls -lh

 - Long format list sorted by size (descending):
   ls -lS

 - Long format list of all files, sorted by modification date (oldest first):
   ls -ltr

 - Only list directories:
   ls -d */
```
# tree 
Donne l'arboresence à partir du dossier actuel
## Exemple (tldr)
```
 - Print files and directories up to 'num' levels of depth (where 1 means the current directory):
   tree -L {{num}}

 - Print directories only:
   tree -d

 - Print hidden files too with colorization on:
   tree -a -C

 - Print the tree without indentation lines, showing the full path instead (use
   "-N" to not escape non-printable characters):
   tree -i -f

 - Print the size of each file and the cumulative size of each directory, in human-readable format:
   tree -s -h --du

 - Print files within the tree hierarchy, using a wildcard (glob) pattern, and pruning out directories that don't contain matching files:
   tree -P '{{*.txt}}' --prune

 - Print directories within the tree hierarchy, using the wildcard (glob) pattern, and pruning out directories that aren't ancestors of the wanted one:
   tree -P {{directory_name}} --matchdirs --prune

 - Print the tree ignoring the given directories:
   tree -I '{{directory_name1|directory_name2}}'
```
