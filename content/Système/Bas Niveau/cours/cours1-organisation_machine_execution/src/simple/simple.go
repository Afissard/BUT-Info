package main

import "fmt"

func add(a, b int) (res int) {
    res = a + b
    return
}

func main() {

    x := add(4, 9)
    fmt.Printf("%d\n", x)
}
