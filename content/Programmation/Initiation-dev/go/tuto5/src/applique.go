package main

import "log"

func applique(t []int, f func(x int) int) {
	for i := 0; i < len(t); i++ {
		t[i] = f(t[i])
	}
}

func main() {

	var t []int = []int{1, 3, 2, 0, -3, 27, 5}
	log.Print(t)

	applique(t, func(x int) int { return 2 * x })
	log.Print(t)

}
