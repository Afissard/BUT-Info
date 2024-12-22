package main

import (
	"fmt"
	"app/demo_pack"
)

func main(){
	fmt.Println("heyo")
	var val string = demo_pack.Test()
	fmt.Println(val)
}