# TD n°4 : Pointeurs

# Objectifs

Les objectifs de cette séance sont :

  - mieux comprendre les pointeurs,
  - mieux comprendre les modes d'adressage direct et indirect.



# Exercice n°1

Soit le programme ci-dessous :

```go
package main

import (
    "fmt"
)

//go:noinline
func f1(x *int, y *int) {
    *x = *x + *y
    return
}

//go:noinline
func f2(x int, y int, t [3]int) {
    t[x] = t[x] + t[y]
}

//go:noinline
func f3(x int, y int, t []int) {
    t[x] = t[x] + t[y]
}

//go:noinline
func f4(x int, y int, t []*int) {
    *(t[x]) = *(t[x]) + *(t[y])
}

func main() {
    t1 := [...]int{0, 1, 2}
    t2 := make([]int,len(t1))
    t3 := make([]*int,len(t1))

    for i:=0; i<len(t1); i++ {
        t2[i] = t1[i]
        t3[i] = &(t1[i])
    }

    f1(&t1[1], &t1[2])
    f2(1, 2, t1) // attention à la nature du tableau donnée : pas les même effet quand passé en paramètres, cf diffrence array / slice.
    f3(1, 2, t2)
    f4(1, 2, t3)

    for cpt:=0; cpt<len(t1); cpt++ {
        fmt.Println(t1[cpt])
    }
    for _, e := range t2 { 
        fmt.Println(e)
    }
    for _, e := range t3 { 
        fmt.Println(*e)
    }
}

```

Et le code x86-64 ci-dessous : 

```nasm

<main.fa> ; fonction f2
  push   %rbp
  mov    %rsp,%rbp
  sub    $0x10,%rsp
  cmp    $0x3,%rax ; vérification taille du tableau = 3 ?
  jae    47adf6
  cmp    $0x3,%rbx
  jae    47ade9 
  mov    0x20(%rsp,%rax,8),%rcx
  add    0x20(%rsp,%rbx,8),%rcx
  mov    %rcx,0x20(%rsp,%rax,8)
  add    $0x10,%rsp
  pop    %rbp
  ret
  mov    %rbx,%rax    # 0x47ade9
  mov    $0x3,%ecx
  call   45cb20 <runtime.panicIndex>
  mov    $0x3,%ecx    # 0x47adf6
  nopl   0x0(%rax,%rax,1)
  call   45cb20 <runtime.panicIndex>

<main.fb> ; fonction : f1
  mov    (%rax),%rcx
  add    (%rbx),%rcx
  mov    %rcx,(%rax)
  ret

<main.fc> ; fonction : f4
  push   %rbp
  mov    %rsp,%rbp
  sub    $0x10,%rsp
  mov    %rcx,0x30(%rsp)
  cmp    %rax,%rdi
  jbe    47ae99 
  mov    (%rcx,%rax,8),%rdx
  cmp    %rbx,%rdi
  jbe    47ae8e 
  mov    (%rdx),%rax
  mov    (%rcx,%rbx,8),%rcx
  add    (%rcx),%rax
  mov    %rax,(%rdx)
  add    $0x10,%rsp
  pop    %rbp
  ret
  mov    %rbx,%rax    # 0x47ae8e
  mov    %rdi,%rcx
  call   45cb20 <runtime.panicIndex>
  mov    %rdi,%rcx    # 0x47ae99
  nopl   0x0(%rax)
  call   45cb20 <runtime.panicIndex>

<main.fd> ; fonction : f3
  push   %rbp
  mov    %rsp,%rbp
  sub    $0x10,%rsp
  mov    %rcx,0x30(%rsp)
  cmp    %rax,%rdi
  jbe    47ae54
  mov    (%rcx,%rax,8),%rdx
  cmp    %rbx,%rdi
  jbe    47ae49
  add    (%rcx,%rbx,8),%rdx
  mov    %rdx,(%rcx,%rax,8)
  add    $0x10,%rsp
  pop    %rbp
  ret
  mov    %rbx,%rax     # 0x47ae49
  mov    %rdi,%rcx 
  call   45cb20 <runtime.panicIndex>
  mov    %rdi,%rcx     # 0x47ae54
  call   45cb20 <runtime.panicIndex>
```

  1. Prédire ce qu'affiche le programme go lorsqu'il est exécuté. 
  2. Le code x86-64 a été obtenu après compilation du programme go, mais l'ordre et le nom des fonctions ont été changés. Identifier les fonctions.
  3. Schématiser l'état de la mémoire au début de chacune des fonctions.
  4. Construire le programme et vérifier les réponses données aux questions précédentes.

# Exercice n°2

Écrire une fonction go qui pourrait correspondre au code x86-64 ci-dessous.

```nasm
<main.f>
  push   %rbp
  mov    %rsp,%rbp
  sub    $0x10,%rsp
  mov    %rax,0x20(%rsp)
  mov    %rbx,0x28(%rsp)
  lea    0x6f41(%rip),%rax
  nop
  call   40bda0 <runtime.newobject>
  mov    0x20(%rsp),%rcx
  mov    (%rcx),%rcx
  mov    0x28(%rsp),%rdx
  cmp    %rcx,%rdx
  jge    47addf <main.f+0x3f>
  sub    %rdx,%rcx
  mov    %rcx,(%rax)
  jmp    47ade5 <main.f+0x45>
  sub    %rcx,%rdx
  mov    %rdx,(%rax)
  add    $0x10,%rsp
  pop    %rbp
  ret
```


