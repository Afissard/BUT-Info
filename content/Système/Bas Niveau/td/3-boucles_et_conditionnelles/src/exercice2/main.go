package main

import "fmt"

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

