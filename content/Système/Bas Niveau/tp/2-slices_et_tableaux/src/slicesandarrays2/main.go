package main

import "fmt"

func print(t []int) {
	fmt.Print("{")
	for i := 0; i < len(t)-1; i++ {
		fmt.Print(t[i], ", ")
	}
	fmt.Printf("%d}\n", t[len(t)-1])
}

func add_and_append(slin []int) (slout []int) {
	l := len(slin)
	slout = append(slin, l)
	for idx := 0; idx < len(slout); idx++ {
		slout[idx] = slout[idx] + 1
	}
	return slout
}

func main() {

	var sli0, sli1, sli2, sli3, sli4, sli5 []int

	sli1 = add_and_append(sli0)
	sli2 = add_and_append(sli1)
	sli3 = add_and_append(sli2)
	sli4 = add_and_append(sli3)
	sli5 = add_and_append(sli4)

	fmt.Println("---sli1---")
	print(sli1)
	fmt.Println("----------")

	fmt.Println("---sli2---")
	print(sli2)
	fmt.Println("----------")
	fmt.Println("---sli3---")
	print(sli3)
	fmt.Println("----------")
	fmt.Println("---sli4---")
	print(sli4)
	fmt.Println("----------")
	fmt.Println("---sli5---")
	print(sli5)
	fmt.Println("----------")
}
