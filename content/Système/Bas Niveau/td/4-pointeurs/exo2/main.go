package main

import "fmt"

/*
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
*/


func f(a *int, b int)(*int){
	var c int
	if *a < b {
		c = *a-b
	} else {
		c = b-*a
	}
	return &c
}

func main() {
	fmt.Println("exo2")
	temp := 1

	res1 := f(&temp, 2)
	res2 := f(&temp, -2)
	fmt.Println(res1, *res1)
	fmt.Println(res2, *res2)
}