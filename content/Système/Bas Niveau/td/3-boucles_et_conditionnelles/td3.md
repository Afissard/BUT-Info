# TD n°3 : Boucles et conditionnelles

# Objectifs

Les objectifs de cette séance sont :

  - comprendre le fonctionnement des registres de codes de condition et des opérations conditionnelles ;
  - comprendre le fonctionnement des opérations de sauts/branchements ;
  - comprendre les schémas de traduction des structures de contrôle usuelles des langages de haut niveau :  conditionnelles et boucles.


# Compléments sur le jeu d'instruction x86-64

## Les registres de codes de condition

Les registres de codes de condition (ou CCR, pour _Condition Code Register_) appartiennent au fichier de registres mais ne sont généralement pas directement accessibles en écriture (ne peuvent être utilisés comme opérande cible dans une instruction).
Chaque bit d'un CCR est associé à une condition.
Lorsque le bit vaut 1, la condition est vraie, lorsqu'il vaut 0, elle est fausse.

Les conditions sont généralement des informations relatives au résultat de la dernière instruction réalisée sur l'unité arithmétique et logique, par exemple sa nullité ou son signe.
Ainsi, les CCR sont donc mis à jour implicitement lors de l'exécution d'une instruction arithmétique et logique.

Certains jeux d'instruction proposent deux versions des opérations arithmétiques et logiques : une qui met à jour le CCR et une qui ne le modifie pas. 
Ce n'est pas le cas du jeu d'instructions x86-64 : toutes les opérations arithmétiques et logiques mettent à jour le CCR.

Par ailleurs, les jeux d'instructions proposent également des opérations qui réalisent des opérations arithmétiques et modifient le CCR, mais n'ont pas de registre destination. 
En x86-64, c'est le cas des opérations suivantes :
  - `cmp a, b` qui réalise l'opération `b - a`, modifie le CCR en conséquence, mais ne modifie pas `b` ;
  - `test a, b` qui réalise l'opération `a && b`, modifie le CCR en conséquence, mais ne modifie pas `b` ;
  - `bt a, n`, qui « copie » le nième bit de `a` dans le bit de retenue du CCR (soit le bit `CF` du registre `eflags`).

## Les opérations de saut

Les opérations de saut permettent de modifier le flot de contrôle naturel du programme, en modifiant la valeur du pointeur d'instruction.
Ainsi, l'instruction exécutée après une instruction de saut n'est pas celle rangée immédiatement après elle en mémoire, mais celle dont l'adresse est passée en opérande du saut (on parle de cible du saut).
Cette cible est soit une adresse statique, soit une adresse indirecte calculée à l'exécution (cas des pointeurs de fonction et de certaines conditionnelles multiples).

On distingue habituellement les sauts conditionnels et les sauts inconditionnels. 
En x86-64, les trois opérations de saut inconditionnel sont :
  - `jmp adr` où `adr` est une adresse quelconque ;
  - `call adr` où ` adr` est normalement l'adresse d'une procédure dans la section `.text` de la mémoire ;
  - `ret`, qui prend en opérande source la valeur située au sommet de la pile.

Les sauts conditionnels sont de la forme `jxxx adr`, où `xxx` désigne une condition, par exemple `a` pour _above_, `g` pour _greater_ ou encore `ne` pour _not equal_.
Une instruction utilisant une opération de saut conditionnel se comporte comme une instruction de saut si et seulement si les bits du CCR sont tels que la condition est vérifiée.
Dans le cas contraire, elle se comporte comme l'instruction `nop` (pour _no operation_, c'est-à-dire l'instruction qui ne fait rien).

## Autres opérations conditionnelles

En plus des sauts, le jeu d'instruction x86-64 propose une série d'opérations conditionnelles, de la forme `setxxx r`, où `xxx` désigne une condition.
Cette instruction positionne la valeur du registre `r` à `0` (resp. 1), si la condition `xxx` est fausse (resp. vraie).


# Exercice n°1

Soit le programme ci-dessous :

```go
func main() {
    f0 := 0 
    f1 := 1
    for cpt := 3; cpt < 20; cpt ++ {
        tmp := f0 + f1
        f0 = f1
        f1 = tmp 
    }
    fmt.Println(f1)
}
```

