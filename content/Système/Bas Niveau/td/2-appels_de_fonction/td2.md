# TD n°2 : Appel de fonction et passage de paramètre en assembleur.

# Objectifs

Les objectifs de cette séance sont :

  - découvrir l'assembleur x86-64,
  - découvrir la mise en oeuvre de l'appel de fonction en Go au niveau assembleur,
  - apprendre à comprendre les correspondances code source et code machine.

# L'assembleur x86-64
Le jeu d'instruction x86-64 est un jeu d'instruction CISC 64 bits, utilisés sur des processeurs Intel et AMD.
C'est un jeu d'instruction très (très) complexe.
Heureusement, pour ce TD, connaître un sous-ensemble très petit est suffisant.

## Les registres

Les registres dont le nom commence par `r` ont une largeur de 8 octets (64 bits).
Les registres dont le nom commence par `e` ont une largeur de 4 octets (32 bits). 
Ils correspondent aux 4 octets de poids faible du registre de 8 octets dont le nom termine par les mêmes 2 caractères.
Ainsi, `eax` correspond aux 4 octets de poids faible de `rax`.
Les registres dont le nom comportent seulement 2 caractères dont le dernier est `x` ont une largeur de 2 octets (16 bits).
Ils correspondent aux 2 octets de poids faible du registre de 4 octets dont le nom termine par les mêmes 2 caractères.
Ainsi, `ax` correspond aux 2 octets de poids faible de `eax` (qui sont également les 2 octets de poids faible de `rax`).
Il existe également des registres qui ont une largeur de 1 octet.

Parmi tous les registres x86-64, ceux utilisés pour cette séance sont :

- `rsp` : le pointeur de pile ;
- `rbp` : le pointeur de base, ou pointeur de cadre.
- `rax`, `rbx`, `rcx` : des registres qui peuvent être utilisés pour stocker des valeurs temporaires de calcul, mais pas uniquement (ce sera plus clair à la fin de la séance).

## Les instructions

Le jeu d'instruction x86-64 relevant du principe de conception CISC, il comporte de nombreuses instructions, dont certaines sont très complexes.
Heureusement, pour cette séance, un sous ensemble très réduit est suffisant.

Une instruction est constituée d'une opération, et de zéro à plusieurs opérandes.

### Opérations

Pour cette séance, les opérations suivantes suffisent.

- `mov`, `movq` : copie l'opérande 1 dans l'opérande 2 ; l'opérande 1 peut être une constante, un registre (copie du contenu du registre), ou une adresse mémoire (copie du contenu de la mémoire à cette adresse) ; l'opérande 2 peut être un registe ou une adresse mémoire ; les 2 opérandes ne peuvent pas être des adresses mémoires. La variante `movq`[^2] travaille *forcément* sur un bloc de 4 * 16 bits (donc 8 octets, ou 64 bits), quelle que soit la taille de l'opérande manipulée.

- `add`, `sub`, `imul`, `xor` : addition, soustraction, multiplication et ou exclusif (bit à bit). Le résultat est écrit dans l'opérande 2 et le registre des codes de condition est mis à jour.

- `lea` (pour "*Load Effective Address*") : charge l'adresse de l'opérande 1 dans l'opérande 2 (comparable dans un sens à l'opérateur `&` de Go). Mais ce n'est pas son unique usage de nos jours ; elle peut aussi servir à combiner des calculs **et** une affectation en une seule instruction[^1].

- `cmp` : comme `sub` mais le résultat n'est pas écrit dans l'opérande 2.

- `test` : réalise un et bit-à-bit entre les 2 opérandes, le résultat est jeté mais le registre des codes de condition est mis à jour.

- `je`, `jbe`, `jne` : saut conditionnels, c'est-à-dire que l'opérande est copiée dans le pointeur d'instruction si une certaine condition est vraie ; la condition est évaluée en lisant les bits du registre des codes de condition (resp. : *equal*, *below*, *not equal*).

- `call` et `ret` : appel et retour de fonction.

Enfin, certaines instructions qui ne font rien sont également insérées dans le code engendré par le compilateur.
Elles servent généralement à adapter l'organisation en mémoire du programme aux différents mécanismes liés à la mémoire et ayant un impact sur les performances (mémoire cache ou pagination, par exemple). 
On trouve ainsi l'instruction `xchg %ax, %ax`, parfois codé `nop`, et l'instruction `nopl`. 
La différence entre ces deux instructions est la taille de leur code : 1 octet pour la première, et 2 pour la seconde.

## Opérandes

Les opérandes peuvent être :

