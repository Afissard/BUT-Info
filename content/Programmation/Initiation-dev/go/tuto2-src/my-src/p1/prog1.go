package main

import "fmt"

func plus1(tab []int){
	for i:=0; i<len(tab); i++{
		tab[i] ++
	}
	return tab
}

func main(){
	var aaa int[] = []int{1,57,545,5445}
	fmt.Println(aaa)
	aaa = plus1(aaa)
	fmt.Println(aaa)
}