Après désassemblage à l'aide de l'outil `objdump`, on obtient le code suivant :


```nasm
47ada0 cmp    0x10(%r14),%rsp
47ada4 jbe    47ae14 <main.main+0x74>
47ada6 push   %rbp
47ada7 mov    %rsp,%rbp
47adaa sub    $0x38,%rsp
47adae mov    $0x3,%eax                 CPT
47adb3 mov    $0x1,%ecx                 1 dans ecx, c'est f1
47adb8 xor    %edx,%edx                 on nullifie edx (il contient 0) c'est donc f0
47adba jmp    47adc9 <main.main+0x29>   on jump à la comp du cpt et de 20
47adbc inc    %rax                      on augmente le cpt
47adbf lea    (%rdx,%rcx,1),%rbx        le calcul de tmp ? (ds rbx ?)
47adc3 mov    %rcx,%rdx                 f0 prend la valeur de f1
47adc6 mov    %rbx,%rcx                 f1 prend la valeur de tmp
47adc9 cmp    $0x14,%rax                comparer le cpt à 20
47adcd jl     47adbc <main.main+0x1c>   on JUMP au début de la boucle si cpt est inf à 14(HEX) soit 20(DEC)
47adcf movups %xmm15,0x28(%rsp)         
47add5 mov    %rcx,%rax
47add8 call   4097a0 <runtime.convT64>
47addd lea    0x6f1c(%rip),%rcx        # 481d00 <type:*+0x6d00>
47ade4 mov    %rcx,0x28(%rsp)
47ade9 mov    %rax,0x30(%rsp)
47adee mov    0xa4833(%rip),%rbx        # 51f628 <os.Stdout>
47adf5 lea    0x3749c(%rip),%rax        # 4b2298 <go:itab.*os.File,io.Writer>
47adfc lea    0x28(%rsp),%rcx
47ae01 mov    $0x1,%edi
47ae06 mov    %rdi,%rsi
47ae09 call   475ca0 <fmt.Fprintln>
47ae0e add    $0x38,%rsp
47ae12 pop    %rbp
47ae13 ret
47ae14 call   45aa60 <runtime.morestack_noctxt.abi0>
47ae19 jmp    47ada0 <main.main>
```

  1. Que calcule ce proramme ? Se renseigner sur la suite de Fibonacci.
  2. Identifier l'implantation des variables `f0`, `f1` et `cpt`.
  3. Analyser le code machine pour comprendre l'implantation de la boucle. Expliquer.
  4. Analyser les deux premières et les deux dernières instructions de la fonction. Proposer une explication sur leur rôle.


# Exercice n°2

Soit le programme ci-dessous :

```go
func main() {
    var x, cpt int
    orig := 10001
    for x, cpt = orig, 0 ; x != 1; cpt++ {
        if (x % 2) == 0 {
            x = x / 2
        } else {
            x = 3 * x + 1
        }
    }
    fmt.Println( orig, "becomes 1 after ", cpt, "iterations")
}
```

Après désassemblage à l'aide de l'outil `objdump`,
le code, construit en désactivant les optimisations et l'inclusion en ligne des fonctions, est le suivant :

```nasm
47adaa	push   %rbp
47adab	mov    %rsp,%rbp
47adae	sub    $0x70,%rsp
47adb2	mov    $0x2711,%eax             orig est initilisé dans eax à 2711(HEX) soit 10001(DEC)
47adb7	xor    %ecx,%ecx                cpt est initialisé dans ecx à 0
47adb9	jmp    47adc1 <main.main+0x21>  on jump à la condition de la boucle (ligne 47adc1)
47adbb	inc    %rcx                     cpt ++
47adbe	mov    %rdx,%rax                
47adc1	cmp    $0x1,%rax                on compare 1 et RAX
47adc5	je     47ade6 <main.main+0x46>  si RAX est égal à 1 on jump à la fin de la fonction (on rend)

47adc7	bt     $0x0,%eax                copie le 1er bit de poid faible de EAX dans le CF (carry flag)
47adcb	jb     47addc <main.main+0x3c>  en fait ça remplace le x % 2 == 0 de la fonction, si c'est pair zou
                                        TOUT PAIR FINI PAR 0 EN BINAIRE
47adcd	mov    %rax,%rdx
47add0	shr    $0x3f,%rax               shift right ? Nullification car on décale de beaucoup de bits ? oui
47add4	add    %rax,%rdx
47add7	sar    %rdx                     shift ARYTHMETIC right diviser par 2 en gardant le bit de signe
                                        (que ça soit positif ou négatif ça marche)
47adda	jmp    47adbb <main.main+0x1b>
47addc	lea    (%rax,%rax,2),%rdx
47ade0	lea    0x1(%rdx),%rdx
47ade4	jmp    47adbb <main.main+0x1b>
47ade6	mov    %rcx,0x28(%rsp)
47adeb	movups %xmm15,0x30(%rsp)
```
  1. Que calcule ce programme ? Se renseigner sur la conjecture de Collatz.
  La conjecture de collatz dit 
  2. Identifier l'implantation des variables `x`, `cpt` et `orig`.
  3. Identifier les lignes du code machine qui correspondent à la boucle.
  4. Identifier le schéma de traduction de la conditionnelle.
  5. Expliquer le schéma de traduction de la division entière par 2.