- un registre, donné par son nom précédé du caractère `%` ;
- une constante entière, donné par sa valeur en décimal ou hexadécimal précédée du caractère `$` ;
- une adresse mémoire (voir ci-dessous).

### Adresses mémoire en tant qu'opérande
`
Une adresse peut être passée comme opérande de différentes façon, qu'on appelle mode d'adressage.
Pour cette séance, les modes d'adressage à connaître sont :

- adresse constante, donnée directement par sa valeur ;

- adresse contenue dans un registre, donné par le nom du registre, précédé du caractère `%` et entre parenthèse, par exemple `(%rax)` ;

- adresse de base plus déplacement : dans ce cas, l'adresse de base est donné par un registre, et le déplacement est précisé sous la forme d'une constante donnée avant l'adresse de base ; ainsi, `0x8(%rax)` est l'adresse située 8 octets après l'adresse mémorisée dans `%rax` ;

- trois opérandes sous la forme de 2 valeurs de registres et un facteur d'échelle, le tout entre parenthèses : on multiplie la deuxième opérande par le facteur d'échelle, puis l'on ajoute le tout à la première opérande. Cela est prévu à la base pour retrouver des cases d'un tableau à partir d'un indice et la taille de chaque case (le facteur d'échelle), mais on peut aussi en faire un usage « détourné » pour gagner en lignes de code machine.

# Exercice n°1

Soit la fonction `RetVal` :

```go
func RetVal() int {
	return 42
}
```

Après désassemblage à l'aide de l'outil `objdump`,
le code, construit en désactivant les optimisations et l'inclusion en ligne des fonctions, est le suivant :

```nasm
sub    $0x10,%rsp
mov    %rbp,0x8(%rsp)
lea    0x8(%rsp),%rbp
movq   $0x0,(%rsp)
movq   $0x2a,(%rsp)
mov    $0x2a,%eax
mov    0x8(%rsp),%rbp
add    $0x10,%rsp
ret
```

1. Exécuter ce code "à la main" et indiquer ce que fait chaque instruction.
Que peut-on dire sur ce code ?

```nasm
sub    $0x10,%rsp     ; rsp -= 0x10 ; ouverture cadre
mov    %rbp,0x8(%rsp) ; cpy rbp -> rsp + 0x8 ; ouverture cadre
lea    0x8(%rsp),%rbp ; cpy &(rsp + 0x8) -> &rbp
movq   $0x0,(%rsp)    ; cpy64bit 0x0 -> rsp ; init à 0
movq   $0x2a,(%rsp)   ; cpy64bit 0x0 -> rsp ; met à 42
mov    $0x2a,%eax     ; cpy 0x2a -> eax ; met 42 dans le registre eax
mov    0x8(%rsp),%rbp ; cpy rsp + 0x8 -> rbp ; fermeture du cadre
add    $0x10,%rsp     ; rsp += 0x10 ; fermeture du cadre
ret                   ; fin de programme
; le code n'est pas très optimisé, certaine valeurs ne sont pas utilisé, exemple ligne 5-6, rsp est initalisé à 0 puis à 42. 
```

2. En utilisant le même programme source, mais en activant les optimisations, on obtient le code ci-dessous.
Que peut-on dire sur ce code ? Sur le rôle du registre `eax` ?

```nasm
mov    $0x2a,%eax
ret              
```

# Exercice n°2

Soit la fonction `AddSub`:

```go
func AddSub(a, b int) (add, sub int) {
	return a + b, a - b
}
```

Après désassemblage à l'aide de l'outil `objdump`,
le code, construit en désactivant les optimisations et l'inclusion en ligne des fonctions, est le suivant :

```nasm
sub    $0x18,%rsp
mov    %rbp,0x10(%rsp)
lea    0x10(%rsp),%rbp
mov    %rax,0x20(%rsp)
mov    %rbx,0x28(%rsp)
movq   $0x0,0x8(%rsp)
movq   $0x0,(%rsp)
mov    0x20(%rsp),%rcx
add    0x28(%rsp),%rcx
mov    %rcx,0x8(%rsp)
mov    0x20(%rsp),%rbx
sub    0x28(%rsp),%rbx
mov    %rbx,(%rsp)
mov    0x8(%rsp),%rax
mov    0x10(%rsp),%rbp
add    $0x18,%rsp
ret
```

1. Exécuter ce code "à la main" et indiquer ce que fait chaque instruction.
```nasm
sub    $0x18,%rsp      ; rsp -= 0x18
mov    %rbp,0x10(%rsp) ; cpy rbp -> rsp + 0x10
lea    0x10(%rsp),%rbp ; cpy &rbp -> &(rsp + 0x10)
mov    %rax,0x20(%rsp) ; cpy rax -> rsp + 0x20
mov    %rbx,0x28(%rsp) ; cpy rbx -> rsp + 0x28 ; sauvegarde des regs
movq   $0x0,0x8(%rsp)  ; cpy64bit 0x0 -> rsp + 0x8
movq   $0x0,(%rsp)     ; cpy64bit 0x0 -> rsp
mov    0x20(%rsp),%rcx ; cpy rsp + 0x20 -> rcx
add    0x28(%rsp),%rcx ; rcx += (rsp + 0x28)
mov    %rcx,0x8(%rsp)  ; cpy rcx -> rsp + 0x8
mov    0x20(%rsp),%rbx ; cpy rsp + 0x20 -> rbx
sub    0x28(%rsp),%rbx ; rbx -= rsp + 0x28
mov    %rbx,(%rsp)     ; cpy rbx -> (rsp) ; ? ; restoration des regs
mov    0x8(%rsp),%rax  ; cpy rsp + 0x8 -> rax
mov    0x10(%rsp),%rbp ; cpy rsp + 0x10 -> rbp
add    $0x18,%rsp      ; rsp += 0x18
ret                    ; retourne rax (ret n°1) et rbx (ret n°2)
; NB :
; lea : comme un mov mais l'adresse est manipulé et non la valeur
```

2. Quelles expressions correspondent aux adresses des variables `a`, `b`, `add`, et `sub` ? Que peut-on conclure concernant l'allocation en mémoire de ces variables.
	- les paramètres ont été passé par la pile 
	- les valeurs de retour sont aussi passé par la pile

3. En utilisant le même programme source, mais en activant les optimisations, on obtient le code ci-dessous.
Comme à la question précédente, exécuter à la main et annoter le code. 
Que peut-on dire concernant l'allication en mémoire des variables `a`, `b`, `add`, et `sub` ?


```nasm
lea    (%rax,%rbx,1),%rcx
sub    %rbx,%rax
mov    %rax,%rbx
mov    %rcx,%rax
ret
```

```nasm
lea    (%rax,%rbx,1),%rcx ; lea est détourné pour : rbx*1+rbx -> rcx
sub    %rbx,%rax          ; rbx -= rax
mov    %rax,%rbx          ; cpy rax -> rbx
mov    %rcx,%rax          ; cpy rcx -> rax
ret                       ; 
```
les inputs et les outputs passe par la pile encore une fois
# Exercice n°3

Soit la fonction `CallAddSub` :

```go
func CallAddSub() (a, b int) {
	a, b = AddSub(4, 5)
	return a, b
}
```

Après désassemblage à l'aide de l'outil `objdump`,
le code, construit en désactivant les optimisations et l'inclusion en ligne des fonctions, est le suivant (les adresses des instructions ont été ajoutées pour permettre d'identifier les cibles des sauts) :

```nasm
4552e0		cmp    0x10(%r14),%rsp ; (r14 + 0x10) = (r14 + 0x10) - rsp
; va à cette endroit si nécessité d'étendre le stack
4552e4		jbe    455347 <main.CallAddSub+0x67> 
4552ea		mov    %rbp,0x30(%rsp) ; cpy rbp-> rsp + 0x30
4552ef		lea    0x30(%rsp),%rbp ; cpy &(rsp + 0x30) -> rbp
4552f4		movq   $0x0,0x18(%rsp) ; cpy64bit 0x0 -> rsp + 0x0
4552fd		movq   $0x0,0x10(%rsp) ; cpy64bit 0x0 -> rsp + 0x10
; passages des variables à la fonction AddSub
455306		mov    $0x4,%eax       ; cpy 0x4 -> eax
45530b		mov    $0x5,%ebx       ; cpy 0x5 -> ebx
; appel de la fonction AddSub
455310		call   455280 <main.AddSub> 
455315		mov    %rax,0x28(%rsp) ; cpy rax -> rsp + 0x28
45531a		mov    %rbx,0x20(%rsp) ; cpy rbx -> rsp + 0x20
45531f		mov    0x28(%rsp),%rcx ; cpy rsp + 0x28 -> rcx
455324		mov    %rcx,0x18(%rsp) ; cpy rcx -> rsp + 0x18
455329		mov    0x20(%rsp),%rcx ; cpy rsp + 0x20 -> rcx
45532e		mov    %rcx,0x10(%rsp) ; cpy rcx -> rsp + 0x10
455333		mov    0x18(%rsp),%rax ; cpy rsp + 0x18 -> rax
455338		mov    0x10(%rsp),%rbx ; cpy rsp + 0x10 -> rbx
45533d		mov    0x30(%rsp),%rbp ; cpy rsp + 0x30 -> rbp
455342		add    $0x38,%rsp      ; rsp += 0x38
455346		ret    
; étend le stack
455347		call   451f60 <runtime.morestack_noctxt.abi0> 
; retourne au code de la function
45534c		jmp    4552e0 <main.CallAddSub> 
```
1. À quoi correspond le paquet `runtime` du langage Go ? 

2. Exécuter ce code "à la main" et indiquer ce que fait chaque instruction.

3. En utilisant le même programme source, mais en activant les optimisations, on obtient le code ci-dessous.
Comme à la question précédente, exécuter à la main et annoter le code. 
Malgré l'activation des optimisations, cette fonction alloue un cadre de pile. Expliquer pourquoi. 

```nasm
455260 	cmp    0x10(%r14),%rsp
455264 	jbe    45528f <main.CallAddSub+0x2f>
455266 	sub    $0x18,%rsp
45526a 	mov    %rbp,0x10(%rsp)
45526f 	lea    0x10(%rsp),%rbp
455274 	mov    $0x4,%eax
455279 	mov    $0x5,%ebx
45527e 	xchg   %ax,%ax ; ne sert à rien (usage equivalent à nop / nopl)
455280 	call   455240 <main.AddSub>
455285 	mov    0x10(%rsp),%rbp
45528a 	add    $0x18,%rsp
45528e 	ret    
45528f 	call   451f60 <runtime.morestack_noctxt.abi0>
455294 	jmp    455260 <main.CallAddSub>
```
La fonction AddSub à besoin d'un cadre pour opéré, CallAddSub le créer, même si AddSub à été optimisé et n'en nécessite pas.

# Exercice n°4

Soit la fonction `ManyArgs` :

```go
func ManyArgs(a, b, c, d, e, f, g, h, i, j, k, l int) (r int) {
	r = a*b + c*d - e*f*g + h + i*j*k - l
	return
}
```

Après désassemblage à l'aide de l'outil `objdump`,
le code, construit en désactivant les optimisations et l'inclusion en ligne des fonctions, est le suivant :

```nasm
; création du cadre
sub    $0x10,%rsp
mov    %rbp,0x8(%rsp)
lea    0x8(%rsp),%rbp
; passage des arguments (mais il manque j, k, l)
mov    %rax,0x30(%rsp)
mov    %rbx,0x38(%rsp)
mov    %rcx,0x40(%rsp)
mov    %rdi,0x48(%rsp)
mov    %rsi,0x50(%rsp)
mov    %r8,0x58(%rsp)
mov    %r9,0x60(%rsp)
mov    %r10,0x68(%rsp)
mov    %r11,0x70(%rsp)
movq   $0x0,(%rsp)

