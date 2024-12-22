package main

import "fmt"

func main() {
	var n int = 3 * 5
	n = n + 10
	fmt.Println(n)

	var b bool = true
	b = !b
	fmt.Println(b)

	var s string = "Bon"
	s = s + "jour"
	fmt.Println(s)

	fmt.Println(3 == 5)
	fmt.Println(3 != 5)
	fmt.Println(3 >= 5)
	fmt.Println(3 <= 5)
}