# Exercice n°3

Soit la fonction `Find` ci-dessous :

```go
func Find (data []int, elem int) (bool, int) {
    lo, hi := 0, len(data)
    mid := lo + (hi - lo) / 2

    for ; data[mid] != elem && hi > lo; {
        if data[mid] > elem {
            hi = mid - 1
            mid = lo + (hi - lo) / 2
        } else {
            lo = mid + 1
            mid = lo + (hi - lo) / 2
        }
    }
    return data[mid] == elem, mid
}

```

Après désassemblage à l'aide de l'outil `objdump`,
le code, construit en désactivant l'inclusion en ligne des fonctions, est le suivant :

```nasm
47c5c0	push   %rbp
47c5c1	mov    %rsp,%rbp
47c5c4	sub    $0x10,%rsp
47c5c8	mov    %rax,0x20(%rsp)
47c5cd	mov    %rbx,%rcx
47c5d0	shr    $0x3f,%rbx ; décalage de 63bit -> nullification
47c5d4	lea    (%rcx,%rbx,1),%rdx ;
47c5d8	sar    %rdx
47c5db	mov    %rcx,%rbx
47c5de	xor    %esi,%esi
47c5e0	cmp    %rdx,%rbx
47c5e3	jbe    47c651 <main.Find+0x91>
47c5e5	mov    (%rax,%rdx,8),%r8
47c5e9	cmp    %r8,%rdi
47c5ec	je     47c645 <main.Find+0x85>
47c5ee	cmp    %rsi,%rcx
47c5f1	jle    47c642 <main.Find+0x82>
47c5f3	cmp    %r8,%rdi
47c5f6	jge    47c619 <main.Find+0x59>
47c5f8	mov    %rdx,%rcx
47c5fb	sub    %rsi,%rdx
47c5fe	lea    -0x1(%rdx),%r8
47c602	shr    $0x3f,%r8
47c606	lea    (%rdx,%r8,1),%rdx
47c60a	lea    -0x1(%rdx),%rdx
47c60e	sar    %rdx
47c611	dec    %rcx
47c614	add    %rsi,%rdx
47c617	jmp    47c5e0 <main.Find+0x20>
47c619	mov    %rcx,%r8
47c61c	sub    %rdx,%r8
47c61f	lea    -0x1(%r8),%r9
47c623	shr    $0x3f,%r9
47c627	lea    (%r8,%r9,1),%r8
47c62b	lea    -0x1(%r8),%r8
47c62f	sar    %r8
47c632	lea    0x1(%rdx),%rsi
47c636	lea    (%rdx,%r8,1),%rdx
47c63a	lea    0x1(%rdx),%rdx
47c63e	xchg   %ax,%ax
47c640	jmp    47c5e0 <main.Find+0x20>
47c642	cmp    %r8,%rdi
47c645	sete   %al
47c648	mov    %rdx,%rbx
47c64b	add    $0x10,%rsp
47c64f	pop    %rbp
47c650	ret
47c651	mov    %rdx,%rax
47c654	mov    %rbx,%rcx
47c657	call   45cb20 <runtime.panicIndex>
```

  1. Que fait la fonction `Find` ?
  2. Étudier le code machine, puis associer chaque ligne du code machine à une ligne du code source.