mov    0x38(%rsp),%rcx
mov    0x30(%rsp),%rdx
imul   %rcx,%rdx
mov    0x40(%rsp),%rcx
mov    0x48(%rsp),%rbx
imul   %rbx,%rcx
add    %rdx,%rcx
mov    0x50(%rsp),%rdx
mov    0x58(%rsp),%rbx
imul   %rbx,%rdx
mov    0x60(%rsp),%rbx
imul   %rbx,%rdx
sub    %rdx,%rcx
add    0x68(%rsp),%rcx
mov    0x70(%rsp),%rdx
mov    0x18(%rsp),%rbx
imul   %rbx,%rdx
mov    0x20(%rsp),%rbx
imul   %rbx,%rdx
lea    (%rcx,%rdx,1),%rax
sub    0x28(%rsp),%rax
mov    %rax,(%rsp)
mov    0x8(%rsp),%rbp
add    $0x10,%rsp
ret    
```

En analysant ce code, identifier comment les valeurs des paramètres sont passées à cette fonction.
Dessiner la pile après l'exécution du prologue de la fonction en incluant tous les emplacements mémoire utilisés par la fonction.
Lorsque c'est possible, indiquer le nom de l'objet Go auquel correspondent ces emplacements.


# Exercice n°5

1. Construire le code du programme `funcs.go` (dans le dossier `src`) en utilisant la commande `go build` (c'est-à-dire en activant les optimisations et l'intégration en ligne des fonctions).
Désassembler le programme et expliquer le code obtenu.

2. Modifier la fonction `main` en ajoutant un appel à `fmt.Println` qui affiche la somme du premier résultat de la fonction `CallAddSub` et le résultat de la fonction `ManyArgs`. Désassembler le programme et expliquer le code obtenu.

---

[^1]: Cf. les explications (*en anglais*) de [cette question/réponse](https://povilasv.me/go-memory-management/).

[^2]: Avec "q" pour "quadword", soit 4 mots de 16 bits.
