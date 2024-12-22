package main

import "fmt"

func f1(a, b int) (r int) {
    if a >= b {
        r = a-b
    } else {
        r = b-a
    }
    return
}

func f2(x  int) (r int) {
    r = 1
    for cpt := 1; cpt < x; cpt+=2 {
        r *= cpt
    }
    return
}

func f3(x int) (r int) {
    x = x % 6
    switch(x) {
    case 0:
        r = 2*x
    case 1:
        r = 1/x
    case 2:
        r = x*x
    case 3:
        r = x/2
    case 4:
        r = x+3
    case 5:
        r = x-7
    }
    return
}

func main() {
    x := 123
    y := 321
    fmt.Printf("f1(%d, %d) = %d\n", x, y, f1(x,y))
    fmt.Printf("f2(%d) = %d\n", x, f2(x))
    fmt.Printf("f3(%d) = %d\n", x, f3(x))
}

