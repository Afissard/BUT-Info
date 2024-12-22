package main

import "fmt"

func f(x *int, y *int, b bool) {
	if b {
		*x += 1
		*y -= 2
	}
}

func main() {
	var n int = 0
	var m int = 10

	fmt.Println("Au début n vaut", n, "et m vaut", m)
	for i := 0; i < 10; i++ {
		f(&n, &m, i%2 == 0)
		fmt.Println("Après l'itération", i, "n vaut", n, "et m vaut", m)
	}
	fmt.Println("À la fin n vaut", n, "et m vaut", m)
}
