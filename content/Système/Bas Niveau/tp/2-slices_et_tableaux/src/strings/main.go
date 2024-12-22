package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "All work and no play makes Jack a dull boy"
	str2 := str1
	str3 := str1[:]
	str4 := strings.Clone(str1)

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4)

}
