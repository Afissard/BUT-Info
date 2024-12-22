# date
Donne la date et l'heure
## Exemple (tldr)
```
 - Display the current date
   date

 - Display the current date using the default locale's format:
   date +%c

 - Display the current date in UTC, using the ISO 8601 format:
   date -u +%Y-%m-%dT%H:%M:%S%Z

 - Display the current date as a Unix timestamp (seconds since the Unix epoch):
   date +%s

 - Convert a date specified as a Unix timestamp to the default format:
   date -d @{{1473305798}}

 - Convert a given date to the Unix timestamp format:
   date -d "{{2018-09-01 00:00}}" +%s --utc

 - Display the current date using the RFC-3339 format (YYYY-MM-DD hh:mm:ss TZ):
   date --rfc-3339=s

 - Set the current date using the format MMDDhhmmYYYY.ss (YYYY and .ss are optional):
   date {{093023592021.59}}

 - Display the current ISO week number:
   date +%V
```
# history
Donne l'historique des commande utilisé
## Exemple (tldr)
```
 - Display the commands history list with line numbers:
   history

 - Display the last 20 commands (in zsh it displays all commands starting from the 20th):
   history {{20}}

 - Clear the commands history list (only for current bash shell):
   history -c

 - Overwrite history file with history of current bash shell (often combined with history -c to purge history):
   history -w

 - Delete the history entry at the specified offset:
   history -d {{offset}}
```
# sleep
Attend x secondes
## Exemples (tldr)
```
 - Delay in seconds:
   sleep {{seconds}}

 - Execute a specific command after 20 seconds delay:
   sleep 20 && {{command}}
```
# which
Donne le chemin vers l'exécutable de la commande donnée
## Exemple (tldr)
```
 - Search the PATH environment variable and display the location of any matching executables:
   which {{executable}}

 - If there are multiple executables which match, display all:
   which -a {{executable}}
```
# xxd
Converti en hexadécimal et inversement
## Exemple (tldr)
```
 - Generate a hexdump from a binary file and display the output:
   xxd {{input_file}}

 - Generate a hexdump from a binary file and save it as a text file:
   xxd {{input_file}} {{output_file}}

 - Display a more compact output, replacing consecutive zeros (if any) with a star:
   xxd -a {{input_file}}

 - Display the output with 10 columns of one octet (byte) each:
   xxd -c {{10}} {{input_file}}

 - Display output only up to a length of 32 bytes:
   xxd -l {{32}} {{input_file}}

 - Display the output in plain mode, without any gaps between the columns:
   xxd -p {{input_file}}

 - Revert a plaintext hexdump back into binary, and save it as a binary file:
   xxd -r -p {{input_file}} {{output_file}}
```