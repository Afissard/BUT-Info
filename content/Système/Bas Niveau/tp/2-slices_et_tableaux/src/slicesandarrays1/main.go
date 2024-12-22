package main

import (
	"fmt"
)

func print(t []int) {
	fmt.Print("{")
	for i := 0; i < len(t)-1; i++ {
		fmt.Print(t[i], ", ")
	}
	fmt.Printf("%d}\n", t[len(t)-1])
}

func swap(v1 *int, v2 *int) (r int) {
	tmp := *v1
	*v1 = *v2
	*v2 = tmp
	r = *v1 + *v2
	return
}

func main() {
	var w = [8]int{2, 3, 5, 7, 11, 13, 17, 19}

	x := w

	y := w[:]

	z := &x

	for cpt := 0; cpt < len(y); cpt++ {
		_ = swap(&((*z)[cpt]), &(y[(cpt+2)%len(y)]))
	}

	print(w[:])
	print(x[:])
	print(y)
	print((*z)[:])
}
