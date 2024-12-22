package main

import "fmt"


func mult7(tab *[]int){
	for i := 0; i < len(tab); i++ {
		tab[i] = i*7
	}
}

func main() {
	fmt.Print("Enter size of the array: ")
    var input int
    fmt.Scanln(&input)

	var itab []int = make([]int, input)
	mult7(&itab)
	fmt.Println(itab)
}