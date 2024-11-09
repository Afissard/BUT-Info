package main

import "fmt"

func echo() {
	fmt.Println("hello")
}

func main() {
	go echo()
}
