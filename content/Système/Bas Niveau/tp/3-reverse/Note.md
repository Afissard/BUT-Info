# Prise de notes Reverse Ingenering

# Tentative Jump By Pass
## Info eflags :
`info registers eflags`

## Comment modifié la valeur d'un registre :
`set $rax = 0 # met rax à 0 (int)`

## Step1 :
Info : "a" est faux
### Préparation :
- `b 25` et `b *0x493cc0` pour directement aller dans l'ASM
- entrer un string aléatoire (probablement faux)
### Observation et analyse
Jump numéro 1 : `jbe 0x493cf8 <main.step1+56>`
Dans le cas ou l'input est (supposé) faux, le jump n'est pas utilisé.
On arrive donc au Jump numéro 2 : `je 0x493ce1 <main.step1+33>`
Dans le cas ou l'input est (supposé) faux, le jump n'est pas utilisé.
-> fin
Info le premier step est bugé, pour passé il suffit d'une chaine de caractère de la bonne taille (22)
### Solution pour passer `je 0x493ce1 <main.step1+33>` :
la ligne décisive est : `cmp 0x16, %rbx`
pour passer le je : `set $rbx = 22`

## Step2 :
### Préparation :
- `start < passorwds.txt` : donne en entré le fichier txt contenant les précedant mot de passe
- `b *0x493eb9` : `cmp $0x5, %rax`
### Solution :
- mettre la valeur de rax à 0x5 : `set $rax = 5`
- `ni`

## Step3 :
### Préparation :
- `start < passorwds.txt` : donne en entré le fichier txt contenant les précedant mot de passe
- `b 41`

# Script gdb pour passer :
```
start < passorwds.txt

# Tout les breakpoint
b *0x493cd3 # cmp 0x16, %rbx # position condition pour passer Step1
b *0x493eb9 # cmp $0x5, %rax # position condition pour passer Step2

# Passage Step1
c
set $rbx = 22 # rempli la condition step 1

# Passage Step2
c
set $rax = 5

# Passage Step3
c
?
```
