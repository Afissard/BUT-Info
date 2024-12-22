package main

import "fmt"

func main() {
	var length int = 3
	var capacity int = 5
	var tab []int = make([]int, length, capacity)
	var t []int = tab
	var max int = 5
	for i := length; i <= max; i++ {
		t = append(t, i)
		t[0] = i
		fmt.Println("Tour de boucle", i)
		fmt.Println("\ttab =", tab, "len(tab) =", len(tab), "cap(tab) =", cap(tab))
		fmt.Println("\tt =", t, "len(t) =", len(t), "cap(t) =", cap(t))
	}
}
