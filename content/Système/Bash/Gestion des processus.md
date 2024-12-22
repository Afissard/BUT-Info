# ps
Affiche les processus en cours
## Exemples (tldr)
```
 - List running processes (principal)
   ps
   
 - List all running processes:
   ps aux

 - List all running processes including the full command string:
   ps auxww

 - Search for a process that matches a string:
   ps aux | grep {{string}}

 - List all processes of the current user in extra full format:
   ps --user $(id -u) -F

 - List all processes of the current user as a tree:
   ps --user $(id -u) f

 - Get the parent PID of a process:
   ps -o ppid= -p {{pid}}

 - Sort processes by memory consumption:
   ps --sort size
```
# kill
Tue un processus
## Exemple (tldr)
```
 - Terminate a program using the default SIGTERM (terminate) signal:
   kill {{process_id}}

 - List available signal names (to be used without the SIG prefix):
   kill -l

 - Terminate a background job:
   kill %{{job_id}}

 - Terminate a program using the SIGHUP (hang up) signal. Many daemons will reload instead of terminating:
   kill -{{1|HUP}} {{process_id}}

 - Terminate a program using the SIGINT (interrupt) signal. This is typically initiated by the user pressing Ctrl + C:
   kill -{{2|INT}} {{process_id}}

 - Signal the operating system to immediately terminate a program (which gets no chance to capture the signal):
   kill -{{9|KILL}} {{process_id}}

 - Signal the operating system to pause a program until a SIGCONT ("continue") signal is received:
   kill -{{17|STOP}} {{process_id}}

 - Send a SIGUSR1 signal to all processes with the given GID (group id):
   kill -{{SIGUSR1}} -{{group_id}}
```
# htop et / ou btop
Deux questionnaire de processus (= gestionnaire des tache de windows) en tui, htop est généralement plus commun car parfois installé par défaut (varie selon l'os). Les deux supporte l'usage de la souris de part leurs utilisation de la librairie curses.