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
    f2(1, 2, t1)
    f3(1, 2, t2)
    f4(1, 2, t3)

    fmt.Print("\nt1 : ")
    for cpt:=0; cpt<len(t1); cpt++ {
        fmt.Print(t1[cpt], " ")
    }
    fmt.Print("\nt2 : ")
    for _, e := range t2 { 
        fmt.Print(e, " ")
    }
    fmt.Print("\nt3 : ")
    for _, e := range t3 { 
        fmt.Print(*e, " ")
    }
}

