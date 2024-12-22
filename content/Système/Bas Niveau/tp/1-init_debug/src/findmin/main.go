package main

import "fmt"

func main() {
	var t []int = []int{12, 34, 9, 27, 14}
	for i := 0; i < len(t)-1; i++ {
		minpos := minimumPosition(t[i:])
		t[i], t[i+minpos] = t[i+minpos], t[i]
	}
	fmt.Println(t)
}

func minimumPosition(t []int) (pos int) {
	for i := 1; i < len(t); i++ {
		if t[i] < t[pos] {
			pos = i
		}
	}
	return pos
}
