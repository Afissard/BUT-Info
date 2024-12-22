package main

import (
	"flag"
	"fmt"
)

func main() {
	var n int
	flag.IntVar(&n, "setn", 27, "Fixe la valeur de n (défaut : 27)")
	flag.Parse()
	fmt.Println("n vaut", n)
}
