# mkdir
Création de dossiers avec leur permissions
## Exemple (tldr)
```
- Create specific directories:
   mkdir {{path/to/directory1 path/to/directory2 ...}}

- Create specific directories and their [p]arents if needed:
   mkdir -p {{path/to/directory1 path/to/directory2 ...}}

- Create directories with specific permissions:
   mkdir -m {{rwxrw-r--}} {{path/to/directory1 path/to/directory2 ...}}
```
# touch
Créer un fichier et paramètre ses permissions
## Exemple (tldr)
```
 - Create specific files:
   touch {{path/to/file1 path/to/file2 ...}}

 - Set the file [a]ccess or [m]odification times to the current one and don't [c]reate file if it doesn't exist:
   touch -c -{{a|m}} {{path/to/file1 path/to/file2 ...}}
```
# rm
Supprime les fichiers / dossiers sélectionné
## Exemple (tldr)
```
 - Remove specific files:
   rm {{path/to/file1 path/to/file2 ...}}

 - Remove specific files ignoring nonexistent ones:
   rm -f {{path/to/file1 path/to/file2 ...}}

 - Remove specific files [i]nteractively prompting before each removal:
   rm -i {{path/to/file1 path/to/file2 ...}}

 - Remove specific files printing info about each removal:
   rm -v {{path/to/file1 path/to/file2 ...}}

 - Remove specific files and directories [r]ecursively:
   rm -r {{path/to/file_or_directory1 path/to/file_or_directory2 ...}}
```
## Si vous êtes sur de ce que vous faites
```bash
rm -rf ./dossier/a/supprimer # supprime le dossier et son contenu
```
# cp
Copie les fichiers et / ou dossier sélectionnés
## Exemples (tldr)
```
 - Copy a file to another location:
   cp {{path/to/source_file.ext}} {{path/to/target_file.ext}}

 - Copy a file into another directory, keeping the filename:
   cp {{path/to/source_file.ext}} {{path/to/target_parent_directory}}

 - Recursively copy a directory's contents to another location (if the destination exists, the directory is copied inside it):
   cp -R {{path/to/source_directory}} {{path/to/target_directory}}

 - Copy a directory recursively, in verbose mode (shows files as they are copied):
   cp -vR {{path/to/source_directory}} {{path/to/target_directory}}

 - Copy multiple files at once to a directory:
   cp -t {{path/to/destination_directory}} {{path/to/file1 path/to/file2 ...}}

 - Copy text files to another location, in interactive mode (prompts user before overwriting):
   cp -i {{*.txt}} {{path/to/target_directory}}
```
# nano
Editeur de texte dans le terminal (plus "user friendly" que vi, vim, neovim ou emacs)
## Exemples (tldr)
```
 - Start the editor:
   nano

 - Open a specific file with the editor:
   nano {{path/to/file}}

 - Open specific files, moving to the next file when closing the previous one:
   nano {{path/to/file1 path/to/file2 ...}}

 - Open a file and create a backup file (path/to/file~) on save:
   nano --backup {{path/to/file}}
```