package main

import (
	"log"
	"sort"
)

func main() {

	var t []int = []int{1, 3, 2, 0, -3, 27, 5}
	log.Print(t)

	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	log.Print(t)

}
