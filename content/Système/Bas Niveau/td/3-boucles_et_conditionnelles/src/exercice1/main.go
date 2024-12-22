package main

import "fmt"

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

