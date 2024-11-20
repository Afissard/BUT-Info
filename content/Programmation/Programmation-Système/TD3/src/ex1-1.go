package main

import "fmt"

// calcule n puissance m
func puissance(n, m int) (res int) {
	res = 1
	for i := 0; i < m; i++ {
		res *= n
	}
	return
}

func main() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, "puissance", j, "vaut", puissance(i, j))
		}
	}

}
