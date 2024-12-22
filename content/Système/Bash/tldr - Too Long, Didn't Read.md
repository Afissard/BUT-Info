# Qu'est-ce que tldr ?
tldr est une commande permettant de résumé l'utilisation de n'importe quelle commande bash (ou autres langages shell (comme fish ou zsh)) avec quelques exemples simples ce qui peut être très pratique si vous ne souhaiter pas lire le man sur la commande.
# Installation de tldr
```bash
sudo apt install tldr -y && tldr -u
```
# Exemples
## tldr
```bash
tldr tldr
```
Output :
```
tldr
Display simple help pages for command-line tools from the tldr-pages project.
Note: the
   --language
 and
   --list
 options are not required by the client specification, but most clients implement them.
More information:
https://github.com/tldr-pages/tldr/blob/main/CLIENT-SPECIFICATION.md#command-line-interface
.

 - Print the tldr page for a specific command (hint: this is how you got here!):
   tldr {{command}}

 - Print the tldr page for a specific subcommand:
   tldr {{command}} {{subcommand}}

 - Print the tldr page for a command in the given [L]anguage (if available, otherwise fall back to English):
   tldr --language {{language_code}} {{command}}

 - Print the tldr page for a command from a specific [p]latform:
   tldr --platform {{android|common|freebsd|linux|osx|netbsd|openbsd|sunos|windows}} {{command}}

 - [u]pdate the local cache of tldr pages:
   tldr --update

 - List all pages for the current platform and
   common
:
   tldr --list
```
## chmod
```bash
tldr chmod
```
Output :
```
chmod
Change the access permissions of a file or directory.
More information:
https://www.gnu.org/software/coreutils/chmod
.

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
```
